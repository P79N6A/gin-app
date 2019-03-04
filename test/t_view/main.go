/**
 * @description t_view
 * @author zhangbingbing@baidu.com
 * @date 2019-02-22
 */
package main

import (
	"github.com/rivo/tview"
)

func main() {
	// box := tview.NewBox().SetBorder(true).SetTitle("hello,world")
	// if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
	// 	panic(err)
	// }
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("List item 1", "some text", 'a', nil).
		AddItem("List item 2", "some text", 'b', nil).
		AddItem("List item 3", "some text", 'c', nil).
		AddItem("List item 4", "some text", 'd', nil).
		AddItem("Quit", "press to exit", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).Run(); err != nil {
		panic(err)
	}

}
