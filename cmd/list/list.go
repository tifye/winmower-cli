package list

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tifye/wim/core"
)

var (
	listStrat ListFormat
	Platform  core.Platform
)

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List different winmower resources.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		types, err := core.FetchBundleTypes()
		cobra.CheckErr(err)

		types = core.FilterBundleTypes(types, Platform)
		switch listStrat {
		case All:
			listAllBundleTypes(types)
		case Partial:
			listPartialBundleTypes(types)
		default:
			listPartialBundleTypes(types)
		}
	},
}

func listAllBundleTypes(types []core.BundleType) {
	for i, t := range types {
		fmt.Printf("%d.\t%s\n", i+1, t.Name)
	}
}

func listPartialBundleTypes(types []core.BundleType) {
	const offset = 5

	if len(types) <= offset*2 {
		listAllBundleTypes(types)
		return
	}

	for i := 0; i < offset; i++ {
		parts := strings.Split(types[i].Name, ".")
		fmt.Printf("%d.\t%s\n", i+1, parts[len(parts)-1])
	}

	fmt.Println("...")

	for i := len(types) - offset; i < len(types); i++ {
		parts := strings.Split(types[i].Name, ".")
		fmt.Printf("%d.\t%s\n", i+1, parts[len(parts)-1])
	}
}

func init() {
	ListCmd.PersistentFlags().VarP(&Platform, "platform", "p", "The winmower platform to list bundle types for.")
	if err := ListCmd.MarkPersistentFlagRequired("platform"); err != nil {
		cobra.CheckErr(err)
	}

	ListCmd.Flags().VarP(&listStrat, "format", "f", "The format to list the bundle types in.")
}
