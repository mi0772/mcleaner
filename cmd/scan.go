package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan system for junk files",
	Long:  `Performs a system scan to identify junk files without removing them`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Scanning for junk files...")
		
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error finding home directory: %v\n", err)
			return
		}

		var cacheSize int64
		var dsStoreCount int
		var tempSize int64

		cacheSize = scanCache(homeDir)
		dsStoreCount = scanDSStore(homeDir)
		tempSize = scanTemp(homeDir)

		fmt.Println("\nScan results:")
		fmt.Printf("Cache found: %s\n", formatSize(cacheSize))
		fmt.Printf(".DS_Store files: %d\n", dsStoreCount)
		fmt.Printf("Temporary files: %s\n", formatSize(tempSize))
		fmt.Printf("Total recoverable space: %s\n", formatSize(cacheSize+tempSize))
	},
}

func scanCache(homeDir string) int64 {
	var totalSize int64
	cacheDirs := []string{
		filepath.Join(homeDir, "Library/Caches"),
		"/Library/Caches",
		"/tmp",
	}

	for _, dir := range cacheDirs {
		size := getDirSize(dir)
		totalSize += size
	}

	return totalSize
}

func scanDSStore(homeDir string) int {
	count := 0
	filepath.Walk(homeDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if strings.HasSuffix(path, ".DS_Store") {
			count++
		}
		return nil
	})
	return count
}

func scanTemp(homeDir string) int64 {
	var totalSize int64
	tempDirs := []string{
		"/tmp",
		"/var/tmp",
		filepath.Join(homeDir, "Downloads"),
	}

	for _, dir := range tempDirs {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if strings.Contains(strings.ToLower(path), "temp") || strings.HasSuffix(path, ".tmp") {
				totalSize += info.Size()
			}
			return nil
		})
	}

	return totalSize
}

func getDirSize(dirPath string) int64 {
	var totalSize int64
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})
	return totalSize
}

func formatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

var rootCmd *cobra.Command

func SetRootCmd(cmd *cobra.Command) {
	rootCmd = cmd
}

func init() {
	if rootCmd != nil {
		rootCmd.AddCommand(scanCmd)
	}
}