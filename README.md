# Enigma Laundry - Go CLI Application

Enigma Laundry adalah aplikasi Command Line Interface (CLI) sederhana berbasis bahasa Go (Golang). Aplikasi ini digunakan untuk mengelola data pelanggan, layanan laundry, dan pemesanan laundry.

---

## Requirements

-   Go 1.20+
-   PostgreSQL (atau sesuaikan dengan database SQL lain jika diperlukan)

---

## Database Schema (DDL)

Gunakan script SQL yang berada di dalam folder/file `db/ddl.sql` untuk membuat struktur database.

---

## Cara Menjalankan Aplikasi

1. Clone repositori ini:

```bash
git clone <repository_url>
cd enigma-laundry
```

2. Edit file `config/config.go` atau sejenisnya untuk menyesuaikan koneksi database:

```go
const (
    DB_USER     = "postgres"
    DB_PASSWORD = "yourpassword"
    DB_NAME     = "enigma_laundry"
    DB_HOST     = "localhost"
    DB_PORT     = "5432"
)
```

3. Jalankan aplikasi:

```bash
go run main.go


## Menu Utama

1. Customer
2. Service
3. Order
4. Exit

```

### Menu Customer

-   Create Customer
-   View List Customer
-   View Details Customer By ID
-   Update Customer
-   Delete Customer

### Menu Service

-   Create Service
-   View List Service
-   View Details Service By ID
-   Update Service
-   Delete Service

### Menu Order

-   Create Order
-   Complete Order
-   View List Order
-   View Order Details By ID

**Note:** Semua entri order akan mencakup order detail di dalam menu order.

---

## Validasi Penting

### Customer:

-   Tidak boleh membuat customer dengan ID yang sudah ada
-   Tidak bisa menghapus customer jika sedang digunakan di orders

### Service:

-   Tidak boleh membuat service dengan ID yang sudah ada
-   Tidak bisa menghapus service jika sedang digunakan di orders

### Order:

-   Tidak bisa membuat order dengan customer ID yang tidak ada
-   Tidak bisa membuat order dengan order ID yang sudah ada
-   Tidak bisa menyelesaikan atau melihat order yang tidak ditemukan

---

## Arsitektur

-   `entity/` - struct dasar (model)
-   `repository/` - interaksi dengan database
-   `service/` - logika bisnis
-   `controller/` - menu CLI
-   `main.go` - integrasi dan pemanggilan menu utama

---

## Lisensi

Free to use for learning purposes only.

---

Dibuat oleh Restu Adi Wahyujati sebagai bagian dari pembelajaran Golang Database.
