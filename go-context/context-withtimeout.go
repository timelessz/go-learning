package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Dog struct {
	Name string
}

func (d *Dog) dogEating(ctx context.Context, food <-chan int, finishChan chan<- struct{}) {
	var allFoodCnt int
	for {
		select {
		case foodCnt := <-food:
			time.Sleep(time.Duration(foodCnt*1) * time.Second)
			allFoodCnt += foodCnt
			fmt.Printf("dog %s eat %v\n", d.Name, foodCnt)
		case <-ctx.Done():
			fmt.Printf("dog %s eat %v totally\n", d.Name, allFoodCnt)
			finishChan <- struct{}{}
			return
		}
	}
}

func main() {
	dogNames := []string{"小黄狗", "小母狗", "小公狗", "小奶狗", "小豆狗"}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	foodChan := make(chan int, 100)
	finishChan := make(chan struct{})
	for _, name := range dogNames {
		dog := Dog{Name: name}
		go dog.dogEating(ctx, foodChan, finishChan)
	}
	allFoodCnt := 100
	for {
		time.Sleep(1 * time.Second)
		count := rand.Intn(10) + 1
		if allFoodCnt < count {
			foodChan <- allFoodCnt
			fmt.Println(allFoodCnt)
			close(foodChan)
			cancel()
			break
		} else {
			allFoodCnt -= count
			fmt.Println(count)
			foodChan <- count
		}
	}
	for i := 0; i < len(dogNames); i++ {
		<-finishChan
	}
	close(finishChan)
}
