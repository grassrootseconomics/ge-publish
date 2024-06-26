# ge-publish

> CLI tool to publish GE related smart contracts to Celo.

![GitHub release (latest by date)](https://img.shields.io/github/v/release/grassrootseconomics/ge-publish)

Supported Smart Contracts:

- [x] SwapPool
- [x] DecimalQuote
- [x] PriceIndexQuote
- [x] Limiter
- [x] LimiterIndex
- [x] TokenRegistry
- [x] ERC20Demurrage

_Note:_ All smart contracts are compiled with v 0.8.19 unless otherwise stated.

## Install

Download and extract the binary for your OS/Arch from the [releases page](https://github.com/grassrootseconomics/ge-publish/releases).

Alternatively you can install with:

```bash
curl -L https://ge-publish.grassecon.net/install.sh | bash
```

## Usage

```bash
ge-publish --help
```

## Examples

### Limiter

```bash
# Set your private key
export PRIVATE_KEY=
ge-publish --testnet publish limiter
```

### SwapPool

```bash
# Set your private key
export PRIVATE_KEY=
ge-publish --testnet publish swap-pool --name MySwapPool --symbol SWP1 --decimals 6 --token-registry 0x000000000000000000000000000000000000dEaD --token-limiter 0x000000000000000000000000000000000000dEaD
```

## License

[AGPL-3.0](LICENSE).