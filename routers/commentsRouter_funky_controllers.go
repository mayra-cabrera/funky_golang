package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["funky/controllers:BeersController"] = append(beego.GlobalControllerRouter["funky/controllers:BeersController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:BeersController"] = append(beego.GlobalControllerRouter["funky/controllers:BeersController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:BeersController"] = append(beego.GlobalControllerRouter["funky/controllers:BeersController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:BeersController"] = append(beego.GlobalControllerRouter["funky/controllers:BeersController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:BeersController"] = append(beego.GlobalControllerRouter["funky/controllers:BeersController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:CatsController"] = append(beego.GlobalControllerRouter["funky/controllers:CatsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:CatsController"] = append(beego.GlobalControllerRouter["funky/controllers:CatsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:CatsController"] = append(beego.GlobalControllerRouter["funky/controllers:CatsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:CatsController"] = append(beego.GlobalControllerRouter["funky/controllers:CatsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:CatsController"] = append(beego.GlobalControllerRouter["funky/controllers:CatsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:StarwarsController"] = append(beego.GlobalControllerRouter["funky/controllers:StarwarsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:StarwarsController"] = append(beego.GlobalControllerRouter["funky/controllers:StarwarsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:StarwarsController"] = append(beego.GlobalControllerRouter["funky/controllers:StarwarsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:StarwarsController"] = append(beego.GlobalControllerRouter["funky/controllers:StarwarsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:StarwarsController"] = append(beego.GlobalControllerRouter["funky/controllers:StarwarsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:SuperheroesController"] = append(beego.GlobalControllerRouter["funky/controllers:SuperheroesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:SuperheroesController"] = append(beego.GlobalControllerRouter["funky/controllers:SuperheroesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:SuperheroesController"] = append(beego.GlobalControllerRouter["funky/controllers:SuperheroesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:SuperheroesController"] = append(beego.GlobalControllerRouter["funky/controllers:SuperheroesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["funky/controllers:SuperheroesController"] = append(beego.GlobalControllerRouter["funky/controllers:SuperheroesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
