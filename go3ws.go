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
	if rename {
		i3socket.Command("rename workspace to " + wStr)
	}
		i3socket.Command("workspace " + wStr)
}



func main() {
	flag.BoolVar(&rename, "rename", false, "rename workpace")
	flag.Parse()
	i3socket, _ = i3ipc.GetIPCSocket()
	setWS()
}

