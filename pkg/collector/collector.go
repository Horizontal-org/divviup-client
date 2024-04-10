package collector

import (
	"bytes"
	"os/exec"
)


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