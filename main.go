package main

import (
	"context"

	"github.com/YumaFuu/ssm-tui/app"
	"github.com/YumaFuu/ssm-tui/aws/ssm"
)

func main() {
	ctx := context.Background()

	client, err := ssm.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	params, err := ssm.List(client, "/")
	if err != nil {
		panic(err)
	}
	app.NewApp(params).Run()

	// tree.SetSelectedFunc(func(node *tview.TreeNode) {
	// 	if len(node.GetChildren()) == 0 {
	// 		// If the node has no children, it might be a file, display its value
	// 		displayNodeInfo(node)
	// 	}
	// 	if node.IsExpanded() {
	// 		node.CollapseAll()
	// 	} else {
	// 		node.ExpandAll()
	// 	}
	// })

	// app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	// 	if event.Key() == tcell.KeyRune && app.GetFocus() == tree {
	// 		// Treeにフォーカスがあるとき
	// 		node := tree.GetCurrentNode()

	// 		switch event.Rune() {
	// 		case 'c':
	// 			if node != nil && len(node.GetChildren()) == 0 {
	// 				param := node.GetReference().(types.Parameter)
	// 				if err := clipboard.WriteAll(*param.Value); err != nil {
	// 					infoView.SetText(
	// 						fmt.Sprintf("[red]Error copying to clipboard: %s", err),
	// 					)
	// 				} else {
	// 					infoView.SetText("[green]Value copied to clipboard")
	// 				}
	// 			}
	// 		case 'i':
	// 			if node != nil && len(node.GetChildren()) == 0 {
	// 				valueView.SetBorderColor(tcell.ColorBlue)
	// 				app.SetFocus(valueView)
	// 				// Show raw value
	// 				param := node.GetReference().(types.Parameter)
	// 				valueView.SetText(*param.Value, true)
	// 			}
	// 		case 'o':
	// 			// CreateViewNewを表示
	// 			layout.ShowPage("new")
	// 			app.SetFocus(createView)

	// 		case 'j':
	// 			tree.Move(1)
	// 		case 'k':
	// 			tree.Move(-1)
	// 		case 'q':
	// 			app.Stop()
	// 		}
	// 		return nil
	// 	} else if app.GetFocus() == valueView {
	// 		// ValueViewにフォーカスがあるとき
	// 		switch event.Key() {
	// 		case tcell.KeyEsc:
	// 			valueView.SetBorderColor(tcell.ColorDefault)
	// 			app.SetFocus(tree)

	// 			node := tree.GetCurrentNode()
	// 			prev := *node.GetReference().(types.Parameter).Value
	// 			new := valueView.GetText()
	// 			if prev != new {
	// 				layout.ShowPage("confirm")
	// 				layout.HidePage("confirm")
	// 				infoView.SetText(
	// 					fmt.Sprintf("[green]Value updated: \n%s -> %s", prev, new),
	// 				)
	// 			}
	// 		}
	// 	} else if app.GetFocus() == createView {
	// 		switch event.Key() {
	// 		case tcell.KeyEsc:
	// 			app.SetFocus(tree)
	// 		}
	// 	}

	// 	// Global
	// 	if event.Key() == tcell.KeyCtrlC {
	// 		app.SetFocus(tree)
	// 		valueView.SetBorderColor(tcell.ColorDefault)
	// 		return nil
	// 	}

	// 	return event
	// })

	// if err := app.SetRoot(layout, true).Run(); err != nil {
	// 	panic(err)
	// }
}
