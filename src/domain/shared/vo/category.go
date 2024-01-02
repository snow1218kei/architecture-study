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
	case Programming.String():
		return Programming, nil
	case Marketing.String():
		return Marketing, nil
	case Design.String():
		return Design, nil
	case Writing.String():
		return Writing, nil
	case Movie.String():
		return Movie, nil
	case Business.String():
		return Business, nil
	case Language.String():
		return Language, nil
	case Lifestyle.String():
		return Lifestyle, nil
	default:
		return "", apperr.BadRequestf("無効なcategory: %s", category)
	}
}

func (category Category) String() string {
	return string(category)
}
