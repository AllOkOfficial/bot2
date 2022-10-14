// (c) Jisin0

package utils

import "github.com/PaulSonOfLars/gotgbot/v2"

var TEXT map[string]string = map[string]string{

	"START": `
<b>Hɪ %v, Nɪᴄᴇ ᴛᴏ ᴍᴇᴇᴛ ʏᴏᴜ !
I'ᴍ %v ᴀ ᴘʀᴇ - ғᴜɴᴄᴛɪᴏɴᴇᴅ Aᴡᴇsᴏᴍᴇ ғɪʟᴛᴇʀ ʙᴏᴛ ᴡɪᴛʜ ɢʟᴏʙᴀʟ ғɪʟᴛᴇʀ sᴜᴘᴘᴏʀᴛ.
I'ᴍ %v ᴀ ᴘʀᴇ - ғᴜɴᴄᴛɪᴏɴᴇᴅ Aᴡᴇsᴏᴍᴇ ғɪʟᴛᴇʀ ʙᴏᴛ ᴡɪᴛʜ ɢʟᴏʙᴀʟ ғɪʟᴛᴇʀ sᴜᴘᴘᴏʀᴛ.

<i>ɪᴛs ᴇᴀsʏ ᴛᴏ ᴜsᴇ ᴍᴇ; ᴊᴜsᴛ ᴀᴅᴅ ᴍᴇ ᴛᴏ ʏᴏᴜʀ ɢʀᴏᴜᴘ ᴀs ᴀᴅᴍɪɴ.<\i></b>
	`,
	"ABOUT": `
<b>Cɾҽα𝜏σɾ</b> : <a href='https://t.me/r2spr'>𝗥𝟮𝘀𝗽𝗿</a>
<b>Լαɳցᥙαցҽ</b> : <a href='https://go.dev'>𝗚𝗢</a>
<b>Ƒɾα𝓶ҽɯσƙ</b> : <a>𝗚𝗼𝘁𝗴𝗯𝗼𝘁</a>
<b>Sҽɾʋҽɾ</b> : <a>𝗖𝗹𝗼𝘂𝗱 𝗟𝗶𝗻𝗼𝗱𝗲</a>
<b>Ɗα𝜏αẞαടҽ</b> : <a href='mongodb.org'>𝗠𝗼𝗻𝗴𝗼𝗗𝗕</a>
<b>Sᥙρρσɾ𝜏</b> : <a href='https://t.me/HombaleCinemasChat'>𝗛𝗲𝗿𝗲</a>
<b>𝓤ρԃα𝜏ҽട</b> : <a href='https://t.me/HombaleCinemas'>𝗝𝗼𝗶𝗻</a>
	`,

	"MF": `
<b>Mᴀɴᴜᴀʟ ғɪʟᴛᴇʀs ᴀʟʟᴏᴡ ʏᴏᴜ ᴛᴏ sᴀᴠᴇ ᴄᴜsᴛᴏᴍ ғɪʟᴛᴇʀs ᴏᴛʜᴇʀ ᴛʜᴀɴ ᴛʜᴇ ᴀᴜᴛᴏᴍᴀᴛɪᴄ ᴏɴᴇs. Fɪʟᴛᴇʀs ᴄᴀɴ ʙᴇ ᴏғ ᴛᴇxᴛ/ᴘʜᴏᴛᴏ/ᴅᴏᴄᴜᴍᴇɴᴛ/ᴀᴜᴅɪᴏ/ᴀɴɪᴍᴀᴛɪᴏɴ/ᴠɪᴅᴇᴏ .</b>

<b><u>Nᴇᴡ ғɪʟᴛᴇʀ :</u></b>

<u>Fᴏʀᴍᴀᴛ</u>
<code>/filter "keyword" text</code> or
Rᴇᴘʟʏ ᴛᴏ ᴀ ᴍᴇssᴀɢᴇ -><code>/filter "keyword"</code>
<u>Usᴀɢᴇ</u>
<code>/filter "hi" hello</code>
[<code>hello</code>] -> Reply -> <code>/filter hi</code>

<b><u>Sᴛᴏᴘ ғɪʟᴛᴇʀ :</u></b>

<u>Fᴏʀᴍᴀᴛ</u>
<code>/stop "keyword"</code>
<u>Usᴀɢᴇ</u>
<code>/stop "hi"</code>

<b><u>Vɪᴇᴡ ғɪʟᴛᴇʀs :</u></b>
<code>/filters</code>
`,

	"GF": `
<b>Gʟᴏʙᴀʟ ғɪʟᴛᴇʀs ᴀʀᴇ ᴍᴀɴᴜᴀʟ ғɪʟᴛᴇʀs sᴀᴠᴇᴅ ʙʏ ʙᴏᴛ ᴀᴅᴍɪɴs ᴛʜᴀᴛ ᴡᴏʀᴋ ɪɴ ᴀʟʟ ᴄʜᴀᴛs. Tʜᴇʏ ᴘʀᴏᴠɪᴅᴇ ʟᴀᴛᴇsᴛ ᴍᴏᴠɪᴇs ɪɴ ᴀ ᴇᴀsʏ ᴛᴏ ᴜsᴇ ғᴏʀᴍᴀᴛ.</b>

<b><u>Sᴛᴏᴘ ғɪʟᴛᴇʀ :</u></b>

<u>Fᴏʀᴍᴀᴛ</u>
<code>/stop "keyword"</code>
<u>Usᴀɢᴇ</u>
<code>/stop "et"</code>

<b><u>Vɪᴇᴡ ғɪʟᴛᴇʀs :</u></b>
/gfilters
`,
	"CONNECT": `
<b>Cᴏɴɴᴇᴄᴛɪᴏɴs ᴀʟʟᴏᴡ ʏᴏᴜ ᴛᴏ ᴍᴀɴᴀɢᴇ ʏᴏᴜʀ ɢʀᴏᴜᴘ ʜᴇʀᴇ ɪɴ ᴘᴍ ɪɴsᴛᴇᴀᴅ ᴏғ sᴇɴᴅɪɴɢ ᴛʜᴏsᴇ ᴄᴏᴍᴍᴀɴᴅs ᴘᴜʙʟɪᴄʟʏ ɪɴ ᴛʜᴇ ɢʀᴏᴜᴘ ⠘⁾</b>

<b><u>Cᴏɴɴᴇᴄᴛ :</u></b>
-> Fɪʀsᴛ ɢᴇᴛ ʏᴏᴜʀ ɢʀᴏᴜᴘ's ɪᴅ ʙʏ sᴇɴᴅɪɴɢ /id ɪɴ ʏᴏᴜʀ ɢʀᴏᴜᴘ
-> <code>/connect [group_id]</code>

<b><u>Dɪsᴄᴏɴɴᴇᴄᴛ :</u></b>
<code>/disconnect</code>
`,

	"BROADCAST": `
<b>The broadcast feature allows bot admins to broadcast a message to all of the bot's users.</b>

<I>Broadcasts are of two types :</i>
 ~ Full Broadcast - Broadcast to all of the bot users <code>/broadcast</code>
 ~ Concast - Broadcast to only users who are connected to a chat <code>/concast</code>
`,
}

