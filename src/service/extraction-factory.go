package service

import (
	"errors"

	"github.com/skumaran141989/web-extraction/src/service/approaches"
	"github.com/skumaran141989/web-extraction/src/utilities/constants"
)

func GetWebExtractorFactory(identifier string, attributes map[string]string) (WebExtractor, error) {
	switch identifier {
	case constants.TEBEKA_SELENIUM:
		return approaches.NewTebekaSelenium(attributes)
	}

	return nil, errors.New(constants.ERROR_UNKNOWN_WEB_EXTRACTOR)
}
