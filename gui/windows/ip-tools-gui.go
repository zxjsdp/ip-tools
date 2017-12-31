package main

import (
	"log"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/zxjsdp/ip-tools/config"
	"github.com/zxjsdp/ip-tools/ip"
)

const (
	Title  = config.GUITitle + " " + config.GUIVersion
	Width  = 850
	Height = 700
)

type MyMainWindow struct {
	*walk.MainWindow

	prettifyIPsButton *walk.PushButton
	getRangeButton    *walk.PushButton
	inputArea         *walk.TextEdit
	outputArea        *walk.TextEdit
}

func main() {
	RunMainWindow()
}

func RunMainWindow() {
	mw := &MyMainWindow{}

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    Title,
		MinSize:  Size{Width: Width, Height: Height},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{
						Text:     "Prettify IPs",
						AssignTo: &mw.prettifyIPsButton,
						OnClicked: func() {
							mw.prettifyIPsButtonTriggered()
						},
					},
					PushButton{
						Text:     "Get IP Range",
						AssignTo: &mw.getRangeButton,
						OnClicked: func() {
							mw.getRangeButtonTriggered()
						},
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					TextEdit{
						AssignTo: &mw.inputArea,
						Text:     "",
						Font:     Font{Family: "Consolas", PointSize: 10},
					},
					TextEdit{
						AssignTo: &mw.outputArea,
						Text:     "",
						Font:     Font{Family: "Consolas", PointSize: 10},
					},
				},
			},
		},
	}.Create()); err != nil {
		log.Println("程序出现致命错误！" + err.Error())
	}

	mw.Run()
}

func (mw *MyMainWindow) prettifyIPsButtonTriggered() {
	input := mw.inputArea.Text()
	input = ip.PrepareInputString(input)
	ips := ip.GetSingleIPsByRegexp(input)
	mw.outputArea.SetText(strings.Join(ips, "\r\n"))
}

func (mw *MyMainWindow) getRangeButtonTriggered() {
	input := mw.inputArea.Text()
	input = ip.PrepareInputString(input)
	ips := ip.GetSingleIPsByRegexp(input)
	result := ip.GetRange(ips)
	mw.outputArea.SetText(strings.Join(result, "\r\n"))
}
