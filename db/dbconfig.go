package db

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"os"
)

type config struct {
	drv, open string
}

func dbConfig() (*config, error) {
	env := "development"

	file, err := yaml.ReadFile("db/dbconf.yml")
	if err != nil {
		return nil, err
	}

	drv, _ := file.Get(fmt.Sprintf("%s.driver", env))
	open, _ := file.Get(fmt.Sprintf("%s.open", env))

	drv = os.ExpandEnv(drv)
	open = os.ExpandEnv(open)

	return &config{drv, open}, nil
}
