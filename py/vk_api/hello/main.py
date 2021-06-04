import vk_api

s = vk_api.VkApi("login", "password")
s.auth()

vk = s.get_api()
print(vk.wall.post(message="Hello, World!"))

