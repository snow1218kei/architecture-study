package requirement

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type Category string

const (
	Programming Category = "プログラミング"
	Marketing   Category = "マーケティング"
	Design      Category = "デザイン"
	Writing     Category = "ライティング"
	Video       Category = "動画・映像"
	Business    Category = "ビジネス"
	Language    Category = "語学"
	Lifestyle   Category = "ライフスタイル"
)

func validateCategory(category Category) error {
	switch category {
	case Programming, Marketing, Design, Writing, Video, Business, Language, Lifestyle:
		return nil
	default:
		return apperr.BadRequestf("categoryはゆう: %d", category)
	}
}
