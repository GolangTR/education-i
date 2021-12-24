package main

import (
	"fmt"
	"github.com/GolangTR/education-i/restorant"
)

func main()  {
	name := "alameddin"
	surname := "çelik"
	restorant.Selamlama(name, surname)
	gelenHesap, hata := restorant.HesapOdeme(150, 5) // %500 =>
	fmt.Printf("Gelen Hesap main fonksiyonunda %.2f'dir\n", gelenHesap)
	fmt.Printf("%v", hata)
}
