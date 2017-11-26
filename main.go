package main

import (
	"fmt"
	"migrator"
)

// need host, port, and login creds
// read/apply sql scripts from ./sql/

// func init() {

// }

// func migrate() {

// }

func main() {
	// we should have the relavent info for accessing the DB here,
	// and pass it into the constroctor, so {{ m }} here should be a pointer
	// to the address of the Migrtor ready to go, i.e. we can just run
	// m.Migrate() to apply all sql files to the DB
	m := migrator.NewMigrator()
	m.Migrate()
	fmt.Println(m.Name)
}
