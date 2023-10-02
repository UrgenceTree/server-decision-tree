# Utilisation d'une image de base avec Go préinstallé
FROM golang:latest

# Creation d'un utilisateur non-root dans le conteneur
RUN useradd -u 10001 -m admin

# Passez à l'utilisateur admin
USER admin

# Creation d'un dossier de travail
WORKDIR /sdt_app

# Changez le propriétaire du dossier
RUN chown admin:admin /sdt_app

# Copie du code vers le ctr
COPY . .

# Initialisation d'un nouveau module Go
RUN rm go.mod && go mod init v0

# Copie d'un repertoire
COPY ./unit_test/time_test/go_part/* ./src

# Move vers un dossier
RUN cd ./src

# Compilation du projet go
RUN go build -o sdt_questions ./src/test_questions.go
