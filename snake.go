package main

import (
	"container/list"
	"github.com/jroimartin/gocui"
	"math/rand"
	"sync"
	"time"
)

const (
	FillBack = ' '
	FillSnake = '+'
	FillApple = '*'
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
	apple			body
)

func snakeSetup(){
	snakeLock.Lock()
	defer snakeLock.Unlock()
	snake.X = 10
	snake.Y = 15
	snakeLoad(snake)
}

func snakeLoad(b body){
	for x := b.X;x < b.X+SnakeLength;x++{
		snakeRecord[b.Y][x] = FillSnake
		snakeList.PushBack(body{b.Y, x})
	}
}

func appleSetup(){
	rand.Seed(time.Now().UnixNano())
	apple.X = rand.Intn(19)
	apple.Y = rand.Intn(34)
	if snakeRecord[apple.X][apple.Y] == FillSnake || apple.X <= 0 || apple.Y <=0 {
		appleSetup()
	}
}

func MoveLeft(g *gocui.Gui, v *gocui.View) error {
	snakeLock.Lock()
	// 取出头尾元素, 并删除尾元素
	headBody := snakeList.Front().Value

	// 判断越界条件
	if headBody.(body).Y-1 > 0{
		// 将新元素放进头部
		snakeList.PushFront(body{headBody.(body).X, headBody.(body).Y-1})

		// 更新snakeRecord
		lastBody := snakeList.Remove(snakeList.Back())
		// 吃苹果
		if headBody.(body).X == apple.X && headBody.(body).Y-1 == apple.Y {
			snakeList.PushFront(apple)
			appleSetup()
		}else{
			snakeRecord[lastBody.(body).X][lastBody.(body).Y] = FillBack
		}
		snakeRecord[headBody.(body).X][headBody.(body).Y-1] = FillSnake

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
		// 吃苹果
		if headBody.(body).X == apple.X && headBody.(body).Y+1 == apple.Y {
			snakeList.PushFront(apple)
			appleSetup()
		}else{
			snakeRecord[lastBody.(body).X][lastBody.(body).Y] = FillBack
		}
		snakeRecord[headBody.(body).X][headBody.(body).Y+1] = FillSnake

		snakeVector = "right"
	}

	defer snakeLock.Unlock()
	return nil
}

func MoveUp(g *gocui.Gui, v *gocui.View) error {
	snakeLock.Lock()
	// 取出头尾元素, 并删除尾元素
	headBody := snakeList.Front().Value

	// 判断越界条件
	if headBody.(body).X > 0 {
		snakeList.PushFront(body{headBody.(body).X-1, headBody.(body).Y})

		// 更新snakeRecord
		lastBody := snakeList.Remove(snakeList.Back())
		// 吃苹果
		if headBody.(body).X == apple.X-1 && headBody.(body).Y == apple.Y {
			snakeList.PushFront(apple)
			appleSetup()
		}else{
			snakeRecord[lastBody.(body).X][lastBody.(body).Y] = FillBack
		}
		snakeRecord[headBody.(body).X-1][headBody.(body).Y] = FillSnake

		snakeVector = "up"
	}

	defer snakeLock.Unlock()
	return nil
}

func MoveDown(g *gocui.Gui, v *gocui.View) error {
	snakeLock.Lock()
	// 取出头尾元素, 并删除尾元素
	headBody := snakeList.Front().Value

	// 判断越界条件
	if headBody.(body).X+1 < 20 {
		snakeList.PushFront(body{headBody.(body).X+1, headBody.(body).Y})

		// 更新snakeRecord
		lastBody := snakeList.Remove(snakeList.Back())
		// 吃苹果
		if headBody.(body).X == apple.X+1 && headBody.(body).Y == apple.Y {
			snakeList.PushFront(apple)
			appleSetup()
		}else{
			snakeRecord[lastBody.(body).X][lastBody.(body).Y] = FillBack
		}
		snakeRecord[headBody.(body).X+1][headBody.(body).Y] = FillSnake

		snakeVector = "down"
	}

	defer snakeLock.Unlock()
	return nil
}
