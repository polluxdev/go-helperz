package helperz

import "fmt"

// JoinBuilder defines the structure for JOIN clause components
type JoinBuilder struct {
	Type            string // Join type ("INNER", "LEFT", etc.) (required)
	Table           string // Table to join (required)
	Column          string // Table column (required)
	ReferenceTable  string // Reference table (required)
	ReferenceColumn string // Reference table column (required)
}

// ConstructJoinClause builds JOIN clauses from JoinBuilder slices
func ConstructJoinClause(builders []JoinBuilder) ([]string, error) {
	var clauses []string

	for _, b := range builders {
		if b.Type == "" || b.Table == "" || b.Column == "" || b.ReferenceTable == "" || b.ReferenceColumn == "" {
			return nil, fmt.Errorf("type, table, column, reference table, and reference column cannot be empty in JoinBuilder")
		}

		clause := fmt.Sprintf("%s JOIN %s ON %s.%s = %s.%s",
			b.Type,
			b.Table,
			b.Table,
			b.Column,
			b.ReferenceTable,
			b.ReferenceColumn,
		)
		clauses = append(clauses, clause)
	}

	return clauses, nil
}
