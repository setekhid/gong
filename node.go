package gong

import (
	goreq "github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
)

type node struct {
	url string
}

func ConnectNode(url string) Node {
	return &node{url: url}
}

func (n *node) Information() (*NodeInfo, error) {

	info := NodeInfo{}

	resp, _, errs := goreq.New().Get(n.url + "/").EndStruct(&info)
	err := combineRequestErrors(resp, errs)
	if err != nil {
		return nil, errors.Wrap(err, "request node information")
	}

	return &info, nil
}

func (n *node) Status() (*NodeStatus, error) {

	status := NodeStatus{}

	resp, _, errs := goreq.New().Get(n.url + "/status").EndStruct(&status)
	err := combineRequestErrors(resp, errs)
	if err != nil {
		return nil, errors.Wrap(err, "request node information")
	}

	return &status, nil
}

func (n *node) NewApiAdmin() ApiAdmin {

	return &apiAdmin{n}
}
