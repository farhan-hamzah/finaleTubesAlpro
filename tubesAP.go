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

func main(){
	var akun arrAkun
	var dompetVirtual arrWalet
	var catatanTransaksi arrTransaksi
	var aplikasiAktif bool = true
	var tipeJualBeli string
	var tanggal, bulan, tahun, id, pass int

	akunFiktif(&akun, &dompetVirtual)
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
			fmt.Println("🙏 Terima kasih telah menggunakan aplikasi kami. Sampai jumpa di transaksi berikutnya!")
			aplikasiAktif = false
		}else{
			fmt.Println("❌ Pilihan tidak valid! Silakan pilih angka 1, 2, atau 3 sesuai menu yang tersedia.")
		}
		for isLogin{
			var dashboardOption int
			DashboardOption()
			fmt.Scan(&dashboardOption)
			switch dashboardOption{
				case 1:
					fmt.Print("🪙 Apa yang ingin Anda lakukan hari ini? (Ketik 'jual' atau 'beli'): ")
					fmt.Scan(&tipeJualBeli)
					jualBeli(tipeJualBeli, &catatanTransaksi, &dompetVirtual, &akun)
				case 2: 
					fmt.Print("🆔 Masukkan ID Anda (angka): ")
					fmt.Scan(&id)
					fmt.Print("🗓️  Masukkan tanggal transaksi (format: DD-MM-YYYY): ")
					fmt.Scan(&tanggal, &bulan, &tahun)
					mencariDataTranksasi(catatanTransaksi, akun, tanggal, bulan, tahun, id, dompetVirtual)
				case 3:
					rangkingKripto(akun)
				case 4:
					tampilanLaporanDataHargaKripto()
				case 5:
					fmt.Print("🆔 Masukkan ID Anda (angka): ")
					fmt.Scan(&id)
					fmt.Print("🔐 Masukkan Password: ")
					fmt.Scan(&pass)
					isLogin = hapusAccount(&akun, id, pass)
				case 6:
					fmt.Print("🆔 Masukkan ID Anda (angka): ")
					fmt.Scan(&id)
					fmt.Print("🔐 Masukkan Password: ")
					fmt.Scan(&pass)
					isLogin = logOut(akun, id, pass)
				case 7:
					aplikasiAktif = false
					isLogin = false
			}
		}
	}
}
func akunFiktif(akun *arrAkun, dompetVirtual *arrWalet){
	akun[0] = account{id: 103012400018,
	username: "farhan",
	password: 201105,
	moneyFiat: 1000000000000,
	saldoVirtual: wallet{id: 103012400018,
		name: "farhan",
		pasokan: 100}}
	dompetVirtual[0] = akun[0].saldoVirtual

akun[1] = account{id: 103012400215,
	username: "Tn. Putih",
	password: 20,
	moneyFiat: 23465432135,
	saldoVirtual: wallet{id: 103012400215,
		name: "Tn. Putih",
		pasokan: 98}}
	dompetVirtual[1] = akun[1].saldoVirtual

akun[2] = account{id: 103012400322,
	username: "fahrezi",
	password: 123456,
	moneyFiat: 999999999,
	saldoVirtual: wallet{id: 103012400322,
		name: "fahrezi",
		pasokan: 85}}
	dompetVirtual[2] = akun[2].saldoVirtual

akun[3] = account{id: 103012400429,
	username: "jotaro",
	password: 543210,
	moneyFiat: 500000000,
	saldoVirtual: wallet{id: 103012400429,
		name: "jotaro",
		pasokan: 120}}
	dompetVirtual[3] = akun[3].saldoVirtual

akun[4] = account{id: 103012400536,
	username: "satoshi",
	password: 111222,
	moneyFiat: 420000000,
	saldoVirtual: wallet{id: 103012400536,
		name: "satoshi",
		pasokan: 76}}
	dompetVirtual[4] = akun[4].saldoVirtual
}
func WelcomeLogin() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║            🌐 SELAMAT DATANG DI CRYPTO TRADE APP           ║")
	fmt.Println("║                💰 Platform Investasi Kripto                ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println("   Silakan pilih menu berikut:")
	fmt.Println("   [1] 🔐 Login")
	fmt.Println("   [2] 📝 Daftar Akun Baru")
	fmt.Println("   [3] ❌ Keluar Aplikasi")
	fmt.Println("-------------------------------------------------------------")
	fmt.Print("📥 Masukkan pilihan Anda : ")
}

