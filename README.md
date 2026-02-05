# go-user-authentication-
This is a complete, step-by-step guide to building about user curd with authentication. I have structured this exactly according to your request, mapping your Python knowledge to Go concepts.


## define project folder with name
 $ mkdir <project_name>
 $ cd project_name

## initialize go mod and get all packages
 $ go mod init <project_name>
 $ go get -u github.com/gin-gonic/gin
 $ go get -u gorm.io/gorm
 $ go get -u gorm.io/driver/postgres
 $ go get -u github.com/golang-jwt/jwt/v5
 $ go get -u golang.org/x/crypto/bcrypt

 $ mkdir -p cmd/api internal/platform internal/user uploads
