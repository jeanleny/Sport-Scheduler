FROM golang:1.27rc2-bookworm

WORKDIR /app

COPY . .

#Ca c'est un banger absolu
#Comme Makefile, chaque ligne est un step mit dans le cache
#Si le cache est différent des fichiers actuels, il le rebuild sinon il passe a l'étape d'après

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
	go build -o api

EXPOSE 5432

CMD ["./api"]
