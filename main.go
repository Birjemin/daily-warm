package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/barryyan/daily-warm/engine"
	"github.com/barryyan/daily-warm/parser"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/barryyan/daily-warm/gomail"

	env "github.com/joho/godotenv"
	cron "github.com/robfig/cron/v3"
)

// User for receive email
type User struct {
	Email string `json:"email"`
	Local string `json:"local"`
}

func isDev() bool {
	return os.Getenv("MAIL_MODE") == "dev"
}

func main() {
	loadConfig()
	if isDev() {
		batchSendMail()
		return
	}
	nyc, _ := time.LoadLocation("Asia/Shanghai")
	cJob := cron.New(cron.WithLocation(nyc))

	cronCfg := os.Getenv("MAIL_CRON")
	if cronCfg == "" {
		batchSendMail()
	} else {
		cJob.AddFunc(cronCfg, func() {
			batchSendMail()
		})
		cJob.Start()
		select {}
	}
}

func loadConfig() {
	err := env.Load()
	if err != nil {
		log.Fatalf("Load .env file error: %s", err)
	}
}

func batchSendMail() {
	loadConfig()

	users := getUsers("MAIL_TO")
	if len(users) == 0 {
		return
	}

	// 批量获取信息
	commonUrls := []parser.IParser{
		parser.NewOne(),
		parser.NewEnglish(),
		parser.NewPoem(),
		parser.NewWallpaper(),
		parser.NewTrivia(),
	}
	data := engine.Run(commonUrls)

	var userUrls []parser.IParser
	for _, user := range users {
		userUrls = append(userUrls, parser.NewWeather(user.Local))
	}
	// 批量获取用户天气
	userWeather := engine.Run(userUrls)

	res := make(chan int)
	defer close(res)

	for _, user := range users {

		data["weather"] = userWeather["weather"+user.Local]
		html := generateHTML(data)

		if isDev() {
			fmt.Println(html)
			return
		}

		go func(email string) {
			sendMail(html, email)
			<-res
		}(user.Email)
		res <- 1
	}
}

func generateHTML(data map[string]interface{}) string {
	var body bytes.Buffer
	t, _ := template.ParseFiles("daily.html")
	if err := t.Execute(&body, data); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
	return string(body.Bytes())
}

func getUsers(envUser string) []User {
	var users []User
	userJSON := os.Getenv(envUser)
	err := json.Unmarshal([]byte(userJSON), &users)
	if err != nil {
		log.Fatalf("Parse users from %s error: %s", userJSON, err)
	}
	return users
}

func sendMail(content string, to string) {
	gomail.Config.Username = os.Getenv("MAIL_USERNAME")
	gomail.Config.Password = os.Getenv("MAIL_PASSWORD")
	gomail.Config.Host = os.Getenv("MAIL_HOST")
	gomail.Config.Port = os.Getenv("MAIL_PORT")
	gomail.Config.From = os.Getenv("MAIL_FROM")

	email := gomail.GoMail{
		To:      []string{to},
		Subject: os.Getenv("MAIL_SUBJECT"),
		Content: content,
	}

	err := email.Send()
	if err != nil {
		log.Printf("Send email fail, error: %s", err)
	} else {
		log.Printf("Send email %s success!", to)
	}
}