func DashboardOption() {
	fmt.Println("\n╔═════════════════════📊 DASHBOARD MENU📊══════════════════════╗")
	fmt.Println("║ [1] 💱 Transaksi Jual/Beli Kripto                            ║")
	fmt.Println("║ [2] 🔍 Cari Riwayat Transaksi                                ║")
	fmt.Println("║ [3] 🏆 Tampilkan Top Akun Kripto                             ║")
	fmt.Println("║ [4] 📈 Laporan Data Harga Kripto                             ║")
	fmt.Println("║ [5] 🗑️ Hapus Akun                                             ║")
	fmt.Println("║ [6] 🔓 Logout                                                ║")
	fmt.Println("║ [7] ❌ Keluar Aplikasi                                       ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")
	fmt.Print("📥 Masukkan pilihan Anda: ")
}

func login(akun *arrAkun, pilihAkun int)bool{
	var i, id int
	var pass int
	var benarLogin bool
	if pilihAkun == 1{
		fmt.Print("🆔 Masukkan ID Anda (angka): ")
		fmt.Scan(&id)
		fmt.Print("🔐 Masukkan Password: ")
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
			fmt.Println("✅ Login berhasil! Selamat datang kembali 👋")
		}else{
			fmt.Println("❌ ID atau kata sandi salah. Coba lagi ya!")
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
		fmt.Println("🆔 Silakan buat ID Anda (gunakan kombinasi angka 0-9):")
		fmt.Print("👉 ID: ")
		fmt.Scan(&id)
		fmt.Println("👤 Masukkan nama pengguna (username) Anda:")
		fmt.Print("👉 Username: ")
		fmt.Scan(&usrn)
		fmt.Println("🔒 Silakan buat password (gunakan kombinasi angka 0-9):")
		fmt.Print("👉 Password: ")
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
			fmt.Print("💵 Masukkan jumlah uang: ")
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
				fmt.Println("🎉 Akun berhasil dibuat! Selamat datang di platform kami.")
			}else{
				fmt.Println("⚠️  ID sudah digunakan. Silakan coba lagi dengan ID lain.")
			}
		}	
	}
	return buatAkun
}
func jualBeli(tipeJualBeli string, dataJualBeli *arrTransaksi, wallet *arrWalet, akun *arrAkun){
	var pasokanKripto, jumlahKriptoYangDijual, jumlahKriptoYangDibeli, nilaiBeliKripto int
	var i, tanggal, bulan, tahun, idxWallet, id, pw, n int
	var nama string
	var nilaiJualKripto float64
	var valid bool
	valid = true
	i = 0
	if tipeJualBeli == "beli" || tipeJualBeli == "jual"{
		fmt.Print("🆔 Masukkan ID Anda (angka): ")
		fmt.Scan(&id)
		fmt.Print("🔐 Masukkan Password: ")
		fmt.Scan(&pw)
	}
	i = 0
	for i < NMAX{
		if akun[i].id == id && akun[i].password == pw{
			pasokanKripto += wallet[i].pasokan
			idxWallet = i
			valid = false
			nama = akun[i].username
			i = NMAX
		}else if akun[i].id == 0{
			i = NMAX
		}else{
			i++
		}
	}
	if valid != false{
		fmt.Println("❌ ID atau kata sandi salah. Coba lagi ya!")
	}
	for valid == false{
		i = 0
		if tipeJualBeli == "jual"{	
			fmt.Print("💸 Masukkan jumlah kripto yang ingin Anda jual: ")
			fmt.Scan(&jumlahKriptoYangDijual)
			if jumlahKriptoYangDijual <= pasokanKripto{
				nilaiJualKripto = float64(jumlahKriptoYangDijual)*21000
				akun[idxWallet].moneyFiat += nilaiJualKripto
				fmt.Print("🗓️  Masukkan tanggal transaksi (format: DD-MM-YYYY): ")
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
				akun[idxWallet].saldoVirtual.id = id
				akun[idxWallet].saldoVirtual.name = nama
				akun[idxWallet].saldoVirtual.pasokan = pasokanKripto-jumlahKriptoYangDijual
				wallet[idxWallet] = akun[idxWallet].saldoVirtual
				valid = true
			}else{
				fmt.Println("🚫 Jumlah kripto yang dimiliki tidak mencukupi. Silakan periksa saldo Anda.")
				valid = true
			}
		}else if tipeJualBeli == "beli"{
			fmt.Print("💰 Masukkan jumlah kripto yang ingin Anda beli: ")
			fmt.Scan(&jumlahKriptoYangDibeli)
			nilaiBeliKripto = jumlahKriptoYangDibeli*21000
			if akun[idxWallet].moneyFiat >= float64(nilaiBeliKripto){
				akun[idxWallet].moneyFiat -=float64(nilaiBeliKripto)
				fmt.Print("🗓️  Masukkan tanggal transaksi (format: DD-MM-YYYY): ")
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
				akun[idxWallet].saldoVirtual.id = id
				akun[idxWallet].saldoVirtual.name = nama
				akun[idxWallet].saldoVirtual.pasokan = pasokanKripto+jumlahKriptoYangDibeli
				wallet[idxWallet] = akun[idxWallet].saldoVirtual
				valid = true
			}else{
				fmt.Println("💸 Maaf, saldo Anda tidak mencukupi untuk melakukan transaksi.")
				valid = true
			}		
		}else{
			fmt.Println("⚠️ Tipe transaksi tidak dikenali. Silakan ketik 'jual' atau 'beli'.")
			valid = true
		}
	}
	
}
func mencariDataTranksasi(dataJualBeli arrTransaksi, akun arrAkun, tanggal, bulan, tahun, id int, wallet arrWalet) {
	var i, temp int
	var ada bool
	ada = false
	temp = -1
	i = 0
	for i < NMAX && ada == false {
		if dataJualBeli[i].userId == 0 {
			i = NMAX
		}else if dataJualBeli[i].tanggal == tanggal && dataJualBeli[i].bulan == bulan && dataJualBeli[i].tahun == tahun && dataJualBeli[i].userId == id {
			temp = i
			ada = true
			i = NMAX
		}else{
			i++
		}
	}
	var idxAkun int
    idxAkun = -1
    var j = 0
    for j < NMAX && idxAkun == -1 {
        if akun[j].id == id {
            idxAkun = j
            j = NMAX  
        } else if akun[j].id == 0 {
            j = NMAX  
        } else {
            j++
        }
	}
	if ada == true {
		fmt.Printf("\n📄 Transaksi Tanggal: %02d/%02d/%04d\n", tanggal, bulan, tahun)
		fmt.Println("═══════════════════════════════════════════════════════════")
		fmt.Printf("%-10s | %-10s | %-10s | %-10s | %-15s\n", "ID", "Username", "Saldo", "Kripto", "Jenis Transaksi")
		fmt.Println("───────────────────────────────────────────────────────────")
		fmt.Printf("%-10d | %-10s | Rp%-9.2f | %-10d | %-15s\n",
		akun[idxAkun].id, akun[idxAkun].username, akun[idxAkun].moneyFiat, wallet[idxAkun].pasokan, dataJualBeli[temp].tipeTransaksi)
		fmt.Println("═══════════════════════════════════════════════════════════")
	} else {
		fmt.Println("📭 Tidak ditemukan riwayat transaksi.")
	}
}
func tampilanLaporanDataHargaKripto(){
    fmt.Println("+----+----------------+--------+----------------+------------------+----------------+----------------+----------------+")
    fmt.Println("| No | Nama Aset      | Jenis  | Harga Saat Ini | Kapitalisasi     | Volume 24 Jam  | Perubahan 24h  | Dominasi Pasar |")
    fmt.Println("+----+----------------+--------+----------------+------------------+----------------+----------------+----------------+")
    fmt.Println("| 1  | Bitcoin (BTC)  | Kripto | Rp12.000,00     | Rp3.500.000.000  | Rp700.000.000  | -0,24%         | 59,14%          |")
    fmt.Println("+----+----------------+--------+----------------+------------------+----------------+----------------+----------------+")
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
	fmt.Println("🔥 === Kripto Papan Atas Saat Ini ===")
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
		fmt.Println("✅ Akun berhasil dihapus. Anda telah logout secara otomatis.")
		return false
	}else{
		fmt.Println("🚫 Indeks yang dimasukkan tidak valid. Silakan coba lagi.")
	}
	fmt.Println("\n👥 Akun yang tersedia saat ini:")
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
		fmt.Println("🔒 Logout berhasil. Silakan login kembali untuk mengakses akun Anda.")
		isLogin = false
	}else{
		fmt.Println("Password atau Id Salah, Logout Gagal.")
		isLogin = true
	}
	return isLogin
}