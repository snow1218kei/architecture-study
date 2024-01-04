package shared

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type ContractType string

const (
	Single     ContractType = "単発"
	Continuous ContractType = "継続"
)

func NewContractType(typ string) (ContractType, error) {
	switch typ {
	case Single.String():
		return Single, nil
	case Continuous.String():
		return Continuous, nil
	default:
		return "", apperr.BadRequestf("無効な契約タイプ: %s", typ)
	}
}

func (contractType ContractType) String() string {
	return string(contractType)
}
