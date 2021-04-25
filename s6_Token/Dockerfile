# Base image 
FROM golang:1.15-alpine AS dev

# Working directory
WORKDIR /ch-token

# Coping All files to 
COPY . .

# Runing dependencies
RUN go get github.com/githubnemo/CompileDaemon && go build -o token

# Exposing Port
EXPOSE 8080

# Command uses if not Overriden
CMD [ "./token" ]

# building Image from dev
FROM dev AS qa

# Testing  working directory
WORKDIR /ch-token-test

# Copy source files from dev
COPY --from=dev /ch-token/ /ch-token-test

# Command to test
CMD ["go test > test_report.log"]

# Building Production image
FROM dev as prod

# Production working directory
WORKDIR /ch-token-prod

# Coping only Executable file from dev
COPY --from=dev /ch-token/token /ch-token-prod/

# Start command
CMD ["./token"]