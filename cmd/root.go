package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zuiurs/viper-test/utils"
)

var rootCmd = &cobra.Command{
	Use: "Test",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("User: %s\n", utils.ShowUser())
		fmt.Printf("Password: %s\n", utils.ShowPassword())
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("vp-user", "", os.Getenv("VIPER_USERNAME"), "User name")
	rootCmd.PersistentFlags().StringP("vp-password", "", os.Getenv("VIPER_PASSWORD"), "User password")

	viper.BindPFlag("global.vp-user", rootCmd.PersistentFlags().Lookup("vp-user"))
	viper.BindPFlag("global.vp-password", rootCmd.PersistentFlags().Lookup("vp-password"))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
