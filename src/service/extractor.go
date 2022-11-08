package service

import (
	"time"
)

type WebExtractor interface {
	Start(domainURL string) error
	GetByType(byType string) string
	SetValue(waitTime time.Duration, by, path string, value string) error
	ClickElement(waitTime time.Duration, by, path string) error
	SubmitElement(waitTime time.Duration, by, path string) error
	GetTextValue(waitTime time.Duration, by, path string) (string, error)
	GetArrayCount(waitTime time.Duration, by, path string) (int, error)
}
