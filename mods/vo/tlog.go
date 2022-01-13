package vo

type TLogMsg struct {
	Ids TLogIds        `json:"ids"`
	TS  UnixTimeMillis `json:"ts"`
	NFO TLogInfo       `json:"nfo"`
	TR  *TLogTR        `json:"tr,omitempty"`
}

type TLogIds struct {
	Evt        string         `json:"evt"`
	EvtTime    UnixTimeMillis `json:"evtTime"`
	DeviceId   string         `json:"deviceId"`
	Token      string         `json:"token"`
	TripId     string         `json:"tripId"`
	TripSerial int            `json:"tripSerial,omitempty"`
}

func (ids TLogIds) TripSerialIndex() int {
	if ids.TripSerial == 0 {
		return 0
	} else {
		return ids.TripSerial - 1
	}
}

type TLogInfo struct {
	Node  string `json:"node"`
	Debug string `json:"debug,omitempty"`
}

type TLogTR struct {
	RC             int            `json:"rc"`
	CellId         int            `json:"cellId"`
	PhysicalCellId int            `json:"physicalCellId"`
	RS             int            `json:"rs"`
	Positions      []TLogPosition `json:"positions"`
}

type TLogPosition struct {
	CT UnixTimeMillis `json:"ct"`
	LT float64        `json:"lt"`
	LN float64        `json:"ln"`
	AC float64        `json:"ac"`
	AL float64        `json:"al"`
	Gx float64        `json:"gx"`
	Gy float64        `json:"gy"`
	Gz float64        `json:"gz"`
	Ax float64        `json:"ax"`
	Ay float64        `json:"ay"`
	Az float64        `json:"az"`
}
