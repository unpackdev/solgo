package observers

import (
	"go.uber.org/zap"
)

type LogsProcessor struct {
	*Manager
	chLogs chan *LogEntry
	hooks  map[HookType]LogHookFn
}

func NewLogsProcessor(manager *Manager) (*LogsProcessor, error) {
	hooks := map[HookType]LogHookFn{}
	for hook, fn := range manager.GetHooks(LogProcessor) {
		hooks[hook] = fn.(LogHookFn)
	}

	return &LogsProcessor{
		Manager: manager,
		chLogs:  make(chan *LogEntry, 100000),
		hooks:   hooks,
	}, nil
}

func (p *LogsProcessor) Worker() error {
	for {
		select {
		case entry := <-p.chLogs:

			if hook, ok := p.hooks[PreHook]; ok {
				entry, err := hook(entry)
				if err != nil {
					zap.L().Error(
						"Pre log hook failed",
						zap.Error(err),
						zap.Any("network_id", entry.NetworkID),
						zap.String("network", entry.Network.String()),
						zap.Any("strategy", entry.Strategy),
						zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
						zap.String("tx_hash", entry.Transaction.Hash().String()),
					)
					continue
				}
			}

			zap.L().Debug(
				"Received new inbound transaction log",
				zap.Any("network_id", entry.NetworkID),
				zap.String("network", entry.Network.String()),
				zap.Any("strategy", entry.Strategy),
				zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
				zap.String("block_hash", entry.BlockHeader.Hash().String()),
				zap.String("tx_hash", entry.Transaction.Hash().String()),
				zap.String("contract_address", entry.ContractAddress.String()),
			)

			if hook, ok := p.hooks[PostHook]; ok {
				entry, err := hook(entry)
				if err != nil {
					zap.L().Error(
						"Post log hook failed",
						zap.Error(err),
						zap.Any("network_id", entry.NetworkID),
						zap.String("network", entry.Network.String()),
						zap.Any("strategy", entry.Strategy),
						zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
						zap.String("tx_hash", entry.Transaction.Hash().String()),
					)
					continue
				}
			}
		case <-p.ctx.Done():
			return nil
		}
	}
}

func (p *LogsProcessor) QueueLog(entry *LogEntry) error {
	p.chLogs <- entry
	return nil
}
