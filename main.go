package main

import (
	"fmt"

	go_say_hello "github.com/alakmal12/go-say-hello"
)

/**
untuk memanggil library dari orang lain atau library yg udah kita bikin
tinggal tulis aja go get namamoduleya

kita udah bikin module untuk di pelajaran 64 yaitu go-say-hello
tinggal panggil aja go get github.com/alakmal12/go-say-hello

nanti decara otomatis golang akan mendownoad librarynya yg versi terbary
dan menambahkan sesuatu di file go.mod


untuk memanggil method atau function tertentu kita tinggal gunakan

namamodulenya.namamethod atau nama function atau nama struct


*/

func main(){

	// panggil method sayHello yg udah dibuat dri libraary go modules yg tadi
	fmt.Println(go_say_hello.SayHello("akmal"))

}