package modules

import (
	"log"
	"sync/atomic"
	"time"
	"github.com/jlaffaye/ftp"
	mdInter "ftp-client/modules/interface"
)

// function to try login
func TryLogin(client *ftp.ServerConn, user, passw string, found *int32, start time.Time) {
	// try login
	err := client.Login(user, passw)
	if err != nil {
		// if fail, log 
		if atomic.LoadInt32(found) == 0 {
			log.Printf("%vFailed%v to login with user: %v and password: %v", mdInter.Red, mdInter.Reset, user, passw)
		}
	} else {
		// if sucess
		if atomic.CompareAndSwapInt32(found, 0, 1) {
			crackTime := time.Since(start)
			crackSec := crackTime.Seconds()
			mdInter.CrackDone(user, passw, crackSec)
		}
	}
}