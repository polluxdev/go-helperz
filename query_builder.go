package helperz

type QueryBuilderHelper interface {
	ConstructSelectClause(builders []SelectBuilder) ([]string, error)
	ConstructConditionalClause(builders []ConditionalBuilder) (string, []interface{}, error)
	ConstructJoinClause(builders []JoinBuilder) ([]string, error)
	ConstructGroupByClause(builders []GroupByBuilder) (string, error)
}
