package apperr

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type appErr struct {
	code  int
	msg   string
	trace error
}

type AppError interface {
	Code() int
	Msg() string
	Trace() error
	Error() string
}

type BadRequestErr struct {
	*appErr
}
type InternalErr struct {
	*appErr
}
type NotFoundErr struct {
	*appErr
}
type UnauthorizedErr struct {
	*appErr
}

func (e *appErr) Code() int {
	return e.code
}
func (e *appErr) Msg() string {
	return e.msg
}
func (e *appErr) Error() string {
	return e.msg
}
func (e *appErr) Trace() error {
	return e.trace
}

func BadRequest(msg string) *BadRequestErr {
	return &BadRequestErr{
		&appErr{
			code:  http.StatusBadRequest,
			msg:   msg,
			trace: errors.New(msg),
		},
	}
}

func BadRequestf(format string, msg ...any) *BadRequestErr {
	message := fmt.Sprintf(format, msg...)

	return &BadRequestErr{
		&appErr{
			code:  http.StatusBadRequest,
			msg:   message,
			trace: errors.Errorf(format, msg...),
		},
	}
}
func BadRequestWrapf(err2 error, format string, msg ...any) *BadRequestErr {
	message := fmt.Sprintf(format, msg...)

	// 1.20からWrapがJoinに変わる
	//err = errors.Join(err, err2)
	return &BadRequestErr{
		&appErr{
			code:  http.StatusBadRequest,
			msg:   message,
			trace: errors.Wrapf(err2, format, msg...),
		},
	}
}

func Internal(msg string) *InternalErr {
	return &InternalErr{
		&appErr{
			code:  http.StatusInternalServerError,
			msg:   msg,
			trace: errors.New(msg),
		},
	}
}

func Internalf(format string, msg ...any) *InternalErr {
	message := fmt.Sprintf(format, msg...)

	return &InternalErr{
		&appErr{
			code:  http.StatusInternalServerError,
			msg:   message,
			trace: errors.Errorf(format, msg...),
		},
	}
}
func InternalWrapf(err2 error, format string, msg ...any) *InternalErr {
	message := fmt.Sprintf(format, msg...)

	// 1.20からWrapがJoinに変わる
	//err = errors.Join(err, err2)
	return &InternalErr{
		&appErr{
			code:  http.StatusInternalServerError,
			msg:   message,
			trace: errors.Wrapf(err2, format, msg...),
		},
	}
}

func NotFound(msg string) *NotFoundErr {
	return &NotFoundErr{
		&appErr{
			code:  http.StatusNotFound,
			msg:   msg,
			trace: errors.New(msg),
		},
	}
}

func NotFoundf(format string, msg ...any) *NotFoundErr {
	message := fmt.Sprintf(format, msg...)

	return &NotFoundErr{
		&appErr{
			code:  http.StatusNotFound,
			msg:   message,
			trace: errors.Errorf(format, msg...),
		},
	}
}
func NotFoundWrapf(err2 error, format string, msg ...any) *NotFoundErr {
	message := fmt.Sprintf(format, msg...)

	// 1.20からWrapがJoinに変わる
	//err = errors.Join(err, err2)
	return &NotFoundErr{
		&appErr{
			code:  http.StatusNotFound,
			msg:   message,
			trace: errors.Wrapf(err2, format, msg...),
		},
	}
}

func Unauthorized(msg string) *UnauthorizedErr {
	return &UnauthorizedErr{
		&appErr{
			code:  http.StatusUnauthorized,
			msg:   msg,
			trace: errors.New(msg),
		},
	}
}

func Unauthorizedf(format string, msg ...any) *UnauthorizedErr {
	message := fmt.Sprintf(format, msg...)

	return &UnauthorizedErr{
		&appErr{
			code:  http.StatusUnauthorized,
			msg:   message,
			trace: errors.Errorf(format, msg...),
		},
	}
}
func UnauthorizedWrapf(err2 error, format string, msg ...any) *UnauthorizedErr {
	message := fmt.Sprintf(format, msg...)

	// 1.20からWrapがJoinに変わる
	//err = errors.Join(err, err2)
	return &UnauthorizedErr{
		&appErr{
			code:  http.StatusUnauthorized,
			msg:   message,
			trace: errors.Wrapf(err2, format, msg...),
		},
	}
}

func Is[T any](err error) bool {
	_, ok := err.(T)
	return ok
}
