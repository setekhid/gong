package gong

import (
	"github.com/hashicorp/go-multierror"
	goreq "github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
	"net/http"
)

func combineRequestErrors(resp goreq.Response, errs []error) error {

	if len(errs) <= 0 && resp.StatusCode/100 != 2 {
		errs = append(errs, errors.New(http.StatusText(resp.StatusCode)))
	}
	if len(errs) > 0 {
		return multierror.Append(nil, errs...)
	}

	return nil
}
