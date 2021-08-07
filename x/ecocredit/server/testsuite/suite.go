package testsuite

import (
	"context"
	"time"

	"github.com/regen-network/regen-ledger/types/testutil"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

type IntegrationTestSuite struct {
	suite.Suite

	fixtureFactory testutil.FixtureFactory
	fixture        testutil.Fixture

	sdkCtx            sdk.Context
	ctx               context.Context
	msgClient         ecocredit.MsgClient
	queryClient       ecocredit.QueryClient
	paramsQueryClient params.QueryClient
	signers           []sdk.AccAddress

	paramSpace paramstypes.Subspace
	bankKeeper bankkeeper.Keeper
}

func NewIntegrationTestSuite(fixtureFactory testutil.FixtureFactory, paramSpace paramstypes.Subspace, bankKeeper bankkeeper.BaseKeeper) *IntegrationTestSuite {
	return &IntegrationTestSuite{
		fixtureFactory: fixtureFactory,
		paramSpace:     paramSpace,
		bankKeeper:     bankKeeper,
	}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.fixture = s.fixtureFactory.Setup()

	// TODO clean up once types.Context merged upstream into sdk.Context
	s.sdkCtx, _ = s.fixture.Context().(types.Context).CacheContext()
	s.ctx = types.Context{Context: s.sdkCtx}

	ecocreditParams := ecocredit.DefaultParams()
	s.paramSpace.SetParamSet(s.sdkCtx, &ecocreditParams)

	s.signers = s.fixture.Signers()
	s.Require().GreaterOrEqual(len(s.signers), 8)
	s.msgClient = ecocredit.NewMsgClient(s.fixture.TxConn())
	s.queryClient = ecocredit.NewQueryClient(s.fixture.QueryConn())
	s.paramsQueryClient = params.NewQueryClient(s.fixture.QueryConn())
}

func fundAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}
	return bankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, amounts)
}

