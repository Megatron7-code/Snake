package main

import (
	"container/list"
	"fmt"
	"github.com/jroimartin/gocui"
	"sync"
)

const (
	FillBack = ' '
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
	snakeList		= list.New()
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
	first := true
	for x := b.X;x < b.X+SnakeLength;x++{
		if first{
			snakeRecord[b.Y][x] = '*'
			first = false
		}else{
			snakeRecord[b.Y][x] = FillSnake
		}
		snakeList.PushBack(body{b.Y, x})
	}

	//fmt.Printf("snakeRecord: %c", snakeRecord[15][10])
}

func MoveLeft(g *gocui.Gui, v *gocui.View) error {
	snakeLock.Lock()
	defer snakeLock.Unlock()
	// 取出头尾元素, 并删除尾元素
	headBody := snakeList.Front().Value
	lastBody := snakeList.Remove(snakeList.Back())
	// 将新元素放进头部
	snakeList.PushFront(body{headBody.(body).X, headBody.(body).Y-1})

	// 更新snakeRecord
	snakeRecord[headBody.(body).X][headBody.(body).Y-1] = FillSnake
	snakeRecord[lastBody.(body).X][lastBody.(body).Y] = FillBack
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
