package singleton

import (
	"sync"
	"testing"
)

func TestGetSingleton(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(5)
	start := make(chan struct{})
	arr := make([]*Singleton, 5)
	for i := 0; i < 5; i++ {
		// 同时起5个goroutine
		go func(index int) {
			<-start                     // 全部阻塞在此
			singleten := GetSingleton() // 获取单例对象
			arr[index] = singleten
			wg.Done() // 将每个goroutine获得的单例对象存储到数组
		}(i)
	}
	close(start) // 关闭通道的同时，5个goroutine同时执行
	wg.Wait()    // 阻塞等待所有goroutine执行完成

	for i := 0; i < 4; i++ {
		if arr[i] != arr[i+1] {
			t.Fatalf("同时有多个实例")
		}
	}
}
