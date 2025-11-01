# go-booking-micro
An booking application based on microservice architecture

### **Structure project:**
- **database**: contain database schema for each service
- **infras**: gateway...
- **serivces**: contain source code for each service
    + **common**: contain util, helper for each service
    + **cmd**: contain file main.go, point start of service
    + **module**: simple architecture: transport, model, repo, biz
        + **biz (service)**: contain business of entity
        + **model**: contain entity schema, dto, response ...
        + **repo**: connect to database
        + **transport(controller)**: ginapi for rest api, grpc for grpc, socket ... 
- **shared**: contain common module for all services
    + **asyncjob**: use to handle side effect task
    + **pubsub**: like rabbitmq, kafka or redis pub/sub
    + **srvctx**: like contex.Context in gin/golang, but use for each service to reuse
    + **tokenprovider**: jwt
    + **uploadprovider**: s3

- **migrations:**
    + create new file migrations: migrate create -ext sql -dir <path_to_migrations_directory> -seq <file_name>
        - example:  migrate create -ext sql -dir migrations -seq delete_type_amenity_table 
    + apply a migration for up/down: migrate -database ${POSTGRESQL_URL} -path <path_to_migrations_directory> <up/down> <index>
  
    + See more: https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md
- **grpc:**
    + create a new proto file in folder sheared/proto
    + run command line: buf.gen.yaml
    + in grpc have client and server
    + in server: using core biz in each service to create 
    + in client: cd repo and create folder rpc folder to create biz for client
    + using dto in shared/dto to share structure data between client and server
    + see example in location/country and property/amenity for detail
    + see more: https://viblo.asia/p/grpc-va-ung-dung-no-trong-microservices-ORNZqo8N50n