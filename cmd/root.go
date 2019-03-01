package cmd

import (
	"fmt"
	"os"

	"github.com/bakito/smc/pkg/mail"
	"github.com/spf13/cobra"
)

var (
	from          string
	to            []string
	cc            []string
	bcc           []string
	subject       string
	body          string
	contentType   string
	encoding      string
	host          string
	noTLS         bool
	skipTLSVerify bool
	port          uint
	username      string
	password      string
	version       string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: version,
	Use:     "smc",
	Short:   "Simple Mail Client",
	Long:    "A simple CLI to send smtp mails",
	RunE: func(cmd *cobra.Command, args []string) error {

		c := mail.NewClient(mail.Config{
			Host:          host,
			Port:          port,
			NoTLS:         noTLS,
			SkipTLSVerify: skipTLSVerify,
			User:          username,
			Password:      password,
		})

		return c.Send(mail.Message{
			From:        from,
			To:          to,
			CC:          cc,
			BCC:         bcc,
			Subject:     subject,
			Body:        body,
			ContentType: contentType,
			Encoding:    encoding,
		})
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringVarP(&from, "from", "f", "", "The sender address (required)")
	rootCmd.MarkPersistentFlagRequired("from")

	rootCmd.PersistentFlags().StringArrayVarP(&to, "to", "t", nil, "The receivers (required)")
	rootCmd.MarkPersistentFlagRequired("to")
	rootCmd.PersistentFlags().StringArrayVar(&cc, "cc", nil, "CC addresses")
	rootCmd.PersistentFlags().StringArrayVar(&bcc, "bcc", nil, "BCC addresses")

	rootCmd.PersistentFlags().StringVarP(&subject, "subject", "s", "", "the mail subject")
	rootCmd.MarkPersistentFlagRequired("subject")
	rootCmd.PersistentFlags().StringVarP(&body, "body", "b", "", "the mail body")
	rootCmd.MarkPersistentFlagRequired("body")
	rootCmd.PersistentFlags().StringVar(&contentType, "content-type", "text/plain", "The content-type of the body")
	rootCmd.PersistentFlags().StringVar(&encoding, "encoding", "UTF-8", "The encoding of the body")

	rootCmd.PersistentFlags().StringVar(&host, "host", "", "The smtp host")
	rootCmd.MarkPersistentFlagRequired("host")
	rootCmd.PersistentFlags().UintVarP(&port, "port", "p", 465, "The smtp port")
	rootCmd.PersistentFlags().BoolVar(&noTLS, "no-tls", false, "Disable TLS")
	rootCmd.PersistentFlags().BoolVar(&skipTLSVerify, "skip-tls-verify", true, "Skip TLS verify")

	rootCmd.PersistentFlags().StringVar(&username, "username", "", "smtp username (if different from from address)")
	rootCmd.PersistentFlags().StringVar(&password, "password", "", "smtp password")
}
