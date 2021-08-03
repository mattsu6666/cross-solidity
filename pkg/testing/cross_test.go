package testing

import (
	"context"
	"fmt"
	"testing"

	simpletypes "github.com/datachainlab/cross/x/core/atomic/protocol/simple/types"
	"github.com/datachainlab/cross/x/core/tx/types"
	xcctypes "github.com/datachainlab/cross/x/core/xcc/types"
	"github.com/datachainlab/cross/x/packets"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"

	"github.com/datachainlab/cross-solidity/pkg/consts"
	"github.com/datachainlab/cross-solidity/pkg/contract/crosssimplemodule"
)

const testMnemonicPhrase = "math razor capable expose worth grape metal sunset metal sudden usage scheme"

type CrossTestSuite struct {
	suite.Suite

	chain *Chain
}

func (suite *CrossTestSuite) SetupTest() {
	chain := NewChain(suite.T(), "http://127.0.0.1:8545", testMnemonicPhrase, consts.Contract)
	suite.chain = chain
}

func (suite *CrossTestSuite) TestRecvPacket() {
	ctx := context.Background()

	txID := []byte(fmt.Sprintf("txid-%v", 0))

	channel := xcctypes.ChannelInfo{}
	xcc, err := xcctypes.PackCrossChainChannel(&channel)
	suite.Require().NoError(err)

	pdc := simpletypes.NewPacketDataCall(txID, types.NewResolvedContractTransaction(xcc, nil, types.ContractCallInfo{}, nil, nil))
	pd := packets.NewPacketData(nil, pdc)
	packetData, err := proto.Marshal(&pd)
	suite.Require().NoError(err)

	suite.Require().NoError(suite.chain.TxSyncIfNoError(ctx)(
		suite.chain.CrossSimpleModule.OnRecvPacket(
			suite.chain.TxOpts(ctx, 0),
			crosssimplemodule.PacketData{
				Data: packetData,
			},
		),
	))
}

func TestChainTestSuite(t *testing.T) {
	suite.Run(t, new(CrossTestSuite))
}
