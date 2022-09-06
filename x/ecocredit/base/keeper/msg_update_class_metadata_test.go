//nolint:revive,stylecheck
package keeper

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/libs/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/base"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

type updateClassMetadata struct {
	*baseSuite
	alice   sdk.AccAddress
	bob     sdk.AccAddress
	res     *types.MsgUpdateClassMetadataResponse
	classId string
	err     error
}

func TestUpdateClassMetadata(t *testing.T) {
	gocuke.NewRunner(t, &updateClassMetadata{}).Path("./features/msg_update_class_metadata.feature").Run()
}

func (s *updateClassMetadata) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.bob = s.addr2
}

func (s *updateClassMetadata) ACreditTypeWithAbbreviation(a string) {
	// TODO: Save for now but credit type should not exist prior to unit test #893
	err := s.k.stateStore.CreditTypeTable().Save(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)
}

func (s *updateClassMetadata) ACreditClassWithClassIdAndAdminAlice(a string) {
	creditTypeAbbrev := base.GetCreditTypeAbbrevFromClassID(a)
	s.classId = a
	err := s.k.stateStore.ClassTable().Insert(s.ctx, &api.Class{
		Id:               a,
		Admin:            s.alice,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *updateClassMetadata) AliceAttemptsToUpdateClassMetadataWithClassId(a string) {
	s.res, s.err = s.k.UpdateClassMetadata(s.ctx, &types.MsgUpdateClassMetadata{
		Admin:   s.alice.String(),
		ClassId: a,
	})
}

func (s *updateClassMetadata) BobAttemptsToUpdateClassMetadataWithClassId(a string) {
	s.res, s.err = s.k.UpdateClassMetadata(s.ctx, &types.MsgUpdateClassMetadata{
		Admin:   s.bob.String(),
		ClassId: a,
	})
}

func (s *updateClassMetadata) AliceAttemptsToUpdateClassMetadataWithClassIdAndNewMetadata(a string, b gocuke.DocString) {
	s.res, s.err = s.k.UpdateClassMetadata(s.ctx, &types.MsgUpdateClassMetadata{
		Admin:       s.alice.String(),
		ClassId:     a,
		NewMetadata: b.Content,
	})
}

func (s *updateClassMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *updateClassMetadata) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *updateClassMetadata) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *updateClassMetadata) ExpectCreditClassWithClassIdAndMetadata(a string, b gocuke.DocString) {
	class, err := s.stateStore.ClassTable().GetById(s.ctx, a)
	require.NoError(s.t, err)

	require.Equal(s.t, b.Content, class.Metadata)
}

func (s *updateClassMetadata) AliceUpdatesTheClassMetadata() {
	newMetadata := rand.Str(5)
	_, err := s.k.UpdateClassMetadata(s.ctx, &types.MsgUpdateClassMetadata{
		Admin:       s.alice.String(),
		ClassId:     s.classId,
		NewMetadata: newMetadata,
	})
	require.NoError(s.t, err)
}

func (s *updateClassMetadata) ExpectEventWithProperties(a gocuke.DocString) {
	var event types.EventUpdateClassMetadata
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	events := s.sdkCtx.EventManager().Events()
	lastEvent := events[len(events)-1]

	require.Equal(s.t, proto.MessageName(&event), lastEvent.Type)
	require.Len(s.t, lastEvent.Attributes, 1) // should only have classID

	classID := strings.Trim(string(lastEvent.Attributes[0].Value), `"`)
	require.Equal(s.t, event.ClassId, classID)
}
