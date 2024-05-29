# MongoDB Create Database and Collection Job

## Description
This is a simple app to do one job: create a database and a collection in MongoDB. Mainly this is going to run in a docker container, to initialize the database and collection in a MongoDB instance.

## Environment Variables
- `MONGO_DB_NAME`: The name of the database to be created.
- `MONGO_COLLECTION`: The name of the collection to be created. If notihing is provided, the collection will be named `delete_me`.
- `MONGO_URL`: The URL of the MongoDB instance.

## How to run
To run this app, you can use the following command:
```bash
docker run -e MONGO_DB_NAME=my_db -e MONGO_COLLECTION=my_collection -e MONGO_URL=mongodb://localhost:27017 kaenova/mongo-init-db-collection
```

ór in docker-compose:
```yaml
version: '3.7'

services:
  
  mongo-db:
    image: mongo
    expose:
      - 27017

  mongo-init-db-collection:
    image: kaenova/mongo-init-db-collection
    environment:
      - MONGO_DB_NAME=my_db
      - MONGO_COLLECTION=my_collection
      - MONGO_URL=mongodb://mongo-db:27017
```

Note to remember that this app will only run once, and it will create the database and collection. After that, it will exit.