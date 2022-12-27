/*
 * @Author: dueb dueb@channelsoft.com
 * @Date: 2022-12-27 09:55:45
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2022-12-27 09:55:48
 * @FilePath: /pkg/utils/autoinc.go
 * @Description:
 *
 * Copyright (c) 2022 by dueb dueb@channelsoft.com, All Rights Reserved.
 */
package utils

const INT_MAX = int(^uint(0) >> 1)

type AutoInc struct {
	start, step int
	queue       chan int
	running     bool
}

func NewAutoInc(start, step int) (ai *AutoInc) {
	ai = &AutoInc{
		start:   start,
		step:    step,
		running: true,
		queue:   make(chan int, 40),
	}
	go ai.process()
	return
}

func (ai *AutoInc) process() {
	defer func() { recover() }()
	for i := ai.start; ai.running; i = i + ai.step {
		if i == INT_MAX {
			i = ai.start
		}
		ai.queue <- i
	}
}

func (ai *AutoInc) Id() int {
	return <-ai.queue
}

func (ai *AutoInc) Close() {
	ai.running = false
	close(ai.queue)
}
