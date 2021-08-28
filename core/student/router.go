package student

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nus-utils/nus-peer-review/models"
	"github.com/nus-utils/nus-peer-review/utils"
	"gorm.io/gorm"
)

type StudentRoute struct {
	DB *gorm.DB
}

func (ur StudentRoute) CreateRouters(route *mux.Router) {
	ur.CreateAuthRouter(route.PathPrefix("/auth").Subrouter())
}

func (ur StudentRoute) CreateAuthRouter(route *mux.Router) {
	signUpRoute := route.NewRoute().Subrouter()
	signUpRoute.Use(utils.DecodeBodyMiddleware(&models.Student{}, "user"))
	signUpRoute.Use(utils.SanitizeDataMiddleware("user"))
	signUpRoute.Use(ur.PasswordHash)
	signUpRoute.HandleFunc("/signup", utils.DBCreateHandleFunc(ur.DB, &models.Student{}, "user", false)).Methods(http.MethodPost)

	loginRoute := route.NewRoute().Subrouter()
	loginRoute.Use(utils.DecodeParamsMiddleware(&models.Student{}, "user"))
	loginRoute.Use(ur.EmailCheck)
	loginRoute.Use(ur.PasswordCheck)
	loginRoute.HandleFunc("/login", utils.LoginHandleFunc(ur.DB, "Student", "user")).Methods(http.MethodGet)
}
