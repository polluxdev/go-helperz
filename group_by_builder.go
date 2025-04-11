package helperz

import (
	"fmt"
	"strings"
)

// GroupByBuilder defines the structure for GROUP BY clause components
type GroupByBuilder struct {
	Table    string // Table name or alias (e.g., "users", "u") (optional)
	Column   string // Column name (required)
	Function string // Function to apply (e.g., "SUM", "COUNT") (optional)
}

// ConstructGroupByClause builds a GROUP BY clause from GroupByBuilder slices
func ConstructGroupByClause(builders []GroupByBuilder) (string, error) {
	var clauses []string

	for _, b := range builders {
		if b.Column == "" {
			return "", fmt.Errorf("column cannot be empty in GroupByBuilder")
		}

		var part strings.Builder

		if b.Table != "" {
			part.WriteString(b.Table + ".")
		}

		part.WriteString(b.Column)

		if b.Function != "" {
			part.WriteString(" " + b.Function)
		}

		clauses = append(clauses, part.String())
	}

	if len(clauses) == 0 {
		return "", nil
	}

	return strings.Join(clauses, ", "), nil
}
