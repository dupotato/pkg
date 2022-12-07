package httpclient

import (
	"aihub2/pkg/logger"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

var Giss2Client *http.Client

func init() {
	Giss2Client = &http.Client{}
}

func SendCallStatus(uuid string, callstatus string) {
	squery := fmt.Sprintf("http://%s/v1/notifycallstatus?uuid=%s&status=%s", viper.GetString("iss2.rpcserver"), uuid, callstatus)
	logger.Infof("[%s][iss2] %s", uuid, squery)
	Giss2Client.Get(squery)
}
