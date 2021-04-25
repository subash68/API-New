# Base image 
FROM golang:1.15-alpine AS dev

# Working directory
WORKDIR /cr-profile

# Coping All files to 
COPY . .

# Runing dependencies
RUN go get github.com/githubnemo/CompileDaemon && go build -o profile

# Exposing Port
EXPOSE 8080

# Command uses if not Overriden
CMD [ "./profile" ]

# building Image from dev
FROM dev AS qa

# Testing  working directory
WORKDIR /cr-profile-test

# Copy source files from dev
COPY --from=dev /cr-profile/ /cr-profile-test

# Command to test
CMD ["go test > test_report.log"]

# Building Production image
FROM dev as prod

# Production working directory
WORKDIR /cr-profile-prod

# Coping only Executable file from dev
COPY --from=dev /cr-profile/profile /cr-profile-prod/

# Start command
CMD ["./profile"]