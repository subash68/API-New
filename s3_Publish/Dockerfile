# Base image 
FROM golang:1.15-alpine AS dev

# Working directory
WORKDIR /cr-publish

# Coping All files to 
COPY . .

# Runing dependencies
RUN go get github.com/githubnemo/CompileDaemon && go build -o publish

# Exposing Port
EXPOSE 8080

# Command uses if not Overriden
CMD [ "./publish" ]

# building Image from dev
FROM dev AS qa

# Testing  working directory
WORKDIR /cr-publish-test

# Copy source files from dev
COPY --from=dev /cr-publish/ /cr-publish-test

# Command to test
CMD ["go test > test_report.log"]

# Building Production image
FROM dev as prod

# Production working directory
WORKDIR /cr-publish-prod

# Coping only Executable file from dev
COPY --from=dev /cr-publish/publish /cr-publish-prod/

# Start command
CMD ["./publish"]