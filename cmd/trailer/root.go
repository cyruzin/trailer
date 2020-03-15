package cmd

import (
	"log"

	tmdb "github.com/cyruzin/golang-tmdb"

	"github.com/spf13/cobra"
)

const (
	youtubeURL = "https://www.youtube.com/watch?v="
	errorFetch = "Oh no! Looks like Thanos snapped his fingers! The Avengers are working to fix this."
)

// Lang string for language flag.
var Lang string

// Trailer client structure.
type Trailer struct {
	client *tmdb.Client
}

// NewTMDBClient initiates the client.
func NewTMDBClient(client *tmdb.Client) *Trailer {
	return &Trailer{client}
}

// RootCmd creates the root command.
func (t *Trailer) RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "trailer",
		Short: "Trailer is a tool that get trailers.",
		Long:  "Trailer is a tool that will quickly bring the link of any movie with a few commands.",
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("Let's search some trailer?")
			log.Printf("")
			log.Printf(`Type "trailer --help" to see the commands available.`)
		},
	}

	movie := t.movieCmd()
	tv := t.tvCmd()

	parseCommonFlags(movie, tv)

	rootCmd.AddCommand(
		movie,
		tv,
		versionCmd(),
	)

	return rootCmd
}
