package restorant

import (
	"errors"
	"fmt"
)

func Selamlama(isim string, soyad string){
	fmt.Printf("Hoş Geldin %s %s\n", isim, soyad)
}

// golang'De özel değişkenlerde boş demek için null -> nil kullanılır
// %5 indirim 100 tl hesap => 100 * (1-0.05)0.95 => 95TL
func HesapOdeme(tutar float64, indirimOrani float64)(float64, error)  {
	if indirimOrani < 0 || indirimOrani > 1{
		return tutar, errors.New("İndirim Oranı yanlıştır")
	}
	odenecekTutar := tutar * (1-indirimOrani)
	fmt.Printf("Odenecek Tutar : %.2f dir\n", odenecekTutar)
	return odenecekTutar, nil
}

func HesapBirlestirme(hesap1 float64, hesap2 float64) float64{
}

func MenuListele()  {
}

func KDVEkle(tutar float64) float64 {
}

func IkramEt(ikramEdilenUrunTutari float64, hesap float64)float64{
}