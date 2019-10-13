package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go-boilerplate-project/config"
)

var rootCmd = &cobra.Command{
	Use:   "bolerplate",
	Short: "Web application server",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error while reading config file: %s\n", err))
	}

	err = viper.Unmarshal(&config.Const)
	if err != nil {
		panic(fmt.Errorf("Couldn't deserialize into config object: %s\n", err))
	}
}
