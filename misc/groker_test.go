/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   groker_test.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/14 15:01:52 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/16 18:45:22 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package groker

import (
	// "fmt"
	"sync"
	"testing"
)

// func TestEmit(t *testing.T) {}

func BenchmarkBroadcastList(b *testing.B) {
	for nBench := 0; nBench < b.N; nBench++ {
		var wg sync.WaitGroup
		n := 100000
		bc := NewBroadcast()
		listeners := make([]Channel, n)
		wg.Add(n + 1)
		for i, _ := range listeners {
			listeners[i] = make(Channel)
			bc.AddListener(listeners[i])
			go func(j int, channel Channel) {
				for {
					select {
					case _ = <-channel:
						// Ignore message for perf precision
						wg.Done()
					}
				}
			}(i, listeners[i])
		}
		go func() {
			for {
				select {
				case _ = <-*bc.out:
					wg.Done()
				}
			}
		}()
		*bc.in <- `broad message`
		wg.Wait()
	}
}

func BenchmarkBroadcastBasic(b *testing.B) {
	for nBench := 0; nBench < b.N; nBench++ {
		var wg sync.WaitGroup
		n := 100000
		listeners := make([]Channel, n)
		wg.Add(n)
		for i, _ := range listeners {
			listeners[i] = make(Channel)
			go func(j int, channel Channel) {
				for {
					select {
					case _ = <-channel:
						// Ignore message for perf precision
						wg.Done()
					}
				}
			}(i, listeners[i])
		}
		in := make(Channel)
		go func() {
			for {
				select {
				case msg := <-in:
					go func(m Message, chans []Channel) {
						for _, v := range chans {
							v <- m
						}
					}(msg, listeners)
				}
			}
		}()
		in <- `broad message`
		wg.Wait()
	}
}
