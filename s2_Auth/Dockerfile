# Base image 
FROM golang:1.15-alpine AS dev

# Working directory
WORKDIR /cr-auth

# Coping All files to 
COPY . .

# Runing dependencies
RUN go get github.com/githubnemo/CompileDaemon && go build -o auth

# Exposing Port
EXPOSE 8080

# Command uses if not Overriden
CMD [ "./auth" ]

# building Image from dev
FROM dev AS qa

# Testing  working directory
WORKDIR /cr-auth-test

# Copy source files from dev
COPY --from=dev /cr-auth/ /cr-auth-test

# Command to test
CMD ["go test"]

# Building Production image
FROM dev as prod

# Production working directory
WORKDIR /cr-auth-prod

# Coping only Executable file from dev
COPY --from=dev /cr-auth/auth /cr-auth-prod/

# Start command
CMD ["./auth"]