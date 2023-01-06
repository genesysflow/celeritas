package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func doAuth() error {
	// migrations
	dbType := cel.DB.DataType
	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro())
	upFile := cel.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
	downFile := cel.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

	err := copyFileFromTemplate("templates/migrations/auth_tables."+dbType+".up.sql", upFile)
	if err != nil {
		return err
	}

	err = copyDataToFile([]byte("DROP TABLE IF EXISTS users CASCADE; DROP TABLE IF EXISTS tokens CASCADE; DROP TABLE IF EXISTS remember_tokens CASCADE;"), downFile)
	if err != nil {
		return err
	}

	// run the migrations
	err = doMigrate("up", "")
	if err != nil {
		return err
	}

	//  copy files over

	err = copyFileFromTemplate("templates/data/user.go.txt", cel.RootPath+"/data/user.go")
	if err != nil {
		return err
	}

	err = copyFileFromTemplate("templates/data/token.go.txt", cel.RootPath+"/data/token.go")
	if err != nil {
		return err
	}

	err = copyFileFromTemplate("templates/data/remember_token.go.txt", cel.RootPath+"/data/remember_token.go")
	if err != nil {
		return err
	}

	err = copyFileFromTemplate("templates/middleware/auth.go.txt", cel.RootPath+"/middleware/auth.go")
	if err != nil {
		return err
	}

	err = copyFileFromTemplate("templates/middleware/auth-token.go.txt", cel.RootPath+"/middleware/auth-token.go")
	if err != nil {
		return err
	}

	err = copyFileFromTemplate("templates/middleware/remember.go.txt", cel.RootPath+"/middleware/remember.go")
	if err != nil {
		return err
	}

	err = copyFileFromTemplate("templates/handlers/auth-handlers.go.txt", cel.RootPath+"/handlers/auth-handlers.go")
	if err != nil {
		return err
	}

	err = copyFileFromTemplate("templates/mailer/password-reset.html.tmpl", cel.RootPath+"/mail/password-reset.html.tmpl")
	if err != nil {
		return err
	}

	err = copyFileFromTemplate("templates/mailer/password-reset.plain.tmpl", cel.RootPath+"/mail/password-reset.plain.tmpl")
	if err != nil {
		return err
	}

	err = copyFileFromTemplate("templates/views/login.jet", cel.RootPath+"/views/login.jet")
	if err != nil {
		return err
	}

	err = copyFileFromTemplate("templates/views/reset-password.jet", cel.RootPath+"/views/reset-password.jet")
	if err != nil {
		return err
	}

	err = copyFileFromTemplate("templates/views/forgot.jet", cel.RootPath+"/views/forgot.jet")
	if err != nil {
		return err
	}

	color.Yellow("   - users, tokens and remember_tokens migrations created and executed")
	color.Yellow("   - user and token models created")
	color.Yellow("   - auth middleware created")
	color.Yellow("")
	color.Yellow(" Don't forget to add user and token models in data/models.go and to add appropriate middleware to your routes!")

	return nil
}
