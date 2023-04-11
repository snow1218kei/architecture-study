package userusecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	mock_user "github.com/yuuki-tsujimura/architecture-study/src/mock"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase/userusecase"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase/userusecase/userinput"
)

func TestSaveUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_user.NewMockUserRepository(ctrl)

	type args struct {
		user     *user.User
		userRepo user.UserRepository
	}
	type testCase struct {
		name          string
		args          args
		mockFunc      func()
		expectedError error
	}

	testCases := []testCase{
		{
			name: "正常系: ユーザの登録に成功する場合",
			args: args{
				user: &user.User{},
				userRepo: mockRepo,
			},
			mockFunc: func() {
				mockRepo.EXPECT().Store(context.Background(), &user.User{}).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "異常系: ユーザの登録に失敗する場合",
			args: args{
				user: &user.User{},
				userRepo: mockRepo,
			},
			mockFunc: func() {
				mockRepo.EXPECT().Store(context.Background(), &user.User{}).Return(errors.New("error"))
			},
			expectedError:  errors.New("ユーザの登録に失敗しました"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			err := userusecase.SaveUser(context.Background(), tc.args.user, tc.args.userRepo)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestCreateUser(t *testing.T) {
	cases := []struct {
		name          string
		input         *userinput.CreateUserInput
		expectedError error
	}{
		{
			name: "正常系：ユーザの作成に成功する場合",
			input: &userinput.CreateUserInput{
				UserInput: userinput.UserInput{
					Name:     "testuser",
					Email:    "testuser@example.com",
					Password: "password",
					Profile:  "test profile",
				},
				CareersInput: []*userinput.CareerInput{
					{
						Detail:    "test career",
						StartYear: 2022,
						EndYear:   2023,
					},
				},
				SkillsInput: []*userinput.SkillInput{
					{
						TagID:      "1",
						Evaluation: 3,
						Years:      1,
					},
				},
			},
			expectedError: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := userusecase.CreateUser(tc.input)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
