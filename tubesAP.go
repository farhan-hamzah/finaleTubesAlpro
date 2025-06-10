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
	tipeKripto string
	quantityFiat int
	quantityKriptoJual int
	quantityKriptoBeli int
	tanggal int
	bulan int
	tahun int
}
type wallet struct{
	id int
	name string
	btc, eth, sol, binance int
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
	var tipeJualBeli, ubahUsrnPw, tipeCoin, rankCripto string
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
					fmt.Print("✏ Ubah Username atau Password: ")
					fmt.Scan(&ubahUsrnPw)
					gantiUsrn(ubahUsrnPw, &akun, &dompetVirtual)
				case 2:
					fmt.Print("🪙 Apa yang ingin Anda lakukan hari ini? (Ketik 'jual' atau 'beli'): ")
					fmt.Scan(&tipeJualBeli)
					jualBeli(tipeJualBeli, &catatanTransaksi, &dompetVirtual, &akun)
				case 3: 
					fmt.Print("🆔 Masukkan ID Anda (angka): ")
					fmt.Scan(&id)
					fmt.Print("🗓️  Masukkan tanggal transaksi (format: DD-MM-YYYY): ")
					fmt.Scan(&tanggal, &bulan, &tahun)
					fmt.Print("🪙 Masukan Tipe Kripto (BTC / ETH / SOL / BNB): ")
					fmt.Scan(&tipeCoin)   	
					mencariDataTranksasi(catatanTransaksi, akun, tanggal, bulan, tahun, id, tipeCoin)
				case 4:
					fmt.Print("🏆 Rangking Kripto Berdasarkan (BTC / ETH / SOL / BNB): ")
					fmt.Scan(&rankCripto)
					rangkingKripto(akun, rankCripto)
				case 5:
					tampilanLaporanDataHargaKripto()
				case 6:
					fmt.Print("🆔 Masukkan ID Anda (angka): ")
					fmt.Scan(&id)
					fmt.Print("🔐 Masukkan Password: ")
					fmt.Scan(&pass)
					isLogin = hapusAccount(&akun, id, pass)
				case 7:
					fmt.Print("🆔 Masukkan ID Anda (angka): ")
					fmt.Scan(&id)
					fmt.Print("🔐 Masukkan Password: ")
					fmt.Scan(&pass)
					isLogin = logOut(akun, id, pass)
				case 8:
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
		btc: 100,
		eth: 25,
		sol: 43,
		binance: 15 }}
	dompetVirtual[0] = akun[0].saldoVirtual

akun[1] = account{id: 103012400215,
	username: "Tn. Putih",
	password: 20,
	moneyFiat: 23465432135,
	saldoVirtual: wallet{id: 103012400215,
		name: "Tn. Putih",
		btc: 98,
		eth: 55,
		sol: 76,
		binance: 21}}
	dompetVirtual[1] = akun[1].saldoVirtual

akun[2] = account{id: 103012400322,
	username: "fahrezi",
	password: 123456,
	moneyFiat: 999999999,
	saldoVirtual: wallet{id: 103012400322,
		name: "fahrezi",
		btc: 85,
		eth: 21,
		sol: 15,
		binance: 10}}
	dompetVirtual[2] = akun[2].saldoVirtual

akun[3] = account{id: 103012400429,
	username: "jotaro",
	password: 543210,
	moneyFiat: 500000000,
	saldoVirtual: wallet{id: 103012400429,
		name: "jotaro",
		btc: 120,
		eth: 75,
		sol: 21,
		binance: 11}}
	dompetVirtual[3] = akun[3].saldoVirtual

