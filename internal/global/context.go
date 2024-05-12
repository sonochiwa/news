package global

import (
	"news/configs"
	"news/internal/instance"
)

type Context interface {
	Config() *configs.Config
	Inst() *instance.Instances
}

type gCtx struct {
	config *configs.Config
	inst   *instance.Instances
}

func (g *gCtx) Config() *configs.Config {
	return g.config
}

func (g *gCtx) Inst() *instance.Instances {
	return g.inst
}

func New(config *configs.Config) Context {
	return &gCtx{
		config: config,
		inst:   &instance.Instances{},
	}
}
