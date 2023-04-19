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

func TestExec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_user.NewMockUserRepository(ctrl)

	testCases := []struct {
		name          string
		input         *userinput.CreateUserInput
		mockFunc      func()
		expectedError error
	}{
		{
			name: "正常系: ユーザの登録に成功する場合",
			input: &userinput.CreateUserInput{
				UserInput: userinput.UserInput{
					Name:     "non_existing_user",
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
			mockFunc: func() {
				mockRepo.EXPECT().Store(context.Background(), &user.User{}).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "異常系: ユーザ名が既に存在する場合",
			input: &userinput.CreateUserInput{
				UserInput: userinput.UserInput{
					Name:     "existing_user",
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
			mockFunc: func() {
				mockRepo.EXPECT().FindByName(context.Background(), "existing_user").Return(&user.User{}, nil)
			},
			expectedError: errors.New("存在しているので他の名前でお願いします"),
		},
		{
			name: "異常系: 存在チェック時にサーバーエラーが起きた場合",
			input: &userinput.CreateUserInput{
				UserInput: userinput.UserInput{
					Name:     "non_existing_user",
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
			mockFunc: func() {
				mockRepo.EXPECT().FindByName(context.Background(), "non_existing_user").Return(nil, errors.New("server error"))
			},
			expectedError: userusecase.ErrInternalServer,
		},
		{
			name: "異常系: ユーザの登録に失敗する場合",
			input: &userinput.CreateUserInput{
				UserInput: userinput.UserInput{
					Name:     "non_existing_user",
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
			mockFunc: func() {
				mockRepo.EXPECT().Store(context.Background(), &user.User{}).Return(errors.New("error"))
			},
			expectedError:  errors.New("ユーザの登録に失敗しました"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()

			userUsecase := userusecase.NewCreateUserUseCase(mockRepo)
			err := userUsecase.Exec(context.Background(), tc.input)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
