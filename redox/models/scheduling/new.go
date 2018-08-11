package scheduling

import (
	"github.com/minhajuddinkhan/gatewaygo/redox/models/common"
)

type New struct {
	Meta    common.Meta
	Patient struct {
		Identifiers  []common.Identifier
		Demographics common.Demographics
	}
	//	AppointmentInfo common.CodeSet
	Visit common.Visit
}
