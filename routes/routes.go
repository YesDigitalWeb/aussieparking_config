package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qor/qor/admin"
	"github.com/theplant/aussie/app/controllers"
	"github.com/theplant/aussie/config"
	"github.com/theplant/aussie/lib/middlewares"
	"github.com/theplant/aussie/resources"
	"github.com/tommy351/gin-sessions"
)

var sessionStore = sessions.NewCookieStore(config.SessionCookieKeyPairs)

func Mux() *http.ServeMux {
	engine := gin.Default()
	engine.Use(config.Cfg.Influx.MonitorGinRequest, middlewares.HandleError, sessions.Middleware("_session", sessionStore))

	engine.StaticFS("/js/", http.Dir("public/js"))
	engine.StaticFS("/css/", http.Dir("public/css"))
	engine.StaticFS("/img/", http.Dir("public/img"))
	engine.StaticFS("/fonts/", http.Dir("public/fonts"))
	engine.StaticFS("/public", http.Dir("public"))
	engine.StaticFile("AussieAirportParking-Ourinmbah-Directions.pdf", "public/AussieAirportParking-Ourinmbah-Directions.pdf")

	home := engine.Group("/", middlewares.HomeViewLayout)
	home.GET("/", controllers.Index)

	engine.POST("/order", controllers.Book)
	engine.POST("/order/confirm", controllers.Confirm)
	engine.POST("/order/pay_on_site", controllers.PayOnSite)
	engine.POST("/order/pay_with_paypal", controllers.PayWithPaypal)
	engine.POST("/order/change_booking", controllers.ChangeBooking)

	home.GET("/thank_you", controllers.ThankYou)

	home.GET("/car_dealing", controllers.CarDealing)
	home.GET("/location", controllers.Location)
	home.GET("/about_us", controllers.AboutUs)
	home.GET("/contact_us", controllers.ContactUs)
	home.GET("/terms-and-conditions", controllers.TermsAndConditions)
	engine.POST("/contact_us", controllers.SendContactUsMessage)

	engine.GET("/clean_range_products/:date", controllers.GetCleanRangeProducts)

	// Require Admin Auth
	adminAuth := engine.Group("/", middlewares.AdminAuth)
	adminAuth.GET("/order/stat/:date", controllers.GetBookingStats)
	adminAuth.GET("/calendar_monthly_bookings", controllers.MonthlyBookings)
	adminAuth.GET("/calendar/:type/:month", controllers.GetDailyReportsByMonth)

	engine.GET("/admin_login", resources.AdminLoginPage)
	engine.POST("/admin_login", resources.AdminLogin)
	engine.GET("/admin_logout", resources.AdminLogout)

	var mux = http.NewServeMux()
	mux.Handle("/", engine)

	qorRouter := resources.Admin.GetRouter()

	// Temporally redirect default admin landing page to Booking List
	qorRouter.Use(func(c *admin.Context, m *admin.Middleware) {
		if c.Request.URL.Path == "/admin" {
			http.Redirect(c.Writer, c.Request, "/admin/order", http.StatusTemporaryRedirect)
			return
		}
		m.Next(c)
	})

	qorRouter.Get("/calendar", controllers.Calendar)
	qorRouter.Post("/range", controllers.CreateRange)
	qorRouter.Post("/discount", controllers.CreateDiscount)
	resources.Admin.MountTo("/admin", mux)

	return mux
}
