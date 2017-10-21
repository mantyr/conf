package tmpl

import (
	"encoding/json"
)

type Json_error_data struct {
	Message string
	Status  string
}

var (
	Json_error_data_en string
	Json_error_data_ru string
)

func init() {
	data := Json_error_data{
		"Bad request, please try again later",
		"ERROR",
	}
	json_srror_data, _ := json.Marshal(data)
	Json_error_data_en = string(json_srror_data)

	data.Message = "Произошла ошибка, попробуйте позже"
	json_srror_data, _ = json.Marshal(data)
	Json_error_data_ru = string(json_srror_data)
}
