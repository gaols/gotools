package httputils

import "testing"

func TestParseHeaders(t *testing.T) {
	headers := `
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9,en;q=0.8
Cache-Control: no-cache
Connection: keep-alive
Cookie: ASL=18752,0000q,24221446; HAList=a-sz-300059-%u4E1C%u65B9%u8D22%u5BCC; ADVC=398fac867fa4f7; ADVS=398fac867fa4f7; em_hq_fls=js; qgqp_b_id=260ab3186e8738a8fc3b2f555a3198d5; st_si=39954275459152; emshistory=%5B%22%E5%8D%8E%E5%A4%8F%E5%9B%9E%E6%8A%A5%E4%BA%8C%E5%8F%B7%22%5D; EMFUND1=null; EMFUND2=null; EMFUND3=null; EMFUND4=null; EMFUND5=null; EMFUND6=null; EMFUND7=null; EMFUND8=null; EMFUND0=null; st_asi=delete; st_pvi=93810733233013; st_sp=2021-05-05%2016%3A48%3A03; st_inirUrl=https%3A%2F%2Fwww.baidu.com%2Flink; st_sn=32; st_psi=20210505192318941-113200301201-6766037322; EMFUND9=05-05 19:29:48@#$%u534E%u590F%u56DE%u62A5%u4E8C%u53F7%u6DF7%u5408@%23%24002021
Host: same.eastmoney.com
Pragma: no-cache
Referer: http://fund.eastmoney.com/
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36
`
	headersMp := parseHeaders(headers)
	assertHasPair := assertHasPairFunc(headersMp, t)
	assertHasPair("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	assertHasPair("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	assertHasPair("Cache-Control", "no-cache")
	assertHasPair("Connection", "keep-alive")
	assertHasPair("Cookie", "ASL=18752,0000q,24221446; HAList=a-sz-300059-%u4E1C%u65B9%u8D22%u5BCC; ADVC=398fac867fa4f7; ADVS=398fac867fa4f7; em_hq_fls=js; qgqp_b_id=260ab3186e8738a8fc3b2f555a3198d5; st_si=39954275459152; emshistory=%5B%22%E5%8D%8E%E5%A4%8F%E5%9B%9E%E6%8A%A5%E4%BA%8C%E5%8F%B7%22%5D; EMFUND1=null; EMFUND2=null; EMFUND3=null; EMFUND4=null; EMFUND5=null; EMFUND6=null; EMFUND7=null; EMFUND8=null; EMFUND0=null; st_asi=delete; st_pvi=93810733233013; st_sp=2021-05-05%2016%3A48%3A03; st_inirUrl=https%3A%2F%2Fwww.baidu.com%2Flink; st_sn=32; st_psi=20210505192318941-113200301201-6766037322; EMFUND9=05-05 19:29:48@#$%u534E%u590F%u56DE%u62A5%u4E8C%u53F7%u6DF7%u5408@%23%24002021")
	assertHasPair("Host", "same.eastmoney.com")
	assertHasPair("Pragma", "no-cache")
	assertHasPair("Referer", "http://fund.eastmoney.com/")
	assertHasPair("Upgrade-Insecure-Requests", "1")
	assertHasPair("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
}

func assertHasPairFunc(headersMp map[string]string, t *testing.T) func(k string, v string) {
	return func(k string, v string) {
		if val, ok := headersMp[k]; !ok || val != v {
			t.FailNow()
		}
	}
}
