package requirement

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type ApplicationPeriod string

const (
	OneDay   ApplicationPeriod = "1日単位"
	TwoWeeks ApplicationPeriod = "最大14日"
)

func newApplicationPeriod(period string) (ApplicationPeriod, error) {
	switch period {
	case OneDay.String():
		return OneDay, nil
	case TwoWeeks.String():
		return TwoWeeks, nil
	default:
		return "", apperr.BadRequestf("無効な申請期間: %s", period)
	}
}

func (applicationPeriod ApplicationPeriod) String() string {
	return string(applicationPeriod)
}
