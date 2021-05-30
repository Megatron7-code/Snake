package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"sync"
)

const (
	FillBack = '*'
	FillSnake = '+'
	MaxX = 50
	MaxY = 35
	SnakeLength = 3
)



type body struct {
	X int
	Y int
}

var (
	snakeLock		sync.Mutex
	snake			body
	snakeRecord		[MaxX][MaxY] byte
	snakeList		[]body
)

func snakeSetup(){
	snakeLock.Lock()
	defer snakeLock.Unlock()
	snake.X = 10
	snake.Y = 15
	snakeLoad(snake)
}

func snakeLoad(b body){
	fmt.Print("snakeLoad...")
	//for x :=10;x<15;x++{
	//	snake[10][x] = '#'
	//}
	for x := b.X;x < b.X+SnakeLength;x++{
		snakeRecord[b.Y][x] = FillSnake
		snakeList = append(snakeList, body{b.Y, x})
	}

	//fmt.Printf("snakeRecord: %c", snakeRecord[15][10])
}

func MoveLeft(g *gocui.Gui, v *gocui.View) error {
	snakeLock.Lock()
	defer snakeLock.Unlock()
	length := len(snakeList)
	headBody := snakeList[0]
	lastBody := snakeList[length-1]
	newItem := body{headBody.X, headBody.Y-1}
	snakeList = append(snakeList[:length-1], newItem)

	snakeRecord[headBody.X][headBody.Y-1] = FillSnake
	snakeRecord[lastBody.X][lastBody.Y] = FillBack
	return nil
}

func MoveRight(g *gocui.Gui, v *gocui.View) error {
	snakeLock.Lock()
	defer snakeLock.Unlock()



	return nil
}

func MoveUp(g *gocui.Gui, v *gocui.View) error {
	snakeLock.Lock()
	defer snakeLock.Unlock()



	return nil
}

func MoveDown(g *gocui.Gui, v *gocui.View) error {
	snakeLock.Lock()
	defer snakeLock.Unlock()



	return nil
}
