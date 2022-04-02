# GO - URL Shortener
Go API created following the excellent tutorial of https://www.eddywm.com/

<details>
  <summary><strong>Table of contents</strong></summary>

  - ⚠️ [Requirement](#-requirement)
  - 🏁 [Getting Started](#-getting-started)
  - 💡 [Size of short url](#-size-of-short-url)

</details>


## ⚠️ Requirement
- Redis server installed 
- Go installed

## 🏁 Getting Started

Copy the project locally

Start the Redis server:
```sh
sudo service redis-server start
```

Start the GO REST API:
```sh
go run ./main.go
```

Send POST request with a REST Client (VSCode for example)
```sh
POST http://localhost:9808/create-short-url HTTP/1.1
content-type: application/json

{
    "long_url": "https://github.com/darkmelcof/GO_URL_SHORTENER",
    "user_id" : "e0dba740-fc4b-4977-872c-d360239e6b10"
}
```

You should reveive the following JSON Response
```sh
{
  "message": "short url created successfully",
  "short_url": "http://localhost:9808/coGaf4Vu"
}
```

## 💡 Size of short url
Go to shorturl_generator.go, change the number in the following code:

```sh
finalString[:8]
```

[Redis]: <https://redis.io/>
[Go]: <https://go.dev/>