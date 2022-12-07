package httpclient

import (
	"aihub2/pkg/logger"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type HttpClient struct {
	httpBase       string
	bInfo          string
	enableTls      int
	clientKey      string
	clientCrt      string
	caCrt          string
	certpath       string
	idleConn       int
	maxConn        int
	sendMss        int
	contentType    string
	boundary       string
	totalTimeout   int
	respTimeout    int
	gClient        *http.Client
	addressByEntid map[string]string
	httpSwitch     int
}

var GhttpClient *HttpClient

func Setup() bool {
	httpbase := viper.GetString("eventpush.address")
	idleconn := viper.GetInt("eventpush.idleconn")
	maxconn := viper.GetInt("eventpush.maxconn")
	contenttype := viper.GetString("eventpush.contenttype")
	totalTimeout := viper.GetInt("eventpush.totaltimeout")
	respTimeout := viper.GetInt("eventpush.resptimeout")
	enabletls := viper.GetInt("eventpush.enabletls")
	cacrt := viper.GetString("eventpush.cacrt")
	clientCrt := viper.GetString("eventpush.clientcrt")
	clientKey := viper.GetString("eventpush.clientkey")

	GhttpClient = &HttpClient{
		httpBase:     httpbase,
		idleConn:     idleconn,
		maxConn:      maxconn,
		contentType:  contenttype,
		totalTimeout: totalTimeout,
		respTimeout:  respTimeout,
		enableTls:    enabletls,
		clientCrt:    clientCrt,
		clientKey:    clientKey,
		caCrt:        cacrt,
		httpSwitch:   viper.GetInt("eventpush.switch"),
	}
	if GhttpClient.httpSwitch == 2 {
		fmt.Println(viper.GetStringMapString("eventpush.addressByEntid"))
		GhttpClient.addressByEntid = viper.GetStringMapString("eventpush.addressByEntid")
	}
	logger.Infof("[TH]Setup %v", GhttpClient)
	return GhttpClient.initHttpClient(idleconn, maxconn)
}

func (g *HttpClient) initHttpClient(idleConn, maxConn int) bool {
	// 参数配置
	var tlsconfig *tls.Config
	logger.Infof("[TH]enable tls =%d idleConn=%d maxConn=%d", g.enableTls, idleConn, maxConn)
	if g.enableTls == 0 {
		tlsconfig = &tls.Config{
			InsecureSkipVerify: true,
		}
		g.gClient = &http.Client{
			Timeout: time.Duration(g.totalTimeout) * time.Second,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
					DualStack: true,
				}).DialContext,
				TLSClientConfig:     tlsconfig,
				MaxIdleConnsPerHost: idleConn,
				MaxConnsPerHost:     maxConn,
				IdleConnTimeout:     90 * time.Second,
			},
		}
		return true
	} else {
		logger.Infof("[TH]clientcrt=%s clientkey=%s cacrt=%s", g.clientCrt, g.clientKey, g.caCrt)

		if g.clientCrt != "" && g.clientKey != "" {
			cert, err := tls.LoadX509KeyPair(g.clientCrt, g.clientKey)
			if err != nil {
				logger.Errorf("[TH]https cert path load error %v", err.Error())
				return false
			}
			certBytes, err := ioutil.ReadFile(g.caCrt)
			if err != nil {
				logger.Errorf("[TH] unable to read cert.pem %s", err.Error())
				return false
			}

			clientCertPool := x509.NewCertPool()
			ok := clientCertPool.AppendCertsFromPEM(certBytes)
			if !ok {
				logger.Errorf("[TH]fial to parse root certificat ")
				return false
			}

			tlsconfig = &tls.Config{
				RootCAs:            clientCertPool,
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: true,
			}
		} else {
			certBytes, err := ioutil.ReadFile(g.caCrt)
			if err != nil {
				logger.Errorf("[TH] unable to read cert.pem %s", err.Error())
				return false
			}

			clientCertPool := x509.NewCertPool()
			ok := clientCertPool.AppendCertsFromPEM(certBytes)
			if !ok {
				logger.Errorf("[TH]fial to parse root certificat ")
				return false
			}
			tlsconfig = &tls.Config{
				RootCAs:            clientCertPool,
				InsecureSkipVerify: true,
			}
		}

		g.gClient = &http.Client{
			Timeout: time.Duration(g.totalTimeout) * time.Second,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
					DualStack: true,
				}).DialContext,
				TLSClientConfig:     tlsconfig,
				MaxIdleConnsPerHost: idleConn,
				MaxConnsPerHost:     maxConn,
				IdleConnTimeout:     90 * time.Second,
			},
		}

		return true
	}
	return false
	//logger.Infof("init HttpClient info ildeConn=%d maxConn=%d", idleConn, maxConn)
}

func SendData(uuid string, senddata []byte, entid string) error {
	ctx, cancel := context.WithCancel(context.TODO())
	timer := time.AfterFunc(time.Duration(GhttpClient.respTimeout)*time.Second, func() {
		cancel()
	})

	defer timer.Stop()
	logger.Debugf("SendData data=%s\n", string(senddata))
	add := GhttpClient.GetSendAddress(entid)
	logger.Debugf("[%s][TH] pushaddr %s=%s", uuid, entid, add)
	req, err := http.NewRequest("POST", add, strings.NewReader(string(senddata)))

	if err != nil {
		logger.Fatalf("[%s][TH] http error %s", uuid, err.Error())
		return err
	}
	req = req.WithContext(ctx)

	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Content-Length", strconv.Itoa(len(senddata)))
	//req.Header.Add("Connection", "Keep-Alive")
	//logger.Debugf("[TEST][%s] request %s", t.ss.Id, string(senddata))
	tbegin := time.Now()
	if resp, err := GhttpClient.gClient.Do(req); err != nil {
		logger.Errorf("[TH][%s] Send error;%s", uuid, err.Error())
		return err
	} else {
		strbody, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		var mapBody map[string]interface{}
		json.Unmarshal(strbody, &mapBody)
		cost := time.Since(tbegin) / 1000000
		logger.Infof("[TH][%s] send success;cost=%d ;body=[%v]", uuid, cost, mapBody)
	}
	return nil
}

func (h *HttpClient) GetSendAddress(entid string) string {
	if h.httpSwitch == 1 {
		return h.httpBase
	} else {
		if v, ok := h.addressByEntid[entid]; ok {
			return v
		} else {
			return h.httpBase
		}
	}
}
