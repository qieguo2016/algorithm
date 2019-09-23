package base

import (
	"fmt"
	"runtime"

	"sync"
	"testing"
	"time"
)

type Single struct {
	Value int32
}

var i int32
var s *Single

var _ = new(sync.Once)
var once = new(Once)

// 获取全局单例
func GetGlobalSingle() *Single {
	once.Do(func() {
		s = &Single{
			Value: i,
		}
		i++ // 方便检查是否单例
	})
	return s
}

func TestGlobalSingle(t *testing.T) {
	n := runtime.NumCPU()
	fmt.Println("cup num=", n)
	runtime.GOMAXPROCS(n)
	fmt.Println("start TestGlobalSingle")
	c := make(chan int32, 5)
	for i := 0; i < 20; i++ {
		go func() {
			s := GetGlobalSingle()
			fmt.Println("send value", s.Value)
			c <- s.Value
			fmt.Println("finish send ", s.Value)
		}()
	}

	/* NOTE: 如果不另外开一个goroutine来消费channel的话，程序会panic，报错是所有goroutine都休眠了
	具体原因涉及到channel的for range和goroutine的调度机制。
	1. for range遍历channel如何判断已经消费到了最后一个？实际上for range等同于 for{v,ok=<-c if!ok{break}}
	其中ok用来标识channel是否关闭，如果没有关闭channel，那么消费的g就会一直堵塞等待消息，也就是asleep状态
	2. 如果没有新开一个g，那么就是主线程一直陷入堵塞状态；如果新开了一个g的话，其实这个g就泄露了，因为他一直在等待channel的消息；
	而channel中还有等待的g，那这个channel也不会被垃圾回收掉
	*/
	go func() {
		for v := range c {
			fmt.Println("receive value ", v)
			i++
		}
	}()
	fmt.Println("end")
	time.Sleep(3 * time.Second)
}

// 新开两个子线程，分别输出1,3,5,7,9...和2,4,6,8,10...，主线程接受子线程的值，输出1,2,3,4,5...
func TestAlternateOutput(t *testing.T) {
	n := runtime.NumCPU()
	fmt.Println("cpu num=", n)
	runtime.GOMAXPROCS(n)

	fmt.Println("====== start ======")
	fmt.Println("stage 0, go num=", runtime.NumGoroutine()) // 默认两个go

	// AlternateOutputViaChannel()
	// AlternateOutputViaAtomic()
	AlternateOutputViaCond()

	fmt.Println("====== end ======")

}
