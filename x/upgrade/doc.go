package upgrade

/*
Package upgrade provides a Cosmos SDK module that can be used for smoothly upgrading a live Cosmos blockchain to a
new software version. It accomplishes this by providing a BeginBlocker hook that prevents the blockchain state
machine from proceeding once a pre-defined upgrade block height or time has been reached. The module does not prescribe
anything regarding how the decision to perform an upgrade is made by the chain's governance, only the mechanism
via which the software upgrade will be coordinated precisely and safely. Without software support for upgrades,
upgrading a live chain is risky because all of the validators need to pause their state machines at exactly the same
block height. Otherwise, there can be state inconsistencies which are hard to deal with.

Integrating With An App

Setup the upgrade.Keeper as usual:
	app.upgradeKeeper = upgrade.NewKeeper(app.upgradeStoreKey, cdc)

The app must define a BeginBlocker that calls the upgrade keepr's BeginBlocker method:
    func (app *myApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
    	app.upgradeKeeper.BeginBlocker(ctx, req)
    	return abci.ResponseBeginBlock{}
    }

The app must then integrate the upgrade keeper with its governance module as appropriate. Once governance has approved
an upgrade, the keeper's ScheduleUpgrade method must be called with the upgrade Plan that has been approved. Governance
should also enable a mechanism to cancel a pending upgrade that will call the keeper's ClearUpgradePlan method.

Performing Upgrades

Upgrades can be scheduled at either a predefined block height or time. Once this block height or time is reached, the
existing software will cease to function and a new version with code that handles the upgrade must be deployed. All
upgrades are coordinated by a unique upgrade name that cannot be reused on the same blockchain. In order for the upgrade
module to know that the upgrade has been safely applied, a handler with the name of the upgrade must be installed.
Here is an example handler for the upgrade named "my-fancy-upgrade":
	app.upgradeKeeper.SetUpgradeHandler("my-fancy-upgrade", func(ctx sdk.Context, plan upgrade.Plan) {
		// Perform any migrations of the state store needed for this upgrade
	})

This upgrade handler performs the dual function of alerting the upgrade module that this named upgrade has been applied
and thus the state machine can safely proceed, as well as providing the opportunity for the upgraded software to
perform any migrations of data in the state store before any new transactions are processed.

Default and Custom Shutdown Behavior

Before "crashing" the ABCI state machine in the BeginBlocker when an upgrade is required, the upgrade module
will log an error that looks like:
	UPGRADE "<Name>" NEEDED at height <NNNN>: <Info>
where Name are Info are the values of the respective fields on the upgrade Plan.

The default behavior for "crashing" the state machine is then to panic with an error which will stop the state machine
but usually not exit the process. This behavior can be overriden using the keeper's SetDoShutdowner method. A custom
shutdown method could perform some other steps such as alerting some other running process to try to install the
new software version automatically, as well as crashing the program with os.Exit() instead of panic.
*/
