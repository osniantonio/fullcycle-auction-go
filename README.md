# Rodar o projeto
osniantonio@Avell:~/aulas/labs/fullcycle-auction-go/cmd/auction$ go run main.go 

# Subir o mongoDB via Docker
docker container run -d -p 27017:27017 --name auctionsDB mongo
