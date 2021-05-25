package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Short:   "This Is a tool for monitoring webistes",
	Long:    "This Is a tool for monitoring webistes tooo",
	Use:     "./monitor",
	Version: "1.0.0",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
