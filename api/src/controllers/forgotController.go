package controllers

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/renasami/svelte-go/api/database"
	"github.com/renasami/svelte-go/api/models"
)

func Forgot(c *fiber.Ctx) error {
	var data map[string]string

	// リクエストデータをパースする
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	token := RandStringRunes(12)
	passwordReset := models.PasswordReset{
		Email: data["email"],
		Token: token,
	}

	// DBに保存
	database.DB.Create(&passwordReset)

	database.DB.Create(&passwordReset)
 
	// SMTPメール送信
	from := "selfnote-owner@yahoo.co.jp"
	to := []string{
		data["email"],
	}
	sendFrom := fmt.Sprintf("From: %s\n", from)
	subject := fmt.Sprintf("Subject; %s\n", "Password Reset")
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	// Vue.jsのアドレス
	url := "http://localhost:3000/reset/" + token
	message := fmt.Sprintf("Click <a href=\"%s\">here</a> to reset password!", url)
	err := smtp.SendMail(
		"smtp:1025", // コンテナサービス名+port
		nil,
		from,
		to,
		[]byte(sendFrom+subject+mime+message),
	)
 
	if err != nil {
		return err
	}
 
	return c.JSON(fiber.Map{
		"message": "SUCCESS",
	})
}

// ランダム文字列を返す関数
func RandStringRunes(n int) string {
	var lettersRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = lettersRunes[rand.Intn(len(lettersRunes))]
	}
	return string(b)
}

func Reset(c *fiber.Ctx) error {
	var data map[string]string
 
	// リクエストデータをパースする
	if err := c.BodyParser(&data); err != nil {
		return err
	}
 
	// パスワードチェック
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match!",
		})
	}
 
	var passwordReset = models.PasswordReset{}
	// JWT Tokenからデータを取得
	err := database.DB.Where("token = ?", data["token"]).Last(&passwordReset)
	if err.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid token!",
		})
	}
 
	// パスワードをエンコード
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	database.DB.Model(&models.User{}).Where("email = ?", passwordReset.Email).Update("password", password)
 
	return c.JSON(fiber.Map{
		"message": "SUCCESS",
	})
}