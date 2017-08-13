package gossm

type Settings struct {
	Monitor       MonitorSettings
	Notifications NotificationSettings
}

type MonitorSettings struct {
	CheckInterval  int `json:"checkInterval"`
	Timeout        int `json:"timeout"`
	MaxConnections int `json:"maxConnections"`
}

type NotificationSettings struct {
	Email []EmailSettings `json:"email"`
	Sms   []SmsSettings   `json:"sms"`
}

type EmailSettings struct {
	SMTP     string `json:"smtp"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SmsSettings struct {
	Sms string `json:"sms"`
}
