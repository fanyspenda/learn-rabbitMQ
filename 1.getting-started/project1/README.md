# Project 1 - getting started

referensi: https://www.rabbitmq.com/tutorials/tutorial-one-go.html

## Deskripsi Projectnya:
Jadi, akan ada 2 server yang berjalan, yaitu `send.go`, dan `receive.go` dimana `send.go` akan mengirimkan pesan ke receive.go melalui RabbitMQ. 

## Langkah Menjalankan Project:
1. Jalankan `send.go` dengan perintah `go run send.go`
2. Akan ada log bahwa pesan telah dipublish (dikirim). Namun, karena server penerima belum nyala, maka RabbitMQ akan menyimpan pesan tersebut untuk dikirim ke penerima setelah server nyala. Untuk check berapa pesan yang sedang dalam Queue, gunakan perintah `sudo rabbitmqctl list_queues`.
3. Jalankan `receive.go` dengan perintah `go run receive.go`. Maka, pesan yang dikirimkan dari `send.go` akan masuk ke terminal `receive.go`.