package plugin

import (
	"github.com/pterodactyl/wings/server"
	"github.com/pterodactyl/wings/server/backup"
)

type Api interface {
	RegisterHandler(name string, handler interface{})
	RegisterBackupStrategy(name string, strategy backup.BackupInterface)
	RegisterEnvironment(name string, environment server.Environment)
	CallHandler(name string)
}
