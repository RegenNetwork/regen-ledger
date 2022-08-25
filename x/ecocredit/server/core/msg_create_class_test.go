package core

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type createClassSuite struct {
	*baseSuite
	alice                sdk.AccAddress
	aliceBalance         sdk.Coin
	creditTypeAbbrev     string
	allowedClassCreators []string
	creditClassFee       []*v1beta1.Coin
	classID              string
	res                  *core.MsgCreateClassResponse
	err                  error
}

func TestCreateClass(t *testing.T) {
	gocuke.NewRunner(t, &createClassSuite{}).Path("./features/msg_create_class.feature").Run()
}

func (s *createClassSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.creditTypeAbbrev = "C"
	s.classID = testClassID
}

func (s *createClassSuite) ACreditType() {
	err := s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *createClassSuite) ACreditTypeWithAbbreviation(a string) {
	err := s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: a,
	})
	require.NoError(s.t, err)
}

func (s *createClassSuite) AllowlistEnabled(a string) {
	allowlistEnabled, err := strconv.ParseBool(a)
	require.NoError(s.t, err)

	err = s.stateStore.AllowListEnabledTable().Save(s.ctx, &api.AllowListEnabled{
		Enabled: allowlistEnabled,
	})
	require.NoError(s.t, err)
}

func (s *createClassSuite) AliceIsAnApprovedCreditClassCreator() {
	s.allowedClassCreators = append(s.allowedClassCreators, s.alice.String())

	for _, creator := range s.allowedClassCreators {
		addr, err := sdk.AccAddressFromBech32(creator)
		require.NoError(s.t, err)

		err = s.stateStore.AllowedClassCreatorTable().Insert(s.ctx, &api.AllowedClassCreator{
			Address: addr,
		})
		require.NoError(s.t, err)
	}
}

func (s *createClassSuite) AllowedCreditClassFee(a string) {
	creditClassFee, err := sdk.ParseCoinsNormalized(a)
	require.NoError(s.t, err)

	for _, fee := range creditClassFee {
		s.creditClassFee = append(s.creditClassFee, &v1beta1.Coin{
			Denom:  fee.Denom,
			Amount: fee.Amount.String(),
		})
	}

	err = s.stateStore.ClassFeesTable().Save(s.ctx, &api.ClassFees{
		Fees: s.creditClassFee,
	})
	require.NoError(s.t, err)
}

func (s *createClassSuite) AliceHasATokenBalance(a string) {
	balance, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	s.aliceBalance = balance
}

func (s *createClassSuite) AClassSequenceWithCreditTypeAndNextSequence(a string, b string) {
	nextSequence, err := strconv.ParseUint(b, 10, 64)
	require.NoError(s.t, err)

	err = s.stateStore.ClassSequenceTable().Insert(s.ctx, &api.ClassSequence{
		CreditTypeAbbrev: a,
		NextSequence:     nextSequence,
	})
	require.NoError(s.t, err)
}

func (s *createClassSuite) AliceAttemptsToCreateACreditClass() {
	s.createClassExpectCalls()

	s.res, s.err = s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.alice.String(),
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
}

func (s *createClassSuite) AliceAttemptsToCreateACreditClassWithCreditType(a string) {
	s.createClassExpectCalls()

	s.res, s.err = s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.alice.String(),
		CreditTypeAbbrev: a,
	})
}

func (s *createClassSuite) AliceAttemptsToCreateACreditClassWithIssuers(a gocuke.DocString) {
	var issuers []string

	err := json.Unmarshal([]byte(a.Content), &issuers)
	require.NoError(s.t, err)

	s.createClassExpectCalls()

	s.res, s.err = s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.alice.String(),
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Issuers:          issuers,
	})
}

func (s *createClassSuite) AliceAttemptsToCreateACreditClassWithProperties(a gocuke.DocString) {
	var msg *core.MsgCreateClass

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	s.createClassExpectCalls()

	s.res, s.err = s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.alice.String(),
		CreditTypeAbbrev: msg.CreditTypeAbbrev,
		Metadata:         msg.Metadata,
	})
}

