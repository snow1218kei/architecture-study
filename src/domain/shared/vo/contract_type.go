package shared

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type ContractType string

const (
	Single     ContractType = "単発"
	Continuous ContractType = "継続"
)

func NewContractType(typ string) (ContractType, error) {
	switch typ {
	case string(Single):
		return Single, nil
	case string(Continuous):
		return Continuous, nil
	default:
		return "", apperr.BadRequestf("無効な契約タイプ: %s", typ)
	}
}
