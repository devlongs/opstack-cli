## OP Stack CLI
OP Stack CLI is a Golang command-line utility for easily spinning up an OP Stack–based rollup. This tool wraps the manual steps normally required to deploy an Optimism-based chain, handle address generation, configure environment variables, build repositories, and run the associated services (execution client, consensus client, batcher, proposer, etc.).

#### Warning: The OP Stack is under active development. 

## Workflow
1. Initialize directories, clone repos, and check dependencies
```bash
opstack-cli init
```

2. Build the OP Stack binaries
```bash
opstack-cli build
```

3. Generate addresses and private keys
```bash
opstack-cli addresses
```

4. Deploy L1 contracts
```bash
opstack-cli deploy-l1
```

5. Create L2 configuration (genesis.json, rollup.json, jwt.txt)
```bash
opstack-cli config-l2
```

6. Start all services
(You can also start them one by one, e.g. `opstack-cli start geth`)
```bash
opstack-cli start all
```