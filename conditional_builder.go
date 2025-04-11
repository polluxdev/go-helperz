package helperz

import (
	"fmt"
	"strings"
)

// ConditionalBuilder defines the structure for WHERE clause components
type ConditionalBuilder struct {
	Table         string      // Table name or alias (e.g., "users", "u") (optional)
	Column        string      // Column name (required)
	Value         interface{} // Comparison value (optional for NULL checks)
	FunctionValue string      // Function to apply to value (e.g., "EXTRACT(MONTH FROM ?)", defaults to "?")
	Logical       string      // Logical operator (e.g., "=", ">", "LIKE") (required)
	Operator      string      // Conditional operator ("AND", "OR") (required)
}

// ConstructConditionalClause builds a WHERE clause from ConditionalBuilder slices
func ConstructConditionalClause(builders []ConditionalBuilder) (string, []interface{}, error) {
	if len(builders) == 0 {
		return "1 = 1", nil, nil
	}

	var (
		clauses []string
		args    []interface{}
	)

	for i, b := range builders {
		clause, clauseArgs, err := buildSingleCondition(b, i == 0)
		if err != nil {
			return "", nil, err
		}
		clauses = append(clauses, clause)
		args = append(args, clauseArgs...)
	}

	return strings.Join(clauses, " "), args, nil
}

func buildSingleCondition(b ConditionalBuilder, isFirst bool) (string, []interface{}, error) {
	if err := validateCondition(b); err != nil {
		return "", nil, err
	}

	if b.FunctionValue == "" {
		b.FunctionValue = "?"
	}

	column := buildColumnName(b.Table, b.Column)

	switch b.Logical {
	case "LIKE":
		return buildLikeCondition(column, b, isFirst)
	case "IN":
		return buildInCondition(column, b, isFirst)
	case "BETWEEN":
		return buildBetweenCondition(column, b, isFirst)
	case "JSON_OVERLAPS":
		return buildJsonOverlapsCondition(column, b, isFirst)
	default:
		return buildDefaultCondition(column, b, isFirst)
	}
}

func validateCondition(b ConditionalBuilder) error {
	if b.Column == "" || b.Logical == "" || b.Operator == "" {
		return fmt.Errorf("column, logical, and operator cannot be empty in ConditionalBuilder")
	}
	return nil
}

func buildColumnName(table, column string) string {
	if table != "" {
		return table + "." + column
	}
	return column
}

func buildLikeCondition(column string, b ConditionalBuilder, isFirst bool) (string, []interface{}, error) {
	operator := ""
	if !isFirst {
		operator = b.Operator + " "
	}
	clause := fmt.Sprintf("%sLOWER(%s) %s %s", operator, column, b.Logical, b.FunctionValue)
	return clause, []interface{}{fmt.Sprintf("%%%s%%", strings.ToLower(fmt.Sprint(b.Value)))}, nil
}

func buildInCondition(column string, b ConditionalBuilder, isFirst bool) (string, []interface{}, error) {
	operator := ""
	if !isFirst {
		operator = b.Operator + " "
	}
	clause := fmt.Sprintf("%s%s %s (%s)", operator, column, b.Logical, b.FunctionValue)
	return clause, []interface{}{b.Value}, nil
}

func buildBetweenCondition(column string, b ConditionalBuilder, isFirst bool) (string, []interface{}, error) {
	values, ok := b.Value.([]interface{})
	if !ok || len(values) != 2 {
		return "", nil, fmt.Errorf("BETWEEN requires a slice of 2 values")
	}

	operator := ""
	if !isFirst {
		operator = b.Operator + " "
	}
	clause := fmt.Sprintf("%s%s %s %s AND %s", operator, column, b.Logical, b.FunctionValue, b.FunctionValue)
	return clause, values, nil
}

func buildJsonOverlapsCondition(column string, b ConditionalBuilder, isFirst bool) (string, []interface{}, error) {
	operator := ""
	if !isFirst {
		operator = b.Operator + " "
	}
	clause := fmt.Sprintf("%s%s(%s, %s)", operator, b.Logical, column, b.FunctionValue)
	return clause, []interface{}{b.Value}, nil
}

func buildDefaultCondition(column string, b ConditionalBuilder, isFirst bool) (string, []interface{}, error) {
	operator := ""
	if !isFirst {
		operator = b.Operator + " "
	}

	var clause string
	var args []interface{}

	if b.Value != nil {
		clause = fmt.Sprintf("%s%s %s %s", operator, column, b.Logical, b.FunctionValue)
		args = []interface{}{b.Value}
	} else {
		clause = fmt.Sprintf("%s%s %s NULL", operator, column, b.Logical)
	}

	return clause, args, nil
}
