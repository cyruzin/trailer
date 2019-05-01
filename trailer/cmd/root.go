package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const youtubeURL = "https://www.youtube.com/watch?v="

// NewRootCmd creates the root command.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "trailer",
		Short: "Trailer is a tool that get trailers.",
		Long:  "Trailer is a tool that will quickly bring the link of any movie with a few commands.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Trailer command")
		},
	}

	rootCmd.AddCommand(movieCmd())
	rootCmd.AddCommand(tvCmd())
	rootCmd.AddCommand(versionCmd())

	return rootCmd
}
