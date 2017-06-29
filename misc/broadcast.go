/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   broadcast.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/14 15:44:14 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/17 13:44:58 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package groker

import ()

/*
	Message type is abstracted behind this type
*/
type Message interface{}

/*
	An alias for Message Channel
*/
type Channel chan Message

/*
	A Node is the minimal dispatching element in the Broadcast chain
|							_______							|
|	-------receiver------->| Node |------replicator-------->|-------receiver------->...
|								|							|
|								|							|
|								|---> emmiter				|
*/
type Node struct {
	receiver   Channel
	emmiter    Channel
	replicator Channel
}

/*
	Enable listening and replication process inside a Node
*/
func (n Node) Listen() {
	for {
		select {
		case msg := <-n.receiver:
			n.replicator <- msg
			n.emmiter <- msg
		}
	}
}

/*
	Get a new disabled Node, use Listen to enable it
*/
func NewNode(emmiter Channel) Node {
	return Node{
		receiver:   make(Channel),
		emmiter:    emmiter,
		replicator: make(Channel),
	}
}

/*
	A broadcast may be seen as a Node list
	Send messages in `*b.in <- msg` and retrieve them in `resp <- *b.out`
	Response is transferred once the message has been transferred to all nodes
	But it's not waiting for any response from Nodes !
	*in Channel = receiver of first node.
	*out Channel = replicator of last node
*/
type Broadcast struct {
	in  *Channel
	out *Channel
}

func NewBroadcast() Broadcast {
	c := make(Channel)
	return Broadcast{
		in:  &c,
		out: &c,
	}
}

/*
	Not thread safe, don't run it conccurently
*/
func (b *Broadcast) AddListener(emmiter Channel) {
	// Create a new tail node
	n := NewNode(emmiter)
	// Plug new tail listener to previous tail replicator
	go func(out Channel, in Channel) {
		for {
			select {
			case msg := <-out:
				in <- msg
			}
		}
	}(*b.out, n.receiver)
	// Make new tail active
	go n.Listen()
	// Broadcast tail is now new tail
	b.out = &n.replicator
}
