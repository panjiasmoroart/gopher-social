# Tech Stack Using

- go1.23.1
- Docker
- Postgres running on Docker
- Swagger for docs
- Golang migrate for migrations

## Directory Structure

- **bin/**: pada dasarnya file ini akan berisi aplikasi kompilasi kita.
- **cmd/**: tempat di mana kita akan memiliki executable atau titik masuk aplikasi yang utama (didalamnya ada API, Migration dll).
- **internal/**: untuk menampung semua paket internal, yang artinya tidak untuk diekspor, untuk digunakan oleh server API kita (auth, db, env, mailer, ratelimiter, store dll).
- **docs/**: digunakan untuk membuat dokumentasi api menggunakan tools swagger or postman
- **scripts/**: digunakan untuk menyiapkan server (setting up server)
- **web/**: file ini akan berisi web application frontend seperti reactjs, vuejs dll
