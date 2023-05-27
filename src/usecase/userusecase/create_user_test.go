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
	ctx := context.TODO()

	testCases := []struct {
		name          string
		input         *userinput.CreateUserInput
		mockFunc      func() *mock_user.MockUserRepository
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
			mockFunc: func() *mock_user.MockUserRepository {
				mockRepo := mock_user.NewMockUserRepository(ctrl)
				mockRepo.EXPECT().Store(ctx, gomock.Any()).Return(nil)
				mockRepo.EXPECT().FindByName(ctx, "non_existing_user").Return(nil, nil)
				return mockRepo
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
			mockFunc: func() *mock_user.MockUserRepository {
				mockRepo := mock_user.NewMockUserRepository(ctrl)
				mockRepo.EXPECT().FindByName(ctx, "existing_user").Return(&user.User{}, nil)
				return mockRepo
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
			mockFunc: func() *mock_user.MockUserRepository {
				mockRepo := mock_user.NewMockUserRepository(ctrl)
				mockRepo.EXPECT().FindByName(ctx, "non_existing_user").Return(nil, errors.New("server error"))
				return mockRepo
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
			mockFunc: func() *mock_user.MockUserRepository {
				mockRepo := mock_user.NewMockUserRepository(ctrl)
				mockRepo.EXPECT().Store(ctx, gomock.Any()).Return(errors.New("error"))
				mockRepo.EXPECT().FindByName(ctx, "non_existing_user").Return(nil, nil)
				return mockRepo
			},
			expectedError: errors.New("ユーザの登録に失敗しました"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userUsecase := userusecase.NewCreateUserUseCase(tc.mockFunc())
			_, err := userUsecase.Exec(ctx, tc.input)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
