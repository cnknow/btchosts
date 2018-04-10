package hosts

import (
	"io/ioutil"
	"strings"
	"regexp"
	"os"
)

func DoIt() {
	var body string;
	pathHosts := `C:\Windows\System32\drivers\etc\hosts`
	ifFile, _ := PathExists(pathHosts)

	if (ifFile) {
		CopyFile(pathHosts+".bak", pathHosts)
		dat, _ := ioutil.ReadFile(pathHosts)
		body = string(dat)
	}
	if (!strings.Contains(body, "#BTC HOSTS")) {
		body = body + GetAddress()

	} else {
		pat := `#BTC HOSTS BEGIN[\S\s]+?#BTC HOSTS END`
		re, _ := regexp.Compile(pat)
		body = re.ReplaceAllString(body, GetAddress())
	}
	println(body)
	ioutil.WriteFile(pathHosts, ([]byte)(body), os.ModeAppend)
}
