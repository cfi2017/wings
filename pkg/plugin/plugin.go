package plugin

import (
	"plugin"
)

type Plugin interface {
	Load(api Api) error
}

func LoadPlugin(api Api, path string) error {
	p, err := plugin.Open(path)
	if err != nil {
		return err
	}
	s, err := p.Lookup("Plugin")
	if err != nil {
		return err
	}
	if plug, ok := s.(Plugin); ok {
		return plug.Load(api)
	} else {
		return err
	}
}
