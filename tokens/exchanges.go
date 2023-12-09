package tokens

import (
	"github.com/unpackdev/solgo/exchanges"
	"github.com/unpackdev/solgo/utils"
)

func (t *Token) GetExchange(exchangeType utils.ExchangeType, simulatorType utils.SimulatorType) (exchanges.Exchange, bool) {
	if manager, found := t.exchanges[exchangeType]; found {
		if exchange, found := manager[simulatorType]; found {
			return exchange, true
		}
	}

	return nil, false
}

func (t *Token) RegisterExchange(exchangeType utils.ExchangeType, simulatorType utils.SimulatorType, exchange exchanges.Exchange) error {
	if _, found := t.exchanges[exchangeType]; !found {
		t.exchanges[exchangeType] = make(map[utils.SimulatorType]exchanges.Exchange)
	}

	if _, found := t.exchanges[exchangeType][simulatorType]; found {
		return nil
	}

	t.exchanges[exchangeType][simulatorType] = exchange
	return nil
}

func (t *Token) GetExchanges() map[utils.ExchangeType]map[utils.SimulatorType]exchanges.Exchange {
	return t.exchanges
}
