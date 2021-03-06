package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	rateString := os.Getenv("RATE")
	rate, err := strconv.ParseFloat(rateString, 32)
	if err != nil {
		panic("RATE could not be converted to a float")
	}

	ticker := time.NewTicker(time.Second / time.Duration(rate))

	gofakeit.Seed(time.Now().UnixNano())

	var ip, httpMethod, path, httpVersion, referrer, userAgent, upstream_ip, host, ssl_protocol, content_type, responseTime string
	var statusCode, bodyBytesSent, port int
	var timeLocal time.Time

	httpVersion = "HTTP/1.1"
	referrer = "-"
        ssl_protocol = "TLSv1.2/ECDHE-RSA-AES256-GCM-SHA384 62b923587b05d59b-BOM"
        content_type = "application/json; charset=utf-8"

	host_list := []string{
                       "api.ppops.com","appone.ppops.com", "apptwo.ppops.com",
                       "apptwo-new.ppops.com", "prod.ppops.pm5",
                      }

	port_list := []int{80, 443, 8443}
	path_list := []string{
			"/myapi/merchants", "/myapi/consumers", "/myapi/kyc",
                        "/check/balance", "/recharge/phone", "/recharge/dth",
			}
       upstream_addr_list := []string{
			"10.77.22.10", "10.77.22.11", "10.77.22.12",
			"10.77.22.13", "10.77.22.14", "10.77.22.15",
			"10.77.23.10", "10.77.23.11", "10.77.23.12",
                        "10.77.27.13", "10.77.27.14", "10.77.27.15",
                        }

	for range ticker.C {
		timeLocal = time.Now()

		ip = gofakeit.IPv4Address()
                upstream_ip = randomUpstreamIp(upstream_addr_list)
		httpMethod = weightedHTTPMethod(50, 20)
		path = randomPath(path_list,1,3)
                host = randomHost(host_list)
		statusCode = weightedStatusCode(80)
                responseTime = randomResponseTime(0.010, 20.000, 80)
		bodyBytesSent = realisticBytesSent(statusCode)
                port = randomPort(port_list)
		userAgent = gofakeit.UserAgent()

		fmt.Printf("%s %s \"%s %s %s\" %v %s %s:%v %v \"%s\" \"%s\" %s %s %s\n", ip, timeLocal.Format("02/Jan/2006:15:04:05 +0530"), httpMethod, path, httpVersion, statusCode, responseTime, upstream_ip, port, bodyBytesSent, referrer, userAgent, ssl_protocol, content_type, host)
	}
}

func realisticBytesSent(statusCode int) int {
	if statusCode != 200 {
		return gofakeit.Number(30, 120)
	}

	return gofakeit.Number(800, 3100)
}

func weightedStatusCode(percentageOk int) int {
	roll := gofakeit.Number(0, 100)
	if roll <= percentageOk {
		return 200
	}

	return gofakeit.HTTPStatusCode()
}

func weightedHTTPMethod(percentageGet, percentagePost int) string {
	if percentageGet+percentagePost >= 100 {
		panic("percentageGet and percentagePost add up to more than 100%")
	}

	roll := gofakeit.Number(0, 100)
	if roll <= percentageGet {
		return "GET"
	} else if roll <= percentagePost {
		return "POST"
	}

	return gofakeit.HTTPMethod()
}

func randomResponseTime(min, max float32, weight int) string {
        roll := gofakeit.Number(0, 100)
        if roll <= weight {
                max = 0.999
        }
        res := gofakeit.Float32Range(min, max)
        return fmt.Sprintf("%.3f", res)
}

func randomPath(path_list []string, min, max int) string {
	var path strings.Builder
	length := gofakeit.Number(min, max)
	min = 0
        max = len(path_list)-1
        randomIndex := gofakeit.Number(min, max)
	path.WriteString(path_list[randomIndex])
	path.WriteString("/")

	for i := 0; i < length; i++ {
		if i > 0 {
			path.WriteString(gofakeit.RandomString([]string{"-", "-", "_", "%20", "/", "/", "/"}))
		}
		path.WriteString(gofakeit.BuzzWord())
	}

	result := path.String()
	return strings.Replace(result, " ", "%20", -1)
}

func randomUpstreamIp(upstream_addr_list []string) string {
        min := 0
        max := len(upstream_addr_list)-1
        randomIndex := gofakeit.Number(min, max)
        return upstream_addr_list[randomIndex]
}

func randomHost(host_list []string) string {
        min := 0
        max := len(host_list)-1
        randomIndex := gofakeit.Number(min, max)
        return host_list[randomIndex]
}

func randomPort(port_list []int) int {
        min := 0
        max := len(port_list)-1
	randomIndex := gofakeit.Number(min, max)
	return port_list[randomIndex]
}
