## Account-Service
#### Build container
- `make docker-build`  
#### Start container
- `make docker-up`  
***
#### To run local development  
- `make run`  
#### To run test  
- `make test`  
***
#### The service will serve as `http://localhost:1400` which contain path below  
- *`GET /account` for getting account data*  
- *`PUT /update-account` for update account data*  
- *`GET /transactions` for getting transaction data*  
- *`GET /health` for service health checking*  
***
##### *See Makefile for other command*