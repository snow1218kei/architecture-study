package shared

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type Status string

const (
	Publish    Status = "公開"
	Suspension Status = "中止"
)

func NewStatus(status string) (Status, error) {
	switch status {
	case Publish.String():
		return Publish, nil
	case Suspension.String():
		return Suspension, nil
	default:
		return "", apperr.BadRequestf("無効なstatusです: %s", status)
	}
}

func (status Status) String() string {
	return string(status)
}
