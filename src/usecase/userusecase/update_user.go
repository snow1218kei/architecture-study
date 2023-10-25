package userusecase

import (
	"context"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase/userusecase/userinput"
)

type UpdateUserUseCase struct {
	userRepo user.UserRepository
}

func NewUpdateUserUseCase(userRepo user.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		userRepo,
	}
}

func (usecase *UpdateUserUseCase) Exec(ctx context.Context, input *userinput.UpdateUserInput) error {
	userID, err := user.NewUserIDByVal(input.UserInput.ID)
	if err != nil {
		return err
	}

	usr, err := usecase.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	userParams := user.UserUpdateParams{
		Name:     input.UserInput.Name,
		Email:    input.UserInput.Email,
		Password: input.UserInput.Password,
		Profile:  input.UserInput.Profile,
	}

	careersParams := make([]user.CareerUpdateParams, 0, len(input.CareersInput))
	for _, careerInput := range input.CareersInput {
		careerParams := user.CareerUpdateParams{
			Detail:    careerInput.Detail,
			StartYear: careerInput.StartYear,
			EndYear:   careerInput.EndYear,
		}
		careersParams = append(careersParams, careerParams)
	}

	skillsParams := make([]user.SkillUpdateParams, 0, len(input.SkillsInput))
	for _, skillInput := range input.SkillsInput {
		skillParams := user.SkillUpdateParams{
			TagID:      skillInput.TagID,
			Evaluation: skillInput.Evaluation,
			Years:      skillInput.Years,
		}
		skillsParams = append(skillsParams, skillParams)
	}

	if err := user.UpdateUserAggregate(usr, userParams, careersParams, skillsParams); err != nil {
		return err
	}

	err = usecase.userRepo.Update(ctx, usr)
	if err != nil {
		return err
	}

	return nil
}
