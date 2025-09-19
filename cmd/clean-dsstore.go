package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var cleanDSStoreCmd = &cobra.Command{
	Use:   "clean-dsstore",
	Short: "Remove all .DS_Store files",
	Long:  `Find and remove all .DS_Store files from the system`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Removing .DS_Store files...")
		
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error finding home directory: %v\n", err)
			return
		}

		var removedCount int
		var totalSize int64

		err = filepath.Walk(homeDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}

			if strings.HasSuffix(path, ".DS_Store") {
				fmt.Printf("Removing: %s\n", path)
				totalSize += info.Size()
				if err := os.Remove(path); err != nil {
					fmt.Printf("Error removing %s: %v\n", path, err)
				} else {
					removedCount++
				}
			}
			return nil
		})

		if err != nil {
			fmt.Printf("Error during scan: %v\n", err)
			return
		}

		fmt.Printf("\nRemoval completed!\n")
		fmt.Printf(".DS_Store files removed: %d\n", removedCount)
		fmt.Printf("Space freed: %s\n", formatSize(totalSize))
	},
}

func init() {
	if rootCmd != nil {
		rootCmd.AddCommand(cleanDSStoreCmd)
	}
}