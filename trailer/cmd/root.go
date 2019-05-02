package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const youtubeURL = "https://www.youtube.com/watch?v="

// RootCmd creates the root command.
func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "trailer",
		Short: "Trailer is a tool that get trailers.",
		Long:  "Trailer is a tool that will quickly bring the link of any movie with a few commands.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Let's search some trailer?")
			fmt.Println("")
			fmt.Println(`Type "trailer --help" to see the commands available.`)
		},
	}

	rootCmd.AddCommand(movieCmd())
	rootCmd.AddCommand(tvCmd())
	rootCmd.AddCommand(versionCmd())

	return rootCmd
}
