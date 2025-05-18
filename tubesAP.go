package main
import (
	"fmt"
)
type account struct{
	id int
	username string
	password int
	moneyFiat float64
	saldoVirtual wallet
}
type transaksi struct{
	userId int
	tipeTransaksi string
	quantityFiat int
	quantityKripto int
	tanggal int
	bulan int
	tahun int
}
type wallet struct{
	id int
	name string
	pasokan int
}
const NMAX int = 100
type arrWalet = [NMAX]wallet
type arrAkun = [NMAX]account
type arrTransaksi = [NMAX]transaksi
var DataTransaksi arrTransaksi
var DataDompet arrWalet
var DataAkun arrAkun

func main(){
	var akun arrAkun
	var dompetVirtual arrWalet
	var catatanTransaksi arrTransaksi
	var aplikasiAktif bool = true
	var tipeJualBeli string
	var tanggal, bulan, tahun, id, pass int
	for aplikasiAktif {
		var accountOption int
		var isLogin bool = false
		WelcomeLogin()
		fmt.Scan(&accountOption)
		if accountOption == 1{
			isLogin = login(&akun, accountOption)
		}else if accountOption == 2{
			isLogin = regis(&akun, accountOption)
		}else if accountOption == 3{
			fmt.Println("Terimakasih Sudah Menggunakan Aplikasi. Sampai Jumpa")
			aplikasiAktif = false
		}else{
			fmt.Print("Pilihan Tidak Valid. Silahkan Pilih 1, 2, atau 3.")
		}
		for isLogin{
			var dashboardOption int
			DashboardOption()
			fmt.Scan(&dashboardOption)
			switch dashboardOption{
				case 1:
					fmt.Print("Pilih Jual atau beli: ")
					fmt.Scan(&tipeJualBeli)
					jualBeli(tipeJualBeli, &catatanTransaksi, &dompetVirtual, &akun)
				case 2: 
					fmt.Print("Masukan Id: ")
					fmt.Scan(&id)
					fmt.Print("Masukan Tanggal, Bulan, dan Tahun: ")
					fmt.Scan(&tanggal, &bulan, &tahun)
					mencariDataTranksasi(catatanTransaksi, akun, tanggal, bulan, tahun, id)
				case 3:
					rangkingKripto(akun)
				case 4:
					fmt.Print("Masukan Id: ")
					fmt.Scan(&id)
					fmt.Print("Masukan Password: ")
					fmt.Scan(&pass)
					isLogin = hapusAccount(&akun, id, pass)
				case 5:
					fmt.Print("Masukan Id: ")
					fmt.Scan(&id)
					fmt.Print("Masukan Password: ")
					fmt.Scan(&pass)
					isLogin = logOut(akun, id, pass)
				case 6:
					aplikasiAktif = false
					isLogin = false
			}
		}
	}
}
func WelcomeLogin() {
	fmt.Println("=========================================================")
	fmt.Println("         SELAMAT DATANG DI APLIKASI JUAL BELI KRIPTO     ")
	fmt.Println("                 INVESTASI ASET DIGITAL                   ")
	fmt.Println("=========================================================")
	fmt.Println("\n🚀 Pilihan Menu:")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("1. Login")
	fmt.Println("2. Daftar Akun")
	fmt.Println("3. Keluar Aplikasi")
	fmt.Println("---------------------------------------------------------")
	fmt.Print("\nMasukkan Pilihanmu 💻: ")
}
func DashboardOption() {
	fmt.Println("\n\n-----------------Welcome To Dashboard------------------")
	fmt.Println()
	fmt.Println("=========================================================")
	fmt.Println("1. Melakukan Transaksi Kripto")
	fmt.Println("2. Search Data Transaksi Kripto")
	fmt.Println("3. Urutkan Top Global Kripto")
	fmt.Println("4. Hapus Akun")
	fmt.Println("5. LogOut")
	fmt.Println("6. Keluar Aplikasi")
	fmt.Println("=========================================================")
	fmt.Print("Masukkan Pilihan: ")
}
func login(akun *arrAkun, pilihAkun int)bool{
	var i, id int
	var pass int
	var benarLogin bool
	if pilihAkun == 1{
		fmt.Print("Masukan id: ")
		fmt.Scan(&id)
		fmt.Print("Masukan Password: ")
		fmt.Scan(&pass)
		i = 0
		benarLogin = false
		for i < NMAX{
			if id == akun[i].id && pass == akun[i].password && benarLogin == false{
				benarLogin = true
			}else if akun[i].id == 0{
				i = NMAX
			}
			i++
		}
		if benarLogin{
			fmt.Println("Login berhasil. ")
		}else{
			fmt.Println("ID atau password salah. ")
		}
	}
	return benarLogin

}
func regis(akun *arrAkun, pilihAkun int)bool{
	var buatAkun, idUnik bool
	var id, pass, i, uangFiat, n int
	var usrn string
	n = 0
	if pilihAkun == 2{
		buatAkun = false
		for buatAkun == false {
			fmt.Print("Buat Id (dalam integer kombinasi 0 - 9): ")
			fmt.Scan(&id)
			fmt.Print("Masukan Username: ")
			fmt.Scan(&usrn)
			fmt.Print("Buat Password (dalam integer kombinasi 0 - 9): ")
			fmt.Scan(&pass)
			i = 0
			idUnik = true
			for i < NMAX{
				if id == akun[i].id{
					idUnik = false
				}else if akun[i].id == 0{
					idUnik = true
					n = i
					i = NMAX
				}else{
					i++
				}
			}
			if idUnik{
				fmt.Print("Masukan nilai uang: ")
				fmt.Scan(&uangFiat)
				akun[n].id = id
				akun[n].username = usrn
				akun[n].password = pass
				akun[n].moneyFiat = float64(uangFiat)
				akun[n].saldoVirtual.id = id
				akun[n].saldoVirtual.name = usrn
				akun[n].saldoVirtual.pasokan = 0
				n = n+1
				buatAkun  = true
				fmt.Println("Akun berhasil dibuat")
			}else{
				fmt.Println("ID sudah digunakan. Silahkan cobalagi")
			}
		}	
	}
	return buatAkun
}
func jualBeli(tipeJualBeli string, dataJualBeli *arrTransaksi, wallet *arrWalet, akun *arrAkun){
	var pasokanKripto, jumlahKriptoYangDijual, jumlahKriptoYangDibeli, nilaiBeliKripto int
	var i, tanggal, bulan, tahun, idxWallet, id, pw, n int
	var nilaiJualKripto float64
	var valid bool
	valid = false
	i = 0
	fmt.Println("Masukan Id: ")
	fmt.Scan(&id)
	fmt.Println("Masukan Password: ")
	fmt.Scan(&pw)
	i = 0
	for i < NMAX{
		if akun[i].id == id && akun[i].password == pw{
			pasokanKripto += wallet[i].pasokan
			idxWallet = i
			i = NMAX
		}else if akun[i].id == 0{
			i = NMAX
		}else{
			i++
		}
	}
	for valid == false{
		i = 0
		if tipeJualBeli == "jual"{	
			fmt.Print("Masukan Jumlah Kripto Yang Ingin Dijual: ")
			fmt.Scan(&jumlahKriptoYangDijual)
			if jumlahKriptoYangDijual <= pasokanKripto{
				nilaiJualKripto = float64(jumlahKriptoYangDijual)*21000
				akun[idxWallet].moneyFiat = nilaiJualKripto
				fmt.Print("Tulis Tanggal, Bulan, dan Tahun: ")
				fmt.Scan(&tanggal, &bulan, &tahun)
				for i < NMAX{
					if dataJualBeli[i].userId == 0{
						dataJualBeli[i].userId = id
						dataJualBeli[i].tipeTransaksi = tipeJualBeli
						dataJualBeli[i].quantityFiat = int(nilaiJualKripto) 
						dataJualBeli[i].quantityKripto = -jumlahKriptoYangDijual
						dataJualBeli[i].tanggal = tanggal
						dataJualBeli[i].bulan = bulan
						dataJualBeli[i].tahun = tahun
						n = n+1
						i = NMAX
					}else{
						i++
					}
				}
				wallet[idxWallet].pasokan = pasokanKripto-jumlahKriptoYangDijual
				akun[idxWallet].saldoVirtual.pasokan = wallet[idxWallet].pasokan
				valid = true
			}else{
				fmt.Print("Jumlah Kripto Yang Dimiliki Tidak Tercukupi")
				valid = true
			}
		}else if tipeJualBeli == "beli"{
			fmt.Print("Masukan Jumlah Kripto Yang Ingin DIbeli: ")
			fmt.Scan(&jumlahKriptoYangDibeli)
			nilaiBeliKripto = jumlahKriptoYangDibeli*21000
			if akun[idxWallet].moneyFiat >= float64(nilaiBeliKripto){
				akun[idxWallet].moneyFiat -=float64(nilaiBeliKripto)
				fmt.Print("Tulis Tanggal, Bulan, dan Tahun: ")
				fmt.Scan(&tanggal, &bulan, &tahun)
				for i < NMAX{
					if dataJualBeli[i].userId == 0{
						dataJualBeli[i].userId = id
						dataJualBeli[i].tipeTransaksi = tipeJualBeli
						dataJualBeli[i].quantityFiat = -nilaiBeliKripto 
						dataJualBeli[i].quantityKripto = jumlahKriptoYangDibeli
						dataJualBeli[i].tanggal = tanggal
						dataJualBeli[i].bulan = bulan
						dataJualBeli[i].tahun = tahun
						n = n+1
						i = NMAX
					}else{
						i++
					}
				}
				wallet[idxWallet].pasokan = pasokanKripto+jumlahKriptoYangDibeli
				akun[idxWallet].saldoVirtual.pasokan = wallet[idxWallet].pasokan
				valid = true
			}else{
				fmt.Println("Uang Tidak Cukup")
				valid = true
			}		
		}else{
			fmt.Println("Tipe transaksi tidak dikenali (harus 'jual' atau 'beli').")
			valid = false
		}
	}
	
}
func mencariDataTranksasi(dataJualBeli arrTransaksi, akun arrAkun, tanggal, bulan, tahun, id int) {
	var i, temp int
	var ada bool
	ada = false
	for i = 0; i < NMAX; i++ {
		if dataJualBeli[i].tanggal == tanggal && dataJualBeli[i].bulan == bulan && dataJualBeli[i].tahun == tahun && dataJualBeli[i].userId == id {
			temp = i
			ada = true
			i = NMAX
		} else if dataJualBeli[i].userId == 0 {
			i = NMAX
		}
	}
	if ada == true {
		fmt.Println("\nAkun Yang Tersedia:")
		fmt.Printf("%-5s %-10s %-10s %-10s\n", "ID", "Username", "Saldo", "Pasokan")
		fmt.Printf("%-5d %-10s Rp%-10.2f %-10d\n", akun[temp].id, akun[temp].username, akun[temp].moneyFiat, akun[temp].saldoVirtual.pasokan)
	} else {
		fmt.Print("Tidak Ada Riwayat Transaksi")
	}
}
func rangkingKripto(A arrAkun){
	var i, idx, j int
	var temp account
	for i = 0; i < NMAX; i++{
		if A[i].id != 0{
			idx = i
			for j = i+1; j < NMAX; j++{
				if A[j].id != 0 && A[i].saldoVirtual.pasokan < A[j].saldoVirtual.pasokan{
					idx = j
				}
			}
			temp = A[i]
			A[i] = A[idx]
			A[idx] = temp
		}else{
			i = NMAX
		}
	}
	fmt.Println("=== Top Tier Kripto ===")
	for i = 0; i < NMAX; i++ {
		if A[i].id != 0 {
			fmt.Printf("ID: %d | Username: %s | Pasokan: %d\n",
				A[i].id, A[i].username, A[i].saldoVirtual.pasokan)
		}
	}
}
func hapusAccount(akun *arrAkun, id, pass int)bool{
	var i, idx, temp int
	var lakukan bool
	lakukan = false
	temp = NMAX
	i = 0
	idx = -1
	for i < NMAX{
		if akun[i].id == id && akun[i].password == pass{
			idx = i
			lakukan = true
			i = NMAX
		}else{
			i++
		}
	}
	if idx >= 0 && idx < temp && lakukan == true{
		for i = idx; i < temp-1; i++{
			akun[i] = akun[i+1]
		}
		temp--
		fmt.Println("akun Berhasil Dihapus, Logout Otomatis")
		return false
	}else{
		fmt.Println("Indeks Tidak Valid")
	}
	fmt.Println("\nAkun Yang Tersedia:")
	fmt.Printf("%-5s %-10s %-10s %-10s\n", "ID", "Username", "Saldo", "Pasokan")
	for i = 0; i < NMAX; i++ {
		if akun[i].id != 0{
			fmt.Printf("%-5d %-10s Rp%-10.2f %-10d\n", akun[i].id, akun[i].username, akun[i].moneyFiat, akun[i].saldoVirtual.pasokan)
		}
	}
	fmt.Println()
	return true
}
func logOut(akun arrAkun, id, pass int)bool{
	var i int
	var keluar, isLogin bool
	keluar = false 
	for i = 0; i < NMAX; i++{
		if akun[i].id == id && akun[i].password == pass{
			keluar = true
			i = NMAX
		}
	}
	if keluar == true{
		 fmt.Println("Berhasil Logout. Silakan login kembali.")
		isLogin = false
	}else{
		fmt.Println("Password atau Id Salah, Logout Gagal.")
		isLogin = true
	}
	return isLogin
}