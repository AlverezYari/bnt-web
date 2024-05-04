# We will use the official Go image for our executable container
FROM golang:1.22.1 as builder

# We will set the working directory to /app
WORKDIR /app

# We then copy our GO bin and our static assests to the /app directory
COPY ./bin/bntweb /app/

# We then copy our ui static assests to the /app directory so our server can server them and render our html templates
COPY ./ui /app/ui

# We then expose our application on port 80
EXPOSE 80

# We then run our application
CMD ["/app/bntweb","-addr",":80"]
