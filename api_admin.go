package gong

import (
	goreq "github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
)

type apiAdmin struct {
	n *node
}

func (aa *apiAdmin) url() string { return aa.n.url + "/apis" }

func (aa *apiAdmin) Add(def *ApiDefinition) (*ApiObject, error) {

	obj := ApiObject{}

	resp, _, errs := goreq.New().Post(aa.url() + "/").
		Send(def).EndStruct(&obj)
	err := combineRequestErrors(resp, errs)
	if err != nil {
		return nil, errors.Wrap(err, "http request add api")
	}

	return &obj, nil
}

func (aa *apiAdmin) Retrieve(name_or_id string) (*ApiObject, error) {

	obj := ApiObject{}

	resp, _, errs := goreq.New().Get(aa.url() + "/" + name_or_id).
		EndStruct(&obj)
	err := combineRequestErrors(resp, errs)
	if err != nil {
		return nil, errors.Wrapf(err,
			"http request to get api of %v", name_or_id)
	}

	return &obj, nil
}

func (aa *apiAdmin) List(filters ...string) ([]*ApiObject, error) {

	// TODO
	return nil, nil
}

func (aa *apiAdmin) Update(
	name_or_id string, def *ApiDefinition) (*ApiObject, error) {

	obj := ApiObject{}

	resp, _, errs := goreq.New().Patch(aa.url() + "/" + name_or_id).
		Send(def).EndStruct(&obj)
	err := combineRequestErrors(resp, errs)
	if err != nil {
		return nil, errors.Wrapf(err,
			"http request to update api of %v", name_or_id)
	}

	return &obj, nil
}

func (aa *apiAdmin) Replace(def *ApiDefinition) (*ApiObject, error) {

	// TODO
	return nil, nil
}

func (aa *apiAdmin) Delete(name_or_id string) error {

	resp, _, errs := goreq.New().Delete(aa.url() + "/" + name_or_id).End()
	err := combineRequestErrors(resp, errs)
	if err != nil {
		return errors.Wrapf(err, "http request delete api %v", name_or_id)
	}

	return nil
}
