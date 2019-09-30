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

			trailers, err := t.client.GetMovieVideos(int(search.Results[0].ID), nil)
			if err != nil {
				log.Println(errorFetch)
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
