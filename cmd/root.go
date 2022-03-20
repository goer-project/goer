package cmd

import (
	goerConfig "github.com/goer-project/goer-core/config"
	"github.com/goer-project/goer/cmd/make"
	"github.com/goer-project/goer/cmd/migrate"
	"github.com/goer-project/goer/config"
	"github.com/goer-project/goer/database"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goer",
	Short: "Api framework in Golang",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

var (
	cfgFile string
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.toml)")

	// Init
	goerConfig.InitConfig(cfgFile, &config.NewConfig) // Init viper
	database.DB = database.Gorm()

	// Add sub command
	rootCmd.AddCommand(
		make.CmdMake,
		migrate.CmdMigrate,
	)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {}
