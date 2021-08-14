package sessionEnd

import (
	"time"

	. "github.com/Contra-Culture/efms/models/base"
)

var sessionEndRecordMarking = []int{
	LOGIN_FLEN,
	TOKEN_FLEN,
	CREATED_AT_FLEN,
}

type sessionEnd struct {
	login     string
	token     string
	createdAt time.Time
}

func sessionEnd2Rec(se *sessionEnd, handleErr func(error)) (rec []byte) {
	var err error
	var bs []byte
	rec, err = Normalize([]byte(se.login), LOGIN_FLEN)
	if err != nil {
		handleErr(err)
		return []byte{}
	}
	bs, err = Normalize([]byte(se.token), TOKEN_FLEN)
	if err != nil {
		handleErr(err)
		return []byte{}
	}
	rec = append(rec, bs...)
	bs, err = Normalize([]byte(se.createdAt.Format(DATETIME_FORMAT)), CREATED_AT_FLEN)
	if err != nil {
		handleErr(err)
		return []byte{}
	}
	rec = append(rec, bs...)
	return
}
func rec2SessionEnd(rec []byte, handlerErr func(error)) *sessionEnd {
	var err error
	var dt time.Time
	splitted := Split(rec, sessionEndRecordMarking)
	dt, err = time.Parse(string(splitted[2]), DATETIME_FORMAT)
	if err != nil {
		handlerErr(err)
		return nil
	}
	se := sessionEnd{
		login:     string(splitted[0]),
		token:     string(splitted[1]),
		createdAt: dt,
	}
	return &se
}
