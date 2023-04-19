package userusecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase/userusecase/userinput"
)

var ErrInternalServer = errors.New("internal server error")

type CreateUserUseCase struct {
	userRepo user.UserRepository
}

func NewCreateUserUseCase(userRepo user.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepo,
	}
}

func (usercase *CreateUserUseCase) Exec(ctx context.Context, input *userinput.CreateUserInput) error {
	err := checkUserNameExistence(ctx, input.UserInput.Name, usercase.userRepo)
	if err != nil {
		return err
	}

	user, err := createUser(input)
	if err != nil {
		return err
	}

	err = saveUser(ctx, user, usercase.userRepo)
	if err != nil {
		return err
	}

	return nil
}

func checkUserNameExistence(ctx context.Context, name string, userRepo user.UserRepository) error {
	isExistByNameService := user.NewIsExistByNameService(userRepo)
	isExist, err := isExistByNameService.Exec(ctx, name)
	if err != nil {
		// サーバ側のエラーとして返す
		return ErrInternalServer
	}
	if isExist {
		// BadRequestを返す
		return fmt.Errorf("存在しているので他の名前でお願いします")
	}
	return nil
}

func createUser(input *userinput.CreateUserInput) (*user.User, error) {
	userParams := user.UserParams{
		Name:     input.UserInput.Name,
		Email:    input.UserInput.Email,
		Password: input.UserInput.Password,
		Profile:  input.UserInput.Profile,
	}

	var careersParams []user.CareerParams
	for _, careerInput := range input.CareersInput {
		careerParams := user.CareerParams{
			Detail:    careerInput.Detail,
			StartYear: careerInput.StartYear,
			EndYear:   careerInput.EndYear,
		}
		careersParams = append(careersParams, careerParams)
	}

	var skillsParams []user.SkillParams
	for _, skillInput := range input.SkillsInput {
		skillParams := user.SkillParams{
			TagID:      skillInput.TagID,
			Evaluation: skillInput.Evaluation,
			Years:      skillInput.Years,
		}
		skillsParams = append(skillsParams, skillParams)
	}

	user, err := user.CreateUserAggregate(userParams, careersParams, skillsParams)

	if err != nil {
		return nil, fmt.Errorf("ユーザの作成に失敗しました")
	}
	return user, nil
}

func saveUser(ctx context.Context, user *user.User, userRepo user.UserRepository) error {
	err := userRepo.Store(ctx, user)

	if err != nil {
		return fmt.Errorf("ユーザの登録に失敗しました")
	}
	return nil
}
