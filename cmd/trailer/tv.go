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

			search, err := t.client.GetSearchTVShow(argsJoin, nil)
			if err != nil {
				log.Println(errorFetch)
				return
			}

			if len(search.Results) <= 0 {
				log.Println("No results for:", argsJoin)
				return
			}

			tvDetails, err := t.client.GetTVDetails(int(search.Results[0].ID), nil)
			if err != nil {
				log.Println(errorFetch)
				return
			}

			firstAiredDate := "NA"

			if tvDetails.FirstAirDate != "" {
				firstAiredDate = parseDate(tvDetails.FirstAirDate)
			}

			trailers, err := t.client.GetTVVideos(int(search.Results[0].ID), nil)
			if err != nil {
				log.Println(errorFetch)
				return
			}



			log.Printf("Results for: %s (First Aired: %s)", argsJoin, firstAiredDate)
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
