package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Cat struct {
	Name string
}

func (d *Cat) catEating(ctx context.Context, food <-chan int, finishChan chan<- struct{}) {
	var allFoodCnt int
label:
	for {
		select {
		case foodCnt := <-food:
			time.Sleep(time.Duration(foodCnt*100) * time.Microsecond)
			allFoodCnt += foodCnt
			fmt.Printf("cat %s eat %v\n", d.Name, foodCnt)
		case <-ctx.Done():
			fmt.Printf("cat %s eat %v totally\n", d.Name, allFoodCnt)
			break label
		}
	}
	finishChan <- struct{}{}
}

func main() {
	dogNames := []string{"小黄狗", "小母狗", "小公狗", "小奶狗", "小豆狗"}
	ctx, cancel := context.WithCancel(context.Background())
	foodChan := make(chan int, 100)
	finishChan := make(chan struct{})
	for _, name := range dogNames {
		cat := Cat{Name: name}
		go cat.catEating(ctx, foodChan, finishChan)
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
