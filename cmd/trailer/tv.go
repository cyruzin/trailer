package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	tvOverview     = "N/A"
	firstAiredDate = "N/A"
	lastAiredDate  = "N/A"
)

func (t *Trailer) tvCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tv [name of the tv show]",
		Short: "Searches for a tv show trailer",
		Long:  "Searches one or more trailers from a tv show.",
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

			search, err := t.client.GetSearchTVShow(argsJoin, options)
			if err != nil {
				fmt.Println(errorFetch)
				return
			}

			if len(search.Results) <= 0 {
				fmt.Println("No results for:", argsJoin)
				return
			}

			tvDetails, err := t.client.GetTVDetails(int(search.Results[0].ID), options)
			if err != nil {
				fmt.Println(errorFetch)
				return
			}

			if tvDetails.Overview != "" {
				tvOverview = tvDetails.Overview
			}

			if tvDetails.FirstAirDate != "" {
				firstAiredDate = parseDate(tvDetails.FirstAirDate)
			}

			if tvDetails.LastAirDate != "" {
				lastAiredDate = parseDate(tvDetails.LastAirDate)
			}

			trailers, err := t.client.GetTVVideos(int(search.Results[0].ID), options)
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
			fmt.Println("Overview:", tvOverview)
			fmt.Println("")
			fmt.Printf("First Aired: %s\n", firstAiredDate)
			fmt.Printf("Last Aired: %s\n", lastAiredDate)
			fmt.Println("")

			for _, trailer := range trailers.Results {
				fmt.Println(trailer.Name)
				fmt.Println(youtubeURL + trailer.Key)
				fmt.Println("")
			}
		},
	}
}
