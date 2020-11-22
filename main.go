package main

import (
	"TWallet/auth"
	"TWallet/category"
	"TWallet/handler"
	"TWallet/helper"
	"TWallet/transactions"
	"TWallet/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/twallet_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := user.NewRepository(db)
	categoryRepository := category.NewRepository(db)
	transactionRepository := transactions.NewRepository(db)

	userService := user.NewService(userRepository) //parsing data input dari user(Struct input) ke Struct user
	categoryService := category.NewService(categoryRepository)
	transactionService := transactions.NewService(transactionRepository)
	authService := auth.NewServiceAuth()

	userHandler := handler.NewuserHandler(userService, authService) //Parsing Input dari user ke user Service
	categoryHandler := handler.NewCategoryHandler(categoryService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	api := router.Group("/api/v1/")

	// Keluarga User
	api.POST("register", userHandler.RegisterUser)
	api.POST("login", userHandler.Login)
	api.POST("check-email", userHandler.CheckEmailAvailablity)

	//Keluarga Category
	api.GET("category", authMiddleware(authService, userService), categoryHandler.GetCategory)
	api.POST("create-category", authMiddleware(authService, userService), categoryHandler.CreateCategory)
	api.GET("category/:id", authMiddleware(authService, userService), categoryHandler.GetCategoryDetail)
	api.PUT("category/:id", authMiddleware(authService, userService), categoryHandler.UpdateCategory)
	api.DELETE("category/:id", authMiddleware(authService, userService), categoryHandler.DeleteCategory)

	//Keluarga Transaksi
	api.GET("transaction", authMiddleware(authService, userService), transactionHandler.GetTransaction)
	api.POST("create-transaction", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.GET("transaction/:id", authMiddleware(authService, userService), transactionHandler.GetTransactionDetail)
	api.DELETE("transaction/:id", authMiddleware(authService, userService), transactionHandler.DeleteTransaction)
	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") { //Cek apakah di auth Header ada kata Bearer atau tidak
			response := helper.APIResponse("Unauthorized #TKN001", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) //Agar proses dihentikan/tidak eksekusi program yang dibungkus middleware
			return
		}

		var tokenString string
		arrayToken := strings.Split(authHeader, " ") //Memisahkan Bearer dengan Token dan kembalian Array(Bearer[0], Token[1])

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1] //Mengambil index ke 1 dari array token
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized #TKN002", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) //Agar proses dihentikan/tidak eksekusi program yang dibungkus middleware
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized #TKN003", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) //Agar proses dihentikan/tidak eksekusi program yang dibungkus middleware
			return
		}

		userID := int(claim["user_id"].(float64)) //Claim type datanya MapClaim, harus diubah ke int(sesuai parameter GetUserByID)

		user, err := userService.GetUserById(userID)

		if err != nil {
			response := helper.APIResponse("Unauthorized #TKN004", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) //Agar proses dihentikan/tidak eksekusi program yang dibungkus middleware
			return
		}

		c.Set("currentUser", user) //Menyimpan current user ke c.CurrentUser
	}
}
