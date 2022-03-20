package migrate

import (
	"github.com/goer-project/goer/config"
	"github.com/goer-project/goer/migrate"
	"github.com/spf13/cobra"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run the database migrations",
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateRollback,
		CmdMigrateRefresh,
		CmdMigrateReset,
		CmdMigrateFresh,
	)
}

func migrator() *migrate.Migrator {
	return migrate.NewMigrator(config.NewDir.Migration + "/")
}
