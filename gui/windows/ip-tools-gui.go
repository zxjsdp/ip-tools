package main

import (
	"log"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/zxjsdp/ip-tools/config"
	"github.com/zxjsdp/ip-tools/ip"
	"golang.org/x/crypto/openpgp/errors"
)

const (
	Width  = 950
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
						HScroll:  true,
						VScroll:  true,
					},
					TextEdit{
						AssignTo: &mw.outputArea,
						Text:     "",
						Font:     Font{Family: "Consolas", PointSize: 10},
						HScroll:  true,
						VScroll:  true,
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
	input, err := mw.getAndCheckInputContent()
	if err == nil {
		input = ip.PrepareInputString(input)
		ips := ip.GetSingleIPsByRegexp(input)
		ipErr := mw.checkIps(ips)
		if ipErr == nil {
			mw.outputArea.SetText(strings.Join(ips, "\r\n"))
		}
	}
}

func (mw *MyMainWindow) getRangeButtonTriggered() {
	input, err := mw.getAndCheckInputContent()
	if err == nil {
		input = ip.PrepareInputString(input)
		ips := ip.GetSingleIPsByRegexp(input)
		ipErr := mw.checkIps(ips)
		if ipErr == nil {
			result := ip.GetRange(ips)
			mw.outputArea.SetText(strings.Join(result, "\r\n"))
		}
	}
}

func (mw *MyMainWindow) helpActionTriggered() {
	walk.MsgBox(mw, "帮助", config.Help, walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) aboutActionTriggered() {
	walk.MsgBox(mw, "关于", config.About, walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) errorActionTriggered(message string) {
	walk.MsgBox(mw, "错误", message, walk.MsgBoxIconError)
}

func (mw *MyMainWindow) getAndCheckInputContent() (string, error) {
	input := mw.inputArea.Text()
	if len(input) == 0 || len(strings.TrimSpace(input)) == 0 {
		mw.errorActionTriggered(config.BlankContentErrorMsg)
		return "", errors.InvalidArgumentError(config.BlankContentErrorMsg)
	}
	return strings.TrimSpace(input), nil
}

func (mw *MyMainWindow) checkIps(ips []string) error {
	if len(ips) == 0 {
		mw.errorActionTriggered(config.InvalidContentErrorMsg)
		return errors.InvalidArgumentError(config.InvalidContentErrorMsg)
	}
	return nil
}
