package shared

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type Category string

const (
	Programming Category = "プログラミング"
	Marketing   Category = "マーケティング"
	Design      Category = "デザイン"
	Writing     Category = "ライティング"
	Movie       Category = "動画・映像"
	Business    Category = "ビジネス"
	Language    Category = "語学"
	Lifestyle   Category = "ライフスタイル"
)

func ValidateCategory(category string) error {
	switch category {
	case string(Programming), string(Marketing), string(Design), string(Writing), string(Movie), string(Business), string(Language), string(Lifestyle):
		return nil
	default:
		return apperr.BadRequestf("無効なcategory: %d", category)
	}
}
