This is a blockchain project based Hyperledger Fabric and its official documentation. This particular segment extends the functionality of the official Test Network by dividing the development into
three parts:

1. Golang C (business logic)
2. Shell Script to interact with the chaincode
3. PHP script as the Web-interface to transport shell commands to the terminal

Prerequisite: Properly install the Fabric Samples Test Network ( https://hyperledger-fabric.readthedocs.io/en/release-2.2/prereqs.html ) .

Part 1: Golang Chaincode
The smartContract.go contains the transaction logic to initialize a ledger with Asset, Owner and ID . Functions with arguments have been built to invoke as needed. You may
replace the official Test Network folder's chaincode-go smartcontract with this chaincode or , go on with it.

Part 2: Shell Script

(...under development)