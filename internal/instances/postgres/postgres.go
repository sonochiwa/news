package postgres

import (
	"github.com/sonochiwa/news/configs"

	_ "github.com/lib/pq"
)

func New(config configs.Postgres) (Instance, error) {
	inst := &Postgres{}

	err := inst.Connect(config)
	if err != nil {
		return nil, err
	}

	err = inst.Ping()
	if err != nil {
		inst.Close()
		return nil, err
	}

	return inst, nil
}
