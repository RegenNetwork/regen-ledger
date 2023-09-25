package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/RegenNetwork/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/RegenNetwork/regen-ledger/x/ecocredit/v3/base"
	types "github.com/RegenNetwork/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// CreateProject creates a new project for a specific credit class.
func (k Keeper) CreateProject(ctx context.Context, req *types.MsgCreateProject) (*types.MsgCreateProjectResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	classInfo, err := k.stateStore.ClassTable().GetById(ctx, req.ClassId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get class with id %s: %s", req.ClassId, err.Error())
	}

	adminAddress, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	err = k.assertClassIssuer(ctx, classInfo.Key, adminAddress)
	if err != nil {
		return nil, err
	}

	projectID, err := k.genProjectID(ctx, classInfo.Key, classInfo.Id)
	if err != nil {
		return nil, err
	}

	// check if non-empty reference id is unique within the scope of the credit class
	err = k.verifyReferenceID(ctx, classInfo.Key, req.ReferenceId)
	if err != nil {
		return nil, err
	}

	if err = k.stateStore.ProjectTable().Insert(ctx, &api.Project{
		Id:           projectID,
		Admin:        adminAddress,
		ClassKey:     classInfo.Key,
		Jurisdiction: req.Jurisdiction,
		Metadata:     req.Metadata,
		ReferenceId:  req.ReferenceId,
	}); err != nil {
		return nil, err
	}

	if err := sdkCtx.EventManager().EmitTypedEvent(&types.EventCreateProject{
		ProjectId: projectID,
	}); err != nil {
		return nil, err
	}

	return &types.MsgCreateProjectResponse{
		ProjectId: projectID,
	}, nil
}

// genProjectID generates a projectID when no projectID was given for CreateProject.
// The ID is generated by concatenating the classID and a sequence number.
func (k Keeper) genProjectID(ctx context.Context, classKey uint64, classID string) (string, error) {
	var nextSeq uint64
	projectSeqNo, err := k.stateStore.ProjectSequenceTable().Get(ctx, classKey)
	switch err {
	case ormerrors.NotFound:
		nextSeq = 1
	case nil:
		nextSeq = projectSeqNo.NextSequence
	default:
		return "", err
	}

	if err = k.stateStore.ProjectSequenceTable().Save(ctx, &api.ProjectSequence{
		ClassKey:     classKey,
		NextSequence: nextSeq + 1,
	}); err != nil {
		return "", err
	}

	return base.FormatProjectID(classID, nextSeq), nil
}

// verifyReferenceID prevents multiple projects from having the same reference id within the
// scope of a credit class. We verify this here at the message server level rather than at the
// ORM level because reference id is optional and therefore multiple projects within the scope
// of a credit class can have an empty reference id (see BridgeReceive for more information)
func (k Keeper) verifyReferenceID(ctx context.Context, classKey uint64, referenceID string) error {
	if referenceID == "" {
		// reference id is optional so an empty reference id is valid
		return nil
	}

	key := api.ProjectClassKeyReferenceIdIndexKey{}.WithClassKeyReferenceId(classKey, referenceID)
	it, err := k.stateStore.ProjectTable().List(ctx, key)
	if err != nil {
		return err
	}
	defer it.Close()
	if it.Next() {
		return sdkerrors.ErrInvalidRequest.Wrapf(
			"a project with reference id %s already exists within this credit class", referenceID,
		)
	}

	return nil
}
