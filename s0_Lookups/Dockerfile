# Base image 
FROM golang:1.15-alpine AS dev

# Working directory
WORKDIR /lut

# Coping All files to 
COPY . .

# Coping configuration
COPY ../config.json /lut/configuration/

# Runing dependencies
RUN go get github.com/githubnemo/CompileDaemon && go build -o lut 

# Exposing Port
EXPOSE 8080

# Command uses if not Overriden
CMD [ "./lut", "info" ]

# building Image from dev
FROM dev AS qa

# Testing  working directory
WORKDIR /lut-test

# Copy source files from dev
COPY --from=dev /lut/ /lut-test

# Command to test
CMD ["go test > test_report.log"]

# Building Production image
FROM dev as prod

# Production working directory
WORKDIR /lut-prod

# Coping only Executable file from dev
COPY --from=dev /lut/lut /lut-prod/

# Start command
CMD ["./lut", "debug"]