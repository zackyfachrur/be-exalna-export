# Gunakan base image golang
FROM golang:1.23.0

# Tentukan working directory di dalam container
WORKDIR /app

# Salin seluruh isi folder chat-services ke dalam container
COPY . .

# Jalankan aplikasi Go
CMD ["go", "run", "."]
