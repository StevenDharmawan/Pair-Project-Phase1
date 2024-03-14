package cli

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
	"PairProjectPhase1/handler"
	"bufio"
	"fmt"
	"net/mail"
	"os"
	"regexp"
	"strconv"
	"time"
)

func RunProgram() {
	scanner := bufio.NewScanner(os.Stdin)
	err := config.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer config.DB.Close()
	fmt.Println("WELCOME TO SY STORE!")
	for {
		fmt.Println("1. Login\n2. Register\n3. Exit")
		input := userInput("Masukkan input berdasarkan angka di menu (1/2/3): ", scanner)
		switch input {
		case "1":
			fmt.Println("LOGIN")
			email := userInput("Masukkan Email Anda: ", scanner)
			password := userInput("Masukkan Password Anda: ", scanner)
			fmt.Println("Mohon menunggu sesaat")
			user, valid := handler.Login(email, password)
			if valid {
				fmt.Println("Login berhasil! Selamat datang", user.Name)
				if user.Role == "Admin" {
					fmt.Println("1. Menampilkan report user\n2. Menampilkan report\n3. Menampilkan report\n4. Menambahkan barang\n5. Menghapus barang\n6. Update stock barang\n7. Exit")
					inputAdmin := userInput("Masukkan input berdasarkan angka di menu (1-7): ", scanner)
					switch inputAdmin {
					case "1":

					case "2":

					case "3":

					case "4":

					case "5":

					case "6":

					case "7":
						return
					default:
						fmt.Println("Input yang dimasukkan tidak valid")
					}
				} else {
					for {
						fmt.Println("1. Menampilkan Barang\n2. Menambahkan barang ke cart\n3. Menghapus barang dari cart\n4. Menambahkan Wishlist\n5. Menghapus Wishlist\n6. Menampilkan Wishlist\n7. History Order\n8. TopUp saldo\n9. Pembayaran\n10. Exit")
						inputCustomer := userInput("Masukkan input berdasarkan angka di menu (1-10): ", scanner)
						switch inputCustomer {
						case "1":
							handler.ShowProduct()
							fmt.Print("\nTekan enter untuk kembali ke menu ")
							fmt.Scanln()
						case "2":
							handler.ShowProduct()
							inputCart := userInput("Masukkan ke Cart berdasarkan angka di menu: ", scanner)
							inputCartInt, err := strconv.Atoi(inputCart)
							if err != nil {
								fmt.Println("Input tidak valid!")
								continue
							}
							inputQuantity := userInput("Masukkan quantity: ", scanner)
							inputQuantityInt, err := strconv.Atoi(inputQuantity)
							if err != nil {
								fmt.Println("Input tidak valid!")
								continue
							}
							product := entity.Products{
								ProductDetailId: inputCartInt,
							}
							quantity := entity.Cart{
								Quantity: inputQuantityInt,
							}
							err = handler.AddToCart(user, product, quantity)
							if err != nil {
								fmt.Println(err)
							} else {
								fmt.Println("Berhasil Add to Cart")
								time.Sleep(1 * time.Second)
							}
						case "3":

						case "4":
							handler.ShowProduct()
							inputWishlist := userInput("Masukkan Wishlist berdasarkan angka di menu: ", scanner)
							inputWishlistInt, err := strconv.Atoi(inputWishlist)
							if err != nil {
								fmt.Println("Input tidak valid!")
								continue
							}
							wishlist := entity.Products{
								ProductDetailId: inputWishlistInt,
							}
							err = handler.AddWishlist(user, wishlist)
							if err != nil {
								fmt.Println(err)
							} else {
								fmt.Println("Wishlist berhasil ditambahkan")
								time.Sleep(1 * time.Second)
							}
						case "5":

						case "6":

						case "7":

						case "8":

						case "9":

						case "10":
							return
						default:
							fmt.Println("Input yang dimasukkan tidak valid")
						}
					}
				}
			} else {
				fmt.Println("Login Gagal!")
			}
		case "2":
			fmt.Println("REGISTER")
			email := userInput("Masukkan Email: ", scanner)
			if len(email) < 1 {
				fmt.Println("Input tidak boleh kosong!")
				continue
			}
			_, err = mail.ParseAddress(email)
			if err != nil {
				fmt.Println("Format Email tidak valid!")
				continue
			}
			password := userInput("Masukkan Password: ", scanner)
			if len(password) < 1 {
				fmt.Println("Input tidak boleh kosong!")
				continue
			}
			name := userInput("Masukkan Nama: ", scanner)
			if len(name) < 1 {
				fmt.Println("Input tidak boleh kosong!")
				continue
			}
			address := userInput("Masukkan Alamat: ", scanner)
			if len(address) < 1 {
				fmt.Println("Input tidak boleh kosong!")
				continue
			}
			phoneNumber := userInput("Masukkan Nomor HP: ", scanner)
			if len(phoneNumber) < 1 {
				fmt.Println("Input tidak boleh kosong!")
			}
			regexPattern := `^(0|\+62)\d{9,12}$`
			re := regexp.MustCompile(regexPattern)
			phoneNumberValid := re.MatchString(phoneNumber)
			if !phoneNumberValid {
				fmt.Println("Nomor telepon valid")
				continue
			}
			newUser := entity.User{
				Email:       email,
				Password:    password,
				Name:        name,
				Address:     address,
				PhoneNumber: phoneNumber,
			}
			err = handler.Register(newUser)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Registrasi Berhasil!")
			}
		case "3":
			return
		default:
			fmt.Println("Input yang dimasukkan tidak valid")
		}
	}
}

func userInput(text string, scanner *bufio.Scanner) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}
