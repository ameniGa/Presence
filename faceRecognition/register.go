package faceRecognition

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/ameniGa/timeTracker/models"
	"github.com/google/uuid"
	"github.com/machinebox/sdk-go/facebox"
	"gocv.io/x/gocv"
	"image/color"
	"log"
	"os"
	"path"
	"runtime"
	"time"
)

var (
	blue          = color.RGBA{0, 0, 255, 0}

	faceAlgorithm = "cascade/haarcascade_frontalface_default.xml"
	fbox          = facebox.New("http://localhost:8080")
)

func Register() models.User {
	// file path
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

	for count := 0; count < 3; count++ {
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
			fbox.Teach(bytes.NewReader(buf), userID, userID)
			imgFace.Close()
			if err != nil {
				log.Printf("unable to encode matrix: %v", err)
				continue
			}
		}
	}

	user := models.User{
		UserId:    userID,
		UserName:  username,
		CreatedAt: uint64(time.Now().Unix()),
	}
	return user
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
