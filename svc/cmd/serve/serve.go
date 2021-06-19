package serve

import (
	"fmt"
	"log"
	"os"
	"web-Scraping-test/svc/configs"
	"web-Scraping-test/svc/rest"

	"github.com/spf13/cobra"
)

var (
	Cmd          *cobra.Command
	argAddress   string
	argCORSHosts string
)

func init() {
	Cmd = &cobra.Command{
		Use:   "serve",
		Short: "Connect to the storage and begin serving requests.",
		Long:  ``,
		Run: func(Cmd *cobra.Command, args []string) {
			if err := serve(Cmd, args); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		},
	}

	Cmd.Flags().StringVarP(&argAddress, "address", "a", ":8080", "address to listen on")
	Cmd.Flags().StringVar(&argCORSHosts, "cors-hosts", "*", "cors hosts, separated by comma")
}

func serve(cmd *cobra.Command, args []string) error {
	svr, err := rest.NewServer(&configs.Config{
		HostPort:  argAddress,
		CORSHosts: argCORSHosts,
	})
	if err != nil {
		return err
	}

	log.Fatalln(svr.Run())

	return nil
}
