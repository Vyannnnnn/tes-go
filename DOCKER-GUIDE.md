# Menjalankan API Tanpa Install Go & PostgreSQL

## 🐳 Menggunakan Docker (Rekomendasi)

Cara paling mudah untuk menjalankan aplikasi tanpa install Go dan PostgreSQL secara lokal.

### Prasyarat
- **Docker Desktop** untuk Windows
  - Download: https://www.docker.com/products/docker-desktop/
  - Install dan pastikan Docker Desktop berjalan

### Cara Menjalankan

1. **Buka Command Prompt atau PowerShell**
   ```bash
   cd "c:\Users\user\Downloads\Tes-Go"
   ```

2. **Jalankan dengan Docker Compose**
   ```bash
   docker-compose up --build
   ```
   
   Atau double-click file `run-docker.bat`

3. **Tunggu proses selesai**
   - Docker akan download image Go dan PostgreSQL
   - Build aplikasi Go dalam container
   - Start database dan API server

4. **Test API**
   - API berjalan di: `http://localhost:8080`
   - Database PostgreSQL di: `localhost:5432`

### Perintah Docker Berguna

```bash
# Jalankan aplikasi
docker-compose up -d

# Lihat logs
docker-compose logs -f api

# Stop aplikasi
docker-compose down

# Hapus semua data (termasuk database)
docker-compose down -v
```

## 🌐 Alternatif Lain

### 1. **Online IDE (Cloud)**
- **Gitpod**: https://gitpod.io/
- **CodeSpaces**: https://github.com/codespaces
- **Replit**: https://replit.com/

Upload kode ke GitHub, lalu buka di online IDE. Environment Go dan PostgreSQL sudah tersedia.

### 2. **PostgreSQL Online**
Gunakan database online gratis:
- **ElephantSQL**: https://www.elephantsql.com/ (gratis)
- **Supabase**: https://supabase.com/ (gratis)
- **Neon**: https://neon.tech/ (gratis)

Ganti connection string di `main.go`:
```go
connStr := "postgresql://username:password@hostname:port/database?sslmode=require"
```

### 3. **Go Playground** (Terbatas)
- https://go.dev/play/
- Hanya untuk testing kode sederhana
- Tidak bisa connect ke database external

## 📱 Testing API dengan Docker

Setelah Docker berjalan, test dengan curl atau Postman:

```bash
# Login
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d "{\"card_number\":\"1234567890\",\"password\":\"password123\"}"

# Create Terminal (ganti YOUR_TOKEN dengan token dari login)
curl -X POST http://localhost:8080/terminals -H "Content-Type: application/json" -H "Authorization: Bearer YOUR_TOKEN" -d "{\"name\":\"Terminal Utama\",\"code\":\"TRM001\",\"location\":\"Jakarta Pusat\"}"
```

## ✅ Keuntungan Docker

1. **Tidak perlu install Go dan PostgreSQL**
2. **Environment yang konsisten**
3. **Mudah deploy ke production**
4. **Isolasi aplikasi**
5. **One-click setup**

## 🚀 Rekomendasi

**Untuk development**: Gunakan Docker
**Untuk production**: Deploy ke cloud (Heroku, Railway, Render, dll)

Docker adalah solusi terbaik karena:
- Setup sekali klik
- Environment yang sama dengan production
- Tidak "mengotori" sistem lokal
- Mudah di-share dengan tim
