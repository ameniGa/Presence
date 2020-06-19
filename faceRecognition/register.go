package faceRecognition

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/ameniGa/timeTracker/config"
	"github.com/ameniGa/timeTracker/database"
	hlp "github.com/ameniGa/timeTracker/helpers"
	ctxUtl "github.com/ameniGa/timeTracker/helpers/context"
	"github.com/ameniGa/timeTracker/notification/slack"
	"github.com/google/uuid"
	"github.com/machinebox/sdk-go/facebox"
	"gocv.io/x/gocv"
	"image/color"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

var (
	red          = color.RGBA{255, 0, 0, 0}
	conf          *config.Config
	faceAlgorithm = "cascade/haarcascade_frontalface_default.xml"
	imgDir        = "img"
	fbox          *facebox.Client
	slackHandler = slack.Slack{}
	entries = map[string]bool{}
)

func init() {
	var err error
	conf, err = config.LoadConfig()
	if err != nil {
		log.Fatalf("cannot load config %v", err)
	}
	// create new fbox client
	fbox = facebox.New(conf.Facebox.Url)
	slackHandler = slack.NewSlackHandler(conf.Notification.Slack)
}

func Register() {
	// cascade file path
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), faceAlgorithm)
	// set user name
	username := getUserInput()
	// set the user ID
	userID := uuid.New().String()
	// open webcam. 0 is the default device ID
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatalf("error opening web cam: %v", err)
	}
	defer webcam.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// open display window
	window := gocv.NewWindow("packagemain")
	defer window.Close()

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(filepath)
	playSound()
	for count := 0; count < conf.Facebox.PictureNumber; count++ {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Print("cannot read image from the cam")
			continue
		}
		// detect faces
		rects := classifier.DetectMultiScale(img)
		for _, r := range rects {
			// Save each found face into the file
			imgFace := img.Region(r)
			imgName := fmt.Sprintf("img/%s.%d.jpg", userID, count)
			imgPath := path.Join(path.Dir(filename), imgName)
			gocv.IMWrite(imgPath, imgFace)
			buf, err := gocv.IMEncode(".jpg", imgFace)
			// train
			// todo cleanup images
			fbox.Teach(bytes.NewReader(buf), userID, username)
			imgFace.Close()
			if err != nil {
				log.Printf("unable to encode matrix: %v", err)
				continue
			}
		}
	}
	saveUser(userID, username)
	cleanup()
}

func getUserInput() string {
	// get username from input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter User Name")
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("cannot get user input %v", err)
	}
	return username
}

func saveUser(userID, userName string) {
	log := hlp.GetLogger()
	handler, err := database.Create(&conf.Database, conf.Server.Deadline, log)
	if err != nil {
		log.Fatalf("cannot create handler %v", err)
	}
	ctx, cancel := ctxUtl.AddTimeoutToCtx(context.Background(), 5)
	defer cancel()
	ch := make(chan error, 1)
	handler.DbAddUser(ctx, userID, userName, ch)
}

func cleanup() {
	_, filename, _, _ := runtime.Caller(0)
	imgPath := path.Join(path.Dir(filename), imgDir)
	files, err := filepath.Glob(filepath.Join(imgPath, "*"))
	if err != nil {
		fmt.Println("cannot read the path", err)
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			fmt.Println("cannot remove the file", err)
		}
	}
}

func playSound() {
	text := "look to the camera"
	cmd := exec.Command("espeak", text)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	time.Sleep(5000)
}