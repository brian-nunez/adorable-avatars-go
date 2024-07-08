# Adorable Avatars

Adorable Avatars is a set of assets when combined generate a random avatar. It is designed to be used by developers who need a quick image placeholder that is unique and adorable. The original idea and code was written in Node.js and is not longer maintained. Check out the original project [here](https://github.com/itsthatguy/avatars-api-middleware).

This project is an API that exposes endpoints to generate avatars.

## Usage

To run the server, execute the following command:

```bash
go run ./cmd/main.go 
```

The server will be running on port 8080.

### Get a random avatar

```bash
GET /random
```

This endpoint will return a random avatar.

### Get a list of available `eyes`, `noses`, and `mouths`

```bash
GET /list
```

The response will be a JSON object with all the available assets.

```json
{
    "data": {
        "eyes": [
            "eyes1",
            "eyes10",
            "eyes2",
            "eyes3",
            "eyes4",
            "eyes5",
            "eyes6",
            "eyes7",
            "eyes9"
        ],
        "noses": [
            "nose2",
            "nose3",
            "nose4",
            "nose5",
            "nose6",
            "nose7",
            "nose8",
            "nose9"
        ],
        "mouths": [
            "mouth1",
            "mouth10",
            "mouth11",
            "mouth3",
            "mouth5",
            "mouth6",
            "mouth7",
            "mouth9"
        ]
    }
}
```

## More to come

This project is a work in progress. More features will be added soon.

