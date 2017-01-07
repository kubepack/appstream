package util

import (
	"strconv"
	"strings"

	"github.com/appscode/api/dtypes"
	"github.com/appscode/errors"
	_env "github.com/appscode/go/env"
	"github.com/appscode/log"
)

func NewStatusOK(message ...string) *dtypes.Status {
	log.Debugln("Sending OK response with message ", message)
	return &dtypes.Status{
		Code:    statusCodeString(int32(dtypes.StatusCode_OK)),
		Status:  dtypes.StatusCode_OK.String(),
		Message: strings.Join(message, ";"),
	}
}

func NewStatusFromError(err error) *dtypes.Status {
	e := errors.Parse(err)
	if e == nil {
		// not an error, so sending ok response
		return NewStatusOK()
	}
	log.Debugln("Sending response ", e.Code(), " with message ", e.Message())
	s := &dtypes.Status{
		Code:    statusCodeString(errorToAPICode(e.Code())),
		Status:  e.Code(),
		Message: statusMessage(e.Code(), e.Messages()),
	}

	if errorHelp := e.Help(); errorHelp != nil {
		s.Help = &dtypes.Help{
			Url:         errorHelp.Url,
			Description: errorHelp.Description,
		}
	}
	if e.Trace() != nil {
		if !_env.FromHost().IsPublic() {
			s.AddDetails(&dtypes.ErrorDetails{
				RequestedResource: e.Error(),
				Stacktrace:        e.TraceString(),
			})
		}
	}
	log.Debugln("Sending EROR response with message", e.Messages())
	return s
}

func NewStatusUnauthorized() *dtypes.Status {
	err := errors.New().Unauthorized()
	return NewStatusFromError(err)
}

func NewStatusBadRequest(message ...string) *dtypes.Status {
	err := errors.New().WithMessage(message...).BadRequest()
	return NewStatusFromError(err)
}

func statusCodeString(code int32) string {
	return strconv.FormatInt(int64(code), 10)
}

func errorToAPICode(code string) int32 {
	c, ok := dtypes.StatusCode_value[code]
	if ok {
		return c
	}
	return -1
}

func statusMessage(code string, msg []string) string {
	switch code {
	case errors.Internal:
		return "Internal Error. Contact support, support@appscode.com"
	case errors.Unauthorized:
		return "Unauthorized. Please Login."
	}
	return strings.Join(msg, ";")
}
