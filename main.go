package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
	"github.com/jlaffaye/ftp"
	modules 	"ftp-client/modules"
	mdLogin 	"ftp-client/modules/login"
)

func main() {
	var found int32
	var wg sync.WaitGroup

	// load configuration
	config_modules := modules.LoadConfig()

	// check if required parameters were provided
	if config_modules.WordlistPath == "" || config_modules.Hostname == "" {
		fmt.Println("Invalid call. Usage: go run main.go -wl <wordlist_path> -h <hostname:21> -tls[OPTIONAL] -lwr[OPTIONAL]")
		return
	}

	// load the wordlist
	start := time.Now()
	wordlist := modules.Loader(config_modules.WordlistPath)
	semaphore := make(chan struct{}, config_modules.MaxConcurrent)

	// loop for login attempts
	for _, user := range wordlist {
		for _, passw := range wordlist {
			semaphore <- struct{}{} // block until space is available
			wg.Add(1)
			go func(user, passw string) {
				defer wg.Done()

				var client *ftp.ServerConn
				var err error
				// connect with TLS if needed
				if config_modules.TlsFlag {
					client, err = ftp.Dial(config_modules.Hostname, ftp.DialWithTLS(&tls.Config{
						InsecureSkipVerify: true, // ignore certificate verification
					}))
					if err != nil {
						log.Println(err)
						<-semaphore // release space
						return
					}
				} else {
					client, err = ftp.Dial(config_modules.Hostname)
				}
				if err != nil {
					log.Println(err)
					<-semaphore // release space
					return
				}
				defer client.Quit()

				// attempt login using the login module function
				mdLogin.TryLogin(client, user, passw, &found, start)
				<-semaphore // release space
			}(user, passw)
			if atomic.LoadInt32(&found) == 1 {
				break
			}
		}
		if atomic.LoadInt32(&found) == 1 {
			break
		}
	}

	// 1 second delay if lwrFlag is enabled
	if config_modules.SleepDuration > 0{
		time.Sleep(time.Duration(config_modules.SleepDuration) * time.Second)
	}

	// wait for all goroutines to finish
	wg.Wait()

	// if no user and password combination was found
	if atomic.LoadInt32(&found) == 0 {
		fmt.Printf("User or password not found in wordlist\n")
	}
}
