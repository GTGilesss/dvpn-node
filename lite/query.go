package lite

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	hubtypes "github.com/sentinel-official/hub/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	plantypes "github.com/sentinel-official/hub/x/plan/types"
	sessiontypes "github.com/sentinel-official/hub/x/session/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
	vpntypes "github.com/sentinel-official/hub/x/vpn/types"

	"github.com/sentinel-official/dvpn-node/utils"
)

func (c *Client) QueryAccount(address sdk.AccAddress) (authtypes.AccountI, error) {
	var (
		account authtypes.AccountI
		qc      = authtypes.NewQueryClient(c.ctx)
	)

	res, err := qc.Account(context.Background(),
		&authtypes.QueryAccountRequest{Address: address.String()})
	if err != nil {
		return nil, utils.IsNotFoundError(err)
	}

	if err := c.ctx.InterfaceRegistry.UnpackAny(res.Account, &account); err != nil {
		return nil, err
	}

	return account, nil
}

func (c *Client) QueryNode(address hubtypes.NodeAddress) (*nodetypes.Node, error) {
	var (
		qc = nodetypes.NewQueryServiceClient(c.ctx)
	)

	res, err := qc.QueryNode(context.Background(),
		nodetypes.NewQueryNodeRequest(address))
	if err != nil {
		return nil, utils.IsNotFoundError(err)
	}

	return &res.Node, nil
}

func (c *Client) QuerySubscription(id uint64) (*subscriptiontypes.Subscription, error) {
	var (
		qc = subscriptiontypes.NewQueryServiceClient(c.ctx)
	)

	res, err := qc.QuerySubscription(context.Background(),
		subscriptiontypes.NewQuerySubscriptionRequest(id))
	if err != nil {
		return nil, utils.IsNotFoundError(err)
	}

	return &res.Subscription, nil
}

func (c *Client) QueryQuota(id uint64, address sdk.AccAddress) (*subscriptiontypes.Quota, error) {
	var (
		qc = subscriptiontypes.NewQueryServiceClient(c.ctx)
	)

	res, err := qc.QueryQuota(context.Background(),
		subscriptiontypes.NewQueryQuotaRequest(id, address))
	if err != nil {
		return nil, utils.IsNotFoundError(err)
	}

	return &res.Quota, nil
}

func (c *Client) QuerySession(id uint64) (*sessiontypes.Session, error) {
	var (
		qc = sessiontypes.NewQueryServiceClient(c.ctx)
	)

	res, err := qc.QuerySession(context.Background(),
		sessiontypes.NewQuerySessionRequest(id))
	if err != nil {
		return nil, utils.IsNotFoundError(err)
	}

	return &res.Session, nil
}

func (c *Client) HasNodeForPlan(id uint64, address hubtypes.NodeAddress) (bool, error) {
	res, _, err := c.ctx.QueryStore(plantypes.NodeForPlanKey(id, address),
		fmt.Sprintf("%s/%s", vpntypes.ModuleName, plantypes.ModuleName))
	if err != nil {
		return false, err
	}
	if res == nil {
		return false, nil
	}

	return true, nil
}
