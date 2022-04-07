package main

import (
	"my_blog/admin/internal/di"
)

// @title myBlog
// @version 1.0
// @description --
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /my-blog
func main() {
	//_ = ScryptPw("123456")

	_, err := di.Init()
	if err != nil {
		panic(err)
	}
}
