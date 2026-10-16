package app

import (
	"://github.com"
	"://github.com"
	"://github.com"
	"://github.com"
)

var (
	DefaultNodeHome = "/.mycustoml1"
	ModuleBasics    = module.NewBasicManager(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
	)
)

type App struct {
	LegacyAmino *codec.LegacyAmino
	AppCodec    codec.Codec
}
