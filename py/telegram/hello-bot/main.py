#!/bin/env python3
import telebot

bot = telebot.TeleBot("your token")

@bot.message_handler(content_types=['text'])
def get_text_messages(message):
	bot.send_message(message.from_user.id, "Hello, World!")

bot.polling(none_stop=True, interval=0)

