package httpclient

import (
	"time"

	"github.com/gojektech/heimdall"
	"github.com/gojektech/heimdall/httpclient"
)

func New(timeout time.Duration, retries int) *httpclient.Client {
	return httpclient.NewClient(
		httpclient.WithHTTPTimeout(timeout),
		httpclient.WithRetrier(
			heimdall.NewRetrier(
				heimdall.NewConstantBackoff(2*time.Millisecond, 5*time.Millisecond),
			),
		),
		httpclient.WithRetryCount(retries),
	)
}
