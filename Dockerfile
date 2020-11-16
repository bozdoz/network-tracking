FROM golang:1.15.5 as build

WORKDIR /app

# had issues installing with GOPATH and it can't be set with go.mod
ENV GOPATH ''

COPY . .

RUN go install

ARG SPREADSHEET_ID

RUN CGO_ENABLED=0 go build \
  -ldflags="-X 'bozdoz.com/spreadsheet.spreadsheetID=$SPREADSHEET_ID'" \
  .


FROM alpine

WORKDIR /app

COPY --from=build /app/network /app
COPY --from=build /app/credentials.json /app

# might not explicitly need this directory
RUN mkdir /etc/periodic/custom

COPY ./network-cron /etc/periodic/custom

RUN chmod 0644 /etc/periodic/custom/network-cron \ 
  && crontab /etc/periodic/custom/network-cron

CMD ["crond", "-f"]