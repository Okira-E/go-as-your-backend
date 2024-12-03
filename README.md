# Go As Your Backend

This is a starter template for using Go as a REST service. It uses [Fiber](https://docs.gofiber.io/) and a bunch of other libraries alongside a sensible backend file structure to form a framework for your app.

## Why?

Because Go is awesome. And will serve more requests per dollar than most other languages.

## What's does the template offer?

- A justfile with commands for most operations
- Database migrations
- Project structure for a web application
- A parser for filter params to construct dynamic SQL queries without needing an ORM or a query builder
- Keeping it simple

## Dependencies

- [gofiber](https://gofiber.io/): Fiber is a Go web framework built on top of Fasthttp, the fastest HTTP engine for Go. It's designed to ease things up for fast development with zero memory allocation and performance in mind.
- [sqlx](github.com/jmoiron/sqlx)
- [swag](https://github.com/swaggo/swag): To generate OpenAPI documentation using `make docs` you need to install [swag](https://github.com/swaggo/swag) system-wide. Installation is detailed in their Github.
- [golang-migrate](https://github.com/golang-migrate/migrate): Database migrations. CLI and Golang library.
- A database Go driver of your liking (comes with pq out of the box)

## Usage

After cloning this template repository & adding database credentials in a `.env` file in the root directory, all the commands are available in the [justfile](justfile):

## Example Request

```js
fetch("http://localhost:3200/api/v1/users?filter={ \"where\": { \"and\": { \"id\": \"126423\" } } }&limit=50", { method: "GET", })
```

## Example Response

```json
{
    "success": true,
    "code": 200,
    "message": "",
    "data": [
        {
            "id": 126423,
            "username": "Elysia3848",
            "email": "Lethia@Nolana3848.com",
            "first_name": "Anet",
            "last_name": "Aurie",
            "is_active": true,
            "created_at": "2024-11-03T20:01:23.768481Z",
            "updated_at": "2024-11-03T20:01:23.768481Z"
        }
    ]
}
```

## Building Dybamic Queries

There are no query builder libraries in this template.
A custom solution is built to handle customizing the clauses of a query from the endpoint.

- You can construct both OR & AND clauses: `/users?filter={ "where": { "and": { "first_name": "Denys", "is_active": "true" }, "or": { "first_name": "Anet" } } }&limit=50&offset=0`
- the `limit` param has a default and a limit of 100
- You can't chain the same column name on the same clause currently: `{ "where": { "or": { "first_name": "Anet", "first_name": "Dosi" } } }`
- For a `LIKE` expression, simply provide a '%' character in the value: `/users?filter={ "where": { "and": { "first_name": "Den%25" } } }&limit=50&offset=0`. Note, '%' has to be encoded in a URL as '%25'

## Contributing

Feel free to suggest better ways for things and open a PR.

## License

This project is licensed under the Unlicensed License - see the [UNLICENSE](UNLICENSE) file for details
