package invitation

import (
	"time"

	. "github.com/Contra-Culture/efms/models/base"
)

var invitationRecordMarking = []int{
	TOKEN_FLEN,

	LOGIN_FLEN,
	EMAIL_FLEN,
	CREATED_AT_FLEN,
	MESSAGE_FLEN,
}

type invitation struct {
	token     string
	login     string
	email     string
	createdAt time.Time
	message   string
}

func invitation2Rec(i *invitation, handleErr func(error)) (rec []byte) {
	var err error
	var bs []byte
	rec, err = Normalize([]byte(i.token), TOKEN_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	bs, err = Normalize([]byte(i.login), LOGIN_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	rec = append(rec, bs...)
	bs, err = Normalize([]byte(i.email), EMAIL_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	rec = append(rec, bs...)
	bs, err = Normalize([]byte(i.createdAt.Format(DATETIME_FORMAT)), DATETIME_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	rec = append(rec, bs...)
	bs, err = Normalize([]byte(i.message), MESSAGE_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	rec = append(rec, bs...)
	return
}
func rec2Invitation(rec []byte, handleErr func(error)) *invitation {
	splitted := Split(rec, invitationRecordMarking)
	dt, err := time.Parse(string(splitted[3]), DATETIME_FORMAT)
	if err != nil {
		handleErr(err)
		return nil
	}
	i := invitation{
		token:     string(splitted[0]),
		login:     string(splitted[1]),
		email:     string(splitted[2]),
		createdAt: dt,
		message:   string(splitted[4]),
	}
	return &i
}