func (s *createClassSuite) AliceAttemptsToCreateACreditClassWithFee(a string) {
	fee, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	s.createClassExpectCalls()

	s.res, s.err = s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.alice.String(),
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Fee:              &fee,
	})
}

func (s *createClassSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *createClassSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *createClassSuite) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *createClassSuite) ExpectClassSequenceWithCreditTypeAndNextSequence(a string, b string) {
	nextSequence, err := strconv.ParseUint(b, 10, 64)
	require.NoError(s.t, err)

	classSequence, err := s.stateStore.ClassSequenceTable().Get(s.ctx, a)
	require.NoError(s.t, err)

	require.Equal(s.t, nextSequence, classSequence.NextSequence)
}

func (s *createClassSuite) ExpectClassIssuers(a gocuke.DocString) {
	var issuers []string

	err := json.Unmarshal([]byte(a.Content), &issuers)
	require.NoError(s.t, err)

	class, err := s.stateStore.ClassTable().GetById(s.ctx, s.classID)
	require.NoError(s.t, err)

	for _, issuer := range issuers {
		account, err := sdk.AccAddressFromBech32(issuer)
		require.NoError(s.t, err)

		exists, err := s.stateStore.ClassIssuerTable().Has(s.ctx, class.Key, account)
		require.NoError(s.t, err)
		require.True(s.t, exists)
	}
}

func (s *createClassSuite) ExpectClassProperties(a gocuke.DocString) {
	var expected *api.Class

	err := json.Unmarshal([]byte(a.Content), &expected)
	require.NoError(s.t, err)

	class, err := s.stateStore.ClassTable().GetById(s.ctx, expected.Id)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.CreditTypeAbbrev, class.CreditTypeAbbrev)
	require.Equal(s.t, expected.Metadata, class.Metadata)
}

func (s *createClassSuite) ExpectTheResponse(a gocuke.DocString) {
	var res *core.MsgCreateClassResponse

	err := json.Unmarshal([]byte(a.Content), &res)
	require.NoError(s.t, err)

	require.Equal(s.t, res, s.res)
}

func (s *createClassSuite) ExpectEventWithProperties(a gocuke.DocString) {
	var event core.EventCreateClass
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *createClassSuite) createClassExpectCalls() {
	var expectedFee sdk.Coin
	var expectedFees sdk.Coins

	if len(s.creditClassFee) == 1 {
		amount, ok := sdk.NewIntFromString(s.creditClassFee[0].Amount)
		require.True(s.t, ok)

		expectedFee = sdk.Coin{
			Denom:  s.creditClassFee[0].Denom,
			Amount: amount,
		}
		expectedFees = sdk.Coins{expectedFee}
	}

	if len(s.creditClassFee) == 2 {
		amount, ok := sdk.NewIntFromString(s.creditClassFee[1].Amount)
		require.True(s.t, ok)

		expectedFee = sdk.Coin{
			Denom:  s.creditClassFee[1].Denom,
			Amount: amount,
		}
		expectedFees = sdk.Coins{expectedFee}
	}

	s.bankKeeper.EXPECT().
		GetBalance(s.sdkCtx, s.alice, expectedFee.Denom).
		Return(s.aliceBalance).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SendCoinsFromAccountToModule(s.sdkCtx, s.alice, ecocredit.ModuleName, expectedFees).
		Do(func(sdk.Context, sdk.AccAddress, string, sdk.Coins) {
			if s.creditClassFee != nil {
				// simulate token balance update unavailable with mocks
				s.aliceBalance = s.aliceBalance.Sub(expectedFee)
			}
		}).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		BurnCoins(s.sdkCtx, ecocredit.ModuleName, expectedFees).
		Return(nil).
		AnyTimes() // not expected on failed attempt
}
