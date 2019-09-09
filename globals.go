package vouched

import (
	"fmt"
	"os"
	"sync"

	"github.com/kthomas/go-logger"
)

const vouchedDefaultEnvironment = "sandbox"

var (
	log           *logger.Logger
	bootstrapOnce sync.Once

	vouchedAPIBaseURL string
	vouchedAPIUser    string
	vouchedAPIToken   string
)

func init() {
	bootstrapOnce.Do(func() {
		lvl := os.Getenv("VOUCHED_LOG_LEVEL")
		if lvl == "" {
			lvl = "INFO"
		}
		log = logger.NewLogger("vouched", lvl, true)

		if os.Getenv("VOUCHED_API_ENVIRONMENT") != "" {
			vouchedAPIBaseURL = fmt.Sprintf("https://%s.vouched.id", os.Getenv("VOUCHED_API_ENVIRONMENT"))
		} else {
			vouchedAPIBaseURL = fmt.Sprintf("https://%s.vouched.id", vouchedDefaultEnvironment)
		}

		if os.Getenv("VOUCHED_API_TOKEN") != "" {
			vouchedAPIToken = os.Getenv("VOUCHED_API_TOKEN")
		}
	})
}
