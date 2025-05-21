# ---- Build Stage ----
# Utilise une image Go officielle comme base pour la compilation
FROM golang:1.24-alpine AS builder

# Définit le répertoire de travail dans le conteneur
WORKDIR /app

# Copie les fichiers go.mod et go.sum pour télécharger les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copie le reste du code source de l'application
COPY . .

# Compile l'application Go.
# -o /app/server spécifie le nom et l'emplacement du binaire de sortie.
# CGO_ENABLED=0 désactive CGO, ce qui est utile pour créer des binaires statiques
# et éviter les dépendances C dans l'image finale.
# cmd/server/main.go est le point d'entrée de notre application.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/server cmd/server/main.go

# ---- Run Stage ----
# Utilise une image Alpine légère comme base pour l'image finale
# Alpine est choisie pour sa petite taille, ce qui réduit la surface d'attaque
# et la taille de l'image.
FROM alpine:latest

# Définit le répertoire de travail dans le conteneur
WORKDIR /app

# Copie uniquement le binaire compilé depuis l'étage 'builder'
COPY --from=builder /app/server /app/server

# (Optionnel) Copie les fichiers de configuration ou les assets s'il y en a
# COPY --from=builder /app/config ./config

# Expose le port sur lequel l'application Go s'exécute à l'intérieur du conteneur
EXPOSE 8080

# Commande pour exécuter l'application lorsque le conteneur démarre
ENTRYPOINT ["/app/server"]