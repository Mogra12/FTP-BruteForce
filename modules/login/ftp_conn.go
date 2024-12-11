package modules

import (
	ftp "github.com/jlaffaye/ftp"
	tls "crypto/tls"
)

func ConnectFTP(hostname string, tlsFlag bool) (*ftp.ServerConn, error) {
	var client *ftp.ServerConn
	var err error

	if tlsFlag {
		client, err = ftp.Dial(hostname, ftp.DialWithTLS(&tls.Config{
			InsecureSkipVerify: true, // ignore certificate validating
		}))
	} else {
		client, err = ftp.Dial(hostname)
	}

	return client, err
}
