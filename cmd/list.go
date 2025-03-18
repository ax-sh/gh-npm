package cmd

import (
	"github.com/ax-sh/gh-npm/api"
	"github.com/k0kubun/pp/v3"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "list all npm packages on github",
	Run: func(cmd *cobra.Command, args []string) {
		packages, err := api.FetchPackages()
		if err != nil {
			return
		}

		cmd.Println(pp.Sprintf("Found %s NPM packages\n", len(packages)))
		for _, pkg := range packages {
			msg := pp.Sprintf("Package: %s (Type: %s)", pkg.Name, pkg.PackageType)
			cmd.Println(msg, pkg.Owner.Login)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
