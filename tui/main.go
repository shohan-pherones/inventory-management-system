package main

import (
	"inventory-management-system/tui/render"
	"log"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	table, err := render.RenderProductList(app)
	if err != nil {
		log.Fatalf("Failed to render product list: %v", err)
	}

	if err := app.SetRoot(table, true).Run(); err != nil {
		log.Fatalf("Error running the TUI application: %v", err)
	}
}
