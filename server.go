package main

import (

	"fmt"
	"net"
)

func server() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0), //1.85.218.81
		Port: 9999,
	})
	if err != nil {
		fmt.Println("监听失败!", err)
		return
	}
	defer conn.Close()
	for {
		// 读取数据
		data := make([]byte, 4096)
		read, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println("读取数据失败!", err)
			continue
		}
		fmt.Println(data[:read])
		fmt.Println("cmd:",string(data[:2]))
		fmt.Println("channelid:",string(data[2:22]))
		fmt.Println("user:",string(data[22:42]))
		fmt.Println("filename:",string(data[42:78]))
		fmt.Println("flag:",string(data[78:96]))
		fmt.Println("token:",string(data[96:296]))
		fmt.Println("data:",data[296:read])
		_,err=conn.WriteToUDP(data,remoteAddr)
		if err != nil {
			fmt.Println("写入数据失败!", err)
			continue
		}
	}
}

