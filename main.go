package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
	"sync"
	"time"
)

const (
	Tick = 120 * time.Millisecond
)

var (
	done = make(chan struct{})
	mainWG sync.WaitGroup
)

func main() {
	// 初始化全局对象
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	// 初始化地图
	g.SetManagerFunc(layout)

	// 键盘事件响应
	if err := keybindings(g); err != nil{
		log.Panicln(err)
	}

	// Init snake, apple
	snakeSetup()
	appleSetup()

	mainWG.Add(1)
	go snakeRun(g, &mainWG)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	mainWG.Wait()
}

func snakeRun(g *gocui.Gui, wg *sync.WaitGroup){
	defer wg.Done()
	for {
		select {
			case <- done:
				return
			case <- time.After(Tick):
				g.Update(func(gui *gocui.Gui) error {
					v, err := g.View("main")
					if err != nil {
						return err
					}
					v.Clear()
					// 默认向右
					snakeMove(snakeVector, g, v)
					// 渲染地图
					for x:=0;x<MaxX;x++{
						for y:=0;y<MaxY;y++{
							if snakeRecord[x][y] == FillSnake {
								fmt.Fprintf(v, "%c", FillSnake)
							}else if x == apple.X && y == apple.Y {
								fmt.Fprintf(v, "%c", FillApple)
							}else{
								fmt.Fprintf(v, "%c", FillBack)
							}
							if y % MaxY == 0 && x != 0{
								fmt.Fprintf(v, "\n")
							}
						}
					}
					return nil
				})
		}
	}
}

func snakeMove(vector string, g *gocui.Gui, v *gocui.View) error {
	switch vector {
	case "right":
		MoveRight(g, v)
	case "left":
		MoveLeft(g, v)
	case "up":
		MoveUp(g, v)
	case "down":
		MoveDown(g, v)
	}

	return nil
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone, MoveLeft); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone, MoveRight); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, MoveUp); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, MoveDown); err != nil {
		log.Panicln(err)
	}
	return nil
}

func layout(g *gocui.Gui) error {
	if _, err := g.SetView("main", 0, 0, MaxX * 0.7, MaxY * 0.6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	if _, err := g.SetView("console", MaxX * 0.8, 0, MaxX + 5, MaxY * 0.2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	return nil
}


func quit(g *gocui.Gui, v *gocui.View) error {
	close(done)
	return gocui.ErrQuit
}