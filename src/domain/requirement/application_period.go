package requirement

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type ApplicationPeriod string

const (
	OneDay   ApplicationPeriod = "1日単位"
	TwoWeeks ApplicationPeriod = "最大14日"
)

func validateApplicationPeriod(applicationPeriod string) error {
	switch applicationPeriod {
	case string(OneDay), string(TwoWeeks):
		return nil
	default:
		return apperr.BadRequestf("無効なapplicationPeriodです: %d", applicationPeriod)
	}
}
