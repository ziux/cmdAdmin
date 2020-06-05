package cmd

var ProgramMap =make(map[int]*Program)


func AddMap(p *Program){
	pid := p.Process.Pid
	ProgramMap[pid] = p

}
func GetProgram(pid int)*Program{
	p, ok :=ProgramMap[pid]
	//fmt.Println(p, ok)
	if !ok {
		return nil
	}
	//if p.Cmd.ProcessState!=nil && p.Cmd.ProcessState.Exited(){
	//	return p
	//}else {
	//	delete(ProgramMap,pid)
	//}
	return p

}
func DelMap(pid int){
	delete(ProgramMap,pid)
}