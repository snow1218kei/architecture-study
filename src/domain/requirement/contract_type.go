package requirement

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type ContractType string

const (
	single     ContractType = "単発"
	continuous ContractType = "継続"
)

func validateContractType(contractType ContractType) error {
	switch contractType {
	case single, continuous:
		return nil
	default:
		return apperr.BadRequestf("categoryはゆう: %d", contractType)
	}
}
