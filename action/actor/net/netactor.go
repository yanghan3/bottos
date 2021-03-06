// Copyright 2017~2022 The Bottos Authors
// This file is part of the Bottos Chain library.
// Created by Rocket Core Team of Bottos.

//This program is free software: you can distribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
// along with bottos.  If not, see <http://www.gnu.org/licenses/>.

/*
 * file description:  net actor
 * @Author:
 * @Date:   2017-12-06
 * @Last Modified by:
 * @Last Modified time:
 */

package netactor

import (
	"fmt"
	//"encoding/json"
	"github.com/bottos-project/bottos/action/message"
	//"github.com/bottos-project/bottos/common/types"
	"github.com/AsynkronIT/protoactor-go/actor"
	p2pserv "github.com/bottos-project/bottos/p2p"
)

//NetActorPid net actor pid
var NetActorPid *actor.PID
var p2p *p2pserv.P2PServer

//NetActor is net actor props
type NetActor struct {
	props *actor.Props
}

//ContructNetActor new a net actor
func ContructNetActor() *NetActor {
	return &NetActor{}
}

//NewNetActor spawn a named actor
func NewNetActor() *actor.PID {

	p2p = p2pserv.NewServ()
	go p2p.Start()

	props := actor.FromProducer(func() actor.Actor { return ContructNetActor() })

	var err error
	NetActorPid, err = actor.SpawnNamed(props, "NetActor")

	if err != nil {
		panic(fmt.Errorf("NetActor SpawnNamed error: %v", err))
	} else {
		return NetActorPid
	}
}

//main loop
func (NetActor *NetActor) handleSystemMsg(context actor.Context) {

	switch context.Message().(type) {
	case *actor.Started:
		fmt.Printf("NetActor received started msg")
	case *actor.Stopping:
		fmt.Printf("NetActor received stopping msg")
	case *actor.Restart:
		fmt.Printf("NetActor received restart msg")
	case *actor.Restarting:
		fmt.Printf("NetActor received restarting msg")
	}
}

//Receive process msg
func (NetActor *NetActor) Receive(context actor.Context) {

	NetActor.handleSystemMsg(context)

	switch msg := context.Message().(type) {
	//case types.Transaction:
	case message.NotifyTrx:
		go p2p.BroadCast(msg.Trx, p2pserv.TRANSACTION)
	//case types.Block:
	case message.NotifyBlock:
		go p2p.BroadCast(msg.Block, p2pserv.BLOCK)
	}
}
