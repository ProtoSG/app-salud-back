package utils

import "fmt"

type EntityNotFound struct {
	value     any
	attribute string
	table     string
}

func NewEntityNotFound(value any, attribute, table string) *EntityNotFound {
	return &EntityNotFound{
		value:     value,
		attribute: attribute,
		table:     table,
	}
}

func (this *EntityNotFound) Error() string {
	return fmt.Sprintf("%s con %s %v no encontrado", this.table, this.attribute, this.value)
}
