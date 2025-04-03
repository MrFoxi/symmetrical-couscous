FROM golang:latest

# Installer bash
RUN apt-get update && apt-get install -y bash

# Installer Python et NumPy via APT (évite l'erreur d'environnement géré)
RUN apt-get update && apt-get install -y python3 python3-numpy

# Installer GCC pour le C
RUN apt-get update && apt-get install -y gcc

# Installer Node.js et mathjs
RUN apt-get update && apt-get install -y nodejs npm
RUN npm install -g mathjs

# Nettoyer les fichiers inutiles pour alléger l'image
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /mathquizzer

COPY . .

# Initialiser le module Go
RUN go mod init math-quizzer

# Télécharger les dépendances et nettoyer le fichier go.mod
RUN go mod tidy

# # Compiler l'application Go
# RUN go build -o /main

# Démarrer un shell interactif pour éviter que le conteneur ne se ferme immédiatement
CMD ["sh", "-c", "tail -f /dev/null"]
