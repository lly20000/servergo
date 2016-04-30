package connector

import (
	"common"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"myproto"
)

func PaserMsg(buf []byte) error {
	msgtype := common.BytesToInt64(buf[0:8])
	fmt.Println("msgtype = ", msgtype)
	switch msgtype {
	case 0: //Login Msg
		loginmsg := &myproto.Login{}
		err := proto.Unmarshal(buf, loginmsg)
		if err != nil {
			return errors.New("0 type msg error")
		}
		fmt.Println("loginmsg = ", loginmsg.Usrname)
	case 1:

	default:

	}
	return errors.New("msg type error")
}
