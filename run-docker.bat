@echo off
echo ==================================================
echo           TES-GO API dengan Docker
echo ==================================================
echo.
echo Memulai aplikasi menggunakan Docker Compose...
echo Ini akan men-download image Go dan PostgreSQL secara otomatis.
echo.
echo Pastikan Docker Desktop sudah berjalan!
echo.
pause

echo Starting containers...
docker-compose up --build

echo.
echo Aplikasi berjalan di: http://localhost:8080
echo.
echo Untuk menghentikan: Ctrl+C atau jalankan 'docker-compose down'
pause
