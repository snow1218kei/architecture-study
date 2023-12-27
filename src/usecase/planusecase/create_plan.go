package planusecase

import (
	"context"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/mentoringplan"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/tag"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
	"strings"
)

type CreatePlanUsecase struct {
	planRepo mentoringplan.PlanRepository
	tagRepo  tag.TagRepository
	userRepo user.UserRepository
}

func NewCreatePlanUsecase(planRepo mentoringplan.PlanRepository, tagRepo tag.TagRepository, userRepo user.UserRepository) *CreatePlanUsecase {
	return &CreatePlanUsecase{
		planRepo,
		tagRepo,
		userRepo,
	}
}

func (u *CreatePlanUsecase) Exec(ctx context.Context, input *CreatePlanInput) error {
	tagExistsService := tag.NewTagIDExistsService(u.tagRepo)
	tagIDs := make([]tag.TagID, len(input.TagIDs))
	for _, tagID := range input.TagIDs {
		newTagID, err := tag.NewTagIDByVal(tagID)
		if err != nil {
			return err
		}
		tagIDs = append(tagIDs, newTagID)
	}
	isTagExist, existTagIDs, err := tagExistsService.Exec(ctx, tagIDs)
	if err != nil {
		return err
	}
	if !isTagExist {
		// existTagIDsをマップに保存し、検索を高速化
		found := make(map[tag.TagID]bool)
		for _, existID := range existTagIDs {
			found[existID] = true
		}

		// 存在しないタグを見つける
		var notExistTags []string
		for _, tagID := range tagIDs {
			if !found[tagID] {
				// TagIDをstringに変換するメソッドを使用して、ここを調整してください。
				notExistTags = append(notExistTags, tagID.String())
			}
		}

		// 存在しないタグのIDを結合し、エラーメッセージとして返す
		notExistTagIDsStr := strings.Join(notExistTags, ", ")
		return apperr.BadRequestWrapf(err, "存在していないtagIDが含まれています: %s", notExistTagIDsStr)
	}

	userExistsService := user.NewIsExistByIDService(u.userRepo)
	userID, err := user.NewUserIDByVal(input.UserID)
	if err != nil {
		return err
	}
	isUserExist, err := userExistsService.Run(ctx, userID)
	if err != nil {
		return err
	}
	if isUserExist {
		return apperr.BadRequestWrapf(err, "存在しているので他の名前でお願いします: %s", input.UserID)
	}

	planParams := &mentoringplan.MentoringPlanParams{
		Title:              input.Title,
		Content:            input.Content,
		Pricing:            input.Pricing,
		Category:           input.Category,
		ConsultationMethod: input.ConsultationMethod,
		Status:             input.Status,
		TagIDs:             tagIDs,
		UserID:             userID,
	}

	plan, err := mentoringplan.NewMentoringPlan(planParams)
	if err != nil {
		return err
	}

	if err := u.planRepo.Store(ctx, plan); err != nil {
		return err
	}

	return nil
}
