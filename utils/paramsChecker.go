package utils

import "fmt"

func CheckOrderAndSortParams(order *string, sort *string) {
	if *order != "ASC" && *order != "DESC" {
		*order = "ASC"
	}
	if *sort == "" {
		*sort = "id"
	}
}

func PrepareWhereExpression(key *string, value *string) string {
	if key == nil || len(*key) == 0 {
		return "TRUE"
	}
	return fmt.Sprintf("%s = %s", *key, *value)
}
