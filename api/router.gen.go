// Code generated by 'malatd gen' command.
// DO NOT EDIT!

package api

import (
	td "github.com/swxctx/malatd"
)

func Route(srv *td.Server, rootGroup string) {
	//自定义路由处理
	routeLogic(srv, rootGroup)

	// APIs...
	{

		v1_test := srv.Group(rootGroup + "/v1/test")
		v1_test.Post("/ping", PingHandle)
		v1_test.Get("/ping", PingHandle)
	}

}
