package opai

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"opai/pkg/opai"
	"os"
	"path/filepath"
)

var (
	cfgFile string
	token   string
	rootCmd = &cobra.Command{
		Use:   "opai",
		Short: "opai - A CLI for OpenAI",
		Long:  "opai - CLI for OpenAI completions",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			res, err := opai.Complete(args[0])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(res)
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.opai/config.yaml)")
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "OpenAI API token (default from config file")
	rootCmd.PersistentFlags().StringP("model", "m", "text-davinci-003", "model name")
	rootCmd.PersistentFlags().IntP("max-tokens", "", 256, "max tokens")
	rootCmd.PersistentFlags().Float32P("temperature", "t", 0.1, "temperature")
	rootCmd.PersistentFlags().VisitAll(func(flag *pflag.Flag) {
		if err := viper.BindPFlag(flag.Name, flag); err != nil {
			log.Fatal(err)
		}
	})
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		userHome, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(filepath.Join(userHome, ".opai"))
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Fatal(err)
	}
	if token != "" {
		viper.Set("api.token", token)
	}
}
