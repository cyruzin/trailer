package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func (t *Trailer) tvCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tv [name of the tv show]",
		Short: "Searches for a tv show trailer",
		Long:  "Searches one or more trailers from a tv show.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			argsJoin := strings.Join(args, " ")

			search, err := t.client.GetSearchTVShow(argsJoin, nil)

			if err != nil {
				fmt.Println(errorFetch)
				return
			}

			if len(search.Results) <= 0 {
				fmt.Println("No results for:", argsJoin)
				return
			}

			trailers, err := t.client.GetTVVideos(int(search.Results[0].ID), nil)
			if err != nil {
				fmt.Println(errorFetch)
				return
			}

			fmt.Println("Results for:", argsJoin)
			fmt.Println("")

			for _, trailer := range trailers.Results {
				fmt.Println(trailer.Name)
				fmt.Println(youtubeURL + trailer.Key)
				fmt.Println("")
			}
		},
	}
}
