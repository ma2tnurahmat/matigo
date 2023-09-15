# Program Hitungan Mundur dengan Go

Program ini adalah contoh aplikasi sederhana yang memungkinkan Anda memulai hitungan mundur dari waktu yang ditentukan dan kemudian mematikan komputer setelah hitungan mundur selesai.

## Penggunaan
Menggunakan argumen command line
Anda dapat menjalankan program ini dari command line dengan memberikan argumen waktu hitungan mundur dalam detik. Contoh:
```
go run mati.go 60
```
Ini akan memulai hitungan mundur dari 60 detik dan kemudian mematikan komputer setelah hitungan mundur selesai.

## Tanpa argumen command line

Jika Anda tidak memberikan argumen command line, program akan meminta Anda memasukkan waktu hitungan mundur secara interaktif. Contoh:
```
go run mati.go
```
Anda akan diminta memasukkan waktu hitungan mundur dalam detik.

## build
```
go build mati.go
```
