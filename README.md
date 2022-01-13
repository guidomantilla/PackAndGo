# Pack and Go

Setup:
1. To Install direnv (MacOS) execute: `brew install direnv`
2. To Download Golang dependencies execute: `make prepare`
3. To Run:
   1. Semi containerized, execute: `make run-local`
   2. Fully containerized, `make deploy-all-local`
4. To Destroy, execute: `make destroy-all-local`
