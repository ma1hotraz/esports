package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/iloginow/esportsdifference/bootstrap"
	controllers "github.com/iloginow/esportsdifference/controller"
	"github.com/iloginow/esportsdifference/database"
	"github.com/iloginow/esportsdifference/logger"
	"github.com/iloginow/esportsdifference/notifier"
	"github.com/iloginow/esportsdifference/repo"
	"github.com/iloginow/esportsdifference/scheduler"
	"github.com/iloginow/esportsdifference/utils"
	"github.com/sirupsen/logrus"
)

const (
	DEFAULT_REDIS_HOST      string = "localhost:6379"
	DEFAULT_VIEWS_DIR       string = "./views"
	DEFAULT_LOGS_DIR        string = "./logs"
	DEFAULT_PUBLIC_DIR      string = "./public"
	DEFAULT_DISCORD_TOKEN   string = ""
	DEFAULT_DISCORD_CHANNEL string = ""
)

var (
	redisHost      = utils.GetEnvOrDefault("REDIS_HOST", DEFAULT_REDIS_HOST)
	viewsDir       = utils.GetEnvOrDefault("VIEWS_DIR", DEFAULT_VIEWS_DIR)
	publicDir      = utils.GetEnvOrDefault("PUBLIC_DIR", DEFAULT_PUBLIC_DIR)
	logsDir        = utils.GetEnvOrDefault("LOGS_DIR", DEFAULT_LOGS_DIR)
	discordToken   = utils.GetEnvOrDefault("DISCORD_TOKEN", DEFAULT_DISCORD_TOKEN)
	discordChannel = strings.Split(utils.GetEnvOrDefault("DISCORD_CHANNEL", DEFAULT_DISCORD_CHANNEL), ",")
)

var reactEndpoints = []string{
	"/",
	"/cod",
	"/cod/slips",
	"/cod/pairs",
	"/csgo",
	"/csgo/slips",
	"/csgo/pairs",
	"/lol",
	"/lol/pairs",
	"/lol/slips",
	"/val",
	"/val/slips",
	"/val/pairs",
	"/dota",
	"/dota/pairs",
	"/dota/slips",
	"/halo",
	"/halo/slips",
	"/halo/pairs",
	"/login",
	"/register",
	"admin/user-manage",
	"admin/invite-code",
	"admin/user-manage/add-user",
	"admin/invite-code/add-code",
	"admin/user-manage/edit-user",
	"admin/message",
	"/user/expired",
	"/user/forgot-password",
	"/user/change-password"}

func init() {
	logger.Init(logsDir)
	repo.Init(redisHost)
	notifier.Init(discordToken, discordChannel)
	scheduler.Init()
	database.Connect()
	bootstrap.RegisterInitAdminAccount()
	utils.InitEmailConfig()
}

func main() {

	engine := html.New(viewsDir, ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
	}))

	// environment := utils.GetEnvOrDefault("ENV", "prod")

	app.Static("/", publicDir)

	renderReactApp(app, reactEndpoints)

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/user/forgot-password", controllers.ForgotPassword)
	app.Post("/api/user/change-password", controllers.ChangeUserPw)

	app.Get("/api/allUser", controllers.GetAllUsers)
	app.Get("/api/allUser/:id", controllers.GetUserById)
	app.Delete("/api/user/:id", controllers.DeleteUser)
	app.Post("/api/user/extend-time", controllers.ApplyNewInviteCode)
	app.Post("/api/user/edit/:id", controllers.EditUser)

	app.Delete("/api/logout", controllers.Logout)
	app.Post("/api/createInviteCode", controllers.CreateInviteCode)
	app.Get("/api/getInviteCode", controllers.GetAllInviteCodes)
	app.Delete("/api/invite-code/:id", controllers.DeleteInviteCode)

	app.Get("/compare", func(c *fiber.Ctx) error {
		result, err := repo.GetCompareResult()
		if err != nil {
			logrus.Error(err)
			c.Status(fiber.StatusInternalServerError).JSON(result)
		}
		return c.Status(fiber.StatusOK).JSON(result)
	})

	app.Get("/slips", func(c *fiber.Ctx) error {
		result, err := repo.GetSlipsResult()
		if err != nil {
			logrus.Error(err)
		}
		return c.Status(fiber.StatusOK).JSON(result)
	})

	app.Get("/pairs", func(c *fiber.Ctx) error {
		result, err := repo.GetPairsResult()
		if err != nil {
			logrus.Error(err)
			c.Status(fiber.StatusInternalServerError).JSON(result)
		}
		return c.Status(fiber.StatusOK).JSON(result)
	})

	app.Get("/api/relevant/underdog", func(c *fiber.Ctx) error {
		result, err := repo.GetUnderdogfantazyRelevant()
		if err != nil {
			logrus.Error(err)
			c.Status(fiber.StatusInternalServerError).JSON(result)
		}
		return c.Status(fiber.StatusOK).JSON(result)
	})

	app.Get("/api/relevant/prize", func(c *fiber.Ctx) error {
		result, err := repo.GetPrizepicksRelevant()
		if err != nil {
			logrus.Error(err)
			c.Status(fiber.StatusInternalServerError).JSON(result)
		}
		return c.Status(fiber.StatusOK).JSON(result)
	})

	app.Get("/api/relevant/sleeper", func(c *fiber.Ctx) error {
		result, err := repo.GetSleeperRelevant()
		if err != nil {
			logrus.Error(err)
			c.Status(fiber.StatusInternalServerError).JSON(result)
		}
		return c.Status(fiber.StatusOK).JSON(result)
	})

	app.Get("/inform-history", func(c *fiber.Ctx) error {
		result, err := repo.GetAllInformedLines()
		if err != nil {
			logrus.Error(err)
			c.Status(fiber.StatusInternalServerError).JSON(result)
		}
		return c.Status(fiber.StatusOK).JSON(result)
	})

	app.Post("/insert-used", controllers.InsertNewUsed)
	app.Post("/remove-stats", controllers.RemoveAllUsed)
	app.Post("/stats", controllers.GetAllUsedRows)
	app.Post("/remove-stat", controllers.RemoveOneUsed)

	// notification
	app.Post("/api/notification", controllers.CreateNewNotification)
	app.Get("/api/notification", controllers.GetNotifications)
	app.Get("/api/my-notification/:userId", controllers.GetMyNotifications)
	app.Post("/api/notification/dismiss", controllers.DismissNotification)
	app.Delete("/api/notification/:id", controllers.DeleteNotification)

	logrus.Error(app.Listen(":3000"))
	defer logger.CloseLogFile()

}

func renderReactApp(app *fiber.App, endponts []string) {
	for _, e := range endponts {
		app.Get(e, func(c *fiber.Ctx) error {
			return c.Render("index", nil)
		})
	}
}
