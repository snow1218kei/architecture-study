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

func NewCategory(category string) (Category, error) {
	switch category {
	case string(Programming):
		return Programming, nil
	case string(Marketing):
		return Marketing, nil
	case string(Design):
		return Design, nil
	case string(Writing):
		return Writing, nil
	case string(Movie):
		return Movie, nil
	case string(Business):
		return Business, nil
	case string(Language):
		return Language, nil
	case string(Lifestyle):
		return Lifestyle, nil
	default:
		return "", apperr.BadRequestf("無効なcategory: %s", category)
	}
}
