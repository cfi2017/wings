package plugin

import (
	"errors"
	"plugin"

	"github.com/cfi2017/wings-api/pkg"
)

func LoadPlugin(api pkg.Api, path string) error {
	p, err := plugin.Open(path)
	if err != nil {
		return err
	}
	s, err := p.Lookup("InitPlugin")
	if err != nil {
		return err
	}
	if plug, ok := s.(func() pkg.Plugin); ok {
		return plug().Load(api)
	} else {
		return errors.New("symbol is not a valid Plugin")
	}
}
