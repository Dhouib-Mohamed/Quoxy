package db

type IncorrectFrequency struct {
	Code int
}

func (e IncorrectFrequency) GetError() (int, string) {
	return e.Code, "invalid frequency, must be daily, monthly or a cron expression of 3 numbers or * separated by spaces"
}

func IncorrectFrequencyError() IncorrectFrequency {
	return IncorrectFrequency{Code: 400}
}
