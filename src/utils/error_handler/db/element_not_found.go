package db

type ElementNotFound struct {
	Code  int
	Table string
}

func (e ElementNotFound) GetError() (int, string) {
	return e.Code, "No Element Found in the " + e.Table + " table"
}

func ElementNotFoundError(table string) ElementNotFound {
	return ElementNotFound{Code: 404, Table: table}
}
