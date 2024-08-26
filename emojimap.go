package logger

var emojiMap = map[string]string{
	"warningTriangle": "⚠️",
	"loopArrow":       "🔄",
	"notePen":         "📝",
	"fileFolder":      "📁",
	"file":            "📄",
	"success":         "✅",
	"fail":            "❌",
	"info":            "ℹ️",
	"question":        "❓",
	"exclamation":     "❗",
	"fire":            "🔥",
	"server":          "🖥️",
	"rocket":          "🚀",
	"package":         "📦",
	"tools":           "🛠️",
	"hammer":          "🔨",
	"gear":            "⚙️",
	"magnifyingGlass": "🔍",
	"lock":            "🔒",
	"key":             "🔑",
	"link":            "🔗",
	"trash":           "🗑️",
	"folder":          "📂",
	"upload":          "📤",
	"git":             "🐙",
	"laptop":          "💻",
	"mobile":          "📱",
}

func P(key string) {
	_ = emojiMap
}
