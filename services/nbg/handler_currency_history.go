package nbg

import "github.com/nats-io/nats.go"

func (s *NBGHistoryProvider) handleCurrencyHistory(msg *nats.Msg) {
	s.logger.Info("Received currency history request", "history", string(msg.Data))
}
