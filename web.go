package main

import (
	"log"
	"net/http"

	"github.com/devilcove/mux"
	"maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

func main() {
	router := mux.NewRouter(mux.Logger)
	router.Get("/{$}", nusak)
	router.Get("/home", home)
	router.Get("/about", about)
	router.Get("/projects", projects)
	router.Get("/contact", contact)
	router.Get("/favicon.ico", favicon)
	router.Group("/assets", serveAssets)
	router.Group("/images", serveImages)
	router.Run(":8080")
}

func homeContent() gomponents.Node {
	return gomponents.Group{
		html.H1(gomponents.Text("Matthew R Kasun")),
		html.Img(html.Src("/images/matt.jpg"), html.Alt("matt"), html.Style("width:200px")),
		html.P(
			gomponents.Text("Systems/Software Engineer"),
			html.Br(),
			gomponents.Text("semi-retired"),
		),
		html.H3(gomponents.Text("complexity;"), html.Br(),
			gomponents.Text("Stupidity is overwhelmed;"), html.Br(),
			gomponents.Text("Mediocrity struggles with it;"), html.Br(),
			gomponents.Text("Intelligence deals with ease;"), html.Br(),
			gomponents.Text("But genius eliminates it "),
		),
		html.Img(html.Src("/images/dontwin.png"), html.Alt("no windows")),
		html.P(gomponents.Text("Open Source Enthusiast")),
	}
}

func home(w http.ResponseWriter, _ *http.Request) {
	if err := html.Div(html.ID("home"), html.Class("content"),
		homeContent(),
	).Render(w); err != nil {
		log.Println("render home", err)
	}
}

func about(w http.ResponseWriter, _ *http.Request) {
	if err := html.Div(html.ID("home"), html.Class("content"),
		html.Div(html.Class("list"),
			html.H1(html.Style("font-size:56px"), gomponents.Text("Matthew R. Kasun")),
			html.P(gomponents.Text("B.Eng, M.Eng, CD")),
			html.H2(gomponents.Text("Employment History")),
			html.Ul(html.Style("text-align:left; list-style-position:inside;"),
				html.Li(gomponents.Text("Semi Retirement: Systems/Software Contracting (since 2020)")),
				html.Ul(
					html.Li(gomponents.Text("Kilobryte")),
					html.Li(gomponents.Text("Netmaker")),
				),
				html.Br(),
				html.Li(gomponents.Text("General Dynamics Misssion Systems (2005 - 2020)")),
				html.Ul(
					html.Li(gomponents.Text("Mk-46 Torpedo Project Engineer")),
					html.Li(gomponents.Text("CP-140 Aurora Project")),
					html.Ul(
						html.Li(gomponents.Text("Chief Engineer")),
						html.Li(gomponents.Text("Software Engineering Manager")),
						html.Li(gomponents.Text("Software Team Lead")),
						html.Li(gomponents.Text("Systems Engineer")),
					),
				),
				html.Br(),
				html.Li(gomponents.Text("L3-Systems (Northrup Gruman) (2004 - 2005)")),
				html.Ul(
					html.Li(gomponents.Text("Avionics Systems Specialist")),
				),
				html.Br(),
				html.Li(gomponents.Text("N-able (2000 - 2004)                       ")),
				html.Ul(
					html.Li(gomponents.Text("Director of Engineering")),
				),
				html.Br(),
				html.Li(gomponents.Text("CompEngServ (1996 - 2000)                   ")),
				html.Ul(
					html.Li(gomponents.Text("Director of Engineering")),
				),
				html.Br(),
				html.Li(gomponents.Text("Royal Canadian Air Force(1975 - 1996)       ")),
				html.Ul(
					html.Li(gomponents.Text("PMO - Maritime Helicopter Program - Ottawa")),
					html.Li(gomponents.Text("Directorate of Avionics, Simulators and Photography, Ottawa")),
					html.Li(gomponents.Text("Post-Graduate Studies, Carleton University, Ottawa")),
					html.Li(gomponents.Text("Directorate of Avionics, Simulators and Photography, Ottawa")),
					html.Li(gomponents.Text("Air Command Headquarters, Winnipeg")),
					html.Li(gomponents.Text("14 Wing Greenwood, Nova Scotia")),
					html.Li(gomponents.Text("CFSOAE, Borden")),
					html.Li(gomponents.Text("Royal Military College, Kingston")),
					html.Li(gomponents.Text("Royal Roads Military College, Victoria")),
					html.Li(gomponents.Text("8th Canadian Hussars, Petawawa")),
				),
			),
		),
	).Render(w); err != nil {
		log.Println("render", err)
	}
}

func projects(w http.ResponseWriter, _ *http.Request) { //nolint:funlen
	if err := html.Div(html.ID("home"), html.Class("content"),
		html.Div(html.Class("list"),
			html.H1(html.Style("font-size:56px"), gomponents.Text("Matthew R. Kasun")),
			html.P(gomponents.Text("B.Eng, M.Eng, CD")),
			html.Ul(html.Style("text-align:left; list-style-type:none;"),
				html.Li(gomponents.Text("Contracted Development")),
				html.Ul(
					html.Li(html.A(html.Href("https://kilobryte.com"), html.Target("_blank"),
						html.Img(html.Src("images/kilobryte.jpg"), html.Style("height:50px"), html.Alt("kilobryte")))),
					html.Ul(
						html.Li(gomponents.Text("Golfball Customization Backend")),
					),
					html.Li(html.A(html.Href("https://netmaker.io"), html.Target("_blank"),
						html.Img(html.Src("images/netmakerlogo.png"), html.Style("height:50px"), html.Alt("netmaker")))),
					html.Ul(
						html.Li(gomponents.Text("WireGuard® VPN")),
					),
				),
				html.Li(gomponents.Text("Personal Projects")),
				html.Ul(
					html.Li(
						html.A(html.Href("https://github.com/devilcove/plexus#readme"), html.Target("_blank"),
							html.Span(html.Class("fab fa-github"), html.Style("color:white;")),
							gomponents.Text("Plexus"))),
					html.Ul(
						html.Li(gomponents.Text("wireguard® management")),
					),
					html.Li(
						html.A(html.Href("https://github.com/devilcove/timetraced#readme"), html.Target("_blank"),
							html.Span(html.Class("fab fa-github"), html.Style("color:white;")),
							gomponents.Text("TimeTrace"))),
					html.Ul(
						html.Li(gomponents.Text("time tracking web application")),
					),
					html.Li(
						html.A(html.Href("https://github.com/devilcove/configuration#readme"), html.Target("_blank"),
							html.Span(html.Class("fab fa-github"), html.Style("color:white;")),
							gomponents.Text("Configuration"))),
					html.Ul(
						html.Li(gomponents.Text("application configuration management")),
					),
					html.Li(
						html.A(html.Href("https://github.com/devilcove/mux#readme"), html.Target("_blank"),
							html.Span(html.Class("fab fa-github"), html.Style("color:white;")),
							gomponents.Text("Mux"))),
					html.Ul(
						html.Li(gomponents.Text("a small, idomatic http router")),
					),
					html.Li(
						html.A(html.Href("https://github.com/devilcove/cookie#readme"), html.Target("_blank"),
							html.Span(html.Class("fab fa-github"), html.Style("color:white;")),
							gomponents.Text("Cookie"))),
					html.Ul(
						html.Li(gomponents.Text("encrypted HTTP cookie management")),
					),
					html.Li(
						html.A(html.Href("https://github.com/devilcove/uptime#readme"), html.Target("_blank"),
							html.Span(html.Class("fab fa-github"), html.Style("color:white;")),
							gomponents.Text("Uptime"))),
					html.Ul(
						html.Li(gomponents.Text("monitor uptime of web servers")),
					),
					html.Li(
						html.A(html.Href("https://github.com/mattkasun/sshlogin#readme"), html.Target("_blank"),
							html.Span(html.Class("fab fa-github"), html.Style("color:white;")),
							gomponents.Text("sshlogin"))),
					html.Ul(
						html.Li(gomponents.Text("Proof of Concept: application login using ssh key")),
					),
					html.Li(
						html.A(html.Href("https://github.com/mattkasun/sshbrowser#readme"), html.Target("_blank"),
							html.Span(html.Class("fab fa-github"), html.Style("color:white;")),
							gomponents.Text("sshbrowser"))),
					html.Ul(
						html.Li(gomponents.Text("Proof of Concept: website login using ssh key")),
					),
					html.Li(
						html.A(html.Href("https://github.com/mattkasun/dyndns#readme"), html.Target("_blank"),
							html.Span(html.Class("fab fa-github"), html.Style("color:white;")),
							gomponents.Text("dyndns"))),
					html.Ul(
						html.Li(gomponents.Text("update dns records at digitalocean")),
					),
				),
			),
		),
	).Render(w); err != nil {
		log.Println("render projects", err)
	}
}

func contact(w http.ResponseWriter, _ *http.Request) {
	if err := html.Div(html.ID("home"), html.Class("content"),
		html.H2(gomponents.Text("Contact Me")),
		html.A(html.Href("https://www.google.com/maps/@45.2496825,-76.1298873,99550m/data=!3m1!1e3?entry=ttu"),
			html.Target("_blank"), html.I(html.Class("fa fa-map-marker fa-fw")),
			gomponents.Text("Ottawa Ont Canada")),
		html.P(html.I(html.Class("fa fa-phone fa-fw")), gomponents.Text("Phone: +01 613-452-0600")),
		html.A(html.Href("mailto:mkasun@nusak.ca?subject=Hello"), html.Target("_blank"),
			html.I(html.Class("fa fa-envelope fa-fw")), gomponents.Text("Email: mkasun@nusak.ca")),
	).Render(w); err != nil {
		log.Println("render contact", err)
	}
}

func nusak(w http.ResponseWriter, _ *http.Request) {
	if err := layout([]gomponents.Node{
		html.Nav(html.Class("sidebar"),
			html.Button(html.Type("button"), hx.Target("#home"), hx.Swap("outerHTML"), hx.Get("/home"),
				html.Style("font-size:48px"), html.Class("fa fa-home")),
			html.P(gomponents.Text("HOME")),
			html.Button(html.Type("button"), hx.Target("#home"), hx.Swap("outerHTML"), hx.Get("/about"),
				html.Style("font-size: 48px"), html.I(html.Class("fa fa-user"))),
			html.P(gomponents.Text("ABOUT")),
			html.Button(html.Type("button"), hx.Target("#home"), hx.Swap("outerHTML"), hx.Get("/projects"),
				html.Style("font-size: 48px"), html.I(html.Class("fa fa-tools"))),
			html.P(gomponents.Text("PROJECTS")),
			html.Button(html.Type("button"), hx.Target("#home"), hx.Swap("outerHTML"), hx.Get("/contact"),
				html.Style("font-size: 48px"), html.I(html.Class("fa fa-envelope"))),
			html.P(gomponents.Text("CONTACT")),
		),
		container(true,
			html.Div(html.Class("content"), html.ID("home"), hx.Target("this"), hx.Swap("outerHTML"),
				homeContent(),
			),
		),
	}).Render(w); err != nil {
		log.Println("render main page", err)
	}
}

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "images/dontwin.svg")
}

