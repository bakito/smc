package cmd

import (
	"fmt"
	"os"

	"github.com/bakito/smc/pkg/mail"
	"github.com/bakito/smc/version"
	"github.com/spf13/cobra"
)

var (
	from        string
	to          []string
	cc          []string
	bcc         []string
	subject     string
	body        string
	contentType string
	encoding    string
	host        string
	port        uint
	username    string
	password    string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: version.Version,
	Use:     "smc",
	Short:   "Simple Mail Client",
	Long:    "A simple CLI to send smtp mails",
	RunE: func(cmd *cobra.Command, args []string) error {

		c := mail.NewClient(mail.Config{
			Host:     host,
			Port:     port,
			User:     username,
			Password: password,
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

	rootCmd.PersistentFlags().StringVarP(&from, "from", "f", "", "The sender addres (required)")
	rootCmd.MarkPersistentFlagRequired("from")

	rootCmd.PersistentFlags().StringSliceVarP(&to, "to", "t", nil, "The receivers (required)")
	rootCmd.MarkPersistentFlagRequired("to")
	rootCmd.PersistentFlags().StringSliceVar(&cc, "cc", nil, "cc")
	rootCmd.PersistentFlags().StringSliceVar(&bcc, "bcc", nil, "bcc")

	rootCmd.PersistentFlags().StringVarP(&subject, "subject", "s", "", "subject")
	rootCmd.MarkPersistentFlagRequired("subject")
	rootCmd.PersistentFlags().StringVarP(&body, "body", "b", "", "body")
	rootCmd.MarkPersistentFlagRequired("body")
	rootCmd.PersistentFlags().StringVar(&contentType, "content-type", "text/plain", "Content-Type")
	rootCmd.PersistentFlags().StringVar(&encoding, "encoding", "UTF-8", "Encoding")

	rootCmd.PersistentFlags().StringVar(&host, "host", "", "host")
	rootCmd.MarkPersistentFlagRequired("host")
	rootCmd.PersistentFlags().UintVarP(&port, "port", "p", 465, "port")

	rootCmd.PersistentFlags().StringVar(&username, "username", "", "username")
	rootCmd.PersistentFlags().StringVar(&password, "password", "", "password")
}
