FROM golang:1.16-alpine
WORKDIR /unsplash_me_workdir
ARG QUERY
ENV QUERY ${QUERY}
ARG PAGE=1
ENV PAGE ${PAGE}

COPY . ./
RUN go mod download

WORKDIR ./src
RUN go build -o /unsplash-me

CMD /unsplash-me ${QUERY} ${PAGE}
