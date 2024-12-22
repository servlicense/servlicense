FROM alpine
COPY --from=golang:1.22-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

# install and cache go dependencies
WORKDIR /api
COPY ./api/go.mod ./api/go.sum ./
RUN go mod download

# install and cache web dependencies
WORKDIR /web
RUN apk --no-cache add nodejs npm
RUN npm install -g yarn
COPY ./web/package.json ./web/yarn.lock ./
RUN yarn install

# build web
ENV NODE_ENV=production
COPY ./web/ .
RUN yarn build

# build api
WORKDIR /api
COPY ./api .
RUN CGO_ENABLED=0 go build -o api .

WORKDIR /
COPY ./start.sh .

CMD ["sh", "start.sh"]
