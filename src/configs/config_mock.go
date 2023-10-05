package configs

func InitConfigMock() {
	Cfg = &Env{
		DiscordWebhook: &DiscordWebhook{
			ID: "1143500994299310092",
			Token: "QjQr2LewR9tqc-BFxjgPFFmyfPfhDe-HHRdvSDluu1rYzGaoTIwuJ70dJDDLko4Aycca",
		},
		Jwt: &Jwt{
			Secret: "secretkub",
		},
	}
}