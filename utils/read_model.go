package utils

import (
	"encoding/json"
	"io/ioutil"
)


func ReadModel(path string)(map[string]interface{}, error){
	var model map[string]interface{}
	// read file in path
	model_json, err := ioutil.ReadFile(path)
	if(err != nil){
		return model, err
	}

	// marshal model_json
	err = json.Unmarshal(model_json, &model)

	return model, err
}