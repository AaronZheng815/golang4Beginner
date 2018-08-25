package main

import (
	"sync/atomic"
	"sync"
	"time"
	"fmt"
	"math/rand"
)

//互斥锁
var lock sync.Mutex

//读写锁，应用场景，读多写少，性能比较高
var rwlock sync.RWMutex

func testMap(){
	var tmap map[int]int
	tmap = make(map[int]int,5)
	tmap[0]=1
	tmap[1]=2
	tmap[2]=3
	tmap[3]=4
	tmap[4]=5

	for i:=0; i<2;i ++{
		go func(bmap map[int]int){
			//写
			lock.Lock()
			bmap[0] = rand.Intn(100)
			lock.Unlock()
		}(tmap)
	}

	//读
	lock.Lock()
	fmt.Println(tmap)
	lock.Unlock()

	time.Sleep(time.Second)
}

func testRWMap() {
	var tmap map[int]int
	tmap = make(map[int]int,5)
	tmap[0]=1
	tmap[1]=2
	tmap[2]=3
	tmap[3]=4
	tmap[4]=5
	var count int32

	for i:=0; i<2;i ++{
		go func(bmap map[int]int){
			//写
			rwlock.Lock()
			//lock.Lock()

			time.Sleep(time.Millisecond*10)
			bmap[0] = rand.Intn(100)
			
			//lock.Unlock()
			rwlock.Unlock()
		}(tmap)
	}

	for i:=0; i<100;i ++{
		go func(bmap map[int]int){
			//读
			for{
				rwlock.RLock()
				//lock.Lock()
				time.Sleep(time.Millisecond)
				test := bmap[0]
				//lock.Unlock()
				rwlock.RUnlock()

				test ++
				atomic.AddInt32(&count,1)
			}
		}(tmap)
	}

	time.Sleep(time.Second*3)
	fmt.Println(atomic.LoadInt32(&count))
}

func main()  {
	rand.Seed(time.Now().Unix())
	
	testMap()
	testRWMap()
}