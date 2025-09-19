package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var maintenanceCmd = &cobra.Command{
	Use:   "maintenance",
	Short: "Run system maintenance tasks",
	Long:  `Run daily, weekly or monthly macOS maintenance tasks`,
}

var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Run daily maintenance",
	Long:  `Run daily system maintenance scripts`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running daily maintenance...")
		runMaintenanceScript("daily")
	},
}

var weeklyCmd = &cobra.Command{
	Use:   "weekly", 
	Short: "Run weekly maintenance",
	Long:  `Run weekly system maintenance scripts`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running weekly maintenance...")
		runMaintenanceScript("weekly")
	},
}

var monthlyCmd = &cobra.Command{
	Use:   "monthly",
	Short: "Run monthly maintenance", 
	Long:  `Run monthly system maintenance scripts`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running monthly maintenance...")
		runMaintenanceScript("monthly")
	},
}

func runMaintenanceScript(period string) {
	scriptPath := fmt.Sprintf("/usr/libexec/periodic/%s", period)
	
	fmt.Printf("Running scripts in: %s\n", scriptPath)
	
	cmd := exec.Command("sudo", "periodic", period)
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		fmt.Printf("Error during execution: %v\n", err)
		fmt.Printf("Output: %s\n", string(output))
		return
	}
	
	fmt.Printf("Maintenance %s completed successfully!\n", period)
	if len(output) > 0 {
		fmt.Printf("Output:\n%s\n", string(output))
	}
	
	fmt.Println("Rebuilding locate cache...")
	locateCmd := exec.Command("sudo", "launchctl", "load", "-w", "/System/Library/LaunchDaemons/com.apple.locate.plist")
	if err := locateCmd.Run(); err != nil {
		fmt.Printf("Warning: cannot rebuild locate cache: %v\n", err)
	} else {
		fmt.Println("Locate cache rebuilt")
	}
}

func init() {
	maintenanceCmd.AddCommand(dailyCmd)
	maintenanceCmd.AddCommand(weeklyCmd) 
	maintenanceCmd.AddCommand(monthlyCmd)
	
	if rootCmd != nil {
		rootCmd.AddCommand(maintenanceCmd)
	}
}