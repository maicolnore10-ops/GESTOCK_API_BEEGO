package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func doReq(method, url string, body []byte) (int, []byte, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return 0, nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, b, nil
}

func extractIDFromList(body []byte) (interface{}, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, err
	}
	// data may be array or object
	data, ok := m["data"]
	if !ok {
		// try lowercase
		data = m["Data"]
		if data == nil {
			return nil, fmt.Errorf("no data field")
		}
	}
	arr, ok := data.([]interface{})
	if !ok {
		// could be null or empty
		return nil, fmt.Errorf("data not array")
	}
	if len(arr) == 0 {
		return nil, fmt.Errorf("empty array")
	}
	first := arr[0]
	obj, ok := first.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("first item not object")
	}
	if id, ok := obj["id"]; ok {
		return id, nil
	}
	if id, ok := obj["Id"]; ok {
		return id, nil
	}
	// try common id names
	for k, v := range obj {
		if k == "Id" || k == "id" || k == "id_backup" || k == "id_notificacion" || k == "id_seguridad" {
			return v, nil
		}
	}
	return nil, fmt.Errorf("id not found in object: %v", obj)
}

func main() {
	base := "http://localhost:8080"
	resources := []struct{
		path string
		post []byte
		put []byte
	}{
		{"/sistema", []byte(`{"NombreEmpresa":"TestCo","Correo":"test@example.com","Descripcion":"desc","Activo":true}`), []byte(`{"NombreEmpresa":"TestCoUpdated","Correo":"u@example.com","Descripcion":"desc2","Activo":false}`)},
		{"/notificacion", []byte(`{"Titulo":"Hola","Tipo":"email","Canal":"smtp","Estado":"pendiente","Activo":true}`), []byte(`{"Titulo":"Hola2","Tipo":"email","Canal":"smtp","Estado":"enviado","Activo":false}`)},
		{"/seguridad", []byte(`{"TipoAutenticacion":"local","RequiereAutenticacion":true,"Activo":true}`), []byte(`{"TipoAutenticacion":"ldap","RequiereAutenticacion":false,"Activo":false}`)},
		{"/backup", []byte(`{"NombreBackup":"bk1","TipoBackup":"full","Estado":"ok","Activo":true}`), []byte(`{"NombreBackup":"bk1-upd","TipoBackup":"inc","Estado":"ok","Activo":false}`)},
	}

	for _, r := range resources {
		fmt.Println("--- Testing", r.path)
		// POST
		status, body, err := doReq("POST", base+r.path, r.post)
		fmt.Println("POST status:", status, "body:", string(body), "err:", err)
		// GET all
		status, body, err = doReq("GET", base+r.path, nil)
		fmt.Println("GET status:", status, "body:", string(body), "err:", err)
		id, err := extractIDFromList(body)
		if err != nil {
			fmt.Println("Could not extract id from list:", err)
			continue
		}
		fmt.Println("Got id:", id)
		// PUT
		putURL := fmt.Sprintf("%s%s/%v", base, r.path, id)
		status, body, err = doReq("PUT", putURL, r.put)
		fmt.Println("PUT status:", status, "body:", string(body), "err:", err)
		// DELETE
		delURL := putURL
		status, body, err = doReq("DELETE", delURL, nil)
		fmt.Println("DELETE status:", status, "body:", string(body), "err:", err)
	}
}
