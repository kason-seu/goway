package concurrency

import (
	"fmt"
	"sync"
	"time"
)

/**
Channels are the pipes that connect concurrent
goroutines. You can send values into channels
from one goroutine and receive those values
into another goroutine

By default sends and receives block
 until both the sender and receiver are ready.
 This property allowed us to wait at the end of
  our program for the "ping" message without
   having to use any other synchronization.
**
*/

// 没有缓存的channel 当发送一个消息时，必须要有一个channel收到这个消息，不然就会阻塞

func channel_unbuffered() {
	// Create a new channel with `make(chan val-type)`.
	// Channels are typed by the values they convey.
	messages := make(chan string)

	// _Send_ a value into a channel using the `channel <-`
	// syntax. Here we send `"ping"`  to the `messages`
	// channel we made above, from a new goroutine.
	go func() { messages <- "ping" }()

	// The `<-channel` syntax _receives_ a value from the
	// channel. Here we'll receive the `"ping"` message
	// we sent above and print it out.
	//msg := <-messages
	//fmt.Println(msg)

	go func() {
		v := <-messages
		fmt.Println(v)
	}()

	time.Sleep(1000)
}

func channel_unbuffered_err() {

	msg := make(chan string)

	msg <- "hello"

	msg <- "ping"

	fmt.Println("main finished")
	time.Sleep(1000)
}
func channel_unbuffered_err2() {

	msg := make(chan string)

	msg <- "hello"

	fmt.Println("main finished")
	time.Sleep(1000)
}

func channel_unbuffered_err3() {

	msg := make(chan string)

	go func() {
		msg <- "hello"
		//msg <- "hello"
		fmt.Println("gofunc println")
	}()

	time.Sleep(10000)
	fmt.Println("main finished")

}

func channe_buffered() {

	/**
	By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values
		**/

	// Here we `make` a channel of strings buffering up to
	// 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these
	// values into the channel without a corresponding
	// concurrent receive.
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func channe_buffered_noerr() {

	/**
	By default channels are unbuffered, meaning that they will only accept sends (chan <-)
	if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered
	channels accept a limited number of values without a corresponding receiver for those values
		**/

	// Here we `make` a channel of strings buffering up to
	// 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these
	// values into the channel without a corresponding
	// concurrent receive.
	messages <- "buffered"
	messages <- "channel"
	//messages <- "channe2"

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we're done.
	done <- true
}

/*
*
We can use channels to synchronize execution across goroutines. Here’s an example of using a blocking receive to wait for a goroutine to finish. When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.

*
*/
func channel_synchronizied() {
	// Start a worker goroutine, giving it the channel to
	// notify on.
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the
	// worker on the channel.
	<-done
}

/**


When using channels as function parameters, you can specify if a channel is meant to only send or receive values. This specificity increases the type-safety of the program
**/
// This `ping` function only accepts a channel for sending
// values. It would be a compile-time error to try to
// receive on this channel.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// The `pong` function accepts one channel for receives
// (`pings`) and a second for sends (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func channel_directions() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

/**

select
Go’s select lets you wait on multiple channel operations.
Combining goroutines
and channels with select is a powerful feature of Go

**/

func select_fn() {
	// For our example we'll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {

		select {
		case <-time.After(2 * time.Second):
			println("测试select")
		}
	}()

	// Each channel will receive a value after some amount
	// of time, to simulate e.g. blocking RPC operations
	// executing in concurrent goroutines.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// We'll use `select` to await both of these values
	// simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

/*
*
timeout
Timeouts are important for programs that connect to external

	resources or that otherwise need to bound execution time.
	 Implementing timeouts in Go is easy and elegant thanks to
	 channels and select

*
*/
func timeout_fn() {
	// For our example, suppose we're executing an external
	// call that returns its result on a channel `c1`
	// after 2s. Note that the channel is buffered, so the
	// send in the goroutine is nonblocking. This is a
	// common pattern to prevent goroutine leaks in case the
	// channel is never read.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Here's the `select` implementing a timeout.
	// `res := <-c1` awaits the result and `<-time.After`
	// awaits a value to be sent after the timeout of
	// 1s. Since `select` proceeds with the first
	// receive that's ready, we'll take the timeout case
	// if the operation takes more than the allowed 1s.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// If we allow a longer timeout of 3s, then the receive
	// from `c2` will succeed and we'll print the result.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

/**

non-blocking channel operations

Basic sends and receives on channels are blocking.
However, we can use select with a default clause to implement
non-blocking sends, receives, and even non-blocking multi-way
selects
**/

func nonBlocking() {
	messages := make(chan string)
	signals := make(chan bool)

	// Here's a non-blocking receive. If a value is
	// available on `messages` then `select` will take
	// the `<-messages` `case` with that value. If not
	// it will immediately take the `default` case.
	go func() {
		messages <- "hello"
	}()

	time.Sleep(1 * time.Second)
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// A non-blocking send works similarly. Here `msg`
	// cannot be sent to the `messages` channel, because
	// the channel has no buffer and there is no receiver.
	// Therefore the `default` case is selected.

	go func() {
		for {
			_ = <-messages
		}
	}()
	time.Sleep(1 * time.Second)

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// We can use multiple `case`s above the `default`
	// clause to implement a multi-way non-blocking
	// select. Here we attempt non-blocking receives
	// on both `messages` and `signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func nonBlocking2() {
	messages := make(chan string)
	//signals := make(chan bool)

	messages <- "helo"

	<-messages
}

/**
closechannel


Closing a channel indicates that no more values
 will be sent on it. This can be useful to
communicate completion to the channel’s receivers.
**/

func close_channel() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			job, state := <-jobs
			if state {
				fmt.Println("receive job", job)
			} else {
				fmt.Println("receive all jobs")
				done <- true
				return
			}

		}
	}()

	for j := 0; j < 10; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}

	close(jobs)
	// close 之后不能再发送
	//jobs <- 100

	<-done
	_, state := <-jobs
	fmt.Println("no more jobs, state = ", state)

}

