package cmd

import (
	"fmt"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "trailer",
	Short: "Trailer is a tool that get trailers.",
	Long:  "Trailer is a tool that will quickly bring the link of any movie with a few commands.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Trailer command")
	},
}

var cmdMovie = &cobra.Command{
	Use:   "movie [name of the movie]",
	Short: "Searches the movie trailer",
	Long:  "Searches the trailer of one or more movies by name.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tmdbClient, err := tmdb.Init(os.Getenv("APIKey"))

		if err != nil {
			fmt.Println(err)
		}

		trailers, err := tmdbClient.GetMovieVideos(500, nil)

		if err != nil {
			fmt.Println(err)
		}

		for _, t := range trailers.Results {
			fmt.Println("https://www.youtube.com/watch?v=" + t.Key)
		}

	},
}

// Execute ...
func Execute() {
	rootCmd.AddCommand(cmdMovie)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
