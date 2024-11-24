package render

import (
	"fmt"
	"inventory-management-system/tui/api"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func RenderProductList(app *tview.Application) (*tview.Table, error) {
	products, err := api.GetAllProducts()
	if err != nil {
		return nil, err
	}

	table := tview.NewTable().
		SetBorders(true).
		SetFixed(1, 1)

	headers := []string{"ID", "Name", "Quantity", "Price"}
	for col, header := range headers {
		table.SetCell(0, col, tview.NewTableCell(header).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	}

	for row, product := range products {
		table.SetCell(row+1, 0, tview.NewTableCell(product.ID).SetAlign(tview.AlignCenter))
		table.SetCell(row+1, 1, tview.NewTableCell(product.Name).SetAlign(tview.AlignCenter))
		table.SetCell(row+1, 2, tview.NewTableCell(fmt.Sprintf("%d", product.Quantity)).SetAlign(tview.AlignCenter))
		table.SetCell(row+1, 3, tview.NewTableCell(fmt.Sprintf("$%.2f", product.Price)).SetAlign(tview.AlignCenter))
	}

	return table, nil
}
