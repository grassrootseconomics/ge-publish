package container

import (
	"math/big"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/crypto"
	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/w3-celo/module/eth"
	"github.com/urfave/cli/v2"
)

type PublishTxResp struct {
	TxHash common.Hash
}

func (c *Container) SendContractPublishTx(cCtx *cli.Context, contractBytecode []byte, gasLimit uint64) (PublishTxResp, error) {
	var (
		nonce  uint64
		txHash common.Hash
	)

	providerOpts := celoutils.ProviderOpts{
		ChainId:     celoutils.MainnetChainId,
		RpcEndpoint: cCtx.String("rpc"),
	}

	if cCtx.Bool("testnet") {
		providerOpts.ChainId = celoutils.TestnetChainId
		providerOpts.RpcEndpoint = cCtx.String("rpc")
	}

	p, err := celoutils.NewProvider(providerOpts)
	if err != nil {
		return PublishTxResp{}, err
	}

	privateKey, err := crypto.HexToECDSA(cCtx.String("private-key"))
	if err != nil {
		return PublishTxResp{}, err
	}

	if err := p.Client.CallCtx(
		cCtx.Context,
		eth.Nonce(crypto.PubkeyToAddress(privateKey.PublicKey), nil).Returns(&nonce),
	); err != nil {
		return PublishTxResp{}, err
	}

	tx, err := p.SignContractPublishTx(
		privateKey,
		celoutils.ContractPublishTxOpts{
			ContractByteCode: contractBytecode,
			GasFeeCap:        big.NewInt(cCtx.Int64("gas-fee-cap")),
			GasTipCap:        big.NewInt(cCtx.Int64("gas-tip-cap")),
			GasLimit:         gasLimit,
			Nonce:            nonce,
		},
	)
	if err != nil {
		return PublishTxResp{}, err
	}

	if err := p.Client.CallCtx(
		cCtx.Context,
		eth.SendTx(tx).Returns(&txHash),
	); err != nil {
		return PublishTxResp{}, err
	}

	return PublishTxResp{
		TxHash: txHash,
	}, nil
}
