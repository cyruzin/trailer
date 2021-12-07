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

			var tvResults []string

			for _, val := range search.Results {
				tvResults = append(
					tvResults,
					fmt.Sprintf("%s - %s", val.Name, parseDate(val.FirstAirDate)),
				)
			}

			prompt := promptui.Select{
				Label:    "Select a tv show",
				Items:    tvResults,
				Size:     10,
				HideHelp: true,
			}

			_, promptResult, err := prompt.Run()

			if err != nil {
				log.Println(errorFetch)
				return
			}

			var tvShowID int64

			for _, val := range search.Results {
				if promptResult == fmt.Sprintf("%s - %s", val.Name, parseDate(val.FirstAirDate)) {
					tvShowID = val.ID
				}
			}

			trailers, err := t.client.GetTVVideos(int(tvShowID), options)
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

			log.Println("Results for:", promptResult)
			log.Println("")

			for _, trailer := range trailers.Results {
				log.Println(trailer.Name)
				log.Println(youtubeURL + trailer.Key)
				log.Println("")
			}
		},
	}
}
