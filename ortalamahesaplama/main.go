package ortalamahesaplama

// Vize Notu 0 ile 100 arasında olabilir //
// Final Notu 0 ile 100 arasında olabilir //
// İnsiyatif 0 ile 1 arasında olabilir //
// Ortalama 0 ile 100 arasında olabilir //

// Ortalama 65'ten küçükse ve ortalama*(1+insiyatif) 65'ten büyükse ortalama 65 olmalı //tamam
// Ortalama 65'ten büyükse ortalama değişmemelidir // tamam
// Çıktı olarak ortalama dönmelidir // tamam
// OrtalamaHesapla fonksiyonu `go test ./...` komutu ile test edilmelidir
// Testlerin hepsi PASS olmalıdır

func OrtalamaHesapla(vizeNot, finalNot int, insiyatif float64) float64 {
	ortalama := (float64(vizeNot) * 0.4) + (float64(finalNot) * 0.6)
	_insiyatif := ortalama * (1 + insiyatif)

	if ortalama < 65 && _insiyatif > 65 {
		ortalama = 65
		return ortalama
	} else if ortalama > 65 {
		return ortalama // test

	}

	return ortalama
}
