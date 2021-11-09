package types_test

import (
	"testing"
	"time"

	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	ibctmtypes "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	"github.com/cosmos/interchain-security/x/ccv/parent/types"

	"github.com/stretchr/testify/require"
)

const (
	chainID                      = "gaia"
	trustingPeriod time.Duration = time.Hour * 24 * 7 * 2
	ubdPeriod      time.Duration = time.Hour * 24 * 7 * 3
	maxClockDrift  time.Duration = time.Second * 10
)

var (
	height      = clienttypes.NewHeight(0, 4)
	upgradePath = []string{"upgrade", "upgradedIBCState"}
)

func TestValidateGenesisState(t *testing.T) {
	testCases := []struct {
		name     string
		genState *types.GenesisState
		expPass  bool
	}{
		{
			"valid initializing parent genesis with nil updates",
			types.NewGenesisState(
				[]types.ChildState{{"chainid-1", "channelid", 1}},
				types.DefaultParams(),
			),
			true,
		},
		{
			"valid validating parent genesis with nil updates",
			types.NewGenesisState(
				[]types.ChildState{{"chainid-1", "channelid", 2}},
				types.DefaultParams(),
			),
			true,
		},
		{
			"valid multiple parent genesis with multiple child chains",
			types.NewGenesisState(
				[]types.ChildState{
					{"chainid-1", "channelid", 2},
					{"chainid-2", "channelid2", 1},
					{"chainid-3", "channelid3", 0},
					{"chainid-4", "channelid4", 3},
				},
				types.DefaultParams(),
			),
			true,
		},
		{
			"valid parent genesis with custom params",
			types.NewGenesisState(
				[]types.ChildState{{"chainid-1", "channelid", 1}},
				types.NewParams(ibctmtypes.NewClientState("", ibctmtypes.DefaultTrustLevel, 0, 0,
					time.Second*40, clienttypes.Height{}, commitmenttypes.GetSDKSpecs(), []string{"ibc", "upgradedIBCState"}, true, false)),
			),
			true,
		},
		{
			"invalid params",
			types.NewGenesisState(
				[]types.ChildState{{"chainid-1", "channelid", 1}},
				types.NewParams(ibctmtypes.NewClientState("", ibctmtypes.DefaultTrustLevel, 0, 0,
					0, clienttypes.Height{}, nil, []string{"ibc", "upgradedIBCState"}, true, false)),
			),
			false,
		},
		{
			"invalid chain id",
			types.NewGenesisState(
				[]types.ChildState{{"invalidid{}", "channelid", 2}},
				types.DefaultParams(),
			),
			false,
		},
		{
			"invalid channel id",
			types.NewGenesisState(
				[]types.ChildState{{"chainid", "invalidchannel{}", 2}},
				types.DefaultParams(),
			),
			false,
		},
	}

	for _, tc := range testCases {
		err := tc.genState.Validate()

		if tc.expPass {
			require.NoError(t, err, "test case: %s must pass", tc.name)
		} else {
			require.Error(t, err, "test case: %s must fail", tc.name)
		}
	}
}
