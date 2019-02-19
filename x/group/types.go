/* GENERATED FROM README.org
   DO NOT EDIT THIS FILE DIRECTLY!!!!! */
package group

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// A group can be used to abstract over users and groups.
// It could be used to group individuals into a group or several groups/users into a larger group.
// It could be used by a single user to manage multiple devices and setup a multisig policy.
type Group struct {
	// The members of the group and their associated weight
	Members []Member `json:"addresses,omitempty"`
	// Specifies the number of votes that must be accumulated in order for a decision to be made by the group.
	// A member gets as many votes as is indicated by their Weight field.
	// A big integer is used here to avoid any potential vulnerabilities from overflow errors
	// where large weight and threshold values are used.
	DecisionThreshold sdk.Int `json:"decision_threshold"`
	// TODO maybe make this something more specific to a domain name or a claim on identity? or Memo leave it generic
	Memo string `json:"memo,omitempty"`
}

// A member specifies a address and a weight for a group member
type Member struct {
	// The address of a group member. Can be another group or a contract
	Address sdk.AccAddress `json:"address"`
	// The integral weight of this member with respect to other members and the decision threshold
	Weight sdk.Int `json:"weight"`
}

// Creates a group on the blockchain
// Should return a tag "group.id" with the bech32 address of the group
type MsgCreateGroup struct {
	Data   Group          `json:"data"`
	Signer sdk.AccAddress `json:"signer"`
}
