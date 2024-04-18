package collector

import (
	"bytes"
	"divviup-client/pkg/common/models"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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



func MockCollector(DB *gorm.DB, vdaf string, divviup string, taskid uint) {

	arguments := CollectorArguments{
		// make manifest env variable
		Manifest: "/home/juan/code/janus-0.7.0-prerelease-2/Cargo.toml",
		TaskId: divviup,
		VdafType: vdaf,
		LeaderUrl: "https://dap-07-1.api.divviup.org/",
	}


	collectorOut, collectorSuccess := RunMockCollector(&arguments)
	
	// do some processing in real case, need to know what divviup actually returns
	taskEvent := models.TaskEvent{TaskID: taskid}
	// {Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	
	if collectorSuccess {
		cleanOut := CleanOutput(collectorOut, vdaf)
		taskEvent.Value = cleanOut
	} else {
		taskEvent.Error = collectorOut
	}

	DB.Create(&taskEvent)

	log.Println("Output: ", collectorOut)
	log.Println("Success: ", collectorSuccess)
}




func RunMockCollector(arg *CollectorArguments) (outString string, outStatus bool) {

	// create a new *Cmd instance
	// here we pass the command as the first argument and the arguments to pass to the command as the
	// remaining arguments in the function
	// argsToString := "-m " + arg.Manifest + " -t " + arg.TaskId + " -l " + arg.LeaderUrl + " -V " + arg.VdafType

	parsedArgs := []string{}
	parsedArgs = append(parsedArgs, "-m")
	parsedArgs = append(parsedArgs, arg.Manifest)	
	parsedArgs = append(parsedArgs, "-t")
	parsedArgs = append(parsedArgs, arg.TaskId)
	parsedArgs = append(parsedArgs, "-l")
	parsedArgs = append(parsedArgs, arg.LeaderUrl)
	parsedArgs = append(parsedArgs, "-V")
	parsedArgs = append(parsedArgs, arg.VdafType)	
	cmd := exec.Command("./scripts/mock.sh", parsedArgs...)



	var outB bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &outB
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
			// fmt.Println("ERROR " + fmt.Sprint(err) + ": " + stderr.String())
			return stderr.String(), false
	}	
	return outB.String(), true
}

// Clean divviup api response
func CleanOutput (output string, vdaf string) (string) {
	data := strings.Split(strings.TrimSpace(strings.TrimSuffix(output, "\n")), ": ")


	return data[1]
}