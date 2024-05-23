package utils

import "strings"

func ModifyQuery(updateQuery strings.Builder) string {
	query := updateQuery.String()
	query = strings.TrimSuffix(query, ", ")
	query = strings.TrimSuffix(query, ",")
	query += " "

	return query
}
