package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/cascadiafoundation/cascadia/x/evm/types"
)

var (
	//go:embed compiled_contracts/VotingEscrow.json
	VotingEscrowJSON []byte // nolint: golint

	// VotingEscrowContract is the compiled erc20 contract
	VotingEscrowContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(VotingEscrowJSON, &VotingEscrowContract)
	if err != nil {
		panic(err)
	}

	if len(VotingEscrowContract.Bin) == 0 {
		panic("load contract failed")
	}
}
