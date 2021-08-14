package sessionStart

import (
	"time"

	. "github.com/Contra-Culture/efms/models/base"
)

var sessionStartRecordMarking = []int{
	LOGIN_FLEN,
	TOKEN_FLEN,
	CREATED_AT_FLEN,
}

type sessionStart struct {
	login     string
	token     string
	createdAt time.Time
}

func sessionStart2Rec(ss *sessionStart, handleErr func(error)) (rec []byte) {
	var err error
	var bs []byte
	rec, err = Normalize([]byte(ss.login), LOGIN_FLEN)
	if err != nil {
		handleErr(err)
		return []byte{}
	}
	bs, err = Normalize([]byte(ss.token), TOKEN_FLEN)
	if err != nil {
		handleErr(err)
		return []byte{}
	}
	rec = append(rec, bs...)
	bs, err = Normalize([]byte(ss.createdAt.Format(DATETIME_FORMAT)), CREATED_AT_FLEN)
	if err != nil {
		handleErr(err)
		return []byte{}
	}
	rec = append(rec, bs...)
	return
}
func rec2SessionStart(rec []byte, handleErr func(error)) *sessionStart {
	var err error
	var dt time.Time
	splitted := Split(rec, sessionStartRecordMarking)
	dt, err = time.Parse(string(splitted[2]), DATETIME_FORMAT)
	if err != nil {
		handleErr(err)
		return nil
	}
	ss := sessionStart{
		login:     string(splitted[0]),
		token:     string(splitted[1]),
		createdAt: dt,
	}
	return &ss
}
