package apperr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

func TestIs(t *testing.T) {

	err := apperr.BadRequest("error")

	assert.True(t, apperr.Is[*apperr.BadRequestErr](err))

}
