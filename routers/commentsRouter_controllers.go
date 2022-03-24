package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/consecutivos_crud/controllers:ConsecutivoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/consecutivos_crud/controllers:ConsecutivoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/consecutivos_crud/controllers:ConsecutivoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/consecutivos_crud/controllers:ConsecutivoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/consecutivos_crud/controllers:ConsecutivoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/consecutivos_crud/controllers:ConsecutivoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
