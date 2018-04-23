package chrome

import (
	"github.com/maruta-giftee/selenium-test/selenium"
	"github.com/sclevine/agouti"
)

type chromeDriver struct {
	driver *agouti.WebDriver
}

func NewDriver() selenium.Driver {
	options := []string{
		"--disable-extensions",
		"--disable-print-preview",
		"--window-size=1280,800",
		"--headless",
	}
	return agouti.ChromeDriver(agouti.ChromeOptions("args", options))
}

func (d chromeDriver) NewPage() (*agouti.Page, error) {
	return d.driver.NewPage(agouti.Browser("chrome"))
}
func (d chromeDriver) Start() error {
	return d.driver.Start()
}

func (d chromeDriver) Stop() error {
	return d.driver.Stop()
}
