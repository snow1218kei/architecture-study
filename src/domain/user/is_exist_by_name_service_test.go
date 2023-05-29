package user_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	mock_user "github.com/yuuki-tsujimura/architecture-study/src/mock"
)

func TestExec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_user.NewMockUserRepository(ctrl)
	isExistByNameService := user.NewIsExistByNameService(mockRepo)

	testCases := []struct {
		name          string
		userName      string
		mockFunc      func()
		expected      bool
		expectedError error
	}{
		{
			name:     "正常系：存在するユーザー名の場合",
			userName: "existing_user",
			mockFunc: func() {
				mockRepo.EXPECT().FindByName(context.Background(), "existing_user").Return(&user.User{}, nil)
			},
			expected:      true,
			expectedError: nil,
		},
		{
			name:     "異常系：存在しないユーザー名（NotFound）の場合",
			userName: "non_existing_user",
			mockFunc: func() {
				mockRepo.EXPECT().FindByName(context.Background(), "non_existing_user").Return(nil, apperr.NotFoundWrapf(sql.ErrNoRows, "RdbUserRepositoryImpl.FindByName failed to find user"))
			},
			expected:      false,
			expectedError: nil,
		},
		{
			name:     "異常系：サーバーエラーの場合",
			userName: "error_user",
			mockFunc: func() {
				mockRepo.EXPECT().FindByName(context.Background(), "error_user").Return(nil, errors.New("server error"))
			},
			expected:      false,
			expectedError: apperr.Internal("server error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			bool, err := isExistByNameService.Exec(context.Background(), tc.userName)
			assert.Equal(t, tc.expected, bool)
			if tc.expectedError != nil {
				assert.EqualError(t, tc.expectedError, err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
