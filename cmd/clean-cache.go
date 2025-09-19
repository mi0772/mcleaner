package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var cleanCacheCmd = &cobra.Command{
	Use:   "clean-cache",
	Short: "Clean system cache",
	Long:  `Remove cache files to free disk space`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cleaning cache...")
		
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error finding home directory: %v\n", err)
			return
		}

		var totalCleaned int64
		
		cacheDirs := []string{
			filepath.Join(homeDir, "Library/Caches"),
			"/Library/Caches",
			"/tmp",
		}

		for _, dir := range cacheDirs {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				continue
			}

			fmt.Printf("Cleaning: %s\n", dir)
			size, err := cleanDirectory(dir, false)
			if err != nil {
				fmt.Printf("Error cleaning %s: %v\n", dir, err)
				continue
			}
			totalCleaned += size
		}

		fmt.Printf("\nCleaning completed!\n")
		fmt.Printf("Space freed: %s\n", formatSize(totalCleaned))
	},
}

func cleanDirectory(dirPath string, removeDir bool) (int64, error) {
	var totalSize int64

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if !info.IsDir() {
			totalSize += info.Size()
			if err := os.Remove(path); err != nil {
				return nil
			}
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	if removeDir && dirPath != "/" && dirPath != "/tmp" {
		filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if info.IsDir() && path != dirPath {
				os.RemoveAll(path)
			}
			return nil
		})
	}

	return totalSize, nil
}

func init() {
	if rootCmd != nil {
		rootCmd.AddCommand(cleanCacheCmd)
	}
}