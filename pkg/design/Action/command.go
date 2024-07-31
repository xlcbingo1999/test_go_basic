package Action

import "fmt"

// 命令模式的核心思想
// 命令可以用链表串联起来, 命令可以被撤回
// 控制器, 用于全局控制所有的命令, 可以包含N个种类的命令, 需要设置一个preCmd用于撤回相关的命令

const labelNoCommand = "Label_No_Command"

type command interface {
	execute()
	undo()
}

type light struct {
	label string
}

func (l *light) on() {
	fmt.Println("light is on")
}

func (l *light) off() {
	fmt.Println("light is off")
}

type Command interface {
	execute()
	undo()
}

type lightOnCommand struct {
	l *light // 需要设置一个Command实际作用的对象
}

func (c *lightOnCommand) execute() {
	c.l.on()
}

func (c *lightOnCommand) undo() {
	c.l.off()
}

type lightOffCommand struct {
	l *light
}

func (c *lightOffCommand) execute() {
	c.l.off()
}

func (c *lightOffCommand) undo() {
	c.l.on()
}

type NoCommand struct{}

func (c *NoCommand) execute() {
	fmt.Println("NoCommand Execute")
}

func (c *NoCommand) undo() {
	fmt.Println("NoCommand Undo")
}

type CommandController struct {
	onCommands  map[string]Command
	offCommands map[string]Command
	preCmd      Command
}

func createController() *CommandController {
	rc := &CommandController{
		onCommands:  make(map[string]Command),
		offCommands: make(map[string]Command),
	}

	noCMD := new(NoCommand)
	rc.onCommands[labelNoCommand] = noCMD
	rc.offCommands[labelNoCommand] = noCMD
	rc.preCmd = noCMD
	return rc
}

func (cc *CommandController) setCommand(label string, onCMD Command, offCMD Command) {
	cc.onCommands[label] = onCMD
	cc.offCommands[label] = offCMD
}

func (cc *CommandController) onButtonWasPressed(label string) {
	// 每次拿出对应label的Command进行执行
	cmd := cc.onCommands[label]
	if cmd == nil {
		cmd = cc.onCommands[labelNoCommand]
	}

	cmd.execute()
	cc.preCmd = cmd
}

func (cc *CommandController) offButtonWasPressed(label string) {
	cmd := cc.offCommands[label]
	if cmd == nil {
		cmd = cc.offCommands[labelNoCommand]
	}
	cmd.execute()
	cc.preCmd = cmd
}

func (cc *CommandController) undo() {
	cc.preCmd.undo()
}

func RunCommand() {
	slight := &light{
		label: "beijing",
	}
	lightOnCmd := &lightOnCommand{l: slight}
	lightOffCmd := &lightOffCommand{l: slight}

	ctl := createController()
	ctl.setCommand(slight.label, lightOnCmd, lightOffCmd)

	ctl.onButtonWasPressed(slight.label)
	ctl.offButtonWasPressed(slight.label)
	ctl.undo()

	ctl.onButtonWasPressed("undefined")
}
