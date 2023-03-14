package tag_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/tag"
)

func TestNewTagID(t *testing.T) {
	tests := []struct {
		name string
		tagID tag.TagID
	}{
		{
			name:   "正常系：NewUserID + String",
			tagID: tag.NewTagID(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := uuid.Parse(tt.tagID.String())
			assert.Nil(t, err, "Expected no error, but got an error")
		})
	}
}

func TestNewTagIDByVal(t *testing.T) {
	testCases := []struct {
		name      string
		val       string
		expected  tag.TagID
		expectErr bool
	}{
		{
			name:      "正常系：有効な値の場合",
			val:       "123",
			expected:  tag.TagID("123"),
			expectErr: false,
		},
		{
			name:      "異常系：空の値の場合",
			val:       "",
			expected:  tag.TagID(""),
			expectErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tagID, err := tag.NewTagIDByVal(tt.val)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expected, tagID)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, tagID)
			}
		})
	}
}

func TestTagID_Equal(t *testing.T) {
	tests := []struct {
		name    string
		tagID tag.TagID
	}{
		{
			name:   "正常系",
			tagID: tag.NewTagID(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tagID1 := tt.tagID
			tagID2 := tagID1
			assert.True(t, tagID1.Equal(tagID2), "Expected tagID1 and tagID2 to be equal, but they were not")
		})
	}
}