/**
range channels

In a previous example we saw how for
 and range provide iteration over basic data structures. We can also
use this syntax to iterate over values received from a channel
**/

func range_channel() {

	queue := make(chan string, 2)

	queue <- "hell0"
	queue <- "sdhjdd"

	close(queue)

	go func() {
		for v := range queue {
			fmt.Println(v)
		}
	}()

	time.Sleep(time.Second)

}

/**
timers
We often want to execute Go code at some point in the future,
or repeatedly at some interval. Go’s built-in timer and ticker
features make both of these tasks easy. We’ll look first at timers
and then at tickers.
**/

func times_fire() {
	// Timers represent a single event in the future. You
	// tell the timer how long you want to wait, and it
	// provides a channel that will be notified at that
	// time. This timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// The `<-timer1.C` blocks on the timer's channel `C`
	// until it sends a value indicating that the timer
	// fired.

	// go func() {
	// 	for {
	// 		select {
	// 		case <-timer1.C:
	// 			fmt.Println("Timer 1 fired")
	// 			return
	// 			// default:
	// 			// 	fmt.Println("default")
	// 		}
	// 	}
	// }()
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// If you just wanted to wait, you could have used
	// `time.Sleep`. One reason a timer may be useful is
	// that you can cancel the timer before it fires.
	// Here's an example of that.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	//time.Sleep(1* time.Second)
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Give the `timer2` enough time to fire, if it ever
	// was going to, to show it is in fact stopped.
	time.Sleep(2 * time.Second)
}

/**
定时间隔多久repeat定时

Timers are for when you want to do something once in the future -
tickers are for when you want to do something repeatedly at regular
intervals. Here’s an example of a ticker that ticks periodically
until we stop it

**/

func ticker_fn() {
	ticker := time.NewTicker(2 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Ding")

			case <-done:
				fmt.Println("overfinieshed")
				return
			}
		}
	}()

	time.Sleep(10 * time.Second)

	ticker.Stop()

	done <- true
	fmt.Print("main finished")
}

/**

use goroutines and channel to implement worker pool

In this example we’ll look at how to implement a worker
pool using goroutines and channels
*/
// 工作执行器
func worker_fn(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {
		fmt.Println("woker id", id, "started handle job", j)
		// 模拟执行job
		time.Sleep(time.Second)
		fmt.Println("worker id", id, "job ", j, "finished")
		results <- j
	}

}
func worker_pool_fn() {
	// 模拟三个worker，实现5个任务的并行处理

	jobsCnt := 5
	wokerCnt := 3
	jobs := make(chan int, jobsCnt)
	results := make(chan int, jobsCnt)
	for j := 0; j < wokerCnt; j++ {
		go worker_fn(j, jobs, results)
	}

	for j := 0; j < jobsCnt; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 0; i < jobsCnt; i++ {
		<-results
	}

}

func worker_pool_goweb() {
	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker_fn(w, jobs, results)
	}

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work.
	// This also ensures that the worker goroutines have
	// finished. An alternative way to wait for multiple
	// goroutines is to use a [WaitGroup](waitgroups).
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

/*
*
waitgroup

*
*/
func wg_worker_fn(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {

	for j := range jobs {
		defer wg.Done()
		fmt.Println("woker id", id, "started handle job", j)
		// 模拟执行job
		time.Sleep(time.Second)
		fmt.Println("worker id", id, "job ", j, "finished")
		results <- j
	}

}
func waitgroups_pool() {

	var wg sync.WaitGroup
	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go wg_worker_fn(w, jobs, results, &wg)
	}

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
}
