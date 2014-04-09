package provider

import "github.com/zinic/dplug/model"

type Credentials struct {
	user     string
	password string
}

type CloudProvider interface {
	List(c *Credentials) []model.Host
	Get(c *Credentials, id string) model.Host
	Delete(c *Credentials, id string) bool
	Create(c *Credentials, requirements *model.Host) bool
}
