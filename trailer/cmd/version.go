package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Print the version number of Trailer",
		Aliases: []string{"v"},
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Trailer Cli v.1.0.1")
		},
	}
}