func serveAssets(_ http.Handler) http.Handler {
	return http.FileServer(http.Dir("assets"))
}

func serveImages(_ http.Handler) http.Handler {
	return http.FileServer(http.Dir("images"))
}

func footer() gomponents.Node {
	return html.Div(
		html.Class("content"),
		html.Style("filter: invert(1)"),
		html.Br(),
		html.A(
			html.Img(html.Src("/images/digitalocean.svg"), html.Style("width:42px;height:42px;")),
			html.Href("https://m.do.co/c/7ecc5675b40b"),
			html.Target("_blank"),
			html.Alt("digital ocean"),
		),
		html.A(html.Img(html.Src("/images/github.svg"), html.Style("width:42px;height:42px;")),
			html.Href("https://github.com/mattkasun"), html.Target("_blank"),
			html.Alt("githbub"),
		),
		html.A(html.Img(html.Src("/images/reddit.svg"), html.Style("width:42px;height:42px;")),
			html.Href("https://www.reddit.com/user/dlrow-olleh"), html.Target("_blank"),
			html.Alt("reddit"),
		),
		html.A(html.Img(html.Src("/images/docker.svg"), html.Style("width:42px;height:42px;")),
			html.Href("https://hub.docker.com/u/nusak"), html.Target("_blank"),
			html.Alt("docker"),
		),
		html.A(html.Img(html.Src("/images/youtube.svg"), html.Style("width:42px;height:42px;")),
			html.Href("https://www.youtube.com/channel/UCpPbdIRA0gpxFOQmdLIUy_w?view_as=public"),
			html.Target("_blank"), html.Alt("youtube"),
		),
		html.Br(),
		html.A(html.Img(html.Src("/images/bmc-button.svg")),
			html.Href("https://www.buymeacoffee.com/mkasun"), html.Target("_blank"),
		),
	)
}
