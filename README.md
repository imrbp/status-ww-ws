# Auto Reload Data Project

Buatlah sebuah microservice untuk meng-update sebuah file json setiap 15 detik dengan angka random antara 1-100 untuk valuewater dan wind. Seperti berikut:
```json
{
  "status" : {
    "water" : 6,
    "wind" : 10
  }
}
```


Kemudian buatlah sebuah halaman untuk menampilkan status tersebut. 

Selain itu kalian harus menentukan statuswater dan wind tersebut. Dengan ketentuan: 
- jika water dibawah 5 maka status aman
- jika water antara 6 - 8 maka status siaga
- jika water diatas 8 maka status bahaya
- jika wind dibawah 6 maka status aman●jika wind antara 7 - 15 maka status siaga
- jika wind diatas 15 maka status bahaya●value water dalam satuan meter

value wind dalam satuan meter per detikBuatlah halaman tersebut semenarik mungkin dengan data selalu up-to-date (auto reload)!!


# Desckripsi Program

Program ini menggunakan go-fiber websocket untuk auto reloadnya. untuk mengaksesnya dapat ke localhost:3000. 

Di CLI ada juga pemberitauan setiap detiknya bahwa file json dibaca. dan setiap 15 detik program akan menulis ke file-nya. 