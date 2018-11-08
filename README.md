# GO Book API

## How to run:

First make sure you have redis up and running on port 6379 with no password.

Run this command to get redis up and running in docker:
```
docker run -d -p 6379:6379 redis
```

Once redis is up and running, run `./runServer.sh` and the api is good to go

## Routes:

`DELETE http://localhost:1323/emptylist`
- Remove all books from the database

`DELETE http://localhost:1323/removebook/:id`
- Remove a book by a particular id

`GET: http://localhost:1323/booklist`
- List out all book Ids

`GET http://localhost:1323/book/:id`
- Get info for a particular book by id

`POST http://localhost:1323/addbook`
```
{
  Author: "Jeff VanderMeer"
  PublishDate: "2014-11-18"
  Publisher: "FSG Originals"
  Rating: "3"
  Status: "Checked-In"
  Title: "Area-X"
}
```
- Add a book with the following info

## Testing

Make sure you have redis up and running on port 6379 with no password.

Then run `./testBookHandlers.sh`