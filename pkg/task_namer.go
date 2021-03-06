package variant

import "strings"

type TaskNamer struct {
	AppName string
}

func NewTaskNamer(appName string) *TaskNamer {
	return &TaskNamer{
		AppName: appName,
	}
}

func (p TaskNamer) FromString(flowKeyStr string) TaskName {
	c := strings.Split(flowKeyStr, ".")
	return TaskName{Components: c}
}

func (p TaskNamer) FromResolvedInput(variable *Input) TaskName {
	return p.FromInputName(variable.Name)
}

func (p TaskNamer) FromInput(input *InputConfig) TaskName {
	return p.FromInputName(input.Name)
}

func (p TaskNamer) FromInputName(inputName string) TaskName {
	c := strings.Split(p.AppName+"."+inputName, ".")
	return TaskName{Components: c}
}
