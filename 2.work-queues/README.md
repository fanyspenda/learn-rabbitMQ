# Project 2 - Work Queues

referensi: https://www.rabbitmq.com/tutorials/tutorial-two-go.html

## Deskripsi Projectnya:
Akan ada 2 server penerima, dan 1 server pengirim. RabbitMQ akan secara otomatis membagi load pekerjaan server penerima. Jika penerima 1 sedang bekerja, maka rabbitMQ akan mengirim ke penerima 2.
> **note:**
> 
> Consumer (penerima) harus berjalan terlebih dahulu sebelum producer (pengirim) mengirimkan pesan. Jika tidak, maka producer hanya akan dihandle oleh 1 consumer yang aktif pertama kali.
> 
> Kecuali jika melakukan 2 kali publish message, maka 2 server penerima akan tetap dapat bekerja bersama.

## Langkah Menjalankan Project:
1. jalankan 2 consumer di 2 terminal berbeda
2. jalankan producer untuk mengirim beberapa pesan dengan parameter (Ex: `go run new_task.go message4......`). Setiap titik (.) mewakilkan 1 detik waktu dibutuhkan untuk menyelesaikan tugas.
3. Consumer 1 dan 2 akan bekerja bersama-sama menyelesaikan pesan yang dikirimkan. Misal, ada 4 pesan diproses. Maka, 4 pesan itu akan dibagi-bagi oleh RabbitMQ ke beberapa consumer agar task lebih cepat selesai (_Round-robin dispatching_).

## Acknowledge
Pada RabbitMQ, proses yang terjadi adalah:
1. RabbitMQ mengirim pesan ke consumer.
2. Consumer menerima pesan.
3. Consumer mengirimkan `acknowledge` ke server RabbitMQ sebagai penanda bahwa pesan berhasil diterima. Jika consumer tidak mengirimkan `acknowledge`, maka RabbitMQ mengira bahwa pesan tidak diterima dengan baik. Sehingga, pesan akan dimasukkan kembali ke daftar antrian (*re-queue*).

Jika menggunakan auto-acknowledge, maka acknowledge akan dikirimkan seketika ketika pesan berhasil diterima oleh consumer. Jika pesan sudah diterima tapi gagal diproses di tengah jalan (misal: server dihentikan `Ctrl + C`), maka pesan akan hilang. Oleh karena itu, kita bisa menggunakan manual acknowledge dan baru mengirimkan kembali ke server rabbitMQ ketika kita selesai memproses pesan (lihat komen di `receiver.go`).

## Message Durability
Mencegah pesan hilang dari Queue ketika server rabbitMQ mati. Langkah membuat pesan menjadi durable adalah dengan mengubah queue, durability menjadi `true`, lalu mengubah `amqp.Publishing.DeliveryMode` ke mode `amqp.presistent` (hanya untuk kasus menggunakan `amqp`).

> **Note:**
> 
> RabbitMQ tidak 100% menjamin pesan akan tersimpan dengan aman karena masih ada celah antara RabbitMQ menerima pesan dan menyimpannya ke disk. Jika ingin jaminan pesan tersimpan lebih aman, bisa menggunakan **publisher confirms**.

## Membatasi agar server hanya bisa menerima max. n message (Fair Dispatch)
RabbitMQ bekerja dengan memberikan semua pesan ke tiap server dengan adil.

Jika ada 2 server (A dan B), dan ada 100 message, maka server A dan B masing-masing akan menerima 50 message dalam 1 waktu. Bagaimana jika server A lebih lambat dari server B?

Masalah ini dapat diatasi dengan `ch.Qos`, dimana fungsi ini dapat membatasi berapa message yang bisa diterima oleh subscriber dalam 1 waktu. Sehingga, kita bisa membuat server A menerima hanya 30 message dalam 1 waktu, dan server B sisanya.