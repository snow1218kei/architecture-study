package requirement

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type Status string

const (
	publish    Status = "公開"
	suspension Status = "中止"
)

func validateStatus(status Status) error {
	switch status {
	case publish, suspension:
		return nil
	default:
		return apperr.BadRequestf("categoryはゆう: %d", status)
	}
}
