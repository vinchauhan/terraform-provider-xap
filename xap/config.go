package xap

import "github.com/vinchauhan/terraform-provider-xap/xap/xapapi"

//Config struct
type Config struct {
	endpoint          string
	User              string
	Password          string
	SkipSslValidation bool
}

func (c *Config) Client() (*xapapi.Session, error) {
	return xapapi.NewSession(c.endpoint, c.User, c.Password, c.SkipSslValidation)
}
