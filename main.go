package main

import (
	"fmt"

	go_say_hello "github.com/alakmal12/go-say-hello/v2"
	// BAWAH UNTUK MENGIMPORT FILE LAIN KE FILE SAAT INI GUNAKAN "NAMA MODULE/NAMAPACKAGE ATAU DIRECTORY"
	"github.com/alakmal12/app-say-hello/nama2"
)

/**
untuk memanggil library dari orang lain atau library yg udah kita bikin
tinggal tulis aja go get namamoduleya

kita udah bikin module untuk di pelajaran 64 yaitu go-say-hello
tinggal panggil aja go get github.com/alakmal12/go-say-hello

nanti decara otomatis golang akan mendownoad librarynya yg versi terbaru
dan menambahkan sesuatu di file go.mod


untuk memanggil method atau function tertentu kita tinggal gunakan
namamodulenya.namamethod atau nama function atau nama struct
 pada saat mengimport nya
tidak perlu kita ngutak ngatik file go modulesnya ok

kalau suatu pakage dan library sudah diimport itu harus dipakai, walaupun yg dipakai cuma satu method atu satu function saja ok



*/

func main(){

	// panggil method sayHello yg udah dibuat dri libraary go modules yg tadi
	fmt.Println(go_say_hello.SayHello("akmal","kamu")) // HASILNYA hello world akmal
	fmt.Println(go_say_hello.SayHello2()) // hasilnya  Halo ini dari m=function SayHello2

	// BAWAH UNTUK MENGIMPORT function atau methid 
	// YG ADA DI STU PROYEK TINGGAL MANGGILNYA namapackage.namamethod atau struct 
	fmt.Println(nama2.Nama2("AKU"))

}