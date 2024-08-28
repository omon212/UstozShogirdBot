package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type User struct {
	State      string
	Name       string
	Age        string
	Technology string
	phone      int64
}

func main() {
	botToken := "7177836345:AAFp94Rm2aXx-0SsccTnx2R6Cz2clmbsuZ0"
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	users := make(map[int64]*User)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		chatID := update.Message.Chat.ID

		if _, ok := users[chatID]; !ok {
			users[chatID] = &User{} // Yangi foydalanuvchi uchun User tuzilmasi yaratish
		}

		switch update.Message.Text {
		case "/start":
			keyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Sherik kerak"),
					tgbotapi.NewKeyboardButton("Ish joyi kerak"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Hodim kerak"),
					tgbotapi.NewKeyboardButton("Ustoz kerak"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Shogird kerak"),
				),
			)
			msg := tgbotapi.NewMessage(chatID, "Assalom alaykum Омон\nUstozShogird kanalining rasmiy botiga xush kelibsiz!\n\n/help yordam buyrugi orqali nimalarga qodir ekanligimni bilib oling!")
			msg.ReplyMarkup = keyboard
			bot.Send(msg)
		case "Ish joyi kerak":
			msg := tgbotapi.NewMessage(chatID, "Ish joyi topish uchun ariza berish\n\nHozir sizga birnecha savollar beriladi. \nHar biriga javob bering. \nOxirida agar hammasi to`g`ri bo`lsa, HA tugmasini bosing va arizangiz Adminga yuboriladi.")
			msg2 := tgbotapi.NewMessage(chatID, "Ism, familiyangizni kiriting?")
			users[chatID].State = "name" // Foydalanuvchi holatini o'zgartirish
			bot.Send(msg)
			bot.Send(msg2)
		default:
			if users[chatID].State == "name" {
				users[chatID].Name = update.Message.Text
				msg := tgbotapi.NewMessage(chatID, "🕑 Yosh: \n\nYoshingizni kiriting?\nMasalan, 19")
				users[chatID].State = "age" // Holatni yangilash
				bot.Send(msg)
			} else if users[chatID].State == "age" {
				users[chatID].Age = update.Message.Text
				msg := tgbotapi.NewMessage(chatID, "📚 Texnologiya:\n\nTalab qilinadigan texnologiyalarni kiriting?\nTexnologiya nomlarini vergul bilan ajrating. Masalan, \n\nJava, C++, C#")
				users[chatID].State = "Technology" // Holatni yangilash
				bot.Send(msg)
			} else if users[chatID].State == "Technology" {
				users[chatID].Technology = update.Message.Text
				msg := tgbotapi.NewMessage(chatID, "📞 Aloqa: \n\nBog`lanish uchun raqamingizni kiriting?\nMasalan, +998996962112")
				users[chatID].State = "phone" // Holatni yangilash
				bot.Send(msg)
			}
		}
	}
}
