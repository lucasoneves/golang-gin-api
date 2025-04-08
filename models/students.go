package models

type Student struct {
	Name string `json:"name"`
	Rg   string `json:"rg"`
	CPF  string `json:"cpf"`
}

var StudentsList []Student
