package requirement

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type ApplicationPeriod string

const (
	OneDay   ApplicationPeriod = "1日単位"
	TwoWeeks ApplicationPeriod = "最大14日"
)

func newApplicationPeriod(period string) (ApplicationPeriod, error) {
	switch period {
	case string(OneDay):
		return OneDay, nil
	case string(TwoWeeks):
		return TwoWeeks, nil
	default:
		return "", apperr.BadRequestf("無効な申請期間: %s", period)
	}
}
