package main

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/bloombits"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/studyzy/simpleJsonrpc/ethapi"
)

type MockBackend struct {
}

func (m MockBackend) SyncProgress() ethereum.SyncProgress {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) FeeHistory(ctx context.Context, blockCount uint64, lastBlock rpc.BlockNumber, rewardPercentiles []float64) (*big.Int, [][]*big.Int, []*big.Int, []float64, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) ChainDb() ethdb.Database {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) AccountManager() *accounts.Manager {
	return nil
}

func (m MockBackend) ExtRPCEnabled() bool {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) RPCGasCap() uint64 {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) RPCEVMTimeout() time.Duration {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) RPCTxFeeCap() float64 {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) UnprotectedAllowed() bool {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) SetHead(number uint64) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) HeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Header, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) CurrentHeader() *types.Header {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) CurrentBlock() *types.Header {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) BlockByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) StateAndHeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) StateAndHeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*state.StateDB, *types.Header, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) Pending() (*types.Block, types.Receipts, *state.StateDB) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) GetTd(ctx context.Context, hash common.Hash) *big.Int {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) GetEVM(ctx context.Context, msg *core.Message, state *state.StateDB, header *types.Header, vmConfig *vm.Config, blockCtx *vm.BlockContext) *vm.EVM {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) GetTransaction(ctx context.Context, txHash common.Hash) (bool, *types.Transaction, common.Hash, uint64, uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) GetPoolTransactions() (types.Transactions, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) Stats() (pending int, queued int) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) TxPoolContent() (map[common.Address][]*types.Transaction, map[common.Address][]*types.Transaction) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) TxPoolContentFrom(addr common.Address) ([]*types.Transaction, []*types.Transaction) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) SubscribeNewTxsEvent(events chan<- core.NewTxsEvent) event.Subscription {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) ChainConfig() *params.ChainConfig {
	return TestChainConfig
}

var TestChainConfig = &params.ChainConfig{
	ChainID:                       big.NewInt(654321),
	HomesteadBlock:                big.NewInt(0),
	DAOForkBlock:                  nil,
	DAOForkSupport:                false,
	EIP150Block:                   big.NewInt(0),
	EIP155Block:                   big.NewInt(0),
	EIP158Block:                   big.NewInt(0),
	ByzantiumBlock:                big.NewInt(0),
	ConstantinopleBlock:           big.NewInt(0),
	PetersburgBlock:               big.NewInt(0),
	IstanbulBlock:                 big.NewInt(0),
	MuirGlacierBlock:              big.NewInt(0),
	BerlinBlock:                   big.NewInt(0),
	LondonBlock:                   big.NewInt(0),
	ArrowGlacierBlock:             big.NewInt(0),
	GrayGlacierBlock:              big.NewInt(0),
	MergeNetsplitBlock:            nil,
	ShanghaiTime:                  newUint64(0),
	CancunTime:                    nil,
	PragueTime:                    nil,
	VerkleTime:                    nil,
	TerminalTotalDifficulty:       nil,
	TerminalTotalDifficultyPassed: false,
	Ethash:                        new(params.EthashConfig),
	Clique:                        nil,
}

func newUint64(val uint64) *uint64 { return &val }

func (m MockBackend) Engine() consensus.Engine {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) GetBody(ctx context.Context, hash common.Hash, number rpc.BlockNumber) (*types.Body, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) GetLogs(ctx context.Context, blockHash common.Hash, number uint64) ([][]*types.Log, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) BloomStatus() (uint64, uint64) {
	//TODO implement me
	panic("implement me")
}

func (m MockBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	//TODO implement me
	panic("implement me")
}

var _ ethapi.Backend = (*MockBackend)(nil)
