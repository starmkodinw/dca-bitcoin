package services

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

const MARKET = "market"
const SYM = "btc_thb"
const AMOUNT = 10.0 // 10 THB
const URL = "https://api.bitkub.com"

var API_KEY = ""
var API_SECRET = ""

func init() {
	_, b, _, _ := runtime.Caller(0)

	projectRootPath := filepath.Join(filepath.Dir(b), "../")

	_ = godotenv.Load(projectRootPath + "/.env")

	API_KEY = os.Getenv("API_KEY")
	API_SECRET = os.Getenv("API_SECRET")
}
