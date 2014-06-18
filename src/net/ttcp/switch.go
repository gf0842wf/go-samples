package main

// 转接handler里network protocol 和 internal IPC 等

import (
	"net/ttcp/types"
)

// network protocol
func SwitchNetProto(sess *types.Session) (result []byte, err error){

}

// internal IPC
func SwitchIPCProto(sess *types.Session) (result []byte, err error){
	
}

// 定时消息
func SwitchTMProto(sess *types.Session) (result []byte, err error){
	
}