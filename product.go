package main

// database adalah variable yang akan menampung data produk
// ini adalah contoh data produk yang disimpan di memory tanpa server database
// map digunakan untuk mempermudah pencarian data produk berdasarkan ID
var database = map[int]Product{}

// lastID adalah variable yang akan menyimpan ID terakhir dari produk yang ada di map products
// ini digunakan untuk memberikan ID unik pada produk yang baru ditambahkan
var lastID = 0

// Product adalah struct yang merepresentasikan sebuah produk
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
