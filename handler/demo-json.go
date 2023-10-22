package handler

import "encoding/json"

type employee struct {
	ID    int
	Name  string
	Tel   string
	Email string
}

func MockJson() []byte {
	emp := employee{101, "Care", "081-3968689", "test@mail.com"}
	data, _ := json.Marshal(emp)
	return data
}

func UnMarshal(emp []byte) (employee, error) {
	unm := employee{}
	err := json.Unmarshal(emp, &unm)
	if err != nil {
		return employee{}, err
	} else {
		return unm, nil
	}
}
