package person

import (
	"time"

	. "github.com/Contra-Culture/efms/models/base"
)

const FIRST_NAME_FLEN = 32
const MIDDLE_NAME_FLEN = 32
const FAMILY_NAME_FLEN = 32
const BIRTH_DATE_FLEN = 64

var personRecordMarking = []int{
	LOGIN_FLEN,
	FIRST_NAME_FLEN,
	MIDDLE_NAME_FLEN,
	FAMILY_NAME_FLEN,
	BIRTH_DATE_FLEN,
}

type person struct {
	login      string
	firstName  string
	middleName string
	familyName string
	birthDate  time.Time
}

func person2Rec(p *person, handleErr func(error)) (rec []byte) {
	var err error
	var bs []byte
	rec, err = Normalize([]byte(p.login), LOGIN_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	bs, err = Normalize([]byte(p.firstName), FIRST_NAME_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	rec = append(rec, bs...)
	bs, err = Normalize([]byte(p.middleName), MIDDLE_NAME_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	rec = append(rec, bs...)
	bs, err = Normalize([]byte(p.familyName), FAMILY_NAME_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	rec = append(rec, bs...)
	bs, err = Normalize([]byte(p.birthDate.Format(DATETIME_FORMAT)), BIRTH_DATE_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	rec = append(rec, bs...)
	return
}
func rec2Person(rec []byte, handleErr func(error)) *person {
	splitted := Split(rec, personRecordMarking)
	dt, err := time.Parse(string(splitted[4]), DATETIME_FORMAT)
	if err != nil {
		handleErr(err)
		return nil
	}
	p := person{
		login:      string(splitted[0]),
		firstName:  string(splitted[1]),
		middleName: string(splitted[2]),
		familyName: string(splitted[3]),
		birthDate:  dt,
	}
	return &p
}
