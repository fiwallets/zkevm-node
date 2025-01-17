package pkg

import (
	"github.com/fiwallets/zkevm-node/test/contracts/bin/ERC20"
	"github.com/fiwallets/zkevm-node/test/contracts/bin/uniswap/v2/core/UniswapV2Factory"
	"github.com/fiwallets/zkevm-node/test/contracts/bin/uniswap/v2/periphery/UniswapV2Router02"
	"github.com/fiwallets/go-ethereum/common"
)

type Deployments struct {
	ACoin     *ERC20.ERC20
	ACoinAddr common.Address
	BCoin     *ERC20.ERC20
	BCoinAddr common.Address
	CCoin     *ERC20.ERC20
	CCoinAddr common.Address
	Router    *UniswapV2Router02.UniswapV2Router02
	Factory   *UniswapV2Factory.UniswapV2Factory
}
