package cmd

import (
	"fmt"
	"os"

	"github.com/mijime/css-demo/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "css-demo",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := server.NewApp(server.AppOptions{
			Debug: viper.GetBool("debug"),
		})
		app.Run(viper.GetString("addr"))
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.css-demo.yaml)")
	RootCmd.Flags().StringP("addr", "a", "127.0.0.1:3000", "Help message for toggle")
	RootCmd.Flags().BoolP("debug", "d", false, "Help message for toggle")
	viper.BindPFlag("addr", RootCmd.Flags().Lookup("addr"))
	viper.BindPFlag("debug", RootCmd.Flags().Lookup("debug"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".css-demo")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
