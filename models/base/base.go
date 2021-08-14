package base

import (
	"fmt"
	"time"
)

const DATETIME_FORMAT = time.RFC3339Nano
const DATETIME_FLEN = len(DATETIME_FORMAT)

const LOGIN_FLEN = 32
const PASSWORD_FLEN = 64
const EMAIL_FLEN = 64
const TOKEN_FLEN = 64
const MESSAGE_FLEN = 255
const CREATED_AT_FLEN = DATETIME_FLEN
const CONFIRMED_AT_FLEN = DATETIME_FLEN
const ALIGNB = byte(' ')

func Normalize(val []byte, flen int) (rec []byte, err error) {
	l := len(val)
	diff := flen - l
	if diff < 0 {
		err = fmt.Errorf("got value of len: %d, max len expected: %d", l, flen)
		return
	}
	rec = append(rec, val...)
	for i := diff; i > 0; i-- {
		rec = append(rec, ALIGNB)
	}
	return
}
func Split(rec []byte, marking []int) (splitted [][]byte) {
	splitted[0] = rec[:marking[1]-1]
	splitted[1] = rec[marking[1] : marking[2]-1]
	splitted[2] = rec[marking[2] : marking[3]-1]
	return
}
