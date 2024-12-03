package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/org/example/opt"
)


func MakeSelectStmt(table string, limit int, offset int, filter opt.Option[Filter]) (string, []any) {
	query := "SELECT * FROM " + table
	var args []any

	if filter.IsSome() {
		f := filter.Unwrap()
		andClauses := []string{}
		orClauses := []string{}
		argIndex := 1

		for key, value := range f.Where.And {
			if strings.Contains(string(value.(string)), "%") {
				andClauses = append(andClauses, key+" LIKE $"+fmt.Sprint(argIndex))
			} else {
				andClauses = append(andClauses, key+" = $"+fmt.Sprint(argIndex))
			}

			args = append(args, value)
			argIndex++
		}

		for key, value := range f.Where.Or {
			if strings.Contains(string(value.(string)), "%") {
				orClauses = append(orClauses, key+" LIKE $"+fmt.Sprint(argIndex))
			} else {
				orClauses = append(orClauses, key+" = $"+fmt.Sprint(argIndex))
			}

			args = append(args, value)
			argIndex++
		}

		if len(andClauses) > 0 || len(orClauses) > 0 {
			query += " WHERE "
		}

		if len(andClauses) > 0 {
			query += "("
			query += strings.Join(andClauses, " AND ")
			query += ")"
		}

		if len(orClauses) > 0 {
			if len(andClauses) > 0 {
				query += " OR "
			}

			query += "("
			query += strings.Join(orClauses, " OR ")
			query += ")"
		}

		if len(f.GroupBy) > 0 {
			query += " GROUP BY " + strings.Join(f.GroupBy, ", ")
		}
		if len(f.OrderBy) > 0 {
			query += " ORDER BY " + strings.Join(f.OrderBy, ", ")
		}
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, limit, offset)

	return query, args
}

func MakeInsertStmt(dataStruct any) (string, []any, error) {
	v := reflect.ValueOf(dataStruct)
	t := reflect.TypeOf(dataStruct)

	// Ensure the input is a struct
	if v.Kind() != reflect.Struct {
		return "", nil, fmt.Errorf("input must be a struct")
	}

	// Gather column names, placeholders, and values
	var columns []string
	var placeholders []string
	var values []any

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		field := t.Field(i)
		column := field.Tag.Get("db")

		// If no "db" tag is specified, use the struct field name as the column name
		if column == "" {
			column = field.Name
		}

		// Skip zero-valued fields
		if fieldValue.IsZero() {
			continue
		}

		columns = append(columns, column)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(placeholders)+1))
		values = append(values, fieldValue.Interface())
	}

	// If no values are present, return an error
	if len(columns) == 0 {
		return "", nil, fmt.Errorf("no non-zero fields to insert")
	}

	// Build the SQL INSERT statement
	columnsStr := strings.Join(columns, ", ")
	placeholdersStr := strings.Join(placeholders, ", ")
	tableName := toSnakeCase(t.Name())
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, columnsStr, placeholdersStr)

	return query, values, nil
}