package util

import (
	"bufio"
	"github.com/qiadapei/Proxyscan/log"
	"github.com/qiadapei/Proxyscan/model"
	"os"
	"strconv"
	"strings"
)

func ReadProxyAddr(filename string) (sliceProxyAddr []model.Input_ProxyAddr) {
	proxyFile, err := os.Open(filename)
	if err != nil {
		log.Log.Fatalf("Open proxy file err, %v", err)
	}

	defer proxyFile.Close()

	scanner := bufio.NewScanner(proxyFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ipPort := scanner.Text()
		t := strings.Split(ipPort, ":")
		ip := t[0]
		port, err := strconv.Atoi(t[1])
		if err == nil {
			proxyAddr := model.Input_ProxyAddr{
				IP:   ip,
				Port: port,
			}
			sliceProxyAddr = append(sliceProxyAddr, proxyAddr)
		}
	}
	return sliceProxyAddr
}
