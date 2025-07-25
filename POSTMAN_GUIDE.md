# 📮 Postman Collection - Tes-Go API

## 📥 **Cara Import ke Postman:**

### **Method 1: Import File (Recommended)**

1. **Download file collection:**
   - `Tes-Go-API.postman_collection.json`
   - `Tes-Go-API.postman_environment.json`

2. **Buka Postman Desktop/Web**

3. **Import Collection:**
   - Klik **"Import"** di pojok kiri atas
   - Drag & drop file `Tes-Go-API.postman_collection.json`
   - Atau klik **"Upload Files"** dan pilih file tersebut

4. **Import Environment:**
   - Klik **"Import"** lagi
   - Upload file `Tes-Go-API.postman_environment.json`
   - Pilih environment **"Tes-Go API Environment"** di dropdown kanan atas

### **Method 2: Share Link (Untuk Tim)**

1. **Upload collection ke GitHub repository Anda**
2. **Share raw link file JSON:**
   ```
   https://raw.githubusercontent.com/USERNAME/REPO/main/Tes-Go-API.postman_collection.json
   ```

3. **Tim bisa import dengan:**
   - Postman → Import → Link → Paste URL

## 🔧 **Setup Environment:**

### **Untuk Testing Lokal:**
- Set `base_url` = `http://localhost:8080`

### **Untuk Testing di Codespaces:**
- Set `base_url` = `https://your-codespace-8080.app.github.dev`
- Ganti `your-codespace` dengan URL Codespaces Anda

## 🧪 **Cara Testing:**

### **1. Sequential Testing (Otomatis):**
1. Klik collection **"Tes-Go API Collection"**
2. Klik **"Run"** 
3. Pilih semua request
4. Klik **"Run Tes-Go API Collection"**
5. Semua test akan jalan otomatis dengan urutan yang benar

### **2. Manual Testing:**
1. **Jalankan "Health Check"** dulu
2. **Jalankan "Login"** - token akan tersimpan otomatis
3. **Jalankan request lainnya** - token otomatis terpakai

## ✅ **Features Collection:**

### **✨ Automatic Token Management:**
- Login otomatis menyimpan JWT token
- Semua request authenticated otomatis pakai token tersimpan

### **🧪 Automated Tests:**
- Status code validation
- Response body validation  
- Error case testing
- Success flow testing

### **📋 Complete API Coverage:**
1. **Health Check** - Test API status
2. **Login** - Get JWT token
3. **Create Terminal** - Test authentication & terminal creation
4. **Get All Terminals** - Test data retrieval
5. **Error Cases** - Test invalid login, unauthorized access, duplicate data

### **🌍 Environment Support:**
- Localhost testing
- Codespaces testing
- Easy switching antar environment

## 🚀 **Sharing dengan Tim:**

### **Option 1: File Sharing**
```bash
# Push ke repository GitHub
git add Tes-Go-API.postman_collection.json Tes-Go-API.postman_environment.json
git commit -m "Add Postman collection for API testing"
git push
```

### **Option 2: Postman Workspace**
1. Buat Postman Team Workspace
2. Publish collection ke workspace
3. Invite anggota tim

### **Option 3: Direct Share**
1. Di Postman, klik collection → Share
2. Copy share link
3. Share link ke tim

## 📖 **Documentation:**

Collection ini include:
- **Pre-request Scripts** untuk setup
- **Test Scripts** untuk validation
- **Example responses** 
- **Detailed descriptions** untuk setiap endpoint

## 🎯 **Testing Workflow:**

```
1. Health Check → ✅ API Running
2. Login → ✅ Get Token  
3. Create Terminal → ✅ Auth Working
4. Get Terminals → ✅ Data Retrieved
5. Error Tests → ✅ Error Handling
```

**Ready to test!** 🎉


