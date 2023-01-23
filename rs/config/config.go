package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Env   string `required:"true" envconfig:"GO_ENV" default:"local"`
	Port  string `required:"true" envconfig:"PORT" default:"8081"`
	DBUrl string `required:"true" envconfig:"DB_URL" default:"postgres://hasura:secret@/postgres?host=/var/run/postgresql"`
}

func (c *Config) IsLocal() bool {
	return c.Env == "local"
}

func Init() (Config, error) {
	c := &Config{}
	err := envconfig.Process("", c)
	return *c, err
}
