// WARNING! Autogenerated by goremote, don't touch.

package main

import (
	"net/rpc"
)

type RPC struct {
}

// wrapper for: serverAutoComplete

type ArgsAutoComplete struct {
	Arg0 []byte
	Arg1 string
	Arg2 int
	Arg3 gocodeEnv
}
type ReplyAutoComplete struct {
	Arg0 []candidate
	Arg1 int
}

func (r *RPC) RPCAutoComplete(args *ArgsAutoComplete, reply *ReplyAutoComplete) error {
	reply.Arg0, reply.Arg1 = serverAutoComplete(args.Arg0, args.Arg1, args.Arg2, args.Arg3)
	return nil
}
func clientAutoComplete(cli *rpc.Client, Arg0 []byte, Arg1 string, Arg2 int, Arg3 gocodeEnv) (c []candidate, d int) {
	var args ArgsAutoComplete
	var reply ReplyAutoComplete
	args.Arg0 = Arg0
	args.Arg1 = Arg1
	args.Arg2 = Arg2
	args.Arg3 = Arg3
	err := cli.Call("RPC.RPCAutoComplete", &args, &reply)
	if err != nil {
		panic(err)
	}
	return reply.Arg0, reply.Arg1
}

// wrapper for: serverCursorTypePkg

type ArgsCursorTypePkg struct {
	Arg0 []byte
	Arg1 string
	Arg2 int
}
type ReplyCursorTypePkg struct {
	Arg0, Arg1 string
}

func (r *RPC) RPCCursorTypePkg(args *ArgsCursorTypePkg, reply *ReplyCursorTypePkg) error {
	reply.Arg0, reply.Arg1 = serverCursorTypePkg(args.Arg0, args.Arg1, args.Arg2)
	return nil
}
func clientCursorTypePkg(cli *rpc.Client, Arg0 []byte, Arg1 string, Arg2 int) (typ, pkg string) {
	var args ArgsCursorTypePkg
	var reply ReplyCursorTypePkg
	args.Arg0 = Arg0
	args.Arg1 = Arg1
	args.Arg2 = Arg2
	err := cli.Call("RPC.RPCCursorTypePkg", &args, &reply)
	if err != nil {
		panic(err)
	}
	return reply.Arg0, reply.Arg1
}

// wrapper for: serverClose

type ArgsClose struct {
	Arg0 int
}
type ReplyClose struct {
	Arg0 int
}

func (r *RPC) RPCClose(args *ArgsClose, reply *ReplyClose) error {
	reply.Arg0 = serverClose(args.Arg0)
	return nil
}
func clientClose(cli *rpc.Client, Arg0 int) int {
	var args ArgsClose
	var reply ReplyClose
	args.Arg0 = Arg0
	err := cli.Call("RPC.RPCClose", &args, &reply)
	if err != nil {
		panic(err)
	}
	return reply.Arg0
}

// wrapper for: serverStatus

type ArgsStatus struct {
	Arg0 int
}
type ReplyStatus struct {
	Arg0 string
}

func (r *RPC) RPCStatus(args *ArgsStatus, reply *ReplyStatus) error {
	reply.Arg0 = serverStatus(args.Arg0)
	return nil
}
func clientStatus(cli *rpc.Client, Arg0 int) string {
	var args ArgsStatus
	var reply ReplyStatus
	args.Arg0 = Arg0
	err := cli.Call("RPC.RPCStatus", &args, &reply)
	if err != nil {
		panic(err)
	}
	return reply.Arg0
}

// wrapper for: serverDropCache

type ArgsDropCache struct {
	Arg0 int
}
type ReplyDropCache struct {
	Arg0 int
}

func (r *RPC) RPCDropCache(args *ArgsDropCache, reply *ReplyDropCache) error {
	reply.Arg0 = serverDropCache(args.Arg0)
	return nil
}
func clientDropCache(cli *rpc.Client, Arg0 int) int {
	var args ArgsDropCache
	var reply ReplyDropCache
	args.Arg0 = Arg0
	err := cli.Call("RPC.RPCDropCache", &args, &reply)
	if err != nil {
		panic(err)
	}
	return reply.Arg0
}

// wrapper for: serverSet

type ArgsSet struct {
	Arg0, Arg1 string
}
type ReplySet struct {
	Arg0 string
}

func (r *RPC) RPCSet(args *ArgsSet, reply *ReplySet) error {
	reply.Arg0 = serverSet(args.Arg0, args.Arg1)
	return nil
}
func clientSet(cli *rpc.Client, Arg0, Arg1 string) string {
	var args ArgsSet
	var reply ReplySet
	args.Arg0 = Arg0
	args.Arg1 = Arg1
	err := cli.Call("RPC.RPCSet", &args, &reply)
	if err != nil {
		panic(err)
	}
	return reply.Arg0
}
