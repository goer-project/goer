package migrate

import (
	goerConfig "github.com/goer-project/goer-core/config"
	"github.com/goer-project/goer/config"
	"github.com/goer-project/goer/database"
	"github.com/goer-project/goer/migrate"
	"github.com/spf13/cobra"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run the database migrations",
}

var (
	cfgFile string
)

func init() {
	CmdMigrate.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.toml)")

	// Init
	goerConfig.InitConfig(cfgFile, &config.NewConfig) // Init viper
	database.DB = database.Gorm()

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
