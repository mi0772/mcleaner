package main

import (
	"fmt"
	"os"

	"cdigiuseppe/mcleaner/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mcleaner",
	Short: "A utility to clean junk files on macOS",
	Long: `mcleaner is a command line tool to clean junk files and cache on macOS.
It can perform scans, clean cache, remove .DS_Store files and handle maintenance tasks.`,
}

func main() {
	cmd.SetRootCmd(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}