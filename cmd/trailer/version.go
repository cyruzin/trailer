package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// TrailerVersion current version
var TrailerVersion = "1.2.5"

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Print the version number of Trailer",
		Aliases: []string{"v"},
		Run: func(cmd *cobra.Command, args []string) {
			log.Println(TrailerVersion)
		},
	}
}
