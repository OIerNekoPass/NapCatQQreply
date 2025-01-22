# 适配NapCat的简易QQ bot消息回复golang库

使用NapCat的HTTP服务器功能进行消息的回复

默认往127.0.0.1:1145发送消息，token默认为bot_token

使用`Set_url(ip string, port string)`设置url和端口，例如`Set_url("127.0.0.1", "1145")`

使用`Set_token(new_token string)`设置token

目前的可用函数为：

1. Group_text(group_id string, text string) 
   
   简单的在群里说一句话。group_id为q群号；string为消息文本
2. Luck_reply(group_id string, user_id string, pic_route string) 
   
   在群里@一个人的同时发一张图片。user_id为@的人的qq号；pic_route为发送的图片文件的本机路径，如"/root/pic/123.png"

3. Group_record(group_id string, record_path string)
   
   在群里发一句语音。record_path为发送的语音文件的本机路径，如"/root/record/123.mp3"


持续完善中（取决于目前写的bot还需要什么新功能OvO）