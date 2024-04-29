package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var apiTarget string
var basePath string
var format string
var insecure bool
var localFile string
var proxy string
var quiet bool
var swaggerURL string
var timeout int64
var noSsl bool

var rootCmd = &cobra.Command{
	Use:   "sj",
	Short: "A tool for auditing documented (swagger/openapi) API endpoints.",
	Long: `The process of reviewing and testing exposed API definition files is often tedious and requires a large investment of time for a thorough review.

sj (swaggerjacker) is a CLI tool that can be used to perform an initial check of API endpoints identified through exposed Swagger/OpenAPI definition files. 
Once you determine what endpoints require authentication and which do not, you can use the "prepare" command to generate command templates for further (manual) testing.

Example usage:

Perform a quick check of endpoints which require authentication:
$ sj automate -u https://petstore.swagger.io/v2/swagger.json

Generate a list of curl commands to use for manual testing:
$ sj prepare -u https://petstore.swagger.io/v2/swagger.json

Generate a list of raw API routes for use with custom scripts:
$ sj endpoints -u https://petstore.swagger.io/v2/swagger.json`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Error("Command not specified. See the --help flag for usage.")
		}
	},
	Version: "1.0.5",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(automateCmd)
	rootCmd.AddCommand(endpointsCmd)
	rootCmd.AddCommand(prepareCmd)
	rootCmd.PersistentFlags().StringVarP(&UserAgent, "agent", "a", "Swagger Jacker (github.com/BishopFox/sj)", "Set the User-Agent string.")
	rootCmd.PersistentFlags().StringVarP(&basePath, "base-path", "b", "", "Set the API base path if not defined in the spec (i.e. /V2/)")
	rootCmd.PersistentFlags().StringVarP(&format, "format", "f", "json", "Declare the format of the documentation (json/yaml/yml/js).")
	rootCmd.PersistentFlags().StringArrayVarP(&Headers, "headers", "H", nil, "Add custom headers, separated by a colon (\"Name: Value\"). Multiple flags are accepted.")
	rootCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "i", false, "Ignores server certificate validation.")
	rootCmd.PersistentFlags().StringVarP(&localFile, "local-file", "l", "", "Loads the documentation from a local file.")
	rootCmd.PersistentFlags().StringVarP(&proxy, "proxy", "p", "NOPROXY", "Proxy host and port. Example: http://127.0.0.1:8080")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Do not prompt for user input - uses default values for all requests.")
	rootCmd.PersistentFlags().StringVarP(&apiTarget, "target", "T", "", "Manually set a target for the requests to be made if separate from the host the documentation resides on.")
	rootCmd.PersistentFlags().Int64VarP(&timeout, "timeout", "t", 30, "Set the request timeout period.")
	rootCmd.PersistentFlags().StringVarP(&swaggerURL, "url", "u", "", "Loads the documentation file from a URL")
	rootCmd.PersistentFlags().BoolVarP(&noSsl, "noSsl", "n", false, "Use http instead of https.")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
