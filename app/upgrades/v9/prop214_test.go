package v9_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/merlins-labs/merlins/v17/app/apptesting"
	v9 "github.com/merlins-labs/merlins/v17/app/upgrades/v9"
)

type UpgradeTestSuite struct {
	apptesting.KeeperTestHelper
}

func (suite *UpgradeTestSuite) SetupTest() {
	suite.Setup()
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (suite *UpgradeTestSuite) TestProp214() {
	poolId := suite.PrepareBalancerPool()
	v9.ExecuteProp214(suite.Ctx, suite.App.GAMMKeeper)

	_, err := suite.App.GAMMKeeper.GetPoolAndPoke(suite.Ctx, poolId)
	suite.Require().NoError(err)

	// Kept as comments for recordkeeping. Since SetPool is now private, the changes being tested for can no longer be made:
	// 		spreadFactor := pool.GetSpreadFactor(suite.Ctx)
	//  	expectedSpreadFactor := sdk.MustNewDecFromStr("0.002")
	//
	//  	suite.Require().Equal(expectedSpreadFactor, spreadFactor)
}
