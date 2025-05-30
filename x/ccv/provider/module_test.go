package provider_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/cosmos/interchain-security/v7/testutil/crypto"
	testkeeper "github.com/cosmos/interchain-security/v7/testutil/keeper"
	"github.com/cosmos/interchain-security/v7/x/ccv/provider"
	"github.com/cosmos/interchain-security/v7/x/ccv/provider/types"
)

// Tests the provider's InitGenesis implementation against the spec.
// See: https://github.com/cosmos/ibc/blob/main/spec/app/ics-028-cross-chain-validation/methods.md#ccv-pcf-initg1
// Spec tag: [CCV-PCF-INITG.1]
//
// Note: Genesis validation for the provider is tested in TestValidateGenesisState
func TestInitGenesis(t *testing.T) {
	type testCase struct {
		name string
		// Whether port capability is already bound to the CCV provider module
		isBound bool
		// Provider's storage of consumer state to test against
		consumerStates []types.ConsumerState
		// Whether method call should panic, default: false
		expPanic bool
	}

	tests := []testCase{
		{
			name:           "already bound port, no consumer states",
			isBound:        true,
			consumerStates: []types.ConsumerState{},
		},
		{
			name:           "no bound port, no consumer states",
			isBound:        false,
			consumerStates: []types.ConsumerState{},
		},
		{
			name:    "no bound port, multiple consumer states",
			isBound: false,
			consumerStates: []types.ConsumerState{
				{
					ChainId:   "chainId1",
					ChannelId: "channelIdToChain1",
				},
				{
					ChainId:   "chainId2",
					ChannelId: "channelIdToChain2",
				},
				{
					ChainId:   "chainId3",
					ChannelId: "channelIdToChain3",
				},
			},
		},
		{
			name:    "already bound port, one consumer state",
			isBound: true,
			consumerStates: []types.ConsumerState{
				{
					ChainId:   "chainId77",
					ChannelId: "channelIdToChain77",
				},
			},
		},
	}

	for _, tc := range tests {
		//
		// Setup
		//
		keeperParams := testkeeper.NewInMemKeeperParams(t)
		providerKeeper, ctx, ctrl, mocks := testkeeper.GetProviderKeeperAndCtx(t, keeperParams)

		appModule := provider.NewAppModule(&providerKeeper, *keeperParams.ParamsSubspace, keeperParams.StoreKey)
		genState := types.NewGenesisState(
			providerKeeper.GetValidatorSetUpdateId(ctx),
			nil,
			tc.consumerStates,
			types.DefaultParams(),
			nil,
			nil,
			nil,
		)

		cdc := keeperParams.Cdc
		jsonBytes := cdc.MustMarshalJSON(genState)

		// Last total power is queried in InitGenesis, only if method has not
		// already panicked from unowned capability.
		if !tc.expPanic {
			// create a mock validator
			cId := crypto.NewCryptoIdentityFromIntSeed(234234)
			validator := cId.SDKStakingValidator()
			valAddr, err := sdk.ValAddressFromBech32(validator.GetOperator())
			require.NoError(t, err)

			orderedCalls := []*gomock.Call{
				mocks.MockStakingKeeper.EXPECT().GetLastTotalPower(
					ctx).Return(math.NewInt(100), nil).Times(1), // Return total voting power as 100
				mocks.MockStakingKeeper.EXPECT().GetBondedValidatorsByPower(
					ctx).Return([]stakingtypes.Validator{validator}, nil).Times(1), // Return a single validator
				mocks.MockStakingKeeper.EXPECT().GetLastValidatorPower(
					ctx, valAddr).Return(int64(100), nil).Times(1), // Return total power as power of the single validator
			}

			gomock.InOrder(orderedCalls...)

			//
			// Execute method, then assert expected results
			//
			if tc.expPanic {
				require.Panics(t, assert.PanicTestFunc(func() {
					appModule.InitGenesis(ctx, cdc, jsonBytes)
				}), tc.name)
				continue // Nothing else to verify
			}

			appModule.InitGenesis(ctx, cdc, jsonBytes)

			numStatesCounted := 0
			for _, state := range tc.consumerStates {
				numStatesCounted += 1
				channelID, found := providerKeeper.GetConsumerIdToChannelId(ctx, state.ChainId)
				require.True(t, found)
				require.Equal(t, state.ChannelId, channelID)

				chainID, found := providerKeeper.GetChannelIdToConsumerId(ctx, state.ChannelId)
				require.True(t, found)
				require.Equal(t, state.ChainId, chainID)
			}
			require.Equal(t, len(tc.consumerStates), numStatesCounted)

			// Expect slash meter to be initialized to it's allowance value
			// (replenish fraction * mocked value defined above)
			slashMeter := providerKeeper.GetSlashMeter(ctx)
			replenishFraction, err := math.LegacyNewDecFromStr(providerKeeper.GetParams(ctx).SlashMeterReplenishFraction)
			require.NoError(t, err)
			expectedSlashMeterValue := math.NewInt(replenishFraction.MulInt(math.NewInt(100)).RoundInt64())
			require.Equal(t, expectedSlashMeterValue, slashMeter)

			// Expect slash meter replenishment time candidate to be set to the current block time + replenish period
			expectedNextReplenishTime := ctx.BlockTime().Add(providerKeeper.GetSlashMeterReplenishPeriod(ctx))
			require.Equal(t, expectedNextReplenishTime, providerKeeper.GetSlashMeterReplenishTimeCandidate(ctx))

			ctrl.Finish()
		}
	}
}
