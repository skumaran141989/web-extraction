package utilities

import "encoding/json"

func GetModelFromJSON[T any](jsonInput string) (T, error) {

	jsonInputByte := []byte(jsonInput)

	var model T
	err := json.Unmarshal(jsonInputByte, &model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func GetJSONFromModel[T any](model T) (string, error) {

	jsonInputByte, err := json.Marshal(model)
	if err != nil {
		return "", err
	}

	jsonString := string(jsonInputByte)

	return jsonString, nil
}
