package main

import (
	"log"
	"net/http"

	"ppid/database"
	"ppid/dtos"
	"ppid/helpers"
	"ppid/models"
	"ppid/repositories"
	"ppid/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// database configs
	dbUser, dbPassword := "xxx", "xxx"

	db, err := database.ConnectToDB(dbUser, dbPassword)

	// unable to connect to database
	if err != nil {
		log.Fatalln(err)
	}

	// ping to database
	err = db.DB().Ping()

	// error ping to database
	if err != nil {
		log.Fatalln(err)
	}

	// migration
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Admin{})
	db.AutoMigrate(&models.Pendaftaran{})
	db.AutoMigrate(&models.Dokumen{})

	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	adminRepository := repositories.NewAdminRepository(db)
	pendaftaranRepository := repositories.NewPendaftaranRepository(db)
	dokumenRepository := repositories.NewDokumenRepository(db)

	// route := configs.SetupRoutesUser(userRepository)
	// route = configs.SetupRoutesAdmin(adminRepository)
	// route = configs.SetupRoutesPendaftaran(pendaftaranRepository)

	router := gin.Default()

	// create route /create endpoint

	// USER ROUTE
	router.POST("/user/create", func(context *gin.Context) {
		// initialization contact model
		var user models.User

		// validate json
		err := context.ShouldBindJSON(&user)

		// validation errors
		if err != nil {
			// generate validation errors response
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		// default http status code = 200
		code := http.StatusOK

		// save contact & get it's response
		response := services.CreateUser(&user, *userRepository)

		// save contact failed
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.GET("/user", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindAllUsers(*userRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.GET("/user/show/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK

		response := services.FindOneUserById(id, *userRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.PUT("/user/update/:id", func(context *gin.Context) {
		id := context.Param("id")

		var user models.User

		err := context.ShouldBindJSON(&user)

		// validation errors
		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := services.UpdateUserById(id, &user, *userRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.DELETE("/user/delete/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK

		response := services.DeleteOneUserById(id, *userRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.POST("/user/delete", func(context *gin.Context) {
		var multiID dtos.MultiID

		err := context.ShouldBindJSON(&multiID)

		// validation errors
		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		if len(multiID.Ids) == 0 {
			response := dtos.Response{Success: false, Message: "IDs cannot be empty."}

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := services.DeleteUserByIds(&multiID, *userRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.GET("/user/pagination", func(context *gin.Context) {
		code := http.StatusOK

		pagination := helpers.GeneratePaginationRequest(context)

		response := services.Pagination(*userRepository, context, pagination)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	// ===================================================================
	// ROUTES ADMIN
	router.POST("/admin/create", func(context *gin.Context) {
		// initialization contact model
		var admin models.Admin

		// validate json
		err := context.ShouldBindJSON(&admin)

		// validation errors
		if err != nil {
			// generate validation errors response
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		// default http status code = 200
		code := http.StatusOK

		// save contact & get it's response
		response := services.CreateAdmin(&admin, *adminRepository)

		// save contact failed
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.GET("/admin", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindAllAdmins(*adminRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.GET("/admin/show/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK

		response := services.FindOneAdminById(id, *adminRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.PUT("/admin/update/:id", func(context *gin.Context) {
		id := context.Param("id")

		var admin models.Admin

		err := context.ShouldBindJSON(&admin)

		// validation errors
		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := services.UpdateAdminById(id, &admin, *adminRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.DELETE("/admin/delete/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK

		response := services.DeleteOneAdminById(id, *adminRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.POST("/admin/delete", func(context *gin.Context) {
		var multiID dtos.MultiID

		err := context.ShouldBindJSON(&multiID)

		// validation errors
		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		if len(multiID.Ids) == 0 {
			response := dtos.Response{Success: false, Message: "IDs cannot be empty."}

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := services.DeleteAdminByIds(&multiID, *adminRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.GET("/admin/pagination", func(context *gin.Context) {
		code := http.StatusOK

		pagination := helpers.GeneratePaginationRequestAdmin(context)

		response := services.PaginationAdmin(*adminRepository, context, pagination)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	// ============================================================
	// ROUTES PENDAFTARAN
	router.POST("/pendaftaran/create", func(context *gin.Context) {
		// initialization contact model
		var pendaftaran models.Pendaftaran

		// validate json
		err := context.ShouldBindJSON(&pendaftaran)

		// validation errors
		if err != nil {
			// generate validation errors response
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		// default http status code = 200
		code := http.StatusOK

		// save contact & get it's response
		response := services.CreatePendaftaran(&pendaftaran, *pendaftaranRepository)

		// save contact failed
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.GET("/pendaftaran", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindAllPendaftarans(*pendaftaranRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.GET("/pendaftaran/show/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK

		response := services.FindOnePendaftaranById(id, *pendaftaranRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.PUT("/pendaftaran/update/:id", func(context *gin.Context) {
		id := context.Param("id")

		var pendaftaran models.Pendaftaran

		err := context.ShouldBindJSON(&pendaftaran)

		// validation errors
		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := services.UpdatePendaftaranById(id, &pendaftaran, *pendaftaranRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.DELETE("/pendaftaran/delete/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK

		response := services.DeleteOnePendaftaranById(id, *pendaftaranRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.POST("/pendaftaran/delete", func(context *gin.Context) {
		var multiID dtos.MultiID

		err := context.ShouldBindJSON(&multiID)

		// validation errors
		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		if len(multiID.Ids) == 0 {
			response := dtos.Response{Success: false, Message: "IDs cannot be empty."}

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := services.DeletePendaftaranByIds(&multiID, *pendaftaranRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.GET("/pendaftaran/pagination", func(context *gin.Context) {
		code := http.StatusOK

		pagination := helpers.GeneratePaginationRequestPendaftaran(context)

		response := services.PaginationPendaftaran(*pendaftaranRepository, context, pagination)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	// ============================================================
	// ROUTES PENDAFTARAN
	router.POST("/dokumen/create", func(context *gin.Context) {
		// initialization contact model
		var dokumen models.Dokumen

		// validate json
		err := context.ShouldBindJSON(&dokumen)

		// validation errors
		if err != nil {
			// generate validation errors response
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		// default http status code = 200
		code := http.StatusOK

		// save contact & get it's response
		response := services.CreateDokumen(&dokumen, *dokumenRepository)

		// save contact failed
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.GET("/dokumen", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindAllDokumens(*dokumenRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.GET("/dokumen/show/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK

		response := services.FindOneDokumenById(id, *dokumenRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.PUT("/dokumen/update/:id", func(context *gin.Context) {
		id := context.Param("id")

		var dokumen models.Dokumen

		err := context.ShouldBindJSON(&dokumen)

		// validation errors
		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := services.UpdateDokumenById(id, &dokumen, *dokumenRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.DELETE("/dokumen/delete/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK

		response := services.DeleteOneDokumenById(id, *dokumenRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.POST("/dokumen/delete", func(context *gin.Context) {
		var multiID dtos.MultiID

		err := context.ShouldBindJSON(&multiID)

		// validation errors
		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		if len(multiID.Ids) == 0 {
			response := dtos.Response{Success: false, Message: "IDs cannot be empty."}

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := services.DeleteDokumenByIds(&multiID, *dokumenRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.GET("/dokumen/pagination", func(context *gin.Context) {
		code := http.StatusOK

		pagination := helpers.GeneratePaginationRequestDokumen(context)

		response := services.PaginationDokumen(*dokumenRepository, context, pagination)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	router.Run(":8000")
}
