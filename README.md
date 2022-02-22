# 24h Regional API

## Build & local run

```
make build
make run
```

## Data Structure
```
{
        "name": string,
        "open": string,
        "gps": {
            "latitude": long,
            "longitude": long
        },
        "location": {
            "street": string,
            "zip": string,
            "city": string,
            "country": string
        },
        "products": []string,
        "source": string
}
```

## API Endpoints

General endpoint is _$(DOMAIN):$(PORT)

| methode | endpoint             | definition                      |
|---------|----------------------|---------------------------------|
| GET     | "/health"            | [link](#health)                 |
| GET     | "/api/v1/stores"     | [link](#get-all-stores)         |
| GET     | "/api/v1/stores/:id" | [link](#get-single-store-by-id) |
| POST    | "/api/v1/stores"     | [link](#add-store)              |
| DELETE  | "/api/v1/stores/:id" | [link](#get-single-store-by-id) |


### Health

* Endpoint: _$(DOMAIN):$(PORT)/health
* Method: GET
* Response: 
```
{
  "status": string
}
```
* Example: 
```
curl --location --request GET 'https://api.24h-regional.de/health/'
status: UP
```

### Get All Stores

* Endpoint: _$(DOMAIN):$(PORT)/api/v1/stores
* Method: GET
* Response:

```
[
    {
            "name": string,
            "open": string,
            "gps": {
                "latitude": long,
                "longitude": long
            },
            "location": {
                "street": string,
                "zip": string,
                "city": string,
                "country": string
            },
            "products": []string,
            "source": string
    }
]
```

* Example:
```
curl --location --request GET 'https://api.24h-regional.de/api/v1/stores/
[
    {
        "name": "Regio Box Wetzendorf",
        "open": "24h",
        "gps": {
            "latitude": 49.47177986951858,
            "longitude": 11.037174273016173
        },
        "location": {
            "street": "Wetzendorfer Straße 278",
            "zip": "90427",
            "city": "Nürnberg",
            "country": "Germany"
        },
        "products": [
            "Eier",
            "Nudeln",
            "Gemüse",
            "Apfelmus",
            "Kaffee"
        ],
        "source":""
    }
    //snip
]
```
### Get single store by ID

* Endpoint: _$(DOMAIN):$(PORT)/api/v1/stores/:id
* Method: GET
* Response:

```
{
    "name": string,
    "open": string,
    "gps": {
        "latitude": long,
        "longitude": long
    },
    "location": {
        "street": string,
        "zip": string,
        "city": string,
        "country": string
    },
    "products": []string,
    "source": string
}
```

* Example:
```
curl --location --request GET 'https://api.24h-regional.de/api/v1/stores/621001e6c372b2c05fce1c2e'
{
    "id": "621001e6c372b2c05fce1c2e",
    "name": "Hornig GbR",
    "gps":
    {
        "latitude": 49.44008648554797,
        "longitude": 10.875766538864738
    },
    "location":
    {
        "street": "Soldnerstraße 47",
        "zip": "90766",
        "city": "Fürth",
        "country": "Germany"
    },
    "open": "24h",
    "products":
    [
        "Eier aus Freilandhaltung",
        "Kartoffeln",
        "Nüsse",
        "Mohn",
        "Gewürze",
        "Fruchtsecco"
    ],
    "source": "https://www.landkreis-fuerth.de/zuhause-im-landkreis/gutes/frische-regionale-produkte-aus-automaten.html"
}
```

### Add Store

* Endpoint: _$(DOMAIN):$(PORT)/api/v1/stores
* Method: POST
* Request:

```
[
    {
            "name": string,
            "open": string,
            "gps": {
                "latitude": long,
                "longitude": long
            },
            "location": {
                "street": string,
                "zip": string,
                "city": string,
                "country": string
            },
            "products": []string,
            "source": string
    }
]
```

* Response
```
[
    {
            "id: string,
            "name": string,
            "open": string,
            "gps": {
                "latitude": long,
                "longitude": long
            },
            "location": {
                "street": string,
                "zip": string,
                "city": string,
                "country": string
            },
            "products": []string,
            "source": string
    }
]
```

* Example:
```
curl --location --request POST 'https://$DOMAIN:$PORT/api/v1/stores/' \
--header 'Content-Type: application/json' \
--data-raw '{
        "name": "Regio Box Wetzendorf",
        "open": "24h",
        "gps": {
            "latitude": 49.47177986951858,
            "longitude": 11.037174273016173
        },
        "location": {
            "street": "Wetzendorfer Straße 278",
            "zip": "90427",
            "city": "Nürnberg",
            "country": "Germany"
        },
        "products": [
            "Eier",
            "Nudeln",
            "Gemüse",
            "Apfelmus",
            "Kaffee"
        ],
        "source":""
    }'
```


### Delete store by ID

* Endpoint: _$(DOMAIN):$(PORT)/api/v1/stores/:id
* Method: DELETE
* Response:

```
Deletion of Store with ID $id: successful
```

* Example:
```
curl --location --request DELETE 'https://$DOMAIN:$PORT/api/v1/stores/6214f9ea4836827453678c14'
Deletion of Store with ID 6214f9ea4836827453678c14: successful
```