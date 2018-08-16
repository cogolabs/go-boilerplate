package observe

import (
	"os"

	"github.com/evalphobia/logrus_sentry"
	raven "github.com/getsentry/raven-go"
	log "github.com/sirupsen/logrus"
)

// InitLogging initializes the logrus sentry logger with level setting and
// sentry's DSN for this project
func InitLogging(debug bool, dsn string) error {
	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	client, err := raven.NewClient(dsn, nil)
	if err != nil {
		log.Errorf("Failure to init raven client %s", err)
		return err
	}

	hook, err := logrus_sentry.NewWithClientSentryHook(client, []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
	})

	if err != nil {
		log.Errorf("Failed to init sentry hook, err: %s", err)
		return err
	} else {
		log.AddHook(hook)
	}
	return nil
}
