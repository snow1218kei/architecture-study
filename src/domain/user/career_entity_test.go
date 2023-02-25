package user

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/shared"
)

func TestNewCareer(t *testing.T) {
	careerID := CareerID("1")
	createdAt := shared.CreatedAt(time.Now())

	t.Run("有効なparamsの場合, career構造体のオブジェクトを返す", func(t *testing.T) {
		params := CareerParams{
			Detail:    "Software Engineer",
			StartYear: 2015,
			EndYear:   2022,
		}

		career, err := NewCareer(params, careerID, createdAt)

		assert.NoError(t, err)
		assert.Equal(t, careerID, career.careerId)
		assert.Equal(t, params.Detail, career.detail)
		assert.Equal(t, params.StartYear, career.startYear)
		assert.Equal(t, params.EndYear, career.endYear)
		assert.Equal(t, createdAt, career.createdAt)
	})

	t.Run("Detailが長過ぎる場合, エラーを返す", func(t *testing.T) {
		params := CareerParams{
			Detail:    strings.Repeat("s", 256),
			StartYear: 2015,
			EndYear:   2022,
		}

		career, err := NewCareer(params, careerID, createdAt)

		assert.Error(t, err)
		assert.Nil(t, career)
	})

	t.Run("StartYearが1970年未満の場合, エラーを返す", func(t *testing.T) {
		params := CareerParams{
			Detail:    "Software Engineer",
			StartYear: 1969,
			EndYear:   2022,
		}

		career, err := NewCareer(params, careerID, createdAt)

		assert.Error(t, err)
		assert.Nil(t, career)
	})

	t.Run("EndYearが1970年未満の場合, エラーを返す", func(t *testing.T) {
		params := CareerParams{
			Detail:    "Software Engineer",
			StartYear: 2015,
			EndYear:   1969,
		}

		career, err := NewCareer(params, careerID, createdAt)

		assert.Error(t, err)
		assert.Nil(t, career)
	})

	t.Run("EndYearがStartYear以下の場合, エラーを返す", func(t *testing.T) {
		params := CareerParams{
			Detail:    "Software Engineer",
			StartYear: 2022,
			EndYear:   2015,
		}

		career, err := NewCareer(params, careerID, createdAt)

		assert.Error(t, err)
		assert.Nil(t, career)
	})
}
