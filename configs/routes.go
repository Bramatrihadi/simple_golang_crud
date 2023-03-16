// package configs

// import (
// 	"net/http"

// 	"ppid/dtos"
// 	"ppid/helpers"
// 	"ppid/models"
// 	"ppid/repositories"
// 	"ppid/services"

// 	"github.com/gin-gonic/gin"
// )

// func SetupRoutesUser(userRepository *repositories.UserRepository) *gin.Engine {
// 	route := gin.Default()

// 	// create route /create endpoint
// 	route.POST("/user/create", func(context *gin.Context) {
// 		// initialization contact model
// 		var user models.User

// 		// validate json
// 		err := context.ShouldBindJSON(&user)

// 		// validation errors
// 		if err != nil {
// 			// generate validation errors response
// 			response := helpers.GenerateValidationResponse(err)

// 			context.JSON(http.StatusBadRequest, response)

// 			return
// 		}

// 		// default http status code = 200
// 		code := http.StatusOK

// 		// save contact & get it's response
// 		response := services.CreateUser(&user, *userRepository)

// 		// save contact failed
// 		if !response.Success {
// 			// change http status code to 400
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.GET("/user", func(context *gin.Context) {
// 		code := http.StatusOK

// 		response := services.FindAllUsers(*userRepository)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.GET("/user/show/:id", func(context *gin.Context) {
// 		id := context.Param("id")

// 		code := http.StatusOK

// 		response := services.FindOneUserById(id, *userRepository)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.PUT("/user/update/:id", func(context *gin.Context) {
// 		id := context.Param("id")

// 		var user models.User

// 		err := context.ShouldBindJSON(&user)

// 		// validation errors
// 		if err != nil {
// 			response := helpers.GenerateValidationResponse(err)

// 			context.JSON(http.StatusBadRequest, response)

// 			return
// 		}

// 		code := http.StatusOK

// 		response := services.UpdateUserById(id, &user, *userRepository)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.DELETE("/user/delete/:id", func(context *gin.Context) {
// 		id := context.Param("id")

// 		code := http.StatusOK

// 		response := services.DeleteOneUserById(id, *userRepository)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.POST("/user/delete", func(context *gin.Context) {
// 		var multiID dtos.MultiID

// 		err := context.ShouldBindJSON(&multiID)

// 		// validation errors
// 		if err != nil {
// 			response := helpers.GenerateValidationResponse(err)

// 			context.JSON(http.StatusBadRequest, response)

// 			return
// 		}

// 		if len(multiID.Ids) == 0 {
// 			response := dtos.Response{Success: false, Message: "IDs cannot be empty."}

// 			context.JSON(http.StatusBadRequest, response)

// 			return
// 		}

// 		code := http.StatusOK

// 		response := services.DeleteUserByIds(&multiID, *userRepository)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.GET("/user/pagination", func(context *gin.Context) {
// 		code := http.StatusOK

// 		pagination := helpers.GeneratePaginationRequest(context)

// 		response := services.Pagination(*userRepository, context, pagination)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	return route
// }

// func SetupRoutesAdmin(adminRepository *repositories.AdminRepository) *gin.Engine {
// 	route := gin.Default()

// 	// create route /create endpoint
// 	route.POST("/admin/create", func(context *gin.Context) {
// 		// initialization contact model
// 		var admin models.Admin

// 		// validate json
// 		err := context.ShouldBindJSON(&admin)

// 		// validation errors
// 		if err != nil {
// 			// generate validation errors response
// 			response := helpers.GenerateValidationResponse(err)

// 			context.JSON(http.StatusBadRequest, response)

// 			return
// 		}

// 		// default http status code = 200
// 		code := http.StatusOK

// 		// save contact & get it's response
// 		response := services.CreateAdmin(&admin, *adminRepository)

// 		// save contact failed
// 		if !response.Success {
// 			// change http status code to 400
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.GET("/admin", func(context *gin.Context) {
// 		code := http.StatusOK

// 		response := services.FindAllAdmins(*adminRepository)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.GET("/admin/show/:id", func(context *gin.Context) {
// 		id := context.Param("id")

// 		code := http.StatusOK

// 		response := services.FindOneAdminById(id, *adminRepository)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.PUT("/admin/update/:id", func(context *gin.Context) {
// 		id := context.Param("id")

// 		var admin models.Admin

// 		err := context.ShouldBindJSON(&admin)

// 		// validation errors
// 		if err != nil {
// 			response := helpers.GenerateValidationResponse(err)

// 			context.JSON(http.StatusBadRequest, response)

// 			return
// 		}

// 		code := http.StatusOK

// 		response := services.UpdateAdminById(id, &admin, *adminRepository)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.DELETE("/admin/delete/:id", func(context *gin.Context) {
// 		id := context.Param("id")

// 		code := http.StatusOK

// 		response := services.DeleteOneAdminById(id, *adminRepository)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.POST("/admin/delete", func(context *gin.Context) {
// 		var multiID dtos.MultiID

// 		err := context.ShouldBindJSON(&multiID)

// 		// validation errors
// 		if err != nil {
// 			response := helpers.GenerateValidationResponse(err)

// 			context.JSON(http.StatusBadRequest, response)

// 			return
// 		}

// 		if len(multiID.Ids) == 0 {
// 			response := dtos.Response{Success: false, Message: "IDs cannot be empty."}

// 			context.JSON(http.StatusBadRequest, response)

// 			return
// 		}

// 		code := http.StatusOK

// 		response := services.DeleteAdminByIds(&multiID, *adminRepository)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.GET("/user/pagination", func(context *gin.Context) {
// 		code := http.StatusOK

// 		pagination := helpers.GeneratePaginationRequest(context)

// 		response := services.Pagination(*adminRepository, context, pagination)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	return route
// }

// func SetupRoutesPendaftaran(pendaftaranRepository *repositories.PendaftaranRepository) *gin.Engine {
// 	route := gin.Default()

// 	// create route /create endpoint
// 	route.POST("/pendaftaran/create", func(context *gin.Context) {
// 		// initialization contact model
// 		var pendaftaran models.Pendaftaran

// 		// validate json
// 		err := context.ShouldBindJSON(&pendaftaran)

// 		// validation errors
// 		if err != nil {
// 			// generate validation errors response
// 			response := helpers.GenerateValidationResponse(err)

// 			context.JSON(http.StatusBadRequest, response)

// 			return
// 		}

// 		// default http status code = 200
// 		code := http.StatusOK

// 		// save contact & get it's response
// 		response := services.CreatePendaftaran(&pendaftaran, *pendaftaranRepository)

// 		// save contact failed
// 		if !response.Success {
// 			// change http status code to 400
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.GET("/pendaftaran", func(context *gin.Context) {
// 		code := http.StatusOK

// 		response := services.FindAllPendaftarans(*pendaftaranRepository)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.GET("/pendaftaran/show/:id", func(context *gin.Context) {
// 		id := context.Param("id")

// 		code := http.StatusOK

// 		response := services.FindOnePendaftaranById(id, *pendaftaranRepository)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.PUT("/pendaftaran/update/:id", func(context *gin.Context) {
// 		id := context.Param("id")

// 		var pendaftaran models.Pendaftaran

// 		err := context.ShouldBindJSON(&pendaftaran)

// 		// validation errors
// 		if err != nil {
// 			response := helpers.GenerateValidationResponse(err)

// 			context.JSON(http.StatusBadRequest, response)

// 			return
// 		}

// 		code := http.StatusOK

// 		response := services.UpdatePendaftaranById(id, &pendaftaran, *pendaftaranRepository)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.DELETE("/pendaftaran/delete/:id", func(context *gin.Context) {
// 		id := context.Param("id")

// 		code := http.StatusOK

// 		response := services.DeleteOnePendaftaranById(id, *pendaftaranRepository)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	route.POST("/pendaftaran/delete", func(context *gin.Context) {
// 		var multiID dtos.MultiID

// 		err := context.ShouldBindJSON(&multiID)

// 		// validation errors
// 		if err != nil {
// 			response := helpers.GenerateValidationResponse(err)

// 			context.JSON(http.StatusBadRequest, response)

// 			return
// 		}

// 		if len(multiID.Ids) == 0 {
// 			response := dtos.Response{Success: false, Message: "IDs cannot be empty."}

// 			context.JSON(http.StatusBadRequest, response)

// 			return
// 		}

// 		code := http.StatusOK

// 		response := services.DeletePendaftaranByIds(&multiID, *pendaftaranRepository)

// 		if !response.Success {
// 			code = http.StatusBadRequest
// 		}

// 		context.JSON(code, response)
// 	})

// 	// route.GET("/user/pagination", func(context *gin.Context) {
// 	// 	code := http.StatusOK

// 	// 	pagination := helpers.GeneratePaginationRequest(context)

// 	// 	response := services.Pagination(*adminRepository, context, pagination)

// 	// 	if !response.Success {
// 	// 		code = http.StatusBadRequest
// 	// 	}

// 	// 	context.JSON(code, response)
// 	// })

// 	return route
// }
