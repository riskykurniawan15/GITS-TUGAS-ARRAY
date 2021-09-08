package main

import (
	"fmt"
	"strconv"
	"strings"
)

var increment int = 1
var id_game = map[int]string{}
var title_game = map[string]string{}
var rate_game = map[string]float64{}
var d_rate_game = map[string]string{}

func main(){
	// Risky Kurniawan - ARS University
	goMenu:
	var pilihan, sub string
	fmt.Println("==========================================")
	fmt.Println("           APLIKASI GAME LIST")
	fmt.Println("==========================================")
	fmt.Println("Silahkan pilih menu dibawah ini :")
	fmt.Println("1. Input data game baru")
	fmt.Println("2. Hapus data game berdasarkan id_game")
	fmt.Println("3. Tampilkan seluruh data game beserta jumlah data yang tersimpan dalam list")
	fmt.Println("4. Cari data game berdasarkan nama")
	fmt.Println("5. Menampilkan top 3 game terfavorit")
	fmt.Println("6. Menampilkan seluruh data dengan rating diatas 4.0")
	fmt.Println("0. Keluar\n")	
	fmt.Print("Masukan Pilihan (nomor) : ")
	fmt.Scan(&pilihan)

	ExitMenu:

	if pilihan == "0"{
		fmt.Println("==========================================")
		fmt.Println("            ** Terimakasih **             ")
		fmt.Println("   By Risky Kurniawan - ARS University    ")
		fmt.Println("==========================================")
	}else if pilihan == "1"{
		InputData()
	}else if pilihan == "2"{
		DeleteData()
	}else if pilihan == "3"{
		SelectData()
	}else if pilihan == "4"{
		FindData()
	}else if pilihan == "5"{
		Top3Data()
	}else if pilihan == "6"{
		Rate4()
	}else{
		fmt.Println("==========================================")
		fmt.Println("       PILIHAN MENU TIDAK TERSEDIA        ")
		fmt.Println("==========================================")
	}

	if pilihan != "0" {
		menu_end:
		fmt.Println("Ketik 9 untuk kembali ke menu")
		fmt.Println("Ketik 0 untuk jika kamu ingin keluar")
		fmt.Print("Pilihan anda : ")
		fmt.Scan(&sub)
		if sub == "9" {
			goto goMenu
		}else if sub=="0"{
			pilihan = "0"
			goto ExitMenu
		}else{
			fmt.Println("==========================================")
			fmt.Println("     KEYWORD YANG ANDA MASUKAN SALAH")
			fmt.Println("==========================================")
			goto menu_end
		}
	}
}

func InputData(){
	var title string
	var rate float64
	fmt.Println("==========================================")
	fmt.Println("          INPUT DATA GAME BARU            ")
	fmt.Println("==========================================")
	fmt.Print("Masukan Judul Game              : ")
	fmt.Scan(&title)
	fmt.Print("Masukan Rating Game (0.0 - 5.0) : ")
	fmt.Scan(&rate)
	id_game[increment] = title
	title_game[title] = title
	rate_game[title] = ParseFloat1(rate)
	d_rate_game[title] = DescRate(rate_game[title])
	increment++
	fmt.Println("==========================================")	
}

func DeleteData(){
	var id int
	fmt.Println("==========================================")
	fmt.Println("            HAPUS DATA GAME             ")
	fmt.Println("==========================================")
	fmt.Print("Masukan ID game   : ")
	fmt.Scan(&id)

	if id_game[id] != "" {
		delete(title_game, id_game[id])
		delete(rate_game, id_game[id])
		delete(d_rate_game, id_game[id])
		delete(id_game, id)
		fmt.Println("Data Game Berhasil Dihapus")	
	}else{
		fmt.Println("ID Tidak Valid")	
	}
	fmt.Println("==========================================")	
}

func SelectData(){
	fmt.Println("==========================================")
	fmt.Println("             TAMPIL DATA GAME             ")
	fmt.Println("==========================================")
	var i float64 = 5.0
  for i >= 0 {  
		for id, title := range id_game {
			if i == rate_game[title] {
				fmt.Println("ID:", id, " -> ",title," ( Rate : ", rate_game[title], " | ", d_rate_game[title], ")")
			}
		}
    i-=0.1
    i = ParseFloat1(i)
		if i == 0 {
			break
		}
  }
	if len(id_game) == 0 {
		fmt.Println("         DATA TIDAK TERSEDIA")
	}
	fmt.Println("==========================================")
	fmt.Println("Jumlah Record Data = ", len(id_game))
	fmt.Println("==========================================")
}

func FindData(){
	var find_key string
	fmt.Println("==========================================")
	fmt.Println("             CARI DATA GAME             ")
	fmt.Println("==========================================")
	fmt.Print("Masukan Judul Game   : ")
	fmt.Scan(&find_key)
	find_key = strings.ToLower(find_key)
	var result bool = false
	for id, title := range id_game {
		if find_key == strings.ToLower(title){
			result = true
			fmt.Println("==========================================")
			fmt.Println("ID:", id, " -> ",title," ( Rate : ", rate_game[title], " | ", d_rate_game[title], ")")
		}
	}

	if result == false {
		fmt.Println("          Data Tidak Ditemukan")
	}
	fmt.Println("==========================================")
}

func Top3Data(){
	fmt.Println("==========================================")
	fmt.Println("             TOP 3 DATA GAME             ")
	fmt.Println("==========================================")
	var counter = 0
	var i float64 = 5.0
  for i >= 0 {  
		for id, title := range id_game {
			if i == rate_game[title] {
				fmt.Println("ID:", id, " -> ",title," ( Rate : ", rate_game[title], " | ", d_rate_game[title], ")")
				counter++
			}

			if counter == 3 {
				break
			}
		}
		if counter == 3 {
			break
		}
    i-=0.1
    i = ParseFloat1(i)
		if i == 0 {
			break
		}
  }
	if counter == 0 {
		fmt.Println("         DATA TIDAK TERSEDIA")
	}
	fmt.Println("==========================================")
}

func Rate4(){
	fmt.Println("==========================================")
	fmt.Println("           RATE > 4.0 DATA GAME             ")
	fmt.Println("==========================================")
	var result bool = false
	var i float64 = 5.0
  for i >= 4.1 {  
		for id, title := range id_game {
			if i == rate_game[title] {
				result = true
				fmt.Println("ID:", id, " -> ",title," ( Rate : ", rate_game[title], " | ", d_rate_game[title], ")")
			}
		}
    i-=0.1
    i = ParseFloat1(i)
		if i == 0 {
			break
		}
  }
	if result == false {
		fmt.Println("         DATA TIDAK TERSEDIA")
	}
	fmt.Println("==========================================")
}

func ParseFloat1(number float64) float64{
	rate, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", number), 64)

	if rate < 0 {
		rate = 0
	}else if rate > 5.0 {
		rate = 5.0
	}
	
	return rate
}

func DescRate(number float64) string{
	if number > 4.0 {
		return "good"
	}else if number < 2.0 {
		return "poor"
	}else{
		return "average"
	}
}