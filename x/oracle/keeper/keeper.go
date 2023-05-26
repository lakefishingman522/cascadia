package keeper

import (
	"fmt"

	"time"

	"github.com/cascadiafoundation/cascadia/x/oracle/cosmosibckeeper"
	"github.com/cascadiafoundation/cascadia/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/bandprotocol/bandchain-packet/obi"
	"github.com/bandprotocol/bandchain-packet/packet"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
)

// Interface declearation
type OracleKeeperI interface {
	// Fetch latest price from asset and source
	GetLatestPriceFromAssetAndSource(sdk.Context, string, string) (types.Price, bool)
	// Fetch latest price from any source
	GetLatestPriceFromAnySource(sdk.Context, string) (types.Price, bool)
}

var _ OracleKeeperI = Keeper{}

type Keeper struct {
	*cosmosibckeeper.Keeper
	cdc        codec.BinaryCodec
	storeKey   storetypes.StoreKey
	memKey     storetypes.StoreKey
	paramstore paramtypes.Subspace
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	channelKeeper cosmosibckeeper.ChannelKeeper,
	portKeeper cosmosibckeeper.PortKeeper,
	scopedKeeper cosmosibckeeper.ScopedKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		Keeper: cosmosibckeeper.NewKeeper(
			types.PortKey,
			storeKey,
			channelKeeper,
			portKeeper,
			scopedKeeper,
		),
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Encode call data, convert the second param as uint8
func (k Keeper) EncodeCalldata(symbols []string, minSourceCnt uint64) ([]byte, error) {
	return obi.Encode(types.Calldata{Symbols: symbols, MinimumSourceCount: (uint8)(minSourceCnt)})
}

func (k Keeper) SendOracleRequest(ctx sdk.Context) {
	params := k.GetParams(ctx)
	if params.BandChannelSource == "" {
		return
	}
	sourcePort := types.PortID
	channelCap, ok := k.ScopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, params.BandChannelSource))
	if !ok {
		return
	}

	assetInfos := k.GetAllAssetInfo(ctx)
	symbols := []string{}
	for _, assetInfo := range assetInfos {
		if assetInfo.BandTicker != "" {
			symbols = append(symbols, assetInfo.BandTicker)
		}
	}

	if len(symbols) == 0 {
		return
	}

	encodedCalldata, _ := k.EncodeCalldata(symbols, params.Multiplier)
	packetData := packet.NewOracleRequestPacketData(
		params.ClientID,
		params.OracleScriptID,
		encodedCalldata,
		params.AskCount,
		params.MinCount,
		params.FeeLimit,
		params.PrepareGas,
		params.ExecuteGas,
	)

	k.ChannelKeeper.SendPacket(ctx, channelCap, sourcePort, params.BandChannelSource, clienttypes.NewHeight(0, 0), uint64(ctx.BlockTime().UnixNano()+int64(10*time.Minute)), packetData.GetBytes())
}
