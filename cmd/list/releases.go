package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tifye/wim/core"
)

// releasesCmd represents the releases command
var releasesCmd = &cobra.Command{
	Use:   "releases",
	Short: "List releases of a winmower.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		types, err := core.FetchBundleTypes()
		cobra.CheckErr(err)

		types = core.FilterBundleTypes(types, Platform)
		latestType := types[0]

		fmt.Printf("Latest %s winmower is %s\n", Platform, latestType.Name)

		latestBuild, err := core.FetchLatestRelease(latestType.Name)
		cobra.CheckErr(err)

		fmt.Printf("Latest build is %s\n", latestBuild.BlobUrl)
	},
}

func init() {
	if err := ListCmd.MarkPersistentFlagRequired("platform"); err != nil {
		cobra.CheckErr(err)
	}

	ListCmd.AddCommand(releasesCmd)
}
