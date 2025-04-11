package helperz

import (
	"fmt"
	"strings"
)

// SelectBuilder defines the structure for SELECT clause components
type SelectBuilder struct {
	Table    string // Table name or alias (e.g., "users", "u") (optional)
	Column   string // Column name (required)
	Function string // Aggregate function (optional)
	Alias    string // Column alias (optional)
}

// ConstructSelectClause builds a SELECT clause from SelectBuilder slices
func ConstructSelectClause(builders []SelectBuilder) ([]string, error) {
	var clauses []string

	for _, b := range builders {
		if b.Column == "" {
			return nil, fmt.Errorf("column cannot be empty in SelectBuilder")
		}

		var part strings.Builder

		// Apply function if specified
		if b.Function != "" {
			part.WriteString(b.Function + "(")
		}

		// Add table prefix if specified
		if b.Table != "" {
			part.WriteString(b.Table + ".")
		}

		// Add Columns
		part.WriteString(b.Column)

		// Close function if applied
		if b.Function != "" {
			part.WriteString(")")
		}

		// Add alias if specified
		if b.Alias != "" {
			part.WriteString(" AS " + b.Alias)
		}

		clauses = append(clauses, part.String())
	}

	return clauses, nil
}
