package concurrency

import (
	"testing"
)

func Test_channel_unbuffered(t *testing.T) {
	channel_unbuffered()
}

func Test_channel_unbuffered_err(t *testing.T) {
	channel_unbuffered_err()
}

func Test_channe_buffered(t *testing.T) {
	channe_buffered()
}

func Test_channe_buffered_noerr(t *testing.T) {
	channe_buffered_noerr()
}

func Test_channel_unbuffered_err2(t *testing.T) {
	channel_unbuffered_err2()
}

func Test_channel_synchronizied(t *testing.T) {
	channel_synchronizied()
}

func Test_channel_directions(t *testing.T) {
	channel_directions()
}

func Test_select_fn(t *testing.T) {
	select_fn()
}

func Test_channel_unbuffered_err3(t *testing.T) {
	channel_unbuffered_err3()
}

func Test_timeout_fn(t *testing.T) {
	timeout_fn()
}

func Test_nonBlocking(t *testing.T) {
	nonBlocking()
}

func Test_nonBlocking2(t *testing.T) {
	nonBlocking2()
}

func Test_close_channel(t *testing.T) {
	close_channel()
}

func Test_range_channel(t *testing.T) {
	range_channel()
}

func Test_times_fire(t *testing.T) {
	times_fire()
}

func Test_ticker_fn(t *testing.T) {
	ticker_fn()
}

func Test_worker_pool_fn(t *testing.T) {
	worker_pool_fn()
}

func Test_waitgroups_pool(t *testing.T) {
	waitgroups_pool()
}
