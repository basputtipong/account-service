# Account-Service
### Build container
- `make docker-build`  
### Start container
- `make docker-up`  
***
### Local Development
- *`make run` To run local development*  
#### The service will serve as `http://localhost:1400` which contain path below  
- *`GET /account` for getting account data*  
- *`PUT /update-account` for update account data*  
- *`GET /transactions` for getting transaction data*  
- *`GET /health` for service health checking*  
***
### Unit test  
- *`make test-service` to run service unit test*  
#### To generate mock file with mockery  
- *First, you need to install mockery by using command
`make mock-install`*  
- *Then use this command to generate repository mock 
`mockery --all --dir=internal/core/port  --output=internal/core/port/mocks --outpkg=mocks`*  
***
#### *See Makefile for other command*