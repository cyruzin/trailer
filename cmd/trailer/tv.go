package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/manifoldco/promptui"
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

			var tvShowResults []string
			var tvShowID int64 = search.Results[0].ID
			var tvShowName string = search.Results[0].Name

			if len(search.Results) > 1 {
				for _, val := range search.Results {
					tvShowResults = append(
						tvShowResults,
						fmt.Sprintf("%s - %s", val.Name, parseDate(val.FirstAirDate)),
					)
				}

				prompt := promptui.Select{
					Label:    "Select a TV Show",
					Items:    tvShowResults,
					Size:     10,
					HideHelp: true,
				}

				_, promptResult, err := prompt.Run()

				tvShowName = promptResult

				if err != nil {
					log.Println(errorFetch)
					return
				}

				for _, val := range search.Results {
					if promptResult == fmt.Sprintf("%s - %s", val.Name, parseDate(val.FirstAirDate)) {
						tvShowID = val.ID
					}
				}
			}

			trailers, err := t.client.GetTVVideos(int(tvShowID), options)
			if err != nil {
				log.Println(errorFetch)
				return
			}

			if len(trailers.Results) <= 0 {
				if Lang != "" {
					log.Printf("No trailers are available for %s in the %s language.\n", tvShowName, Lang)
					return
				}
				log.Printf("No trailers available for: %s.\n", tvShowName)
				return
			}

			log.Println("Results for:", tvShowName)
			log.Println("")

			for _, trailer := range trailers.Results {
				log.Println(trailer.Name)
				log.Println(youtubeURL + trailer.Key)
				log.Println("")
			}
		},
	}
}
