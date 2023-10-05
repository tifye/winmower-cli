package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tifye/wim/core"
)

// platformsCmd represents the platforms command
var platformsCmd = &cobra.Command{
	Use:   "platforms",
	Short: "List all mower platforms.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		platforms := core.GetPlatforms()
		for _, p := range platforms {
			fmt.Println(p)
		}
	},
}

func init() {
	ListCmd.AddCommand(platformsCmd)
}
