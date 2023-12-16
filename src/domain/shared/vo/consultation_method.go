package shared

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type ConsultationMethod string

const (
	Chat  ConsultationMethod = "チャット"
	Video ConsultationMethod = "ビデオ"
)

func ValidateConsultationMethod(consultationMethod string) error {
	switch consultationMethod {
	case string(Chat), string(Video):
		return nil
	default:
		return apperr.BadRequestf("無効なconsultationMethodです: %d", consultationMethod)
	}
}
