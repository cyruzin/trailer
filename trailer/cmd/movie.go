package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func (t *Trailer) movieCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "movie [name of the movie]",
		Short: "Searches for a movie trailer",
		Long:  "Searches one or more trailers from a movie.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			argsJoin := strings.Join(args, " ")

			search, err := t.client.GetSearchMovies(argsJoin, nil)
			if err != nil {
				log.Println(errorFetch)
				return
			}

			if len(search.Results) <= 0 {
				log.Println("No results for:", argsJoin)
				return
			}

			releaseDates, err := t.client.GetMovieReleaseDates(int(search.Results[0].ID))
			if err != nil {
				log.Println(errorFetch)
				return
			}

			usReleaseDate := "NA"

			for _, res := range releaseDates.Results {
				if res.Iso3166_1 == "US" {
					usReleaseDate = parseDate(res.ReleaseDates[0].ReleaseDate)
				}
			}


			trailers, err := t.client.GetMovieVideos(int(search.Results[0].ID), nil)
			if err != nil {
				log.Println(errorFetch)
				return
			}

			log.Printf("Results for: %s (US Release: %s)", argsJoin, usReleaseDate)
			log.Println("")

			for _, trailer := range trailers.Results {
				log.Println(trailer.Name)
				log.Println(youtubeURL + trailer.Key)
				log.Println("")
			}
		},
	}
}
