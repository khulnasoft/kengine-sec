package cmd

import (
	"log"
	"os"

	"github.com/khulnasoft/kengine/khulnasoft_ctl/http"
	"github.com/spf13/cobra"
)

var (
	console_ip string
)

var (
	rootCmd = &cobra.Command{
		Use:   "khulnasoftctl",
		Short: "A khulnasoft controller CLI",
		Long:  `A simple CLI alternative to khulnasoft UI`,
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	console_ip = os.Getenv("KHULNASOFT_URL")
	if console_ip == "" {
		log.Fatal("KHULNASOFT_URL not specified. Please provie Console IP.")
	}
	http.InjectConsoleIp(console_ip)
}
