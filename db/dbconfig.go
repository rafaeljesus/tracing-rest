package db

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"os"
)

type config struct {
	drv  string
	open string
}

func dbConfig() (*config, error) {
	env := "development"

	file, err := yaml.ReadFile("db/dbconf.yml")
	if err != nil {
		return nil, err
	}

	drv, err := file.Get(fmt.Sprintf("%s.driver", env))
	if err != nil {
		return nil, err
	}

	open, err := file.Get(fmt.Sprintf("%s.open", env))
	if err != nil {
		return nil, err
	}

	drv = os.ExpandEnv(drv)
	open = os.ExpandEnv(open)

	return &config{drv, open}, nil
}
