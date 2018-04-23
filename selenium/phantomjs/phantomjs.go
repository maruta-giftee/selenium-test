package phantomjs

import (
	"github.com/maruta-giftee/selenium-test/selenium"
	"github.com/sclevine/agouti"
)

type phantomjsDriver struct {
	driver *agouti.WebDriver
}

func NewDriver() selenium.Driver {
	options := []string{
		"--disable-extensions",
		"--disable-print-preview",
		"--headless",
	}
	return agouti.PhantomJS(agouti.ChromeOptions("args", options))
}

func (d phantomjsDriver) NewPage() (*agouti.Page, error) {
	return d.driver.NewPage(agouti.Browser("phantomjs"))
}
func (d phantomjsDriver) Start() error {
	return d.driver.Start()
}

func (d phantomjsDriver) Stop() error {
	return d.driver.Stop()
}
