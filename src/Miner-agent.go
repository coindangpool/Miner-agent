package main

import (
	"net/url"
	"io/ioutil"
	"net/http"
	"time"
	"net"
	"os"
	"flag"
	"encoding/json"
	"github.com/Elbandi/go-ccminer-api"
	"github.com/Nitron/go-cgminer-api"
	"github.com/yinhm/sgminer"
	"github.com/murlokswarm/log"
)

func main() {
	workername := flag.String("workername", "", "workername (Required)")
	miner := flag.String("miner", "", "miner {claymore|ewbfminer|ccminer|cgminer|sgminer} (Optional),\n If empty find miner automatically")
	currency := flag.String("currency", "", "currency {zcash|ethereerum} (Optional)")

	log.Info("Coindang Pool Monitoring V0.1 (beta)")

	if *workername == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	ok := false
	if *miner == "claymore" ||
		*miner == "ewbfminer" ||
		*miner == "ccminer" ||
		*miner == "cgminer" ||
		*miner == "sgminer" {
		ok = true
	} else {
		log.Warn("use miner option {claymore|ewbfminer|ccminer|cgminer|sgminer}")
	}

	//var miner string

	for {
		time.Sleep(3 * time.Second)
		var raw string
		if !ok || *miner == "claymore" {
			ok, raw = getClaymoreStatus()
			if ok {
				*miner = "claymore"
			} else {
				*miner = ""
				log.Infof("%s coundn't connect", "claymore")
			}
		}

		if !ok || *miner == "ewbfminer" {
			ok, raw = getEWBFStatus()
			if ok {
				*miner = "ewbfminer"
			} else {
				*miner = ""
				log.Infof("%s coundn't connect", "ewbfminer")
			}
		}

		if !ok || *miner == "ccminer" {
			ok, raw = getCCMinerStatus()
			if ok {
				*miner = "ccminer"
			} else {
				*miner = ""
				log.Infof("%s coundn't connect", "ccminer")
			}
		}

		if !ok || *miner == "cgminer" {
			ok, raw = getCGMinerStatus()
			if ok {
				*miner = "cgminer"
			} else {
				*miner = ""
				log.Infof("%s coundn't connect", "cgminer")
			}
		}

		if !ok || *miner == "sgminer" {
			ok, raw = getSGMinerStatus()
			if ok {
				*miner = "sgminer"
			} else {
				*miner = ""
				log.Infof("%s coundn't connect", "sgminer")
			}
		}

		if ok {
			log.Infof("Current found miner is %s", miner)
			setWorkerStatus(*workername, *miner, *currency, raw)
		} else {
			log.Warnf("Any miner coudn't find")
		}

		time.Sleep(10 * time.Second)

	}
}

func getClaymoreStatus() (bool, string) {

	strEcho := "{\"id\":0,\"jsonrpc\":\"2.0\",\"method\":\"miner_getstat1\"}"
	servAddr := ":3333"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		return false, err.Error()
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return false, err.Error()
	}

	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		return false, err.Error()
	}

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		return false, err.Error()
	}

	conn.Close()

	return true, string(reply)
}

func getEWBFStatus() (bool, string) {

	resp, err := http.Get("http://localhost:42000/getstat")
	if err != nil {
		return false, err.Error()
		//panic(err)
	}

	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		return true, str
	}

	return false, "-1"
}

func getCCMinerStatus() (bool, string) {
	miner := ccminer.New("localhost", 4068)
	summary, err := miner.Devs()

	if err != nil {
		return false, err.Error()
	}

	status, err := json.Marshal(&summary)
	return true, string(status)
}

func getCGMinerStatus() (bool, string) {

	miner := cgminer.New("localhost", 4068)
	summary, err := miner.Devs()

	if err != nil {
		return false, err.Error()
	}

	status, err := json.Marshal(&summary)
	return true, string(status)
}

func getSGMinerStatus() (bool, string) {
	miner := sgminer.New("localhost", 4028)
	summary, err := miner.Devs()
	if err != nil {
		return false, err.Error()
	}
	status, err := json.Marshal(&summary)
	return true, string(status)
}

func setWorkerStatus(workername string, miner string, currency string, raw string) {
	resp, err := http.PostForm("http://test.coindangpool.com/index.php?page=api&action=setworkerstatus",
		url.Values{"workername": {workername}, "miner": {miner}, "currency": {currency}, "raw": {raw}})
	if err != nil {
		print("setworker error")
		//panic(err)
	}

	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		print(str)
	}
}
