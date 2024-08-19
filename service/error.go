package service

import "fmt"

var (
	ErrMissingName        = Generic{Msg: "missing Name", Code: 1}
	ErrMissingProjectType = Generic{Msg: "missing project type", Code: 2}
	ErrNotInserted        = Generic{Msg: "insert error", Code: 3}
	ErrUpdate             = Generic{Msg: "update error", Code: 4}
	ErrMalformedRequest   = Generic{Code: 5, Msg: "malformed request"}
	ErrDelete             = Generic{Code: 6, Msg: "unable to perform delete"}
	ErrMissingId          = Generic{Code: 7, Msg: "err missing id"}
	ErrInvalidEmail       = Generic{Code: 8, Msg: "invalid email"}
)

type Generic struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

func (g Generic) Error() string {
	return fmt.Sprintf("code: %d, msg: %s\n", g.Code, g.Msg)
}
