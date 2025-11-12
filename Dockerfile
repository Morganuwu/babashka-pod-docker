# Etapa 1: Compilación del ejecutable Go
FROM golang:1.22-alpine AS build

WORKDIR /app

# Copiar dependencias
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copiar código fuente
COPY main2.go ./
COPY docker/ ./docker/
COPY babashka/ ./babashka/

# Compilar los dos programas: Babashka pod y main2.go
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o babashka-pod-docker
RUN go build -ldflags "-s -w" -o buscador main2.go

# Etapa 2: Imagen final ligera
FROM alpine:3.17
ARG version

# Directorio de trabajo
WORKDIR /root/.babashka/pods/repository/docker/docker-tools/0.1.0

# Copiar repositorios Babashka
COPY repository/ /root/.babashka/pods/repository

# Copiar ejecutables compilados
COPY --from=build /app/babashka-pod-docker .
COPY --from=build /app/buscador /usr/local/bin/buscador

# Dar permisos de ejecución
RUN chmod 755 /root/.babashka/pods/repository/docker/docker-tools/0.1.0/babashka-pod-docker \
    && chmod 755 /usr/local/bin/buscador

# Comando por defecto (puedes ejecutar el buscador o abrir shell)
CMD ["/bin/sh"]
