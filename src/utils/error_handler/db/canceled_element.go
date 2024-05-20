package db

type CanceledElement struct {
	Code  int
	Table string
}

func (e CanceledElement) GetError() (int, string) {
	return e.Code, e.Table + " is out of service"
}

func CanceledElementError(table string) CanceledElement {
	return CanceledElement{Code: 410, Table: table}
}
