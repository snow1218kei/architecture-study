package presenter

type Presenter interface {
	JSON(code int, obj any)
}
