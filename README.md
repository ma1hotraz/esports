## Readme

### Run redis

```bash
docker run -dp 6379:6379 redis
```

### Build client

```bash
cd js
npm run esbuild
```

### Run server

```bash
go run main.go

```

### Compose

```bash
docker compose up -d --no-deps --build app
```

### Air

https://stackoverflow.com/questions/71643902/how-to-reload-go-fiber-in-the-terminal


pip3 install datamodel-code-generator
datamodel-codegen --input underdog.json --output underdog.py --input-file-type=json


