package shared

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type ConsultationMethod string

const (
	Chat  ConsultationMethod = "チャット"
	Video ConsultationMethod = "ビデオ"
)

func NewConsultationMethod(method string) (ConsultationMethod, error) {
	switch method {
	case Chat.String():
		return Chat, nil
	case Video.String():
		return Video, nil
	default:
		return "", apperr.BadRequestf("無効な相談方法: %s", method)
	}
}

func (consultationMethod ConsultationMethod) String() string {
	return string(consultationMethod)
}
