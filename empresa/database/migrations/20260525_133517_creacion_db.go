package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionDb_20260525_133517 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionDb_20260525_133517{}
	m.Created = "20260525_133517"

	migration.Register("CreacionDb_20260525_133517", m)
}

// Run the migrations
func (m *CreacionDb_20260525_133517) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CreacionDb_20260525_133517) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
