package main

import "net/http"

func main() {
	// buat mux baru
	mux := http.NewServeMux()

	// fungsi-fungsi handler dibuat di handler.go

	// tambahkan handler ke mux
	mux.HandleFunc("GET /products", listProduct)
	mux.HandleFunc("POST /products/", createProduct)
	mux.HandleFunc("PUT /products/{id}", updateProduct)
	mux.HandleFunc("DELETE /products/{id}", deleteProduct)

	// buat http server
	server := &http.Server{
		Addr:    ":8080", // server berjalan di port 8080
		Handler: mux,     // server melayani route yang telah ditentukan oleh mux
	}

	// jalankan server
	server.ListenAndServe()
}
