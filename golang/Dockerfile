FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./src /go/src/github.com/Di0niz/cyber-backend
WORKDIR /go/src/github.com/Di0niz/cyber-backend

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	app; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi
	
EXPOSE 8080