var BUTTONS map[string][][]gotgbot.InlineKeyboardButton = map[string][][]gotgbot.InlineKeyboardButton{
	"START": {
		{
			{Text: "☂ Aʙᴏᴜᴛ ☂", CallbackData: "edit(ABOUT)"},
			{Text: "🧭 Help 🧭", CallbackData: "edit(HELP)"},
			{Text: "🫂 Sᴜᴘᴘᴏʀᴛ 🫂", Url: "t.me/r2spr"},
		},
	},
	"ABOUT": {
		{
			{Text: "𝙷𝙾𝙼𝙴", CallbackData: "edit(START)"},
			{Text: "𝚂𝚃𝙰𝚃𝚂", CallbackData: "stats"},
		},
	},
	"STATS": {
		{
			{Text: "𝙱𝙰𝙲𝙺", CallbackData: "edit(ABOUT)"},
			{Text: "𝚁𝙴𝙵𝚁𝙴𝚂𝙷", CallbackData: "stats"},
		},
	},
	"HELP": {
		{{Text: "Fɪʟᴛᴇʀ", CallbackData: "edit(MF)"},
			{Text: "Gʟᴏʙᴀʟ", CallbackData: "edit(GF)"},
		}, {
			{Text: "Cᴏɴɴᴇᴄᴛ", CallbackData: "edit(CONNECT)"}, {Text: "Broadcast", CallbackData: "edit(BROADCAST)"},
		},
		{{Text: "Bᴀᴄᴋ ➔", CallbackData: "edit(START)"}},
	},
}
