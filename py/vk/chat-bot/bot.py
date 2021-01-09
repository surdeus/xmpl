#!/bin/env python3

import vk_api, vk, json
from vk_api.keyboard import VkKeyboard, VkKeyboardColor
from vk_api.utils import get_random_id
from vk_api.bot_longpoll import VkBotLongPoll, VkBotEventType
from vk_api.longpoll import VkLongPoll, VkEventType

config_path = "../config.json"

with open(config_path, "r") as read_file :
    config = json.load(read_file)
token = config["token"]
group_id = int(config["group_id"])
server = config["server"]
key = config["key"]
ts = config["ts"]


vk_session = vk_api.VkApi(token=token)
longpoll = VkBotLongPoll(vk_session, group_id)
vk = vk_session.get_api()

for event in longpoll.listen() :
    if event.type == VkBotEventType.MESSAGE_NEW and event.from_chat :
        print(event.chat_id)
        #vk.messages.send(key=key, server=server, ts=ts, random_id=get_random_id(),
        #    message='Yes, it works. Stop typing.', chat_id = event.chat_id)
