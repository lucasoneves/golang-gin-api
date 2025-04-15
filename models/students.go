package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string `json:"name"`
	Rg   string `json:"rg"`
	CPF  string `json:"cpf"`
}
