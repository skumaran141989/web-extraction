package approaches

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

type TebekaSelenium struct {
	driver selenium.WebDriver
}

func NewTebekaSelenium(attributes map[string]string) (*TebekaSelenium, error) {
	width, _ := strconv.Atoi(attributes["width"])
	height, _ := strconv.Atoi(attributes["height"])
	server_url := attributes["server_url"]

	args := []string{
		fmt.Sprintf("--window-size=%d,%d", width, height),
		"--ignore-certificate-errors",
		"--disable-extensions",
		"--no-sandbox",
		"--disable-dev-shm-usage",
	}
	capabilities := selenium.Capabilities{
		"browserName":   "chrome",
		"chromeOptions": args,
	}

	driver, err := selenium.NewRemote(capabilities, server_url)
	if err != nil {
		return nil, err
	}

	return &TebekaSelenium{
		driver: driver,
	}, nil

}

func (tebekaSelenium *TebekaSelenium) Start(domainURL string) error {
	return tebekaSelenium.driver.Get(domainURL)
}

func (tebekaSelenium *TebekaSelenium) GetByType(byType string) string {
	upperByType := strings.ToUpper(byType)

	switch upperByType {
	case "CSS":
		return selenium.ByCSSSelector
	case "ID":
		return selenium.ByID
	case "XPATH":
		return selenium.ByXPATH
	case "TEXT":
		return selenium.ByPartialLinkText
	}

	return ""
}

func (tebekaSelenium *TebekaSelenium) SetValue(waitTime time.Duration, by, path string, value string) error {
	element, err := waitFindElement(tebekaSelenium.driver, waitTime, by, path)
	if err != nil {
		return err
	}
	element.SendKeys(value)
	return nil
}

func (tebekaSelenium *TebekaSelenium) ClickElement(waitTime time.Duration, by, path string) error {
	element, err := waitFindElement(tebekaSelenium.driver, waitTime, by, path)
	if err != nil {
		return err
	}
	element.Click()

	return nil
}

func (tebekaSelenium *TebekaSelenium) SubmitElement(waitTime time.Duration, by, path string) error {
	element, err := waitFindElement(tebekaSelenium.driver, waitTime, by, path)
	if err != nil {
		return err
	}
	element.Submit()

	return nil
}

func (tebekaSelenium *TebekaSelenium) GetArrayCount(waitTime time.Duration, by, path string) (int, error) {
	elements, err := waitFindElements(tebekaSelenium.driver, waitTime, by, path)
	if err != nil {
		return -1, err
	}

	return len(elements), nil
}

func (tebekaSelenium *TebekaSelenium) GetTextValue(waitTime time.Duration, by, path string) (string, error) {
	element, err := waitFindElement(tebekaSelenium.driver, waitTime, by, path)
	if err != nil {
		return "", err
	}

	value, err := element.Text()
	if err != nil {
		return "", err
	}

	return value, nil
}

func waitFindElement(driver selenium.WebDriver, waitTime time.Duration, by, path string) (selenium.WebElement, error) {
	driver.SetImplicitWaitTimeout(time.Second)
	defer driver.SetImplicitWaitTimeout(waitTime)
	element, err := driver.FindElement(by, path)
	if err != nil {
		return nil, err
	}

	return element, nil
}

func waitFindElements(driver selenium.WebDriver, waitTime time.Duration, by, path string) ([]selenium.WebElement, error) {
	driver.SetImplicitWaitTimeout(time.Second)
	defer driver.SetImplicitWaitTimeout(waitTime)
	elements, err := driver.FindElements(by, path)
	if err != nil {
		return nil, err
	}

	return elements, nil
}
