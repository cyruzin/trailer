package cmd

import "github.com/spf13/cobra"

func parseCommonFlags(c ...*cobra.Command) {
	for _, f := range c {
		f.Flags().StringVarP(&Lang, "lang", "l", "en-US", "language ouput")
	}
}
