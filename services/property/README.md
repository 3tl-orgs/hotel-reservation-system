# Build property-service-local for testing in localhost
- Step 1: Sit in root project run command to build prj: (When you change code -> build it again) 
`docker build -f services/property/Dockerfile -t property-service .`.
- Step 2: Cd to property service
`cd services/property`
- Step 3: run command:
 ** Need .env file **
`docker-compose up`
- Step 4: Testing
    + Config database
    + test: http://localhost:3004/api/v1/properties


