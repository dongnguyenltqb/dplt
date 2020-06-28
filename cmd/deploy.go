package cmd

import (
	"dplt/pkg/deploy"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// environment variables
var env string
var pem string

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "a code deploy tool, written in go",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := deploy.Config{
			Env:  env,
			Pem:  pem,
			Ips:  viper.GetViper().GetStringSlice("ips"),
			Cmds: viper.GetViper().GetStringSlice("cmds"),
		}
		deploy.Deploy(cfg)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	deployCmd.PersistentFlags().StringVar(&env, "env", "dev", "node environment")
	deployCmd.PersistentFlags().StringVar(&pem, "pem", "pem.pem", "path to pem file")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
