package shared

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type ConsultationMethod string

const (
	Chat  ConsultationMethod = "チャット"
	Video ConsultationMethod = "ビデオ"
)

func NewConsultationMethod(method string) (ConsultationMethod, error) {
	switch method {
	case string(Chat):
		return Chat, nil
	case string(Video):
		return Video, nil
	default:
		return "", apperr.BadRequestf("無効な相談方法: %s", method)
	}
}
