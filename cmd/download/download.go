package download

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tifye/wim/core"
)

var (
	platform core.Platform
)

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download is a palette of commands for downloading winmowers.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Downloading latest %s winmower...\n", platform)

		types, err := core.FetchBundleTypes()
		cobra.CheckErr(err)

		types = core.FilterBundleTypes(types, platform)
		latestType := types[0]
		fmt.Printf("Latest %s winmower is %s\n", platform, latestType.Name)

		latestBuild, err := core.FetchLatestRelease(latestType.Name)
		cobra.CheckErr(err)
		fmt.Printf("Latest build is %s\n", latestBuild.BlobUrl)

		err = core.DownloadAndUnpack(latestBuild.BlobUrl, filepath.Join("./tmp", latestType.Name))
		cobra.CheckErr(err)
	},
}

func init() {
	DownloadCmd.Flags().VarP(&platform, "platform", "p", "Winmower platform to download.")
	if err := DownloadCmd.MarkFlagRequired("platform"); err != nil {
		cobra.CheckErr(err)
	}
}
