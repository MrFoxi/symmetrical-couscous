FROM golang:latest

# Installer bash et autres outils utiles
RUN apt-get update && apt-get install -y bash

WORKDIR /mathquizzer

COPY . .

# Initialiser le module Go
RUN go mod init math-quizzer

# Télécharger les dépendances et nettoyer le fichier go.mod
RUN go mod tidy

# Compiler l'application Go
RUN go build -o main

# Démarrer un shell interactif pour éviter que le conteneur ne se ferme immédiatement
CMD ["sh", "-c", "tail -f /dev/null"]