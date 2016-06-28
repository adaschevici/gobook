package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/couchbase/gocb"
	"github.com/davecgh/go-spew/spew"
	"github.com/howeyc/fsnotify"
)

type Client struct {
	ID         int
	Connection *net.Conn
}

type File struct {
	Hash             string `json:hash`
	Name             string `json:file_name`
	Created          int64  `json:created`
	CreatedUser      int    `json:created_user`
	LastModified     int64  `json:last_modified`
	LastModifiedUser int    `json:last_modified_user`
	Revisions        int    `json:revisions`
	Version          int    `json:version`
}

type Message struct {
	Hash     string `json:hash`
	Action   string `json:action`
	Location string `json:location`
	Name     string `json:name`
	Version  int    `json:version`
}

func generateHash(name string) string {
	hash := md5.New()
	io.WriteString(hash, name)

	hashString := hex.EncodeToString(hash.Sum(nil))

	return hashString
}

func alertServers(hash string, name string, action string, location string, version int) {
	msg := Message{Hash: hash, Action: action, Location: location, Name: name, Version: version}
	msgJSON, _ := json.Marshal(msg)

	fmt.Println(string(msgJSON))

	for i := range Clients {
		fmt.Println("Sending to clients")
		fmt.Fprintln(*Clients[i].Connection, string(msgJSON))
	}
}

func startServer(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {

		}
		currentClient := Client{ID: 1, Connection: &conn}
		Clients = append(Clients, currentClient)
		for i := range Clients {
			fmt.Println("Client", Clients[i].ID)
		}
	}

}

func removeFile(name string, bucket *gocb.Bucket) {
	bucket.Remove(generateHash(name), 0)
}

func updateExistingFile(name string, bucket *gocb.Bucket) int {
	fmt.Println(name, "updated")
	hashString := generateHash(name)

	thisFile := Files[hashString]
	thisFile.Hash = hashString
	thisFile.Name = name
	log.Printf("The NAME: %v", name)
	thisFile.Version = thisFile.Version + 1
	thisFile.LastModified = time.Now().Unix()
	Files[hashString] = thisFile
	var file map[string]interface{}
	cas, err := bucket.Get(hashString, &file)
	if err != nil {
		log.Printf("Error is : %v", err)
	}
	bucket.Replace(hashString, &file, cas, 0)
	return thisFile.Version
}

func evalFile(event *fsnotify.FileEvent, bucket *gocb.Bucket) {
	fmt.Println(event.Name, "changed")
	create := event.IsCreate()
	fileComponents := strings.Split(event.Name, "\\")
	fileComponentSize := len(fileComponents)
	trueFileName := fileComponents[fileComponentSize-1]
	hashString := generateHash(trueFileName)

	if create == true {
		log.Printf("CREATE")
		updateFile(trueFileName, bucket)
		alertServers(hashString, event.Name, "CREATE", event.Name, 0)
	}

	delete := event.IsDelete()
	if delete == true {
		log.Printf("DELETE")
		removeFile(trueFileName, bucket)
		alertServers(hashString, event.Name, "DELETE", event.Name, 0)
	}

	modify := event.IsModify()
	if modify == true {
		log.Printf("MODIFY")
		newVersion := updateExistingFile(trueFileName, bucket)
		fmt.Println(newVersion)
		alertServers(hashString, trueFileName, "MODIFY", event.Name, newVersion)
	}

	rename := event.IsRename()
	if rename == true {
	}
}

func updateFile(name string, bucket *gocb.Bucket) {
	thisFile := File{}
	hashString := generateHash(name)

	thisFile.Hash = hashString
	thisFile.Name = name
	thisFile.Created = time.Now().Unix()
	thisFile.CreatedUser = 0
	thisFile.LastModified = time.Now().Unix()
	thisFile.LastModifiedUser = 0
	thisFile.Revisions = 0
	thisFile.Version = 1

	Files[hashString] = thisFile

	checkFile := File{}
	cas, err := bucket.Get(hashString, &checkFile)
	fmt.Printf("Got document so file is being watched %v.\n", cas)
	if err != nil {
		spew.Dump(err)
		fmt.Println("New File Added", name)
		bucket.Insert(hashString, thisFile, 0)
	}
	bucket.Get(hashString, &checkFile)
	log.Printf("File inserted: %s", checkFile.Name)
}

var Clients []Client
var Files map[string]File

func main() {
	Files = make(map[string]File)
	endScript := make(chan bool)
	someUser, _ := user.Current()
	currentUser := someUser.Username
	var listenFolder = fmt.Sprintf("/Users/%s/couchbase_volume/", currentUser)
	log.SetOutput(os.Stdout)
	log.Printf(listenFolder)
	cluster, err := gocb.Connect("couchbase://localhost/")
	if err != nil {
		fmt.Println("Error connecting to Couchbase", err)
	}
	bucket, err := cluster.OpenBucket("file_manager", "")
	if err != nil {
		fmt.Println("Error getting bucket", err)
	}

	files, _ := ioutil.ReadDir(listenFolder)
	for _, file := range files {
		updateFile(file.Name(), bucket)
	}

	dirSpy, err := fsnotify.NewWatcher()
	defer dirSpy.Close()

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Could not start server!", err)
	}

	go func() {
		for {
			select {
			case ev := <-dirSpy.Event:
				evalFile(ev, bucket)
			case err := <-dirSpy.Error:
				fmt.Println("error:", err)
			}
		}
	}()
	err = dirSpy.Watch(listenFolder)
	startServer(listener)

	<-endScript
}
