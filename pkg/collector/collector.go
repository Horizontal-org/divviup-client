package collector

import (
	"bytes"
	"divviup-client/pkg/common/models"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)



func ScheduledCollector(DB *gorm.DB, vdaf string, divviup string, taskid uint) {
	log.Print(vdaf)


	arguments := CollectorArguments{
		// make manifest env variable
		Manifest: viper.Get("JANUS_MANIFEST").(string),
		TaskId: divviup,
		VdafType: VdafParse(vdaf),
		LeaderUrl: viper.Get("DIVVIUP_LEADER_URL").(string),
		CredentialFile: viper.Get("COLLECTOR_CREDENTIAL_FILE").(string),
	}


	collectorOut, collectorSuccess := RunCollector(&arguments)
	log.Println("Output: ", collectorOut)
	log.Println("Success: ", collectorSuccess)
	// do some processing in real case, need to know what divviup actually returns
	taskEvent := models.TaskEvent{TaskID: taskid}
	taskEvent.Success = collectorSuccess
	taskEvent.Output = collectorOut

	if collectorSuccess {
		cleanOut := CleanOutput(collectorOut, vdaf)
		taskEvent.Value = cleanOut
	}

	DB.Create(&taskEvent)

	log.Println("----FINISHED----")
}



func RunCollector(arg *CollectorArguments) (outString string, outStatus bool) {

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
	parsedArgs = append(parsedArgs, "-c")
	parsedArgs = append(parsedArgs, arg.CredentialFile)	
	cmd := exec.Command("./scripts/collect.sh", parsedArgs...)



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

// Add vdaf specific params
func VdafParse (vdaf string) (string) {
	if vdaf == "sum" {
		return  "sum --bits=16"
	}
	return vdaf
} 

// Clean divviup api response
func CleanOutput (output string, vdaf string) (string) {
	firstStep := strings.Split(strings.TrimSpace(strings.TrimSuffix(output, "\n")), "Aggregation result: ")
	firstLine := strings.Split(firstStep[1], "\n")
	return firstLine[0]
}