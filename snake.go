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
	snakeVector		= "right"
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
	for x := b.X;x < b.X+SnakeLength;x++{
		snakeRecord[b.Y][x] = FillSnake
		snakeList.PushBack(body{b.Y, x})
	}
}

func MoveLeft(g *gocui.Gui, v *gocui.View) error {
	snakeLock.Lock()
	headBody := snakeList.Front().Value
	// 取出头尾元素, 并删除尾元素
	//defer fmt.Printf("x: %d, y: %d\n", headBody.(body).X, headBody.(body).Y)

	// 判断越界条件
	if headBody.(body).Y-1 > 0{
		// 将新元素放进头部
		snakeList.PushFront(body{headBody.(body).X, headBody.(body).Y-1})

		// 更新snakeRecord
		lastBody := snakeList.Remove(snakeList.Back())
		snakeRecord[headBody.(body).X][headBody.(body).Y-1] = FillSnake
		snakeRecord[lastBody.(body).X][lastBody.(body).Y] = FillBack

		snakeVector = "left"
	}

	defer snakeLock.Unlock()
	return nil
}

func MoveRight(g *gocui.Gui, v *gocui.View) error {
	snakeLock.Lock()
	headBody := snakeList.Front().Value

	// 判断越界条件
	if headBody.(body).Y+1 < MaxY{
		// 将新元素放进头部
		snakeList.PushFront(body{headBody.(body).X, headBody.(body).Y+1})

		// 更新snakeRecord
		lastBody := snakeList.Remove(snakeList.Back())
		snakeRecord[headBody.(body).X][headBody.(body).Y+1] = FillSnake
		snakeRecord[lastBody.(body).X][lastBody.(body).Y] = FillBack

		snakeVector = "right"
	}

	defer snakeLock.Unlock()
	return nil
}

func MoveUp(g *gocui.Gui, v *gocui.View) error {
	snakeLock.Lock()
	headBody := snakeList.Front().Value

	// 取出头尾元素, 并删除尾元素
	//defer fmt.Printf("x: %d, y: %d\n", headBody.(body).X, headBody.(body).Y)

	// 判断越界条件
	if headBody.(body).X > 0 {
		snakeList.PushFront(body{headBody.(body).X-1, headBody.(body).Y})

		// 更新snakeRecord
		lastBody := snakeList.Remove(snakeList.Back())
		snakeRecord[headBody.(body).X-1][headBody.(body).Y] = FillSnake
		snakeRecord[lastBody.(body).X][lastBody.(body).Y] = FillBack

		snakeVector = "up"
	}

	defer snakeLock.Unlock()
	return nil
}

func MoveDown(g *gocui.Gui, v *gocui.View) error {
	snakeLock.Lock()
	headBody := snakeList.Front().Value

	// 取出头尾元素, 并删除尾元素
	//defer fmt.Printf("x: %d, y: %d\n", headBody.(body).X, headBody.(body).Y)

	// 判断越界条件
	if headBody.(body).X+1 < 20 {
		snakeList.PushFront(body{headBody.(body).X+1, headBody.(body).Y})

		// 更新snakeRecord
		lastBody := snakeList.Remove(snakeList.Back())
		snakeRecord[headBody.(body).X+1][headBody.(body).Y] = FillSnake
		snakeRecord[lastBody.(body).X][lastBody.(body).Y] = FillBack

		snakeVector = "down"
	}

	defer snakeLock.Unlock()
	return nil
}
