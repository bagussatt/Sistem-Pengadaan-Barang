# Fleetify - Purchasing Management System

Sistem manajemen purchasing untuk mengelola items, suppliers, dan purchase orders dengan integrasi webhook notification.

## Spesifikasi

### Backend
- **Framework**: Go dengan Fiber v2
- **Database**: PostgreSQL 
- **ORM**: GORM
- **Authentication**: JWT (HMAC-SHA256)
- **API Documentation**: Swagger UI

### Frontend
- **HTML5**, **CSS3**, **JavaScript **
- **jQuery** 3.7.1
- **Bootstrap** 5.3.3
- **SweetAlert2** untuk notifikasi

### Fitur Utama
1. **Authentication**: Register & Login dengan JWT token
2. **Item Management**: CRUD items dengan stock tracking
3. **Supplier Management**: CRUD suppliers
4. **Purchase Management**: 
   - Create purchase order
   - Auto update stock setelah purchase
   - Delete purchase (restore stock otomatis)
   - Webhook notification ke URL eksternal

## Prasyarat

### Software yang Diperlukan
- Go 1.21 atau lebih tinggi
- PostgreSQL 13 atau lebih tinggi
- Docker & Docker Compose (untuk database)
- Browser modern (Chrome, Firefox, Edge)

## Cara Menjalankan Aplikasi

### 1. Setup Database

Jalankan PostgreSQL menggunakan Docker Compose:
```bash
cd purchasing-backend
docker-compose up -d
```

Atau gunakan PostgreSQL yang sudah terinstall.

### 2. Konfigurasi Environment

Buat file `.env` di folder `purchasing-backend`:
```env
APP_PORT=3000

WEBHOOK_URL=https://webhook.site/id-webhoook

DB_HOST=localhost
DB_PORT=5433
DB_USER=appuser
DB_PASSWORD=apppassword
DB_NAME=purchasing_db

JWT_SECRET=your-secret-key-here
```

### 3. Jalankan Backend Server

```bash
cd purchasing-backend

# Build server
go build -o server.exe ./cmd/server

# Jalankan server
./server.exe
```

Server akan berjalan di `http://localhost:3000`

### 4. Akses Frontend

Buka file HTML berikut di browser:
- Login: `frontend/login.html`
- Register: `frontend/register.html`
- Dashboard: `frontend/dashboard.html`
- Purchase: `frontend/purchase.html`

### 5. Generate Swagger Documentation (Opsional)

```bash
cd purchasing-backend
swag init -g cmd/server/main.go -o docs
```

Swagger UI tersedia di: `http://localhost:3000/swagger/index.html`

## API Documentation

### Authentication Endpoints

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/register` | Register user baru |
| POST | `/api/login` | Login dan dapatkan JWT token |

### Item Endpoints

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/items` | Get semua items |
| POST | `/api/items` | Create item baru |
| GET | `/api/items/:id` | Get item by ID |
| PUT | `/api/items/:id` | Update item |
| DELETE | `/api/items/:id` | Delete item |

### Supplier Endpoints

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/suppliers` | Get semua suppliers |
| POST | `/api/suppliers` | Create supplier baru |
| GET | `/api/suppliers/:id` | Get supplier by ID |
| PUT | `/api/suppliers/:id` | Update supplier |
| DELETE | `/api/suppliers/:id` | Delete supplier |

### Purchasing Endpoints

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/purchasings` | Get semua purchases |
| POST | `/api/purchasings` | Create purchase order |
| GET | `/api/purchasings/:id` | Get purchase by ID |
| DELETE | `/api/purchasings/:id` | Delete purchase |

## Contoh Penggunaan

### 1. Register User

```bash
curl -X POST http://localhost:3000/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "password123",
    "role": "admin"
  }'
```

### 2. Login

```bash
curl -X POST http://localhost:3000/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "password123"
  }'
```

Response:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 3. Create Item (dengan Token)

```bash
curl -X POST http://localhost:3000/api/items \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{
    "name": "Laptop ASUS",
    "stock": 10,
    "price": 15000000
  }'
```

### 4. Create Purchase

```bash
curl -X POST http://localhost:3000/api/purchasings \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{
    "supplier_id": 1,
    "items": [
      {"item_id": 1, "qty": 5}
    ]
  }'
```

## Struktur Database

### users
- id (Primary Key)
- username (Unique)
- password (Hashed)
- role (admin/user/staff)
- created_at

### suppliers
- id (Primary Key)
- name
- email
- address
- created_at

### items
- id (Primary Key)
- name
- stock
- price
- created_at

### purchasings
- id (Primary Key)
- date
- supplier_id (Foreign Key)
- user_id (Foreign Key)
- grand_total
- created_at

### purchasing_details
- id (Primary Key)
- purchasing_id (Foreign Key)
- item_id (Foreign Key)
- qty
- sub_total

## Catatan

- Stock item akan otomatis berkurang saat purchase dibuat
- Stock item akan otomatis dikembalikan saat purchase dihapus
- Webhook notification dikirim secara asynchronous setelah purchase berhasil
- JWT token expire dalam 24 jam
- Pastikan PostgreSQL berjalan sebelum menjalankan backend server

## Troubleshooting

### Server tidak bisa start
1. Pastikan PostgreSQL berjalan
2. Cek koneksi database di file `.env`
3. Pastikan port 3000 tidak digunakan aplikasi lain

### Frontend tidak bisa akses API
1. Pastikan backend server berjalan
2. Cek browser console untuk error message
3. Pastikan sudah login dan token valid

### Port sudah digunakan
Windows:
```bash
netstat -ano | findstr :3000
taskkill /F /PID [PID_NUMBER]
```

Linux/Mac:
```bash
lsof -ti:3000 | xargs kill -9
```
