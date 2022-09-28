# Firehose on StarkNet

Firehose implementation for [StarkNet](https://starknet.io/), bootstrapped from [streamingfast/firehose-acme](https://github.com/streamingfast/firehose-acme/commit/6966e1a3aaf49d2d398686333967299e97bde05b). This project is maintained by the [zkLend](https://zklend.com/) team.

## Docker images

Docker images are available on [Docker Hub](https://hub.docker.com/r/starknet/firestark). Two types of tags are provided:

### starknet/firestark:${FIRESTARK_VERSION}

These images are pure `firestark` images with no StarkNet client bundled. These are suitable for using as non-reader services (e.g. `relayer`) that are not client-dependent.

### starknet/firestark:${FIRESTARK_VERSION}-pathfinder-${PATHFINDER_VERSION}

These images are `firestark` images bundled with [instrumented pathfinder](https://github.com/starknet-graph/pathfinder) builds. They work as an out-of-the-box solution that you can use to quickly spin up a working cluster.

## Running Firehose

### Quickstart with `docker-compose`

An [example `docker-compose.yml` file](./docker/docker-compose.yml) is available for getting the whole system to run in no time. Make sure to modify the [.env file](./docker/.env) with the L1 Ethereum RPC URL, and then:

```console
$ cd ./docker
$ sudo docker-compose up -d
```

Once everything is up and running, you should be able to subscribe to the block stream with [grpcurl](https://github.com/fullstorydev/grpcurl):

```console
$ grpcurl -plaintext -d '{"start_block_num": 0}' localhost:18015 sf.firehose.v2.Stream.Blocks
```

As noted in the `docker-compose.yml` file, it only serves as an example of running different services separately. For production workload you might want to deploy a setup with high-availability, taking the file as a starting point.

### Running without Docker

_Note that these instructions are only suitable for development purposes. For production workload, you'd probably want to run a high-availability setup._

To run a Firehose-enabled `pathfinder` node, you need to install the [instrumented fork version](https://github.com/starknet-graph/pathfinder) we also maintain:

```console
$ cargo install --locked --git https://github.com/starknet-graph/pathfinder --branch patch pathfinder
```

_Alternatively, you can also build the binary without installing. Just make sure to modify [standard.yaml](./devel/standard/standard.yaml) to point to the resulting binary._

Then, replace `ETHEREUM_URL` in [standard.yaml](./devel/standard/standard.yaml) with a valid Ethereum JSON-RPC URL for the L1 network you're going to sync.

After which you can just run:

```console
$ cd ./devel/standard
$ mkdir -p ./pathfinder-data
$ ./start.sh
```

The instrumented `pathfinder` node should now start syncing. You can test the setup by subscribing to the block stream with [grpcurl](https://github.com/fullstorydev/grpcurl):

```console
$ grpcurl -plaintext -d '{"start_block_num": 0}' localhost:18015 sf.firehose.v2.Stream.Blocks
```

## Generating protobuf types

A [script](./types/pb/generate.sh) is available for generating Go types from protobuf files. Unless you have a specific reason to use another version, make sure you have the exact versions of the following tools installed:

```console
$ protoc --version
libprotoc 3.19.4
$ protoc-gen-go --version
protoc-gen-go v1.27.1
```

You will also need to have [streamingfast/proto](https://github.com/streamingfast/proto) cloned in the same folder where the current repository lives.

Then, invoke the script to generate the types:

```console
$ ./types/pb/generate.sh
```

If any diff is found, make sure to also commit the `last_generate.txt` file so that others know what exact code versions were used to generate the types. If you're making changes to the proto files, make sure to commit those proto changes first before running the script, otherwise checking out to the revision specified in `last_generate.txt` would yield the proto files _before_ the changes.

## License

[Apache 2.0](./LICENSE)
