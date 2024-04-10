package collector

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestCollectorResponse struct {
	success bool
	output string
}

func (h handler) TestCollector(c *gin.Context) {

	arguments := CollectorArguments{
		// make manifest env variable
		Manifest: "/home/juan/code/janus-0.7.0-prerelease-2/Cargo.toml",
		TaskId: "19tPTIP7gYq2mcDvyq62aUGg_PSWB17QDvaVbGb5tFI",
		VdafType: "count",
		LeaderUrl: "https://dap-07-1.api.divviup.org/",
	}


	collectorOut, collectorSuccess := RunCollector(&arguments)

	log.Println("Output: ", collectorOut)
	log.Println("Success: ", collectorSuccess)

	response := TestCollectorResponse{
		success: collectorSuccess,
		output: collectorOut,
	}

	c.JSON(http.StatusOK, response)

}
