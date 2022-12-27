## Petunjuk Instalasi

```bash
$ make proto
```

## Buat 3 database baru `auth_svc`, `product_svc`, dan `order_svc`

## Menjalankan aplikasi

```bash
$ make server
```

## Menjalankan semua service menggunakan Makefile
Pertama jalankan [API gateway](https://github.com/allrested/product/api-gateway)
Kedua jalankan [Auth service](https://github.com/allrested/product/auth-svc)
Ketiga jalankan [Product service](https://github.com/allrested/product/product-svc)
Keempat jalankan [Order service](https://github.com/allrested/product/order-svc)

## Curl commands
Register User
```bash
curl --request POST --url http://localhost:3000/auth/register --header 'Content-Type:application/json' --data '{"email": "allrested@gmail.com","password": "admin"}'
```
Login User
```bash
curl --request POST --url http://localhost:3000/auth/login --header 'Content-Type:application/json' --data '{"email": "allrested@gmail.com","password": "admin"}'
```
Create Product
```bash
curl --request POST --url http://localhost:3000/product --header 'Authorization: Bearer YOUR_TOKEN' --header 'Content-Type: application/json' --data '{ "name": "Product A", "stock": 5, "price": 15}'
```
Find One Product
```bash
curl --request POST --url http://localhost:3000/product/1 --header 'Authorization: Bearer YOUR_TOKEN'
```
Create Order
```bash
curl --request POST --url http://localhost:3000/order --header 'Authorization: Bearer YOUR_TOKEN' --header 'Content-Type:application/json' --data '{"productId": "1","quantity": "1"}'
```
