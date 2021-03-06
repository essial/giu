package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
)

var (
	names []string
)

func buildRows() []*g.RowWidget {
	rows := make([]*g.RowWidget, len(names))

	for i := range rows {
		rows[i] = g.Row(
			g.Label(fmt.Sprintf("%d", i)),
			g.Label(names[i]),
		)
	}

	return rows
}

func loop() {
	g.SingleWindow("Huge list demo").Layout(g.Layout{
		g.Label("Use FastTable to display huge amount of rows"),
		g.Label("Note: FastTable only works if all rows have same height"),
		g.Child("Container").Layout(g.Layout{
			g.FastTable("Fast table").Rows(buildRows()),
		}),
	})
}

func main() {
	names = make([]string, 10000)
	for i := range names {
		names[i] = fmt.Sprintf("Huge list name demo %d", i)
	}

	wnd := g.NewMasterWindow("Huge list demo", 800, 600, 0, nil)
	wnd.Run(loop)
}
