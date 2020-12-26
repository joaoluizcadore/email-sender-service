# Go E-mail Sender Service

## About
This is an e-mail sender service written in Golang. Integrated with Simple Message Broker (I'm using RabbitMQ)

## Installation

```bash
go get github.com/joaoluizcadore/email-sender-service
cd $GOPATH/src/github.com/joaoluiz/email-sender-service/
go get
go install
```

The binary will be in: $GOPATH/bin/

Create a folder where you want, move the binary in there (or make sure you have the bin folder in your PATH varivable)
Inside the folder, create a file called **config.json**

```javascript
{
    "Queue": {
        "Host": "amqp://guest:guest@localhost:5672/",
        "Name": "EMAIL_SENDER_QUEUE"    
    },
    "SMTP": {
        "Host": "smtp.gmail.com",
        "Port": 587,
        "From": "your_email@gmail.com",
        "Password": "your_app_password"
    },
    "Application": {
        "LogFile": "./application.log",
        "TemplateDir": "./templates/",
        "TemplateOnError": "error_message.html",
        "SendEmailOnError": false
    }
}
```
## Setup the RabbitMQ via Docker
Install the docker app in your computer.
Add the RabbitMQ docker: 
```bash
docker run -d — hostname my-rabbit — name rabbit13 -p 8080:15672 -p 5672:5672 -p 25676:25676 rabbitmq:3-management
```
The management web portal will be at: http://localhost:8080 (username: guest, password: guest)







