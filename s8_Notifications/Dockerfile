# Base image 
FROM golang:1.15-alpine AS dev

# Working directory
WORKDIR /notifications

# Coping All files to 
COPY . .

# Runing dependencies
RUN go get github.com/githubnemo/CompileDaemon && go build -o notifications

# Exposing Port
EXPOSE 8080

# Command uses if not Overriden
CMD [ "./notifications" ]

# building Image from dev
FROM dev AS qa

# Testing  working directory
WORKDIR /notifications-test

# Copy source files from dev
COPY --from=dev /notifications/ /notifications-test

# Command to test
CMD ["go test > test_report.log"]

# Building Production image
FROM dev as prod

# Production working directory
WORKDIR /notifications-prod

# Coping only Executable file from dev
COPY --from=dev /notifications/notifications /notifications-prod/

# Start command
CMD ["./notifications"]