# To-Do App

Yavuzlar 3. hafta ödevi kapsamında geliştirilen yapılacaklar listesi uygulaması.

İlk hafta yapılan HTML/CSS/JS tabanlı frontend, bu haftada Go backend ve PostgreSQL veritabanı ile birleştirildi. Tüm sistem Docker üzerinde çalışıyor.

## Teknolojiler

- **Backend:** Go (net/http + GORM)
- **Frontend:** HTML, CSS, JavaScript
- **Veritabanı:** PostgreSQL
- **Auth:** JWT
- **Container:** Docker & Docker Compose

## Nasıl Çalıştırılır

```bash
docker-compose up --build
```

Uygulama `http://localhost:3000` adresinde açılır.

## API Endpoint'leri

Base path: `/api/v1`

| Method | Endpoint | Açıklama | Auth |
|--------|----------|----------|:----:|
| POST | `/api/v1/auth/register` | Kayıt ol | - |
| POST | `/api/v1/auth/login` | Giriş yap | - |
| GET | `/api/v1/todos` | Görevleri listele | ✓ |
| POST | `/api/v1/todos` | Yeni görev ekle | ✓ |
| PUT | `/api/v1/todos/{id}` | Görev güncelle | ✓ |
| DELETE | `/api/v1/todos/{id}` | Görev sil | ✓ |

Auth gereken endpoint'lere istek atarken header'a `Authorization: Bearer <token>` eklenmeli.

## Proje Yapısı

```
├── backend/
│   ├── handlers/        # API handler'ları
│   ├── middleware/       # Auth ve CORS middleware
│   ├── models/          # Veritabanı modelleri
│   ├── database/        # DB bağlantısı
│   ├── utils/           # JWT işlemleri
│   ├── main.go
│   └── dev.backend.Dockerfile
├── frontend/
│   ├── main.html
│   ├── style.css
│   ├── nginx.conf
│   └── dev.frontend.Dockerfile
└── docker-compose.yaml
```

## Notlar

- Şifreler bcrypt ile hashlenir.
- JWT token süresi 24 saattir.
- Veritabanı verileri Docker volume sayesinde container yeniden başlatılsa bile korunur.
- Frontend tarafında localStorage sadece JWT token saklamak için kullanılır, görev verileri backend'den çekilir.
