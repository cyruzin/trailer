package cmd

import (
	"log"
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

			options := make(map[string]string)
			if Lang != "" {
				options["language"] = Lang
			}

			search, err := t.client.GetSearchTVShow(argsJoin, options)
			if err != nil {
				log.Println(errorFetch)
				return
			}

			if len(search.Results) <= 0 {
				log.Println("No results for:", argsJoin)
				return
			}

			trailers, err := t.client.GetTVVideos(int(search.Results[0].ID), options)
			if err != nil {
				log.Println(errorFetch)
				return
			}

			if len(trailers.Results) <= 0 {
				if Lang != "" {
					log.Printf("No trailers are available for %s in the %s language.\n", argsJoin, Lang)
					return
				}
				log.Printf("No trailers available for: %s.\n", argsJoin)
				return
			}

			log.Println("Results for:", argsJoin)
			log.Println("")

			for _, trailer := range trailers.Results {
				log.Println(trailer.Name)
				log.Println(youtubeURL + trailer.Key)
				log.Println("")
			}
		},
	}
}
