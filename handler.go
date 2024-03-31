package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// listProduct akan menampilkan semua produk
func listProduct(w http.ResponseWriter, r *http.Request) {
	// ubah database menjadi array produk
	products := []Product{}
	for _, product := range database {
		products = append(products, product)
	}

	// ubah array produk menjadi JSON
	data, err := json.Marshal(products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Gagal mengubah data produk menjadi JSON"))
		return
	}

	// set header content-type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// kirim response
	w.Write(data)
}

// createProduct akan menambahkan produk baru
func createProduct(w http.ResponseWriter, r *http.Request) {
	// ambil nama dan harga produk dari json request
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Gagal membaca data request produk"))
		return
	}

	// tambahkan 1 pada ID produk terakhir
	lastID++

	// gunakan ID terakhir untuk produk yang baru
	product.ID = lastID

	// simpan produk ke database
	database[product.ID] = product

	// ubah produk menjadi JSON
	data, err := json.Marshal(product)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Gagal mengubah data produk menjadi JSON"))
		return
	}

	// set header content-type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// kirim response
	w.Write(data)
}

// updateProduct akan mengubah data produk
func updateProduct(w http.ResponseWriter, r *http.Request) {
	// ambil ID produk dari URL
	productID := r.PathValue("id") // "id" sesuai dengan penulisan di route yaitu {id}

	// ubah menjadi integer
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID produk harus berupa angka"))
		return
	}

	// cek apakah ID produk tersedia di database
	product, ok := database[productIDInt]
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Produck tidak ditemukan"))
		return
	}

	// jika ada, baca nama dan harga baru
	var newProduct Product
	err = json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Gagal membaca data request produk"))
		return
	}

	// jika tidak ada nama atau harga baru, kirim response bad request
	if newProduct.Name == "" && newProduct.Price == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Nama atau harga produk harus diisi"))
		return
	}

	// ubah nama produk jika ada
	if newProduct.Name != "" {
		product.Name = newProduct.Name
	}

	// ubah harga produk jika ada
	if newProduct.Price != 0 {
		product.Price = newProduct.Price
	}

	// simpan data produk yang telah diubah
	database[productIDInt] = product

	// ubah data produk menjadi JSON
	data, err := json.Marshal(product)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Gagal mengubah data produk menjadi JSON"))
		return
	}

	// set header content-type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// kirim response
	w.Write(data)
}

// deleteProduct akan menghapus produk
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// ambil ID produk dari URL
	productID := r.PathValue("id") // "id" sesuai dengan penulisan di route yaitu {id}

	// ubah menjadi integer
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID produk harus berupa angka"))
		return
	}

	// cek apakah ID produk tersedia di database
	_, ok := database[productIDInt]
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Produck tidak ditemukan"))
		return
	}

	// hapus produk dari database
	delete(database, productIDInt)

	// kirim header status no content
	w.WriteHeader(http.StatusNoContent)

	// tidak perlu kirim response body
}
