# Base image 
FROM golang:1.15-alpine AS dev

# Working directory
WORKDIR /Onboarding

# Coping All files to 
COPY . .

# Runing dependencies
RUN go get github.com/githubnemo/CompileDaemon && go build -o onboarding

# Exposing Port
EXPOSE 8080

# Command uses if not Overriden
CMD [ "./Onboarding" ]

# building Image from dev
FROM dev AS qa

# Testing  working directory
WORKDIR /Onboarding-test

# Copy source files from dev
COPY --from=dev /Onboarding/ /Onboarding-test

# Command to test
CMD ["go test"]

# Building Production image
FROM dev as prod

# Production working directory
WORKDIR /Onboarding-prod

# Coping only Executable file from dev
COPY --from=dev /Onboarding/onboarding /Onboarding-prod/

# Start command
CMD ["./onboarding"]