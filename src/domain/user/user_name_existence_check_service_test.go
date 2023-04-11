package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	mock_user "github.com/yuuki-tsujimura/architecture-study/src/mock"
)

func TestCheckUserNameExistence(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_user.NewMockUserRepository(ctrl)

	testCases := []struct {
		name          string
		userName      string
		mockFunc      func()
		expectedError error
	}{
		{
			name:     "正常系：存在しないユーザー名の場合",
			userName: "non_existing_user",
			mockFunc: func() {
				mockRepo.EXPECT().FindByName(context.Background(), "non_existing_user").Return(nil, nil)
			},
			expectedError: nil,
		},
		{
			name:     "異常系：存在するユーザー名の場合",
			userName: "existing_user",
			mockFunc: func() {
				mockRepo.EXPECT().FindByName(context.Background(), "existing_user").Return(&user.User{}, errors.New(""))
			},
			expectedError: errors.New("既に存在するユーザ名です"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			err := user.CheckUserNameExistence(context.Background(), tc.userName, mockRepo)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
