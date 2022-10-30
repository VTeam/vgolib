package mr

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
	"time"
)

var (
	WorkerTimeOutDuration = time.Duration * time.Minute
)

type Coordinator struct {
	sync.RWMutex
	Files []string
	// Your definitions here.

}

type MapedFile struct {
	FileName  string
	StartTime time.Time
}

// Your code here -- RPC handlers for the worker to call.

//
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
//
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

func (c *Coordinator) NextMap(args *MapReq, reply *MapResp) error {
	f, err := os.Open("already_map")
	if err != nil {
		return err
	}
	defer f.Close()

	c.Lock()
	defer c.Unlock()

	var HasMapedFiles []MapedFile

	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	for fs.Scan() {
		HasMapedFiles = append(HasMapedFiles, MapedFile{FileName: fs.Text(), StartTime: time.Now()})
	}

	for _, file := range c.Files {
		if !SliceHasStringAndInTime(HasMapedFiles, file) {
			b, err := json.Marshal(HasMapedFiles)
			if err != nil {
				return err
			}
			fmt.Fprintf(f, "%s\n", string(b))
			reply.MapInputFileName = file
			return nil
		}
	}
	return nil
}

func SliceHasStringAndInTime(all []MapedFile, target string) bool {
	for _, one := range all {
		if one.FileName == target || one.StartTime+WorkerTimeOutDuration < time.Now() {
			return true
		}
	}
	return false
}

func (c *Coordinator) NextReduce(args *ReduceReq, reply *ReduceResp) error {

	return nil
}

//
// start a thread that listens for RPCs from worker.go
//
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
//
func (c *Coordinator) Done() bool {
	ret := false

	// Your code here.

	return ret
}

//
// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
//
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{
		Files: files,
	}

	// Your code here.

	c.server()
	return &c
}
