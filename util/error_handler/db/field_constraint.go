package db

type FieldConstraint struct {
	Code       int
	Table      string
	Field      string
	Constraint string
}

func (e FieldConstraint) GetError() (int, string) {
	return e.Code, "constraint violation on field " + e.Field + " in the " + e.Table + " table - " + e.Field + " " + e.Constraint
}

func FieldConstraintError(table string, field string, constraint string) FieldConstraint {
	return FieldConstraint{Code: 409, Table: table, Field: field, Constraint: constraint}
}
