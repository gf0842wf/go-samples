// TODO: 这个不应该在main里
package main

// TODO: 这里定义Session
/* Session应该包含的字段
IP net.IP
MQ chan IPCObj // Player's Internal Message Queue  (IPCObj包括发送/接受放ID, 消息(json string), 时间等)
Encoder
Decoder
User // User包含玩家基本信息(区别Player, Player包含User和游戏内部具体信息)
LoggedIn bool
KickOut bool
各种时间信息
*/
type Session struct{}
