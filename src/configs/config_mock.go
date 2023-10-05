package configs

func InitConfigMock() {
	Cfg = &Env{
		DiscordWebhook: &DiscordWebhook{
			ID: "1159508574125428776",
			Token: "peWVu2CWq7ZnLyZ4KCab_MGsw9-EmX6X4nA43FHrffk_ITAXqGDOA0Cwct3d8agenfsO",
		},
		Jwt: &Jwt{
			Secret: "secretkub",
		},
	}
}