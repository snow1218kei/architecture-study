package userusecase

import (
	"context"
	"fmt"

	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase/userusecase/userinput"
)

type CreateUserUseCase struct {
  userRepo user.UserRepository
}

func NewCreateUserUseCase(userRepo user.UserRepository) *CreateUserUseCase {
  return &CreateUserUseCase {
    userRepo,
  }
}

func (usercase *CreateUserUseCase) Exec(ctx context.Context, input *userinput.CreateUserInput) error {
	err := user.CheckUserNameExistence(input.UserInput.Name, usercase.userRepo)
  user, err := createUser(input)
	err = saveUser(user, usercase.userRepo)

	if err != nil {
		return err
	}
	return nil
}

func createUser(input *userinput.CreateUserInput) (*user.User, error) {
	userParams := user.UserParams{
		Name: input.UserInput.Name,
		Email: input.UserInput.Email,
		Password: input.UserInput.Password,
		Profile: input.UserInput.Profile,
	}

	var careersParams []user.CareerParams
	for _, careerInput := range input.CareersInput {
		careerParams :=	user.CareerParams{
			Detail: careerInput.Detail, 
			StartYear: careerInput.StartYear, 
			EndYear: careerInput.EndYear,
		}
		careersParams = append(careersParams, careerParams)
	}

	var skillsParams []user.SkillParams
	for _, skillInput := range input.SkillsInput {
		skillParams :=	user.SkillParams{
			TagID: skillInput.TagID, 
			Evaluation: skillInput.Evaluation, 
			Years: skillInput.Years,
		}
		skillsParams = append(skillsParams, skillParams)
	}

	user, err := user.CreateUserAggregate(userParams, careersParams, skillsParams)
	
	if err != nil {
		return nil, fmt.Errorf("ユーザの作成に失敗しました")
	}
	return user, nil
}

func saveUser(user *user.User, userRepo user.UserRepository) error {
	err := userRepo.Store(user)

	if err != nil {
		return fmt.Errorf("ユーザの登録に失敗しました")
	}
	return nil
}
