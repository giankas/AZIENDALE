# Usa un'immagine di base ufficiale di Go
FROM golang:1.18-alpine

# Imposta la directory di lavoro
WORKDIR /app

# Copia il codice sorgente nel container
COPY . .

# Installa le dipendenze Go
RUN go mod tidy

# Costruisce il progetto Go
RUN go build -o main .

# Esponi la porta 8080
EXPOSE 8080

# Esegui l'applicazione Go
CMD ["./main"]
