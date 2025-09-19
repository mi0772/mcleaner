package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var cleanTempCmd = &cobra.Command{
	Use:   "clean-temp",
	Short: "Remove temporary files",
	Long:  `Find and remove temporary files from the system`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Removing temporary files...")
		
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error finding home directory: %v\n", err)
			return
		}

		var removedCount int
		var totalSize int64

		tempDirs := []string{
			"/tmp",
			"/var/tmp",
			filepath.Join(homeDir, "Downloads"),
			filepath.Join(homeDir, "Library/Caches/Temporary Items"),
		}

		for _, dir := range tempDirs {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				continue
			}

			fmt.Printf("Scanning directory: %s\n", dir)
			count, size := cleanTempFiles(dir)
			removedCount += count
			totalSize += size
		}

		fmt.Printf("\nRemoval completed!\n")
		fmt.Printf("Temporary files removed: %d\n", removedCount)
		fmt.Printf("Space freed: %s\n", formatSize(totalSize))
	},
}

func cleanTempFiles(dirPath string) (int, int64) {
	var removedCount int
	var totalSize int64

	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		shouldRemove := false
		fileName := strings.ToLower(info.Name())

		if strings.HasSuffix(fileName, ".tmp") ||
		   strings.HasSuffix(fileName, ".temp") ||
		   strings.Contains(fileName, "temp") ||
		   strings.HasPrefix(fileName, "tmp") {
			shouldRemove = true
		}

		if strings.Contains(dirPath, "Downloads") {
			if time.Since(info.ModTime()) > 30*24*time.Hour {
				if strings.HasSuffix(fileName, ".dmg") ||
				   strings.HasSuffix(fileName, ".zip") ||
				   strings.HasSuffix(fileName, ".tar.gz") ||
				   strings.Contains(fileName, "installer") {
					shouldRemove = true
				}
			}
		}

		if shouldRemove {
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

	return removedCount, totalSize
}

func init() {
	if rootCmd != nil {
		rootCmd.AddCommand(cleanTempCmd)
	}
}