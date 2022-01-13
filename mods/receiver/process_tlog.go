package receiver

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/heeseung815/zed-tard/mods/vo"
)

func (r *receiver) processTLog(msg *message.Message) {
	msg.Ack()

	tlogMsg := vo.TLogMsg{}
	if err := json.Unmarshal(msg.Payload, &tlogMsg); err != nil {
		r.log.Warnf("consumer failed to unmarshal tlog: %v\n%+v", err, msg.Payload)
		return
	}

}
