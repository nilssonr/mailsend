package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/mail.v2"
)

var (
	host        string
	port        int
	from        string
	to          []string
	subject     string
	body        string
	attachments []string
)

func init() {
	sendCmd.Flags().StringVar(&host, "host", "", "SMTP server address")
	sendCmd.Flags().IntVar(&port, "port", 25, "SMTP server port")
	sendCmd.Flags().StringVar(&from, "from", "", "Sender address (user@domain.tld)")
	sendCmd.Flags().StringSliceVar(&to, "to", []string{}, "Recipient addresses. Use multiple --to for multiple recipients")
	sendCmd.Flags().StringVar(&subject, "subject", "", "Subject")
	sendCmd.Flags().StringVar(&body, "body", "", "Email HTML body (<h1>Hello</h1>)")
	sendCmd.Flags().StringSliceVar(&attachments, "attachment", []string{}, "File path to attachment (/home/user/file.png)")

	sendCmd.MarkFlagRequired("host")
	sendCmd.MarkFlagRequired("from")
	sendCmd.MarkFlagRequired("to")
	rootCmd.AddCommand(sendCmd)
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send email",
	Example: `mailsend send --host 127.0.0.1 \
		--port 25 \
		--from example@example.org \
		--to recipient@domaind.tld \
		--to other.recipient@domain.tld \
		--subject "Hello" \
		--body "<h1>Hello</h1>" \
		--attachment "/path/to/file.png" \
		--attachment "/path/to/other.file.png"`,
	Run: func(cmd *cobra.Command, args []string) {
		d := mail.Dialer{
			Host:           host,
			Port:           port,
			SSL:            false,
			StartTLSPolicy: mail.NoStartTLS,
		}

		m := mail.NewMessage()
		m.SetHeader("From", from)
		m.SetHeader("To", to...)
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", body)

		if len(attachments) > 0 {
			for _, attachment := range attachments {
				m.Attach(attachment)
			}
		}

		if err := d.DialAndSend(m); err != nil {
			fmt.Println(fmt.Errorf("failed to send message: %w", err))
			os.Exit(1)
		}

		fmt.Println("message sent successfully")
	},
}
