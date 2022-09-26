# Firehose on StarkNet

Firehose implementation for [StarkNet](https://starknet.io/), bootstrapped from [streamingfast/firehose-acme](https://github.com/streamingfast/firehose-acme/commit/6966e1a3aaf49d2d398686333967299e97bde05b). This project is maintained by the [zkLend](https://zklend.com/) team.

## Generating protobuf types

A [script](./types/pb/generate.sh) is available for generating Go types from protobuf files. Unless you have a specific reason to use another version, make sure you have the exact versions of the following tools installed:

```bash
$ protoc --version
libprotoc 3.19.4
$ protoc-gen-go --version
protoc-gen-go v1.27.1
```

You will also need to have [streamingfast/proto](https://github.com/streamingfast/proto) cloned in the same folder where the current repository lives.

Then, invoke the script to generate the types:

```bash
$ ./types/pb/generate.sh
```

If any diff is found, make sure to also commit the `last_generate.txt` file so that others know what exact code versions were used to generate the types.

## License

[Apache 2.0](./LICENSE)
