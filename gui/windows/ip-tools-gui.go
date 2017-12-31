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
	Width  = 850
	Height = 700
)

type MyMainWindow struct {
	*walk.MainWindow

	titleLabel        *walk.Label
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
		Title:    config.Title,
		MinSize:  Size{Width: Width, Height: Height},
		Layout:   VBox{},

		MenuItems: []MenuItem{
			Menu{
				Text: "&File",
				Items: []MenuItem{
					Separator{},
					Action{
						Text:        "退出",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						Text:        "帮助",
						OnTriggered: mw.helpActionTriggered,
					},
					Separator{},
					Action{
						Text:        "关于",
						OnTriggered: mw.aboutActionTriggered,
					},
				},
			},
		},


		Children: []Widget{
			Label{
				AssignTo: &mw.titleLabel,
				Text:     config.Title,
				Font:     Font{Family: "Microsoft Yahei", PointSize: 15},
			},
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

func (mw *MyMainWindow) helpActionTriggered() {
	walk.MsgBox(mw, "帮助", config.Help, walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) aboutActionTriggered() {
	walk.MsgBox(mw, "关于", config.About, walk.MsgBoxIconInformation)
}