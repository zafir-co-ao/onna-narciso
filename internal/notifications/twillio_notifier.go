package notifications

import (
	"log/slog"
	"os"

	"github.com/kindalus/godx/pkg/assert"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func NewTwillioSMSNotifier() Notifier {
	return NotifierFunc(

		func(n, msg string) error {

			// Find your Account SID and Auth Token at twilio.com/console
			// and set the environment variables. See http://twil.io/secure
			// Make sure TWILIO_ACCOUNT_SID, TWILIO_AUTH_TOKEN, TWILIO_NUMBER exists in your environment
			client := twilio.NewRestClient()

			assert.NotNil(os.Getenv("TWILIO_ACCOUNT_SID"), "TWILIO_ACCOUNT_SID env not found")
			assert.NotNil(os.Getenv("TWILIO_AUTH_TOKEN"), "TWILIO_AUTH_TOKEN env not found")
			assert.NotNil(os.Getenv("TWILIO_NUMBER"), "TWILIO_NUMBER env not found")

			from := os.Getenv("TWILIO_NUMBER")

			params := &api.CreateMessageParams{}
			params.SetBody(msg)
			params.SetFrom(from)
			//params.SetFrom("+15017122661")
			params.SetTo(n)

			resp, err := client.Api.CreateMessage(params)
			if err != nil {
				slog.Error(err.Error())
				return err
			}

			if resp.Body != nil {
				slog.Info(*resp.Body)
			}

			return nil
		},
	)
}
