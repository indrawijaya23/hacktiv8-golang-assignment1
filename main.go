package main

import (
	"fmt"
	"os"
)

type Student struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	var students = make(map[string]Student, 0)
	students["1"] = Student{nama: "Indra Wijaya", alamat: "Jakarta", pekerjaan: "Software Engineer", alasan: "Golang Keren !"}
	students["2"] = Student{nama: "Ricky Romansyah", alamat: "Bogor", pekerjaan: "UI/UX Designer", alasan: "Hebat banget Golang"}
	students["3"] = Student{nama: "Handi", alamat: "Bandung", pekerjaan: "Developer", alasan: "Menambah pengetahuan"}
	students["4"] = Student{nama: "Muhammad Syaibani Al Hakim", alamat: "Yogyakarta", pekerjaan: "Data Scientist", alasan: "Mencoba bahasa baru"}
	students["5"] = Student{nama: "Muhammad Faturrohman", alamat: "Bali", pekerjaan: "Cloud Engineer", alasan: "Memahamai bahasa Golang"}
	students["6"] = Student{nama: "Khairul", alamat: "Tangerang", pekerjaan: "Frontend Developer", alasan: "Supaya bisa jadi fullstack"}
	students["7"] = Student{nama: "Felix Ivander Ganumba", alamat: "Solo", pekerjaan: "Backend Developer", alasan: "Menambah opsi pemrograman backend"}
	students["8"] = Student{nama: "Tata Redha Al Fath", alamat: "Palembang", pekerjaan: "Web Developer", alasan: "Penasaran cara kerja Golang"}
	students["9"] = Student{nama: "Altaf Syahrastani", alamat: "Jambi", pekerjaan: "Security Analyst", alasan: "Mau coba Golang"}
	students["10"] = Student{nama: "Pebrido Sihaloho", alamat: "Padang", pekerjaan: "Database Administrator", alasan: "Belajar bahasa pemrograman baru"}

	printStudentDetail(students, os.Args[1])
}

func printStudentDetail(students map[string]Student, index string) {
	var s = students[index]
	fmt.Println("Nama:", s.nama)
	fmt.Println("Alamat:", s.alamat)
	fmt.Println("Pekerjaan:", s.pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", s.alasan)
}