akun[4] = account{id: 103012400536,
	username: "satoshi",
	password: 111222,
	moneyFiat: 420000000,
	saldoVirtual: wallet{id: 103012400536,
		name: "satoshi",
		btc: 76,
		eth: 17,
		sol: 21,
		binance: 98}}
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
	fmt.Println("║ [1] ✏  Ubah Username/Password                                ║")
	fmt.Println("║ [2] 💱 Transaksi Jual/Beli Kripto                            ║")
	fmt.Println("║ [3] 🔍 Cari Riwayat Transaksi                                ║")
	fmt.Println("║ [4] 🏆 Tampilkan Top Akun Kripto                             ║")
	fmt.Println("║ [5] 📈 Laporan Data Harga Kripto                             ║")
	fmt.Println("║ [6] 🗑️  Hapus Akun                                            ║")
	fmt.Println("║ [7] 🔓 Logout                                                ║")
	fmt.Println("║ [8] ❌ Keluar Aplikasi                                       ║")
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
				akun[n].saldoVirtual.btc = 0
				akun[n].saldoVirtual.eth = 0
				akun[n].saldoVirtual.sol = 0
				akun[n].saldoVirtual.binance = 0
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
func gantiUsrn(gantiUsrn string, A *arrAkun, B *arrWalet){
	var id, i, idx, pass, newPass int
	var user, pw bool
	var newUsrn string
	user = false
	pw = false
	if gantiUsrn == "username"{
		fmt.Print("🆔 Masukkan ID Anda (angka): ")
		fmt.Scan(&id)
		fmt.Print("🔐 Masukkan Password: ")
		fmt.Scan(&pass)
		for i = 0; i < NMAX; i++{
			if A[i].id == id && A[i].password == pass{
				idx = i
				user = true
			}
		}
	}else if gantiUsrn == "password"{
		fmt.Print("🆔 Masukkan ID Anda (angka): ")
		fmt.Scan(&id)
		fmt.Print("🔐 Masukkan Password: ")
		fmt.Scan(&pass)
		for i = 0; i < NMAX; i++{
			if A[i].id == id && A[i].password == pass{
				idx = i
				pw = true
			}
		}
	}
	if user == true{
		fmt.Print("📝👤 Masukan Username Baru: ")
		fmt.Scan(&newUsrn)
		A[idx].username = newUsrn
		A[idx].saldoVirtual.name = newUsrn
		B[idx].name = newUsrn
		fmt.Println("✅ Username Berhasil Diubah")
	}else if pw == true{
		fmt.Print("🔑🛡 Masukan Password Baru: ")
		fmt.Scan(&newPass)
		A[idx].password = newPass
		fmt.Println("✅ Password Berhasil Diubah")
	}else{
		fmt.Print("❌ Id atau Kata Sandi Salah")
	}
}
func jualBeli(tipeJualBeli string, dataJualBeli *arrTransaksi, wallet *arrWalet, akun *arrAkun){
	var pasokanBTC, pasokanETH, pasokanSOL, pasokanBNB, jumlahKriptoYangDijual, jumlahKriptoYangDibeli, nilaiBeliKripto int
	var i, tanggal, bulan, tahun, idxWallet, id, pw, n int
	var nama string
	var nilaiJualKripto float64
	var valid bool
	var beliKoin, jualKoin string
	valid = false
	i = 0
	if tipeJualBeli == "jual"{
		fmt.Print("🆔 Masukkan ID Anda (angka): ")
		fmt.Scan(&id)
		fmt.Print("🔐 Masukkan Password: ")
		fmt.Scan(&pw)
		fmt.Print("🪙 Tentukan Aset Yang Mau Dijual (BTC / ETH / SOL / BNB): ")
		fmt.Scan(&jualKoin)
	}
	if tipeJualBeli == "beli"{
		fmt.Print("🆔 Masukkan ID Anda (angka): ")
		fmt.Scan(&id)
		fmt.Print("🔐 Masukkan Password: ")
		fmt.Scan(&pw)
		fmt.Print("🪙 Tentukan Aset Yang Mau Dibeli (BTC / ETH / SOL / BNB): ")
		fmt.Scan(&beliKoin)
	}
	for i < NMAX && akun[i].id != 0{
		if akun[i].id == id && akun[i].password == pw {
			idxWallet = i
			nama = akun[i].username
			pasokanBTC = wallet[i].btc
			pasokanETH = wallet[i].eth
			pasokanSOL = wallet[i].sol
			pasokanBNB = wallet[i].binance
			valid = true
		}
		i++
	}
	if valid == false{
		fmt.Println("❌ ID atau kata sandi salah. Coba lagi ya!")
	}
	for valid == true{
		i = 0
		if tipeJualBeli == "jual"{	
			fmt.Print("💸 Masukkan jumlah kripto ", jualKoin, " yang ingin Anda jual: ")
			fmt.Scan(&jumlahKriptoYangDijual)
			switch jualKoin {
				case "BTC":
					if jumlahKriptoYangDijual <= pasokanBTC{
						nilaiJualKripto = float64(jumlahKriptoYangDijual)*21000
						akun[idxWallet].moneyFiat += nilaiJualKripto
						fmt.Print("🗓️  Masukkan tanggal transaksi (format: DD-MM-YYYY): ")
						fmt.Scan(&tanggal, &bulan, &tahun)
						for i < NMAX{
							if dataJualBeli[i].userId == 0{
								dataJualBeli[i].tipeKripto = jualKoin
								dataJualBeli[i].userId = id
								dataJualBeli[i].tipeTransaksi = tipeJualBeli
								dataJualBeli[i].quantityFiat = int(nilaiJualKripto) 
								dataJualBeli[i].quantityKriptoJual = jumlahKriptoYangDijual
								dataJualBeli[i].quantityKriptoBeli = 0
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
						akun[idxWallet].saldoVirtual.btc = pasokanBTC-jumlahKriptoYangDijual
						wallet[idxWallet] = akun[idxWallet].saldoVirtual
						fmt.Print("✅ Kripto ", jualKoin, " Berhasil Dijual")
						valid = false
					}else{
						fmt.Println("🚫 Jumlah kripto yang dimiliki tidak mencukupi. Silakan periksa saldo Anda.")
						valid = false
					}
				case "ETH":
					if jumlahKriptoYangDijual <= pasokanBTC{
						nilaiJualKripto = float64(jumlahKriptoYangDijual)*19000
						akun[idxWallet].moneyFiat += nilaiJualKripto
						fmt.Print("🗓️  Masukkan tanggal transaksi (format: DD-MM-YYYY): ")
						fmt.Scan(&tanggal, &bulan, &tahun)
						for i < NMAX{
							if dataJualBeli[i].userId == 0{
								dataJualBeli[i].tipeKripto = jualKoin
								dataJualBeli[i].userId = id
								dataJualBeli[i].tipeTransaksi = tipeJualBeli
								dataJualBeli[i].quantityFiat = int(nilaiJualKripto) 
								dataJualBeli[i].quantityKriptoJual = jumlahKriptoYangDijual
								dataJualBeli[i].quantityKriptoBeli = 0
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
						akun[idxWallet].saldoVirtual.eth = pasokanETH-jumlahKriptoYangDijual
						wallet[idxWallet] = akun[idxWallet].saldoVirtual
						fmt.Print("✅ Kripto ", jualKoin, " Berhasil Dijual")
						valid = false
					}else{
						fmt.Println("🚫 Jumlah kripto yang dimiliki tidak mencukupi. Silakan periksa saldo Anda.")
						valid = false
					}
				case "SOL":
					if jumlahKriptoYangDijual <= pasokanBTC{
						nilaiJualKripto = float64(jumlahKriptoYangDijual)*17000
						akun[idxWallet].moneyFiat += nilaiJualKripto
						fmt.Print("🗓️  Masukkan tanggal transaksi (format: DD-MM-YYYY): ")
						fmt.Scan(&tanggal, &bulan, &tahun)
						for i < NMAX{
							if dataJualBeli[i].userId == 0{
								dataJualBeli[i].tipeKripto = jualKoin
								dataJualBeli[i].userId = id
								dataJualBeli[i].tipeTransaksi = tipeJualBeli
								dataJualBeli[i].quantityFiat = int(nilaiJualKripto) 
								dataJualBeli[i].quantityKriptoJual = jumlahKriptoYangDijual
								dataJualBeli[i].quantityKriptoBeli = 0
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
						akun[idxWallet].saldoVirtual.sol = pasokanBTC-jumlahKriptoYangDijual
						wallet[idxWallet] = akun[idxWallet].saldoVirtual
						fmt.Print("✅ Kripto ", jualKoin, " Berhasil Dijual")
						valid = false
					}else{
						fmt.Println("🚫 Jumlah kripto yang dimiliki tidak mencukupi. Silakan periksa saldo Anda.")
						valid = false
					}
				case "BNB":
					if jumlahKriptoYangDijual <= pasokanBTC{
						nilaiJualKripto = float64(jumlahKriptoYangDijual)*13000
						akun[idxWallet].moneyFiat += nilaiJualKripto
						fmt.Print("🗓️  Masukkan tanggal transaksi (format: DD-MM-YYYY): ")
						fmt.Scan(&tanggal, &bulan, &tahun)
						for i < NMAX{
							if dataJualBeli[i].userId == 0{
								dataJualBeli[i].tipeKripto = jualKoin
								dataJualBeli[i].userId = id
								dataJualBeli[i].tipeTransaksi = tipeJualBeli
								dataJualBeli[i].quantityFiat = int(nilaiJualKripto) 
								dataJualBeli[i].quantityKriptoJual = jumlahKriptoYangDijual
								dataJualBeli[i].quantityKriptoBeli = 0
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
						akun[idxWallet].saldoVirtual.binance = pasokanBTC-jumlahKriptoYangDijual
						wallet[idxWallet] = akun[idxWallet].saldoVirtual
						fmt.Print("✅ Kripto ", jualKoin, " Berhasil Dijual")
						valid = false
					}else{
						fmt.Println("🚫 Jumlah kripto yang dimiliki tidak mencukupi. Silakan periksa saldo Anda.")
						valid = false
					}
				} 
		}else if tipeJualBeli == "beli"{
			fmt.Print("💰 Masukkan jumlah kripto ", beliKoin," yang ingin Anda beli: ")
			fmt.Scan(&jumlahKriptoYangDibeli)
			switch beliKoin {
				case "BTC":
					nilaiBeliKripto = jumlahKriptoYangDibeli*21000
					if akun[idxWallet].moneyFiat >= float64(nilaiBeliKripto){
						akun[idxWallet].moneyFiat -=float64(nilaiBeliKripto)
						fmt.Print("🗓️  Masukkan tanggal transaksi (format: DD-MM-YYYY): ")
						fmt.Scan(&tanggal, &bulan, &tahun)
						for i < NMAX{
							if dataJualBeli[i].userId == 0{
								dataJualBeli[i].tipeKripto = beliKoin
								dataJualBeli[i].userId = id
								dataJualBeli[i].tipeTransaksi = tipeJualBeli
								dataJualBeli[i].quantityFiat = -nilaiBeliKripto 	
								dataJualBeli[i].quantityKriptoBeli = jumlahKriptoYangDibeli
								dataJualBeli[i].quantityKriptoJual = 0
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
						akun[idxWallet].saldoVirtual.btc = pasokanBTC+jumlahKriptoYangDibeli
						wallet[idxWallet] = akun[idxWallet].saldoVirtual
						fmt.Print("✅ Kripto ", beliKoin, " Berhasil Dibeli")
						valid = false
					}else if akun[idxWallet].moneyFiat < float64(nilaiBeliKripto) {
						fmt.Println("💸 Maaf, saldo Anda tidak mencukupi untuk melakukan transaksi.")
						valid = false		
					}
				case "ETH":
					nilaiBeliKripto = jumlahKriptoYangDibeli*19000
					if akun[idxWallet].moneyFiat >= float64(nilaiBeliKripto){
						akun[idxWallet].moneyFiat -=float64(nilaiBeliKripto)
						fmt.Print("🗓️  Masukkan tanggal transaksi (format: DD-MM-YYYY): ")
						fmt.Scan(&tanggal, &bulan, &tahun)
						for i < NMAX{
							if dataJualBeli[i].userId == 0{
								dataJualBeli[i].tipeKripto = beliKoin
								dataJualBeli[i].userId = id
								dataJualBeli[i].tipeTransaksi = tipeJualBeli
								dataJualBeli[i].quantityFiat = -nilaiBeliKripto
								dataJualBeli[i].quantityKriptoBeli = jumlahKriptoYangDibeli
								dataJualBeli[i].quantityKriptoJual = 0
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
						akun[idxWallet].saldoVirtual.eth = pasokanETH+jumlahKriptoYangDibeli
						wallet[idxWallet] = akun[idxWallet].saldoVirtual
						fmt.Print("✅ Kripto ", beliKoin, " Berhasil Dibeli")
						valid = false
					}else{
						fmt.Println("💸 Maaf, saldo Anda tidak mencukupi untuk melakukan transaksi.")
						valid = false		
					}
				case "SOL":
					nilaiBeliKripto = jumlahKriptoYangDibeli*17000
					if akun[idxWallet].moneyFiat >= float64(nilaiBeliKripto){
						akun[idxWallet].moneyFiat -=float64(nilaiBeliKripto)
						fmt.Print("🗓️  Masukkan tanggal transaksi (format: DD-MM-YYYY): ")
						fmt.Scan(&tanggal, &bulan, &tahun)
						for i < NMAX{
							if dataJualBeli[i].userId == 0{
								dataJualBeli[i].tipeKripto = beliKoin
								dataJualBeli[i].userId = id
								dataJualBeli[i].tipeTransaksi = tipeJualBeli
								dataJualBeli[i].quantityFiat = -nilaiBeliKripto 	
								dataJualBeli[i].quantityKriptoBeli = jumlahKriptoYangDibeli
								dataJualBeli[i].quantityKriptoJual = 0
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
						akun[idxWallet].saldoVirtual.sol = pasokanSOL+jumlahKriptoYangDibeli
						wallet[idxWallet] = akun[idxWallet].saldoVirtual
						fmt.Print("✅ Kripto ", beliKoin, " Berhasil Dibeli")
						valid = false
					}else{
						fmt.Println("💸 Maaf, saldo Anda tidak mencukupi untuk melakukan transaksi.")
						valid = false		
					}
				case "BNB":
					nilaiBeliKripto = jumlahKriptoYangDibeli*13000
					if akun[idxWallet].moneyFiat >= float64(nilaiBeliKripto){
						akun[idxWallet].moneyFiat -=float64(nilaiBeliKripto)
						fmt.Print("🗓️  Masukkan tanggal transaksi (format: DD-MM-YYYY): ")
						fmt.Scan(&tanggal, &bulan, &tahun)
						for i < NMAX{
							if dataJualBeli[i].userId == 0{
								dataJualBeli[i].tipeKripto = beliKoin
								dataJualBeli[i].userId = id
								dataJualBeli[i].tipeTransaksi = tipeJualBeli
								dataJualBeli[i].quantityFiat = -nilaiBeliKripto 	
								dataJualBeli[i].quantityKriptoBeli = jumlahKriptoYangDibeli
								dataJualBeli[i].quantityKriptoJual = 0
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
						akun[idxWallet].saldoVirtual.binance = pasokanBNB+jumlahKriptoYangDibeli
						wallet[idxWallet] = akun[idxWallet].saldoVirtual
						fmt.Print("✅ Kripto ", beliKoin, " Berhasil Dibeli")
						valid = false
					}else{
						fmt.Println("💸 Maaf, saldo Anda tidak mencukupi untuk melakukan transaksi.")
						valid = false		
					}
				}	
			
		}else{
			fmt.Println("⚠️ Tipe transaksi tidak dikenali. Silakan ketik 'jual' atau 'beli'.")
			valid = false
		}
	}

}
func mencariDataTranksasi(dataJualBeli arrTransaksi, akun arrAkun, tanggal, bulan, tahun, id int, tipeKoin string) {
	var i, temp int
	var ada bool
	ada = false
	temp = -1
	i = 0
	for i < NMAX && ada == false {
		if dataJualBeli[i].userId == 0 {
			i = NMAX
		}else if dataJualBeli[i].tanggal == tanggal && dataJualBeli[i].bulan == bulan && dataJualBeli[i].tahun == tahun && dataJualBeli[i].userId == id && dataJualBeli[i].tipeKripto == tipeKoin {
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
	switch tipeKoin{
		case "BTC":
			if ada == true {
				fmt.Printf("\n📄 Transaksi Tanggal: %02d/%02d/%04d\n", tanggal, bulan, tahun)
				fmt.Println("═════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════")
				fmt.Printf("%-5s | %-12s | %-18s | %-22s | %-22s | %-18s | %-15s\n", "ID", "Username", "Saldo (Fiat)", "Jumlah Kripto Dibeli", "Jumlah Kripto Dijual", "Jenis Transaksi", "Tipe Kripto")
				fmt.Println("─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────")
				fmt.Printf("%-5d | %-12s | Rp%-15.2f | %-22d | %-22d | %-18s | %-15s\n", akun[idxAkun].id, akun[idxAkun].username, akun[idxAkun].moneyFiat, dataJualBeli[temp].quantityKriptoBeli, dataJualBeli[temp].quantityKriptoJual, dataJualBeli[temp].tipeTransaksi, tipeKoin)
				fmt.Println("═════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════")
			}else {	
				fmt.Println("📭 Tidak ditemukan riwayat transaksi.")
			}		
		case "ETH":
			if ada == true {
				fmt.Printf("\n📄 Transaksi Tanggal: %02d/%02d/%04d\n", tanggal, bulan, tahun)
				fmt.Println("═════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════")
				fmt.Printf("%-5s | %-12s | %-18s | %-22s | %-22s | %-18s | %-15s\n", "ID", "Username", "Saldo (Fiat)", "Jumlah Kripto Dibeli", "Jumlah Kripto Dijual", "Jenis Transaksi", "Tipe Kripto")
				fmt.Println("─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────")
				fmt.Printf("%-5d | %-12s | Rp%-15.2f | %-22d | %-22d | %-18s | %-15s\n", akun[idxAkun].id, akun[idxAkun].username, akun[idxAkun].moneyFiat, dataJualBeli[temp].quantityKriptoBeli, dataJualBeli[temp].quantityKriptoJual, dataJualBeli[temp].tipeTransaksi, tipeKoin)
				fmt.Println("═════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════")
			}else {	
				fmt.Println("📭 Tidak ditemukan riwayat transaksi.")
			}		
		case "SOL":
			if ada == true {
				fmt.Printf("\n📄 Transaksi Tanggal: %02d/%02d/%04d\n", tanggal, bulan, tahun)
				fmt.Println("═════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════")
				fmt.Printf("%-5s | %-12s | %-18s | %-22s | %-22s | %-18s | %-15s\n", "ID", "Username", "Saldo (Fiat)", "Jumlah Kripto Dibeli", "Jumlah Kripto Dijual", "Jenis Transaksi", "Tipe Kripto")
				fmt.Println("─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────")
				fmt.Printf("%-5d | %-12s | Rp%-15.2f | %-22d | %-22d | %-18s | %-15s\n", akun[idxAkun].id, akun[idxAkun].username, akun[idxAkun].moneyFiat, dataJualBeli[temp].quantityKriptoBeli, dataJualBeli[temp].quantityKriptoJual, dataJualBeli[temp].tipeTransaksi, tipeKoin)
				fmt.Println("═════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════")
			}else {	
				fmt.Println("📭 Tidak ditemukan riwayat transaksi.")
			}	
		case "BNB":
			if ada == true {
				fmt.Printf("\n📄 Transaksi Tanggal: %02d/%02d/%04d\n", tanggal, bulan, tahun)
				fmt.Println("═════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════")
				fmt.Printf("%-5s | %-12s | %-18s | %-22s | %-22s | %-18s | %-15s\n", "ID", "Username", "Saldo (Fiat)", "Jumlah Kripto Dibeli", "Jumlah Kripto Dijual", "Jenis Transaksi", "Tipe Kripto")
				fmt.Println("─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────")
				fmt.Printf("%-5d | %-12s | Rp%-15.2f | %-22d | %-22d | %-18s | %-15s\n", akun[idxAkun].id, akun[idxAkun].username, akun[idxAkun].moneyFiat, dataJualBeli[temp].quantityKriptoBeli, dataJualBeli[temp].quantityKriptoJual, dataJualBeli[temp].tipeTransaksi, tipeKoin)
				fmt.Println("═════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════")
			}else {	
				fmt.Println("📭 Tidak ditemukan riwayat transaksi.")
			}	
	}

}
func tampilanLaporanDataHargaKripto() {
	fmt.Println("+----+----------------+--------+----------------+------------------+----------------+----------------+----------------+")
	fmt.Println("| No | Nama Aset      | Jenis  | Harga Saat Ini | Kapitalisasi     | Volume 24 Jam  | Perubahan 24h  | Dominasi Pasar |")
	fmt.Println("+----+----------------+--------+----------------+------------------+----------------+----------------+----------------+")
	fmt.Println("| 1  | Bitcoin (BTC)  | Kripto | Rp21.000,00     | Rp3.500.000.000  | Rp700.000.000  | -0,24%         | 59,14%          |")
	fmt.Println("| 2  | Ethereum (ETH) | Kripto | Rp19.000,00     | Rp1.800.000.000  | Rp350.000.000  | +0,56%         | 18,21%          |")
	fmt.Println("| 3  | Solana (SOL)   | Kripto | Rp17.000,00     | Rp650.000.000    | Rp120.000.000  | +1,12%         | 4,98%           |")
	fmt.Println("| 4  | Binance (BNB)   | Kripto | Rp13.000,00     | Rp500.000.000    | Rp100.000.000  | -0,75%         | 3,10%           |")
	fmt.Println("+----+----------------+--------+----------------+------------------+----------------+----------------+----------------+")
}


func rangkingKripto(A arrAkun, rankCripto string){
	var i, idx, j int
	var temp account
	if rankCripto == "BTC"{
		for i = 0; i < NMAX; i++{
			if A[i].id != 0{
				idx = i
				for j = i+1; j < NMAX; j++{
					if A[j].id != 0 && A[i].saldoVirtual.btc < A[j].saldoVirtual.btc{
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
		fmt.Println("🔥 === Kripto ",rankCripto, " Papan Atas Saat Ini ===")
		for i = 0; i < NMAX; i++ {
			if A[i].id != 0 {
				fmt.Printf("ID: %d | Username: %s | Pasokan BTC: %d\n", A[i].id, A[i].username, A[i].saldoVirtual.btc)
			}
		}
	}else if rankCripto == "ETH"{
		for i = 0; i < NMAX; i++{
			if A[i].id != 0{
				idx = i
				for j = i+1; j < NMAX; j++{
					if A[j].id != 0 && A[i].saldoVirtual.eth < A[j].saldoVirtual.eth{
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
		fmt.Println("🔥 === Kripto ",rankCripto, " Papan Atas Saat Ini ===")
		for i = 0; i < NMAX; i++ {
			if A[i].id != 0 {
				fmt.Printf("ID: %d | Username: %s | Pasokan ETH: %d\n", A[i].id, A[i].username, A[i].saldoVirtual.eth)
			}
		}
	}else if rankCripto == "SOL"{
		for i = 0; i < NMAX; i++{
			if A[i].id != 0{
				idx = i
				for j = i+1; j < NMAX; j++{
					if A[j].id != 0 && A[i].saldoVirtual.sol < A[j].saldoVirtual.sol{
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
		fmt.Println("🔥 === Kripto ",rankCripto, " Papan Atas Saat Ini ===")
		for i = 0; i < NMAX; i++ {
			if A[i].id != 0 {
				fmt.Printf("ID: %d | Username: %s | Pasokan SOL: %d\n", A[i].id, A[i].username, A[i].saldoVirtual.sol)
			}
		}
	}else if rankCripto == "BNB"{
		for i = 0; i < NMAX; i++{
			if A[i].id != 0{
				idx = i
				for j = i+1; j < NMAX; j++{
					if A[j].id != 0 && A[i].saldoVirtual.binance < A[j].saldoVirtual.binance{
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
		fmt.Println("🔥 === Kripto ",rankCripto, " Papan Atas Saat Ini ===")
		for i = 0; i < NMAX; i++ {
			if A[i].id != 0 {
				fmt.Printf("ID: %d | Username: %s | Pasokan BNB: %d\n", A[i].id, A[i].username, A[i].saldoVirtual.binance)
			}
		}
	}else{
		fmt.Print("❌ Masukan Kripto Tidak Valid")
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
	fmt.Printf("%-5s %-12s %-15s %-10s %-10s %-10s %-10s\n", "ID", "Username", "Saldo", "BTC", "ETH", "SOL", "BNB")
	for i = 0; i < NMAX; i++ {
		if akun[i].id != 0 {
			fmt.Printf("%-5d %-12s Rp%-13.2f %-10d %-10d %-10d %-10d\n", akun[i].id, akun[i].username, akun[i].moneyFiat, akun[i].saldoVirtual.btc, akun[i].saldoVirtual.eth, akun[i].saldoVirtual.sol, akun[i].saldoVirtual.binance)
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
		fmt.Println("❌ Password atau Id Salah, Logout Gagal.")
		isLogin = true
	}
	return isLogin
}