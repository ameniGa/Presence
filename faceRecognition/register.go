package faceRecognition

import (
	"bytes"
	"fmt"
	"github.com/machinebox/sdk-go/facebox"
	"gocv.io/x/gocv"
	"log"
)
var (
	faceAlgorithm = "cascade/haarcascade_frontalface_default.xml"
	fbox          = facebox.New("http://localhost:8080")
)

func Register(){
	// get username from input



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
	classifier.Load(faceAlgorithm)

	for count:=0; count<3; count++ {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Print("cannot read image from the cam")
			continue
		}

		// detect faces
		rects := classifier.DetectMultiScale(img)
		for _, r := range rects {
			// Save each found face into the file
			imgFace := img.Region(r)
			imgName := fmt.Sprintf("%d.jpg", count)
			gocv.IMWrite(imgName, imgFace)
			buf, err := gocv.IMEncode(".jpg", imgFace)
			fbox.Teach(bytes.NewReader(buf),"1","moez")
			imgFace.Close()
			if err != nil {
				log.Printf("unable to encode matrix: %v", err)
				continue
			}
		}
	}
}
