package main

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/components"
	"maragu.dev/gomponents/html"
)

func layout(nodes []gomponents.Node) gomponents.Node { //nolint:ireturn
	return components.HTML5(components.HTML5Props{ //nolint:exhaustruct
		Title:    "Nusak",
		Language: "en",
		Head: []gomponents.Node{
			html.Meta(html.Name("viewport"), html.Content("width=device-width, initial-scale=1.0")),
			html.Link(html.Rel("icon"), html.Href("favicon.ico"), html.Type("image.svg")),
			html.Link(html.Rel("stylesheet"), html.Href("/assets/css/styles.css")),
			html.Link(html.Rel("stylesheet"), html.Href("https://fonts.googleapis.com/css?family=Tangerine|Indie+Flower|Roboto")), //nolint:lll
			html.Link(html.Rel("stylesheet"), html.Href("/assets/css/fontawesome.min.css")),
			html.Link(html.Rel("stylesheet"), html.Href("/assets/css/brands.min.css")),
			html.Link(html.Rel("stylesheet"), html.Href("/assets/css/solid.min.css")),
			html.Script(html.Src("https://unpkg.com/htmx.org@1.9.9"),
				html.Integrity("sha384-QFjmbokDn2DjBjq+fM+8LUIVrAgqcNW2s0PjAxHETgRn9l4fvX31ZxDxvwQnyMOX"),
				html.CrossOrigin("anonymous")),
		},
		Body: []gomponents.Node{
			html.Style("background-color: black"),
			html.Div(
				container(false, nodes...),
			),
			footer(),
		},
	})
}

func container(center bool, children ...gomponents.Node) gomponents.Node { //nolint:ireturn
	style := ""
	if center {
		style = "text-align:center; margin-left:auto; margin-right:auto"
	}
	return html.Div(
		html.Style(style),
		gomponents.Group(children),
	)
}
