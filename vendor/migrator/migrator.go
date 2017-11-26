package migrator

import "fmt"

type Migrator struct {
	Name string
	Id   int
}

func NewMigrator() *Migrator {
	m := Migrator{"test", 1}

	return &m
}

func (m *Migrator) Migrate() {
	fmt.Println(m.Name)
}
