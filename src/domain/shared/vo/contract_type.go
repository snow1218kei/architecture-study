package shared

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type ContractType string

const (
	Single     ContractType = "単発"
	Continuous ContractType = "継続"
)

func ValidateContractType(contractType string) error {
	switch contractType {
	case string(Single), string(Continuous):
		return nil
	default:
		return apperr.BadRequestf("無効なcontractType: %d", contractType)
	}
}
