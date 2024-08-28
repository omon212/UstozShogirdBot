package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type User struct {
	State       string
	Name        string
	Age         string
	Technology  string
	Phone       string
	Narxi       string
	Hudud       string
	Kasbi       string
	MurojatVaqt string
	Maqsad      string
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
			users[chatID] = &User{}
		}

		switch update.Message.Text {
		case "/start":
			// Initial greeting and keyboard setup
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
			msg2 := tgbotapi.NewMessage(chatID, "👨‍💼 Ism, familiyangizni kiriting?")
			users[chatID].State = "name"
			bot.Send(msg)
			bot.Send(msg2)
		default:
			handleUserInput(update, users, bot)
		}
	}
}

func handleUserInput(update tgbotapi.Update, users map[int64]*User, bot *tgbotapi.BotAPI) {
	chatID := update.Message.Chat.ID
	user := users[chatID]
	text := update.Message.Text

	switch user.State {
	case "name":
		user.Name = text
		user.State = "age"
		msg := tgbotapi.NewMessage(chatID, "🕑 Yosh: \n\nYoshingizni kiriting?\nMasalan, 19")
		bot.Send(msg)
	case "age":
		user.Age = text
		user.State = "Technology"
		msg := tgbotapi.NewMessage(chatID, "📚 Texnologiya:\n\nTalab qilinadigan texnologiyalarni kiriting?\nTexnologiya nomlarini vergul bilan ajrating. Masalan, \n\nJava, C++, C#")
		bot.Send(msg)
	case "Technology":
		user.Technology = text
		user.State = "Phone"
		msg := tgbotapi.NewMessage(chatID, "📞 Aloqa: \n\nBog`lanish uchun raqamingizni kiriting?\nMasalan, +998 90 123 45 67")
		bot.Send(msg)
	case "Phone":
		user.Phone = text
		user.State = "Hudud"
		msg := tgbotapi.NewMessage(chatID, "🌐 Hudud: \n\nQaysi hududdansiz?\nViloyat nomi, Toshkent shahar yoki Respublikani kiriting.")
		bot.Send(msg)
	case "Hudud":
		user.Hudud = text
		user.State = "Narxi"
		msg := tgbotapi.NewMessage(chatID, "💰 Narxi:\n\nTolov qilasizmi yoki Tekinmi?\nKerak bo`lsa, Summani kiriting?")
		bot.Send(msg)
	case "Narxi":
		user.Narxi = text
		user.State = "Kasbi"
		msg := tgbotapi.NewMessage(chatID, "👨🏻‍💻 Kasbi: \n\nIshlaysizmi yoki o`qiysizmi?\nMasalan, Talaba")
		bot.Send(msg)
	case "Kasbi":
		user.Kasbi = text
		user.State = "Murojat_Vaqt"
		msg := tgbotapi.NewMessage(chatID, "🕰 Murojaat qilish vaqti: \n\nQaysi vaqtda murojaat qilish mumkin?\nMasalan, 9:00 - 18:00")
		bot.Send(msg)
	case "Murojat_Vaqt":
		user.MurojatVaqt = text
		user.State = "MaqSad"
		msg := tgbotapi.NewMessage(chatID, "🔎 Maqsad: \n\nMaqsadingizni qisqacha yozib bering")
		bot.Send(msg)
	case "MaqSad":
		user.Maqsad = text
		user.State = ""
		summary := fmt.Sprintf("Ish joyi kerak:\n\n👨‍💼 Xodim: %s\n🕑 Yosh: %s\n📚 Texnologiya: %s\n🇺🇿  Aloqa: %s\n🌐 Hudud: %s\n💰 Narxi: %s\n👨🏻‍💻 Kasbi: %s\n🕰 Murojaat qilish vaqti: %s\n🔎 Maqsad: %s",
			user.Name, user.Age, user.Technology, user.Phone, user.Hudud, user.Narxi, user.Kasbi, user.MurojatVaqt, update.Message.Text)
		msg := tgbotapi.NewMessage(chatID, summary)
		bot.Send(msg)
	}
}
