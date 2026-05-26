package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Backup_20260525_210652 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Backup_20260525_210652{}
	m.Created = "20260525_210652"

	migration.Register("Backup_20260525_210652", m)
}

// Run the migrations
func (m *Backup_20260525_210652) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Backup_20260525_210652) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
