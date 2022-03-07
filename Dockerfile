
FROM golang:1.16-alpine

WORKDIR /app

COPY ./main ./main

CMD [ "./main" ]
