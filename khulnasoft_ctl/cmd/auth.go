package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/khulnasoft/kengine/khulnasoft_ctl/http"
	"github.com/khulnasoft/kengine/khulnasoft_ctl/output"
	"github.com/khulnasoft/kengine/khulnasoft_utils/log"
	oahttp "github.com/khulnasoft-lab/golang_sdk/utils/http"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate",
	Long:  `This subcommand authenticate with remote server`,
	Run: func(cmd *cobra.Command, args []string) {
		api_token, _ := cmd.Flags().GetString("api-token")
		if api_token == "" {
			log.Fatal().Msg("Please provide an api-token")
		}

		https_client := oahttp.NewHttpsConsoleClient(console_ip, "443")
		err := https_client.APITokenAuthenticate(api_token)
		if err != nil {
			log.Fatal().Msgf("Failed to authenticate %v\n", err)
		}

		access, refresh := https_client.DumpTokens()
		tokens := http.AuthTokens{
			AccessToken:  access,
			RefreshToken: refresh,
		}

		b, err := json.Marshal(tokens)
		if err != nil {
			log.Fatal().Msgf("Failed to authenticate %v\n", err)
		}

		err = os.WriteFile(fmt.Sprintf("%s/%s", os.TempDir(), http.TokensFilename), b, 0600)
		if err != nil {
			log.Fatal().Msgf("Failed to authenticate %v\n", err)
		}

		output.Out(map[string]string{"login": "successful"})
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.PersistentFlags().String("api-token", "", "Khulnasoft console API token")
}
