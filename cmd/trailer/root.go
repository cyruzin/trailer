package cmd

import (
	"log"

	tmdb "github.com/cyruzin/golang-tmdb"

	"github.com/spf13/cobra"
)

const youtubeURL = "https://www.youtube.com/watch?v="

const errorFetch = "Oh no! Looks like Thanos snapped his fingers! The Avengers are working to fix this."

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
			log.Println("Let's search some trailer?")
			log.Println("")
			log.Println(`Type "trailer --help" to see the commands available.`)
		},
	}

	rootCmd.AddCommand(t.movieCmd())
	rootCmd.AddCommand(t.tvCmd())
	rootCmd.AddCommand(versionCmd())

	return rootCmd
}
