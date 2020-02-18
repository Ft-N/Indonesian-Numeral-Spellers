# Indonesian Numeral Spellers

## Latar Belakang
Mengeja angka merupakan salah satu kegiatan dasar yang dilakukan setiap harinya. Contoh kegiatan tersebut yaitu mengeja harga barang, nilai data, tanggal dan tahun, serta masih banyak lagi. Meskipun terkesan hal sepele, berdasarkan penelitian dari para dokter di Indonesia, seorang anak baru bisa membaca dan mengeja angka pada umur 4-6 tahun. Rentang usia tersebut tentunya terasa kurang cepat. Padahal, semakin cepat seorang bisa membaca dan mengeja angka, maka semakin cepat pula anak tersebut dapat belajar berhitung dan mempelajari hal-hal lainnya, bahkan termasuk belajar pemrograman.

Dari permasalah di atas, maka diperlukanlah suatu sarana pembelajaran yang dapat membantu anak-anak balita di Indonesia untuk membaca dan mengeja angka. Dengan adanya solusi tersebut, diharapkan anak-anak dapat membaca dan mengeja angka lebih cepat sehingga mampu segera mempelajari hal-hal lebih besar lainnya dan tentunya meningkatkan tingkat pendidikan di Indonesia.

## Spesifikasi
Buatlah dalam bahasa pemrograman **_Go_**, sebuah web service berupa **_REST API_**, yang dapat mengeja (dalam bahasa Indonesia) dari angka yang diberikan serta menuliskan angka yang tepat dari masukkan ejaan angka (dalam bahasa Indonesia juga), dengan ketentuan-ketentuan sebagai berikut :

1. Terdapat 2 buah endpoint API yang perlu dibuat, yaitu '**GET** /spell' yang menerima parameter angka, serta '**POST** /read' yang menerima body/payload berupa text/ejaan. Jika input parameter atau body/payload tidak valid, maka berikan response keterangan error/gagal dengan format dibebaskan.

2. Sebagai REST API, maka response harus berupa JSON. Struktur data response JSON dibebaskan.

3. Program dibuat dengan mengikuti standar development resmi Go (lihat referensi #3), yaitu environment kode program berada pada ```$GOPATH/src/```, misalkan ```$GOPATH/src/github.com/Indonesian-Numeral-Spellers```.

4. Arsitektur program dibebaskan (boleh mengikuti referensi-referensi _REST API with Go_ dari internet), namun harus tetap tersusun dengan rapi dan mengerti apa kegunaan setiap fungsi, file, serta package.

5. Batasan kasus uji : 2000000000 (dua milyar)

## Contoh Kasus Uji
### Contoh Kasus Uji 1 : Pengejaan
Request :
```
GET '/spell?number=123456'
```
Response :
```JSON
STATUS CODE 200
{
    "status" : "OK",
    "text" : "seratus dua puluh tiga ribu empat ratus lima puluh enam"
}
```
### Contoh Kasus Uji 2 : Pembacaan
Request:
```JSON
POST '/read'
{
    "text" : "seribu sembilan ratus sembilan puluh tujuh"
}
```
Response :
```JSON
STATUS CODE 200
{
    "status" : "OK",
    "number" : 1997
}
```

## Referensi Pengerjaan _(sangat disarankan untuk diikuti dengan baik)_
1. https://golang.org/doc/install
2. https://github.com/golang/go/wiki/SettingGOPATH
3. https://golang.org/doc/code.html#Introduction
4. https://tour.golang.org/welcome/1
5. https://openclassrooms.com/en/courses/3432056-build-your-web-projects-with-rest-apis/3496011-identify-examples-of-rest-apis
6. https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
