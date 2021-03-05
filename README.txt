This is a blockchain project based Hyperledger Fabric and its official documentation. This particular segment extends the functionality of the official Test Network by dividing the development into
three parts:

1. Golang Chaincode (business logic)
2. Shell Script to interact with the chaincode
3. PHP script as the Web-interface to transport shell commands to the terminal

Prerequisite: Properly install the Fabric Samples Test Network ( https://hyperledger-fabric.readthedocs.io/en/release-2.2/prereqs.html ) and PHP .

Part 1: Golang Chaincode
The smartContract.go contains the transaction logic to initialize a ledger with Asset, Owner and ID . Functions with arguments have been built to invoke as needed. You may replace the official Test Network folder's chaincode-go smartcontract with this chaincode or , go on with it.

Part 2: Shell Script
Shell commands are embedded into PHP's shell_exec() function to make the whole process much user-friendly from a browser. Backticks have played a key role in running the commands into the terminal instead of displaying the output status on the browser.

Part 3: PHP Script
Starting with any simple HTML login page, PHP here receives the user-credentials and then processes, sanitizes, inserts (into Shell command as variable) and finally transports the command from the browser to the terminal to run. The outcome of the commands are well-displayed on the browser as soon as they are ready for the user.
