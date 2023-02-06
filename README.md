# Galaxy 

EVM-compatible chain secured by the Lachesis consensus algorithm.

## Galaxy transaction tracing node

Node with implemented OpenEthereum(Parity) transaction tracing API.

```shell
make galaxy
```
The build output is ```build/galaxy``` executable.

## Running `galaxy`

Going through all the possible command line flags is out of scope here,
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `galaxy` instance.

### Launching a network

Before first launching of your mainnet node you will need appropriate genesis file:

```shell
wget https://github.com/Flag-Blockchain/chain-binary/raw/master/mainnet.g
```

Or for testnet node:
```shell
wget https://github.com/Flag-Blockchain/chain-binary/raw/master/testnet.g
```

Launching `galaxy` readonly (non-validator) node for network specified by the genesis file:

```shell
$ galaxy --genesis mainnet.g
```

### Configuration

As an alternative to passing the numerous flags to the `galaxy` binary, you can also pass a
configuration file via:

```shell
$ galaxy --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to
export your existing configuration:

```shell
$ galaxy --your-favourite-flags dumpconfig
```

#### Validator

New validator private key may be created with `galaxy validator new` command.

To launch a validator, you have to use `--validator.id` and `--validator.pubkey` flags to enable events emitter.

```shell
$ galaxy --nousb --validator.id YOUR_ID --validator.pubkey 0xYOUR_PUBKEY
```

`galaxy` will prompt you for a password to decrypt your validator private key. Optionally, you can
specify password with a file using `--validator.password` flag.

This branch is for a transaction tracing node. It's not recommended to use this branch for an Galaxy validator.

### Available methods

```shell
trace_get
trace_transaction
trace_block
trace_filter
```

### Building the source

Building `galaxy` requires a Go (version 1.15 or later) and a C compiler. Once the dependencies are installed, run:

```shell
make galaxy
```
The build output is ```build/galaxy``` executable.

### Running `go-galaxy` node with tracing ability

It's recommended to launch new node from scratch with CLI option `--tracenode` as this flag has to be set to use stored transaction traces.

```shell
$ galaxy --genesis /path/to/genesis.g --tracenode
```

### Tracing pre-genesis blocks
If you want to use tracing for pre-genesis blocks, you'll need to import evm history. You can find instructions here: [importing evm history](https://github.com/deamchain/lachesis_launch/blob/master/docs/import-evm-history.sh)

Then enable JSON-RPC API with option `trace` (`--http.api=eth,web3,net,txpool,flag,sfc,trace"`)

After that, you have to connect your nodes. Either connect them statically or specify a bootnode:
```shell
$ galaxy --fakenet 1/5 --bootnodes "enode://verylonghex@1.2.3.4:5050"
```

### Running the demo

For the testing purposes, the full demo may be launched using:
```shell
cd demo/
./start.sh # start the Galaxy processes
./stop.sh # stop the demo
./clean.sh # erase the chain data
```
Check README.md in the demo directory for more information.
