
FROM golang:1.18-alpine

WORKDIR /app

COPY ./main ./main

CMD [ "./main" ]
