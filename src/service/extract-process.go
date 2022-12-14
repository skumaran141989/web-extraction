package service

import (
	"errors"
	"fmt"

	"github.com/skumaran141989/web-extraction/src/service/models"
	"github.com/skumaran141989/web-extraction/src/utilities"
	"github.com/skumaran141989/web-extraction/src/utilities/constants"
)

const (
	WAIT_TIME = 10000
)

type WebExtraction struct {
	extractor WebExtractor
}

func NewWebExtraction(WebExtractorType string, attributes map[string]string) *WebExtraction {
	extractor, _ := GetWebExtractorFactory(WebExtractorType, attributes)

	return &WebExtraction{extractor: extractor}
}

func (extraction *WebExtraction) Extract(url string, fieldString string) (string, error) {
	var err error

	fields, err := utilities.GetModelFromJSON(fieldString)
	if err != nil {
		return "", err
	}

	var output_fields []models.Field

	err = extraction.extractor.Start(url)
	if err != nil {
		return "", err
	}

	for _, field := range fields {
		byType := extraction.extractor.GetByType(field.Type)
		path := field.Path
		array_fields = nil

		switch field.ActionType {
		case constants.SET_VALUE:
			err = extraction.extractor.SetValue(WAIT_TIME, byType, path, field.Value)
		case constants.GET_TEXT_VALUE:
			field.Value, err = extraction.extractor.GetTextValue(WAIT_TIME, byType, path)
		case constants.CLICK:
			err = extraction.extractor.ClickElement(WAIT_TIME, byType, path)
		case constants.SUBMIT:
			err = extraction.extractor.SubmitElement(WAIT_TIME, byType, path)
		case constants.GET_ARRAY:
			err = getArrayValue(extraction, field, output_fields)
		}

		if err != nil {
			field.Error = err
		}

		output_fields = append(output_fields, field)
	}

	outputFieldsJSON, err := utilities.GetJSONFromModel[[]models.Field](output_fields)
	if err != nil {
		return "", err
	}

	return outputFieldsJSON, nil
}

func getArrayValue(extraction *WebExtraction, field models.Field, outputFields []models.Field) error {
	var err error
	if field.ArrayPath != "" {
		err = errors.New(constants.ERROR_ARRAY_DOES_NOT_EXIST)
		field.Error = err
		return nil, err
	}

	array_len, err := extraction.extractor.GetArrayCount(WAIT_TIME, field.Type, field.ArrayPath)
	if field.ArrayPath != "" {
		return nil, err
	}
	if array_len == 0 {
		err = errors.New(constants.ERROR_ARRAY_PATH_MISSING)
		field.Error = err
		return nil, err
	}

	fields := make([]models.Field, array_len)

	i := 0
	for i < array_len {

		field_name := fmt.Sprintf("%s[%d]", field.Name, i)
		path := fmt.Sprintf("%s[%d]", field.Path, i)
		field := models.Field{
			Name:      field_name,
			Path:      path,
			ArrayPath: field.ArrayPath,
		}

		field.Value, err = extraction.extractor.GetTextValue(WAIT_TIME, constants.GET_TEXT_VALUE, fmt.Sprintf("%s/%s", field.Path, field.ArrayPath))
		if err != nil {
			field.Error = err
		}

		fields = append(outputFields, field)

		i++
	}

	return fields, nil
}
