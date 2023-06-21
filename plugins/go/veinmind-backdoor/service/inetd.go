package service

import (
	api "github.com/chaitin/libveinmind/go"
	"github.com/chaitin/veinmind-common-go/service/report/event"
)

func inetdBackdoorCheck(fs api.FileSystem) (bool, *event.BackdoorDetail) {
	// TODO
	return false, nil
}

func init() {
	ImageCheckFuncMap["inetd"] = inetdBackdoorCheck
	ContainerCheckFuncMap["inetd"] = inetdBackdoorCheck
}
