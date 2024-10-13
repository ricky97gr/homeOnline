package conf

import (
	"fmt"
	"os"
	"testing"

	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
)

func TestLoad(t *testing.T) {
	newlog.InitLogger("", os.Stdout)
	configPath = "config.yaml"
	fmt.Println(GetConfig())
}
