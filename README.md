# get Docbase Articles to Markdown

## go get

use **.env**. need to go get.

```
$ go get github.com/joho/godotenv
```

## .env Edit

### 1) Create .env
```
$ cp env-example .env
```

### 2) get API Token

https://{team-domain}.docbase.io/settings

### 3) env Edit

```
# Please get access token from your team's settings.
ACCESS_TOKEN=

# Example https://{TEAM_DOMAIN}.docbase.io/
TEAM_DOMAIN=

# Your AUTHOR ID.
AUTHOR_ID=

# Save location.
SAVE_DIR=

# docbase api get limit 20 items. 1hour limit 30 items.
# [pages = 1] -> 1~19 items. [pages = 2] -> 20~39 items.
# [pages = 0] All Articles
# get for descending order.
PAGES=
```

## go run

```
$ go run docbase.go struct.go

or 

$ go build docbase.go struct.go
$ go run docbase
```

