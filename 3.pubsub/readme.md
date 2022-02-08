# Publish/Subscribe

Jika di project sebelumnya, konsepnya adalah ada 1 pengirim -->  1 Queue --> 1 penerima (1 publisher, 1 JENIS receiver). Dalam pubsub, 1 publisher, bisa memiliki lebih dari 1 receiver atau biasa disebut _subscriber_.

Pada dasarnya, pesan yang dipublish akan dibroadcast ke semua subscriber.

## Exchange
Jadi, sebenernya di rabbitMQ, publisher nggk pernah ngirim pesan secara langsung ke queue, melainkan ada penengah diantaranya, yang disebut **exchange**.

dalam exchange ini, dia harus tahu messagenya harus diapakan (dikirim ke Queue A, B, atau dibuang)