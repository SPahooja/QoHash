package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Files",
	Short: "Enter folder location to list files",
}
