package requirement

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type Status string

const (
	Publish    Status = "公開"
	Suspension Status = "中止"
)

func validateStatus(status string) error {
	switch status {
	case string(Publish), string(Suspension):
		return nil
	default:
		return apperr.BadRequestf("無効なstatusです: %d", status)
	}
}
