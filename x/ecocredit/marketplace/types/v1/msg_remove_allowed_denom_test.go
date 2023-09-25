package v1

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/RegenNetwork/gocuke"
	"github.com/stretchr/testify/require"
)

type msgRemoveAllowedDenomSuite struct {
	t         gocuke.TestingT
	msg       *MsgRemoveAllowedDenom
	err       error
	signBytes string
}

func TestMsgRemoveAllowedDenomSuite(t *testing.T) {
	gocuke.NewRunner(t, &msgRemoveAllowedDenomSuite{}).Path("./features/msg_remove_allowed_denom.feature").Run()
}

func (s *msgRemoveAllowedDenomSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgRemoveAllowedDenomSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgRemoveAllowedDenom{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgRemoveAllowedDenomSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgRemoveAllowedDenomSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgRemoveAllowedDenomSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgRemoveAllowedDenomSuite) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgRemoveAllowedDenomSuite) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
