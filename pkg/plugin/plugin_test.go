package plugin

import (
	"testing"

	"github.com/pterodactyl/wings/server"
	"github.com/pterodactyl/wings/server/backup"
)

type MockApi struct {
	handlers map[string][]interface{}
	ponged   bool
}

func (m *MockApi) RegisterHandler(name string, handler interface{}) {
	if _, ok := m.handlers[name]; !ok {
		m.handlers[name] = make([]interface{}, 0)
	}
	m.handlers[name] = append(m.handlers[name], handler)
}

func (m *MockApi) RegisterBackupStrategy(_ string, _ backup.BackupInterface) {
	panic("implement me")
}

func (m *MockApi) RegisterEnvironment(_ string, _ server.Environment) {
	panic("implement me")
}

func (m *MockApi) CallHandler(name string) {
	if hs, ok := m.handlers[name]; ok {
		for _, h := range hs {
			if f, ok := h.(func(Api)); ok {
				f(m)
			}
		}
	}
}

func TestLoadPlugin(t *testing.T) {
	type args struct {
		api  *MockApi
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "events good path", args: struct {
			api  *MockApi
			path string
		}{api: &MockApi{handlers: map[string][]interface{}{
			"pong": {pong},
		}}, path: "examples/events/events.so"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadPlugin(tt.args.api, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("LoadPlugin() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.args.api.CallHandler("ping")
			if !tt.args.api.ponged {
				t.Errorf("LoadPlugin() want ponged = true")
			}
		})
	}
}

func pong(api Api) {
	if ma, ok := api.(*MockApi); ok {
		ma.ponged = true
	}
}
