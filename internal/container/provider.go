package container

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/grassrootseconomics/ethutils"
	"github.com/lmittmann/w3/module/eth"
	"github.com/urfave/cli/v2"
)

type PublishTxResp struct {
	TxHash          common.Hash
	ContractAddress common.Address
}

func (c *Container) SendContractPublishTx(cCtx *cli.Context, contractBytecode []byte, gasLimit uint64) (PublishTxResp, error) {
	var (
		nonce  uint64
		txHash common.Hash
	)

	p := ethutils.NewProvider(cCtx.String("rpc"), cCtx.Int64("chainid"), ethutils.WithDivviConsumerAddress(cCtx.String("divvi")))

	privateKey, err := crypto.HexToECDSA(cCtx.String("private-key"))
	if err != nil {
		return PublishTxResp{}, err
	}

	publicKey := crypto.PubkeyToAddress(privateKey.PublicKey)

	if err := p.Client.CallCtx(
		cCtx.Context,
		eth.Nonce(publicKey, nil).Returns(&nonce),
	); err != nil {
		return PublishTxResp{}, err
	}

	divviRefferal := p.GetReferalTag(publicKey)
	tx, err := p.SignContractPublishTx(
		privateKey,
		ethutils.ContractPublishTxOpts{
			ContractByteCode: ethutils.ConcatBytes(contractBytecode, divviRefferal),
			GasFeeCap:        big.NewInt(cCtx.Int64("gas-fee-cap")),
			GasTipCap:        big.NewInt(cCtx.Int64("gas-tip-cap")),
			GasLimit:         gasLimit,
			Nonce:            nonce,
		},
	)
	if err != nil {
		return PublishTxResp{}, err
	}
	rawTx, err := tx.MarshalBinary()
	if err != nil {
		return PublishTxResp{}, err
	}

	c.Logg.Debug("tx", "raw_tx", hexutil.Encode(rawTx))

	if err := p.Client.CallCtx(
		cCtx.Context,
		eth.SendTx(tx).Returns(&txHash),
	); err != nil {
		return PublishTxResp{}, err
	}

	_ = p.SubmitReferral(cCtx.Context, txHash)

	return PublishTxResp{
		TxHash:          txHash,
		ContractAddress: crypto.CreateAddress(publicKey, nonce),
	}, nil
}
