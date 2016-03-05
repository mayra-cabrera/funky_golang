// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"funky/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/beers",
			beego.NSInclude(
				&controllers.BeersController{},
			),
		),

		beego.NSNamespace("/cats",
			beego.NSInclude(
				&controllers.CatsController{},
			),
		),

		beego.NSNamespace("/starwars",
			beego.NSInclude(
				&controllers.StarwarsController{},
			),
		),

		beego.NSNamespace("/superheroes",
			beego.NSInclude(
				&controllers.SuperheroesController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
