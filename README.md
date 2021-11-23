# 24h Regional API

## Build & local run

```
make build
make run
```

## API Endpoints

General endpoint is _$(DOMAIN):$(PORT)/api/v1_

### Health

* Endpoint: /health
* Method: GET
* Response: 

```
{
  "status": "UP"
}
```