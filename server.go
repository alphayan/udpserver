package main

import (
	"fmt"
	"net"
)

var remotes = make(map[string]*net.UDPAddr)

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
		//fmt.Println("cmd:", string(data[:2]))
		//fmt.Println("channelid:", string(data[2:22]))
		//fmt.Println("user:", string(data[22:42]))
		//fmt.Println("filename:", string(data[42:78]))
		//fmt.Println("flag:", string(data[78:86]))
		//fmt.Println("token:", string(data[86:286]))
		//fmt.Println("data:", data[296:read])
		if _, ok := remotes[remoteAddr.String()]; ok {

		} else {
			remotes[remoteAddr.String()] = remoteAddr
		}
		fmt.Println("所有的客户端信息：",remotes)
		for _, v := range remotes {
			_, err = conn.WriteToUDP(data, v)
			if err != nil {
				fmt.Println("写入数据失败!", err)
				continue
			}
		}

	}
}
