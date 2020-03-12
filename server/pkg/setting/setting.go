package setting

import (
	"log"
	"strconv"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	URL          string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	HTTPS_CRT    string
	HTTPS_KEY    string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts string

	RelativePath string
	Root         string

	PageSize        int
	ReCaptchaSecret string

	EmailUser string
	EmailPass string
	EmailHost string
	EmailPort int
)

func init() {
	var err error
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadStatic()
	LoadEmail()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	URL = sec.Key("URL").MustString("127.0.0.1")
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	HTTPS_CRT = sec.Key("HTTPS_CRT").MustString("crt.pem")
	HTTPS_KEY = sec.Key("HTTPS_KEY").MustString("key.pem")
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	ImagePrefixUrl = sec.Key("ImagePrefixUrl").MustString("http://127.0.0.1:" + strconv.Itoa(HTTPPort))
	ImageSavePath = sec.Key("ImageSavePath").MustString("assets/teams/")
	ImageMaxSize = sec.Key("ImageMaxSize").MustInt(2) * 1024 * 1024
	ImageAllowExts = sec.Key("ImageAllowExts").MustString(".jpg,.jpeg,.png")

	ReCaptchaSecret = sec.Key("RECAPTCHA_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func LoadStatic() {
	sec, err := Cfg.GetSection("static")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	RelativePath = sec.Key("RelativePath").MustString("/assets")
	Root = sec.Key("Root").MustString("./assets")
}

func LoadEmail() {
	sec, err := Cfg.GetSection("email")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	EmailUser = sec.Key("USER").MustString("z.jiang@pokeshare.monster")
	EmailPass = sec.Key("PASS").MustString("your_pass")
	EmailHost = sec.Key("HOST").MustString("smtp.mail.ru")
	EmailPort = sec.Key("PORT").MustInt(465)
}
