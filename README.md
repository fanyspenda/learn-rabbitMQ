# Belajar RabbitMQ

RabbitMQ = message broker = makelar pesan = tukang ngirimin pesan

Artinya, pesan yang akan dikirimkan dari 1 service ke service lain, bakal dikirimin lewat RabbitMQ.

## Analoginya:
1. Kalau service 1 ngirim pesan langsung ke service lain, itu sama kayak kita nganterin surat langsung ke rumah temen kita.
2. Kalau pakai message broker, ibarat kita minta tolong ke pak pos buat ngirimin pesan kita ke temen kita.

## Kenapa harus pake Message Broker?
### 1. Menjaga pesan agar tetap ada jika service yang dituju mati.
Jadi, dari video yang saya tonton, _message broker_ punya kemampuan untuk menyimpan pesan yang dikirimkan suatu service ke service lain.

Misal ada service A dan B. Suatu ketika, A mengirim ke B, tapi si B mati. Kalau tanpa Message Broker, pesan yang dibawa dari A akan hilang dan ngereturn Error. Kalau pakai Message Broker, pesan yang dibawa akan disimpen oleh Message Broker dan akan langsung dikirimkan ke si B begitu si B hidup lagi.

### 2. Menyeimbangkan workload antar server.
misal A berkemampuan 3 proses/detik, sedangkan si B sanggup 1 proses/detik. Kalau tanpa Message Broker, A bakal jadi bottleneck karena masih harus nungguin response dari service B. Kalau dengan Message Broker, service A tetep bakal jalan karena nggk perlu nungguin si B.

### 3. Ketergantungan 1 service ke banyak service lain.
Kalau tanpa Message Broker, service A bakal error ketika B juga error. Bayangin kalo udh ketergantungan sama banyak service. 1 service mati ajah bisa error semuanya.

Kalau dengan Message Broker, service A bakal tetep jalan meski service B nggk halan.

## Negative Side
### 1. Ketergantungan sama Message Broker
Memang service kita jadi nggak ketergantungan sama service lain, tapi jadinya ketergantungan sama Message Broker. Message Broker mati, semua service mati.

### 2. Merubah paradigma berpikir kita.
Ini belum aku riset lebih jauh. Tapi, karena kita nggk bergantung dengan service lain, artinya kita nggk bisa mengandalkan response dari service yang kita tembak. Melainkan, mengandalkan response yang diberikan RabbitMQ, ya kan? Tapi, kalo salah, mohon koreksinya ajah, bikin issue sana wkwk..

## Istilah di RabbitMQ
### 1. _Producing_
mengirim pesan. Pengirimnya disebut _Producer_.
### 2. _Queue_
Tempat dimana pesan disimpan. 1 Queue bisa berisi banyak pesan dari macam-macam _producer_.
### 3. _Consuming_
Menerima pesan. Penerimanya disebut _Consumer_.



