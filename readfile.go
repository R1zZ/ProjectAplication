package main

//menambahkan package yang digunakan
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//melakukan cek jika terjadi kesalahan operasi
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//membaca file
func read(bc string) {
	bs, err := ioutil.ReadFile(bc)
	check(err) //dilakukan cek jika file yang dibaca tidak dikenali
	str := string(bs)
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Isi dari File :" + bc + " adalah")
	fmt.Println("------------------------------------------------------------------")
	fmt.Println(str) //menampilkan isi file

}

//membuat file baru
func new() {
	lokasi()
	fmt.Println("Masukkan Nama File :")
	var nmfile string
	fmt.Scanln(&nmfile)
	fmt.Println("Masukkan isi file / content :")
	var isifile string
	fmt.Scanln(&isifile)
	tostring := []byte(isifile)
	file := ioutil.WriteFile(nmfile, tostring, 0644) //membuat file baru dengan nama dari nmfile dan memasukkan isi file kemudian memberi hak akses ke file dengan permissin 0644
	//user read + write
	//group read
	//other read

	check(file)

}

//hapus file
func hapus() {
	lokasi()
	fmt.Println("Masukkan Nama File :")
	var nmfile string
	fmt.Scanln(&nmfile)
	err := os.Remove(nmfile) //menghapus file
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("File " + nmfile + " berhasil dihapus !!")
	}

}

//list direktrori
func direktori(input string) {
	lokasi()
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Direktori di folder " + input)
	fmt.Println("------------------------------------------------------------------")
	dir, err := os.Open(input)
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

//cek lokasi sekarang
func lokasi() {
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Direktori Aktif di :")
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)
	fmt.Println("------------------------------------------------------------------")
}

func main() {

	fmt.Println("1. Buat Sebuah File")
	fmt.Println("2. Baca Sebuah File")
	fmt.Println("3. Hapus Sebuah File")
	fmt.Println("4. List Direktori")
	fmt.Print("Masukkan Pilihan :")
	var pilihan int
	fmt.Scanln(&pilihan)
	if pilihan == 1 {
		new()
	} else if pilihan == 2 {
		lokasi()
		fmt.Println("Masukkan Nama dan Path File:")
		fmt.Println("example:/home/root/hallo.txt")
		var input string
		fmt.Scanln(&input)
		read(input)
	} else if pilihan == 3 {
		hapus()
	} else if pilihan == 4 {
		fmt.Println("Masukkan Path Folder")
		fmt.Println("example:/home/root/")
		var input string
		fmt.Scanln(&input)
		direktori(input)
	} else {
		fmt.Println("Masukkan Salah")
	}

}
