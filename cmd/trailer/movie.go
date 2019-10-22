package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	movieOverview = "N/A"
	usReleaseDate = "N/A"
)

func (t *Trailer) movieCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "movie [name of the movie]",
		Short: "Searches for a movie trailer",
		Long:  "Searches one or more trailers from a movie.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			argsJoin := strings.Join(args, " ")

			language, err := cmd.Flags().GetString("lang")
			if err != nil {
				fmt.Println(err)
			}

			options := make(map[string]string)
			if language != "" {
				options["language"] = language
			}

			search, err := t.client.GetSearchMovies(argsJoin, options)
			if err != nil {
				fmt.Println(errorFetch)
				return
			}

			if len(search.Results) <= 0 {
				fmt.Println("No results for:", argsJoin)
				return
			}

			if search.Results[0].Overview != "" {
				movieOverview = search.Results[0].Overview
			}

			releaseDates, err := t.client.GetMovieReleaseDates(int(search.Results[0].ID))
			if err != nil {
				fmt.Println(errorFetch)
				return
			}

			for _, res := range releaseDates.Results {
				if res.Iso3166_1 == "US" {
					usReleaseDate = parseDate(res.ReleaseDates[0].ReleaseDate)
				}
			}

			trailers, err := t.client.GetMovieVideos(int(search.Results[0].ID), options)
			if err != nil {
				fmt.Println(errorFetch)
				return
			}

			if len(trailers.Results) <= 0 {
				if language != "" {
					fmt.Printf("No trailers are available for %s in the %s language.\n", argsJoin, language)
					return
				}
				fmt.Printf("No trailers available for: %s.\n", argsJoin)
				return
			}

			fmt.Println("Results for: ", argsJoin)
			fmt.Println("")
			fmt.Println("Overview: ", movieOverview)
			fmt.Println("")
			fmt.Printf("US Release: %s\n", usReleaseDate)
			fmt.Println("")

			for _, trailer := range trailers.Results {
				fmt.Println(trailer.Name)
				fmt.Println(youtubeURL + trailer.Key)
				fmt.Println("")
			}
		},
	}
}
