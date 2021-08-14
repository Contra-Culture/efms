package invitationConfirmation

import (
	"time"

	. "github.com/Contra-Culture/efms/models/base"
)

var invitationConfirmationRecordMarking = []int{
	TOKEN_FLEN,
	CREATED_AT_FLEN,
}

type invitationConfirmation struct {
	token     string
	createdAt time.Time
}

func invitationConfirmation2Rec(ic *invitationConfirmation, handleErr func(error)) (rec []byte) {
	var err error
	var bs []byte
	rec, err = Normalize([]byte(ic.token), TOKEN_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	bs, err = Normalize([]byte(ic.createdAt.Format(DATETIME_FORMAT)), CREATED_AT_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	rec = append(rec, bs...)
	return
}
func rec2InvitationConfirmation(rec []byte, handleErr func(error)) *invitationConfirmation {
	splitted := Split(rec, invitationConfirmationRecordMarking)
	dt, err := time.Parse(string(splitted[1]), DATETIME_FORMAT)
	if err != nil {
		handleErr(err)
		return nil
	}
	ic := invitationConfirmation{
		token:     string(splitted[0]),
		createdAt: dt,
	}
	return &ic
}
