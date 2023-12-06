package contracts

import (
	"fmt"
	"math/big"

	"github.com/ic-n/ethcontractbind/pkg/tools"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

var _ = Describe("Contract test", Label("sol"), func() {
	var (
		auth *bind.TransactOpts
		cli  *backends.SimulatedBackend
		c    *CrowdFund
	)

	BeforeEach(func() {
		var err error
		auth, cli, c, err = testContract(common.Address{})
		Expect(err)
	})

	It("...", func() {
		_ = auth
		_ = cli
		_ = c
		// TODO: implement contract test
	})
})

const (
	simChainID = 1337
	gasLim     = 4712388
)

func testContract(token common.Address) (*bind.TransactOpts, *backends.SimulatedBackend, *CrowdFund, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to generate key: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(simChainID))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create transactor: %w", err)
	}
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: tools.Wei(1000)}
	client := backends.NewSimulatedBackend(alloc, gasLim)

	_, _, contract, err := DeployCrowdFund(auth, client, token)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to deploy contract: %w", err)
	}

	client.Commit()

	return auth, client, contract, nil
}
