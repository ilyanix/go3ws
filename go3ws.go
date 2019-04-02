package main

import ( 
	"fmt"
	"os/exec"
	"io"
	"github.com/mdirkse/i3ipc"
	"flag"
)

var (
	i3socket *i3ipc.IPCSocket
	rename bool
	move bool
)


func selector(s string) (res string) {
	rofiCmd := exec.Command("rofi", "-dmenu")
	rofiIn, _ := rofiCmd.StdinPipe()
	go func() {
		defer rofiIn.Close()
		io.WriteString(rofiIn, s)
	}()
	out, _ := rofiCmd.CombinedOutput()
	res = fmt.Sprintf("%s", out)
	return
}

func getWSstr(ws []i3ipc.Workspace) (list string) {
	for _, w := range ws {
		list += w.Name + "\n"
	}
	return
}

func setWS() {
	ws, _ := i3socket.GetWorkspaces()
	wStrAll := getWSstr(ws)
	wStr := selector(wStrAll)
	switch {
	case rename == true:
		i3socket.Command("rename workspace to " + wStr)
	case move == true:
		i3socket.Command("move container to workspace " + wStr)
	default:
		i3socket.Command("workspace " + wStr)
	}
}



func main() {
	flag.BoolVar(&rename, "rename", false, "rename workpace")
	flag.BoolVar(&move, "move", false, "move to workpace")
	flag.Parse()
	i3socket, _ = i3ipc.GetIPCSocket()
	setWS()
}

