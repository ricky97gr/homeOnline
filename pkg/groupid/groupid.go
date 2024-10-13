package groupid

import (
	"math/rand"
	"strconv"
	"time"
)

const SuperAdministratorGroup = 10000000

func GetGroupID() string {
	rand.NewSource(time.Now().Unix())
	id := rand.Intn(89999999)
	return "g_" + strconv.Itoa(id+SuperAdministratorGroup)
}
