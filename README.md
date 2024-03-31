# RESTfult API Example
## Cara Menjalankan
1. Clone repository ini
```bash
git clone git@github.com:fastcampus-backend-golang/restful-api-example.git
```

2. Masuk ke direktori
```bash
cd restful-api-example
```

3. Jalankan dengan perintah
```bash
go run .
```

## Konten
- main.go : file utama yang berisi konfigurasi server dan routing
- product.go : file yang berisi model produk
- handler.go : file yang berisi fungsi-fungsi yang akan dijalankan ketika ada request

## Route
- GET /products: mengambil daftar produk yang tersedia
- POST /products: membuat sebuah produk baru
- PUT /product/{id}: memperbarui informasi sebuah produk
- DELETE /product/{id}: menghapus sebuah produk
 