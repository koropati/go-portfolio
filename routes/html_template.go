package routes

func ConfigHTMLTemplates(cfg *SetupConfig) {
	cfg.Gin.Static("assets", "./templates/assets")
	cfg.Gin.LoadHTMLGlob("./templates/html/*.html")
	cfg.Gin.LoadHTMLGlob("./templates/html/register/*.html")
}