func (s *IntegrationTestSuite) TestScenario() {
	designer := s.signers[0]
	issuer1 := s.signers[1].String()
	issuer2 := s.signers[2].String()
	addr1 := s.signers[3].String()
	addr2 := s.signers[4].String()
	addr3 := s.signers[5].String()
	addr4 := s.signers[6].String()
	addr5 := s.signers[7].String()

	// create class with insufficient funds and it should fail
	createClsRes, err := s.msgClient.CreateClass(s.ctx, &ecocredit.MsgCreateClass{
		Designer: designer.String(),
		Issuers:  []string{issuer1, issuer2},
		Metadata: nil,
	})
	s.Require().Error(err)
	s.Require().Nil(createClsRes)

	// create class with sufficient funds and it should succeed
	s.Require().NoError(fundAccount(s.bankKeeper, s.sdkCtx, designer, sdk.NewCoins(sdk.NewInt64Coin("stake", 10000))))

	createClsRes, err = s.msgClient.CreateClass(s.ctx, &ecocredit.MsgCreateClass{
		Designer: designer.String(),
		Issuers:  []string{issuer1, issuer2},
		Metadata: nil,
	})
	s.Require().NoError(err)
	s.Require().NotNil(createClsRes)

	clsID := createClsRes.ClassId
	s.Require().NotEmpty(clsID)

	// designer should have no funds remaining
	s.Require().Equal(s.bankKeeper.GetBalance(s.sdkCtx, designer, "stake"), sdk.NewInt64Coin("stake", 0))

	// create batch
	t0, t1, t2 := "10.37", "1007.3869", "100"
	tSupply0 := "1117.7569"
	r0, r1, r2 := "4.286", "10000.4589902", "0"
	rSupply0 := "10004.7449902"

	time1 := time.Now()
	time2 := time.Now()

	// Batch creation should succeed with StartDate before EndDate, and valid data
	createBatchRes, err := s.msgClient.CreateBatch(s.ctx, &ecocredit.MsgCreateBatch{
		Issuer:          issuer1,
		ClassId:         clsID,
		StartDate:       &time1,
		EndDate:         &time2,
		ProjectLocation: "AB",
		Issuance: []*ecocredit.MsgCreateBatch_BatchIssuance{
			{
				Recipient:          addr1,
				TradableAmount:     t0,
				RetiredAmount:      r0,
				RetirementLocation: "GB",
			},
			{
				Recipient:          addr2,
				TradableAmount:     t1,
				RetiredAmount:      r1,
				RetirementLocation: "BF",
			},
			{
				Recipient:          addr4,
				TradableAmount:     t2,
				RetiredAmount:      r2,
				RetirementLocation: "",
			},
			{
				Recipient:          addr5,
				RetirementLocation: "",
			},
		},
	})
	s.Require().NoError(err)
	s.Require().NotNil(createBatchRes)

	batchDenom := createBatchRes.BatchDenom
	s.Require().NotEmpty(batchDenom)

	// query balances
	queryBalanceRes, err := s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
		Account:    addr1,
		BatchDenom: batchDenom,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryBalanceRes)
	s.Require().Equal(t0, queryBalanceRes.TradableAmount)
	s.Require().Equal(r0, queryBalanceRes.RetiredAmount)

	queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
		Account:    addr2,
		BatchDenom: batchDenom,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryBalanceRes)
	s.Require().Equal(t1, queryBalanceRes.TradableAmount)
	s.Require().Equal(r1, queryBalanceRes.RetiredAmount)

	queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
		Account:    addr4,
		BatchDenom: batchDenom,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryBalanceRes)
	s.Require().Equal(t2, queryBalanceRes.TradableAmount)
	s.Require().Equal(r2, queryBalanceRes.RetiredAmount)

	// if we didn't issue tradable or retired balances, they'll be default to zero.
	queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
		Account:    addr5,
		BatchDenom: batchDenom,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryBalanceRes)
	s.Require().Equal("0", queryBalanceRes.TradableAmount)
	s.Require().Equal("0", queryBalanceRes.RetiredAmount)

	// query supply
	querySupplyRes, err := s.queryClient.Supply(s.ctx, &ecocredit.QuerySupplyRequest{BatchDenom: batchDenom})
	s.Require().NoError(err)
	s.Require().NotNil(querySupplyRes)
	s.Require().Equal(tSupply0, querySupplyRes.TradableSupply)
	s.Require().Equal(rSupply0, querySupplyRes.RetiredSupply)

	// cancel credits
	cancelCases := []struct {
		name               string
		holder             string
		toCancel           string
		expectErr          bool
		expTradable        string
		expTradableSupply  string
		expRetired         string
		expTotalAmount     string
		expAmountCancelled string
	}{
		{
			name:      "can't cancel more credits than are tradable",
			holder:    addr4,
			toCancel:  "101",
			expectErr: true,
		},
		{
			name:      "can't cancel no credits",
			holder:    addr4,
			toCancel:  "0",
			expectErr: true,
		},
		{
			name:      "can't cancel beyond precision of batch",
			holder:    addr4,
			toCancel:  "0.00000001",
			expectErr: true,
		},
		{
			name:               "can cancel a small amount of credits",
			holder:             addr4,
			toCancel:           "2.0002",
			expectErr:          false,
			expTradable:        "97.9998",
			expTradableSupply:  "1115.7567",
			expRetired:         "0",
			expTotalAmount:     "11120.5016902",
			expAmountCancelled: "2.0002",
		},
		{
			name:               "can cancel all remaining credits",
			holder:             addr4,
			toCancel:           "97.9998",
			expectErr:          false,
			expTradable:        "0",
			expTradableSupply:  "1017.7569",
			expRetired:         "0",
			expTotalAmount:     "11022.5018902",
			expAmountCancelled: "100.0000",
		},
		{
			name:      "can't cancel anymore credits",
			holder:    addr4,
			toCancel:  "1",
			expectErr: true,
		},
		{
			name:               "can cancel from account with positive retired balance",
			holder:             addr1,
			toCancel:           "1",
			expectErr:          false,
			expTradable:        "9.37",
			expTradableSupply:  "1016.7569",
			expRetired:         "4.286",
			expTotalAmount:     "11021.5018902",
			expAmountCancelled: "101.0000",
		},
	}

	for _, tc := range cancelCases {
		s.Run(tc.name, func() {
			_, err := s.msgClient.Cancel(s.ctx, &ecocredit.MsgCancel{
				Holder: tc.holder,
				Credits: []*ecocredit.MsgCancel_CancelCredits{
					{
						BatchDenom: batchDenom,
						Amount:     tc.toCancel,
					},
				},
			})

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				// query balance
				queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
					Account:    tc.holder,
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.Require().Equal(tc.expTradable, queryBalanceRes.TradableAmount)
				s.Require().Equal(tc.expRetired, queryBalanceRes.RetiredAmount)

				// query supply
				querySupplyRes, err = s.queryClient.Supply(s.ctx, &ecocredit.QuerySupplyRequest{BatchDenom: batchDenom})
				s.Require().NoError(err)
				s.Require().NotNil(querySupplyRes)
				s.Require().Equal(tc.expTradableSupply, querySupplyRes.TradableSupply)
				s.Require().Equal(rSupply0, querySupplyRes.RetiredSupply)

				// query batchInfo
				queryBatchInfoRes, err := s.queryClient.BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: batchDenom})
				s.Require().NoError(err)
				s.Require().NotNil(queryBatchInfoRes)
				s.Require().Equal(tc.expTotalAmount, queryBatchInfoRes.Info.TotalAmount)
				s.Require().Equal(tc.expAmountCancelled, queryBatchInfoRes.Info.AmountCancelled)
			}
		})
	}

	// retire credits
	retireCases := []struct {
		name               string
		toRetire           string
		retirementLocation string
		expectErr          bool
		expTradable        string
		expRetired         string
		expTradableSupply  string
		expRetiredSupply   string
	}{
		{
			name:               "cannot retire more credits than are tradable",
			toRetire:           "10.371",
			retirementLocation: "AF",
			expectErr:          true,
		},
		{
			name:               "can't use more than 7 decimal places",
			toRetire:           "10.00000001",
			retirementLocation: "AF",
			expectErr:          true,
		},
		{
			name:               "can't retire to an invalid country",
			toRetire:           "0.0001",
			retirementLocation: "ZZZ",
			expectErr:          true,
		},
		{
			name:               "can't retire to an invalid region",
			toRetire:           "0.0001",
			retirementLocation: "AF-ZZZZ",
			expectErr:          true,
		},
		{
			name:               "can't retire to an invalid postal code",
			toRetire:           "0.0001",
			retirementLocation: "AF-BDS 0123456789012345678901234567890123456789012345678901234567890123456789",
			expectErr:          true,
		},
		{
			name:               "can't retire without a location",
			toRetire:           "0.0001",
			retirementLocation: "",
			expectErr:          true,
		},
		{
			name:               "can retire a small amount of credits",
			toRetire:           "0.0001",
			retirementLocation: "AF",
			expectErr:          false,
			expTradable:        "9.3699",
			expRetired:         "4.2861",
			expTradableSupply:  "1016.7568",
			expRetiredSupply:   "10004.7450902",
		},
		{
			name:               "can retire more credits",
			toRetire:           "9",
			retirementLocation: "AF-BDS",
			expectErr:          false,
			expTradable:        "0.3699",
			expRetired:         "13.2861",
			expTradableSupply:  "1007.7568",
			expRetiredSupply:   "10013.7450902",
		},
		{
			name:               "can retire all credits",
			toRetire:           "0.3699",
			retirementLocation: "AF-BDS 12345",
			expectErr:          false,
			expTradable:        "0",
			expRetired:         "13.656",
			expTradableSupply:  "1007.3869",
			expRetiredSupply:   "10014.1149902",
		},
		{
			name:      "can't retire any more credits",
			toRetire:  "1",
			expectErr: true,
		},
	}

	for _, tc := range retireCases {
		tc := tc
		s.Run(tc.name, func() {
			_, err := s.msgClient.Retire(s.ctx, &ecocredit.MsgRetire{
				Holder: addr1,
				Credits: []*ecocredit.MsgRetire_RetireCredits{
					{
						BatchDenom: batchDenom,
						Amount:     tc.toRetire,
					},
				},
				Location: tc.retirementLocation,
			})

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				// query balance
				queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
					Account:    addr1,
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.Require().Equal(tc.expTradable, queryBalanceRes.TradableAmount)
				s.Require().Equal(tc.expRetired, queryBalanceRes.RetiredAmount)

				// query supply
				querySupplyRes, err = s.queryClient.Supply(s.ctx, &ecocredit.QuerySupplyRequest{BatchDenom: batchDenom})
				s.Require().NoError(err)
				s.Require().NotNil(querySupplyRes)
				s.Require().Equal(tc.expTradableSupply, querySupplyRes.TradableSupply)
				s.Require().Equal(tc.expRetiredSupply, querySupplyRes.RetiredSupply)
			}
		})
	}

	sendCases := []struct {
		name                 string
		sendTradable         string
		sendRetired          string
		retirementLocation   string
		expectErr            bool
		expTradableSender    string
		expRetiredSender     string
		expTradableRecipient string
		expRetiredRecipient  string
		expTradableSupply    string
		expRetiredSupply     string
	}{
		{
			name:               "can't send more tradable than is tradable",
			sendTradable:       "2000",
			sendRetired:        "10",
			retirementLocation: "AF",
			expectErr:          true,
		},
		{
			name:               "can't send more retired than is tradable",
			sendTradable:       "10",
			sendRetired:        "2000",
			retirementLocation: "AF",
			expectErr:          true,
		},
		{
			name:               "can't send to an invalid country",
			sendTradable:       "10",
			sendRetired:        "20",
			retirementLocation: "ZZZ",
			expectErr:          true,
		},
		{
			name:               "can't send to an invalid region",
			sendTradable:       "10",
			sendRetired:        "20",
			retirementLocation: "AF-ZZZZ",
			expectErr:          true,
		},
		{
			name:               "can't send to an invalid postal code",
			sendTradable:       "10",
			sendRetired:        "20",
			retirementLocation: "AF-BDS 0123456789012345678901234567890123456789012345678901234567890123456789",
			expectErr:          true,
		},
		{
			name:                 "can send some",
			sendTradable:         "10",
			sendRetired:          "20",
			retirementLocation:   "AF",
			expectErr:            false,
			expTradableSender:    "977.3869",
			expRetiredSender:     "10000.4589902",
			expTradableRecipient: "10",
			expRetiredRecipient:  "20",
			expTradableSupply:    "987.3869",
			expRetiredSupply:     "10034.1149902",
		},
		{
			name:                 "can send with no retirement location",
			sendTradable:         "10",
			sendRetired:          "0",
			retirementLocation:   "",
			expectErr:            false,
			expTradableSender:    "967.3869",
			expRetiredSender:     "10000.4589902",
			expTradableRecipient: "20",
			expRetiredRecipient:  "20",
			expTradableSupply:    "987.3869",
			expRetiredSupply:     "10034.1149902",
		},
		{
			name:                 "can send all tradable",
			sendTradable:         "67.3869",
			sendRetired:          "900",
			retirementLocation:   "AF",
			expectErr:            false,
			expTradableSender:    "0",
			expRetiredSender:     "10000.4589902",
			expTradableRecipient: "87.3869",
			expRetiredRecipient:  "920",
			expTradableSupply:    "87.3869",
			expRetiredSupply:     "10934.1149902",
		},
		{
			name:         "can't send any more",
			sendTradable: "1",
			sendRetired:  "1",
			expectErr:    true,
		},
	}

	for _, tc := range sendCases {
		tc := tc
		s.Run(tc.name, func() {
			_, err := s.msgClient.Send(s.ctx, &ecocredit.MsgSend{
				Sender:    addr2,
				Recipient: addr3,
				Credits: []*ecocredit.MsgSend_SendCredits{
					{
						BatchDenom:         batchDenom,
						TradableAmount:     tc.sendTradable,
						RetiredAmount:      tc.sendRetired,
						RetirementLocation: tc.retirementLocation,
					},
				},
			})

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				// query sender balance
				queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
					Account:    addr2,
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.Require().Equal(tc.expTradableSender, queryBalanceRes.TradableAmount)
				s.Require().Equal(tc.expRetiredSender, queryBalanceRes.RetiredAmount)

				// query recipient balance
				queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
					Account:    addr3,
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.Require().Equal(tc.expTradableRecipient, queryBalanceRes.TradableAmount)
				s.Require().Equal(tc.expRetiredRecipient, queryBalanceRes.RetiredAmount)

				// query supply
				querySupplyRes, err = s.queryClient.Supply(s.ctx, &ecocredit.QuerySupplyRequest{BatchDenom: batchDenom})
				s.Require().NoError(err)
				s.Require().NotNil(querySupplyRes)
				s.Require().Equal(tc.expTradableSupply, querySupplyRes.TradableSupply)
				s.Require().Equal(tc.expRetiredSupply, querySupplyRes.RetiredSupply)
			}
		})
	}

	/****   TEST SET PRECISION   ****/
	precisionCases := []struct {
		name string
		msg  ecocredit.MsgSetPrecision
		ok   bool
	}{
		{
			"can NOT decrease the decimals", ecocredit.MsgSetPrecision{
				Issuer: issuer1, BatchDenom: batchDenom, MaxDecimalPlaces: 2},
			false,
		}, {
			"can NOT set to the same value", ecocredit.MsgSetPrecision{
				Issuer: issuer1, BatchDenom: batchDenom, MaxDecimalPlaces: 7},
			false,
		}, {
			"can increase", ecocredit.MsgSetPrecision{
				Issuer: issuer1, BatchDenom: batchDenom, MaxDecimalPlaces: 8},
			true,
		}, {
			"can NOT change precision of not existing denom", ecocredit.MsgSetPrecision{
				Issuer: issuer1, BatchDenom: "not/existing", MaxDecimalPlaces: 1},
			false,
		},
	}
	require := s.Require()
	for _, tc := range precisionCases {
		tc := tc
		s.Run(tc.name, func() {
			_, err := s.msgClient.SetPrecision(s.ctx, &tc.msg)

			if !tc.ok {
				require.Error(err)
			} else {
				require.NoError(err)
				res, err := s.queryClient.Precision(s.ctx,
					&ecocredit.QueryPrecisionRequest{
						BatchDenom: tc.msg.BatchDenom})
				require.NoError(err)
				require.Equal(tc.msg.MaxDecimalPlaces, res.MaxDecimalPlaces)
			}
		})
	}

	/****   TEST ALLOWLIST CREDIT DESIGNERS   ****/
	allowlistCases := []struct {
		name             string
		designerAcc      sdk.AccAddress
		allowlist        []string
		allowlistEnabled bool
		wantErr          bool
	}{
		{
			name:             "valid allowlist and enabled",
			allowlist:        []string{s.signers[0].String()},
			designerAcc:      s.signers[0],
			allowlistEnabled: true,
			wantErr:          false,
		},
		{
			name:             "valid multi addrs in allowlist",
			allowlist:        []string{s.signers[0].String(), s.signers[1].String(), s.signers[2].String()},
			designerAcc:      s.signers[0],
			allowlistEnabled: true,
			wantErr:          false,
		},
		{
			name:             "designer is not part of the allowlist",
			allowlist:        []string{s.signers[0].String()},
			designerAcc:      s.signers[1],
			allowlistEnabled: true,
			wantErr:          true,
		},
		{
			name:             "valid allowlist but disabled - anyone can create credits",
			allowlist:        []string{s.signers[0].String()},
			designerAcc:      s.signers[0],
			allowlistEnabled: false,
			wantErr:          false,
		},
		{
			name:             "empty and enabled allowlist - nobody can create credits",
			allowlist:        []string{},
			designerAcc:      s.signers[0],
			allowlistEnabled: true,
			wantErr:          true,
		},
	}

	for _, tc := range allowlistCases {
		tc := tc
		s.Run(tc.name, func() {
			s.paramSpace.Set(s.sdkCtx, ecocredit.KeyAllowedClassDesigners, tc.allowlist)
			s.paramSpace.Set(s.sdkCtx, ecocredit.KeyAllowlistEnabled, tc.allowlistEnabled)

			// fund the designer account
			s.Require().NoError(fundAccount(s.bankKeeper, s.sdkCtx, tc.designerAcc, sdk.NewCoins(sdk.NewInt64Coin("stake", 10000))))

			createClsRes, err = s.msgClient.CreateClass(s.ctx, &ecocredit.MsgCreateClass{
				Designer: tc.designerAcc.String(),
				Issuers:  []string{issuer1, issuer2},
				Metadata: nil,
			})
			if tc.wantErr {
				s.Require().Error(err)
				s.Require().Nil(createClsRes)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(createClsRes)
			}
		})
	}
}
