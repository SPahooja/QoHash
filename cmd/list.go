package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/SPahooja/QoHash/files"
	"github.com/spf13/cobra"
)

// listCmd represents the do command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Finds the files in the given location",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Enter the file location to search in")
			os.Exit(1)
		}
		addr := args[0]
		files, size, err := files.FindFiles(addr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Total Files:", size, "in", addr)
		fmt.Println(files)

	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
