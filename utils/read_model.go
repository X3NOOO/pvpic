package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/X3NOOO/pvpic/pvpic"
)


func ReadModel(path string)(pvpic.Model, error){
	var model pvpic.Model
	// read file in path
	model_json, err := ioutil.ReadFile(path)
	if(err != nil){
		return model, err
	}

	// marshal model_json
	err = json.Unmarshal(model_json, &model)

	return model, err
}