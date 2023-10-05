package download

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	urlPath string
	client  = http.Client{
		Timeout: time.Second * 2,
	}
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Pings the target url.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := ping(urlPath); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func ping(domain string) (string, error) {
	url := "http://" + domain
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	return resp.Status, nil
}

func init() {
	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "api.joshuadematas.me", "The url to ping.")

	DownloadCmd.AddCommand(pingCmd)
}
