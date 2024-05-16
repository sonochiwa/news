package global

import (
	"github.com/sonochiwa/news/configs"
	"github.com/sonochiwa/news/internal/instances"
)

type Context interface {
	Config() *configs.Config
	Inst() *instances.Instances
}

type gCtx struct {
	config *configs.Config
	inst   *instances.Instances
}

func (g *gCtx) Config() *configs.Config {
	return g.config
}

func (g *gCtx) Inst() *instances.Instances {
	return g.inst
}

func New(config *configs.Config) Context {
	return &gCtx{
		config: config,
		inst:   &instances.Instances{},
	}
}
