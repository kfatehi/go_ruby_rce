package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func getRuby() string {
	return getEnv("RUBY", "ruby")
}

func checkRuby() {
	rubyCmd := exec.Command(getRuby(), "-e", "printf RUBY_VERSION")
	rubyOut, err := rubyCmd.Output()
	if err != nil {
		fmt.Println("No working ruby:", err, string(rubyOut))
		os.Exit(1)
		return
	}
	fmt.Println("Tested working ruby", string(rubyOut))
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	ruby := r.Group("/ruby")

	// curl -XPOST localhost:8080/ruby/validate -F file=script.rb
	ruby.POST("/validate", func(ctx *gin.Context) {
		rubyCmd := exec.Command(getRuby(), "-e", `
			require 'ripper'
			require 'json'
			sexp = Ripper.sexp(ARGF)
			puts JSON.generate sexp
		`)
		stdin, err := rubyCmd.StdinPipe()
		if err != nil {
			ctx.String(http.StatusInternalServerError, "stdin creation error: %s", err.Error())
			return
		}
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}
		openedFile, err := file.Open()
		if err != nil {
			ctx.String(http.StatusBadRequest, "open file error: %s", err.Error())
			return
		}
		sourceCode, _ := ioutil.ReadAll(openedFile)
		io.WriteString(stdin, string(sourceCode))
		stdin.Close()
		rubyOut, err := rubyCmd.Output()
		if err != nil {
			ctx.String(http.StatusInternalServerError, "ruby invocation error: %s\n%s", err.Error(), string(rubyOut))
			return
		}
		ctx.Data(http.StatusOK, "application/json", rubyOut)
	})
	return r
}

func main() {
	checkRuby()

	gin.SetMode(gin.ReleaseMode)

	router := setupRouter()

	listenStr := getEnv("LISTEN_ADDR", "localhost") + ":" + getEnv("LISTEN_PORT", "8080")
	fmt.Println("Listening on", listenStr)
	router.Run(listenStr)
}
