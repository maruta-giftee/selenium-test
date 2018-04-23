package selenium

import "github.com/sclevine/agouti"

type Driver interface {
	NewPage(options ...agouti.Option) (*agouti.Page, error)
	Start() error
	Stop() error
}
