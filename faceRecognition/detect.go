package faceRecognition

import (
"bytes"
	"context"
	"fmt"
	"github.com/ameniGa/timeTracker/database"
	hlp "github.com/ameniGa/timeTracker/helpers"
	ctxUtl "github.com/ameniGa/timeTracker/helpers/context"
	"gocv.io/x/gocv"
"image"
"log"
)
func Dedect() {
	// open webcam. 0 is the default device ID, change it if your device ID is different
	webcam, err := gocv.VideoCaptureDevice(conf.Camera.DeviceID)
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
	faceAlgorithm := "faceRecognition/cascade/haarcascade_frontalface_default.xml"
	classifier.Load(faceAlgorithm)
	for {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Print("cannot read webcam")
			continue
		}

		// detect faces
		rects := classifier.DetectMultiScale(img)

		for _, r := range rects {
			// Save each found face into the file
			imgFace := img.Region(r)
			buf, err := gocv.IMEncode(".jpg", imgFace)
			if err != nil {
				log.Printf("unable to encode matrix: %v", err)
				continue
			}
			faces, err := fbox.Check(bytes.NewReader(buf))
			if err != nil {
				log.Printf("unable to recognize face: %v", err)
			}

			var caption = "I don't know you"
			if len(faces) > 0 {
				caption = fmt.Sprintf("%s", faces[0].Name)
				userID := fmt.Sprintf("%s", faces[0].ID)
				saveEntry(userID)
				// todo
				sendToSlack(userID)
			}

			// draw rectangle for the face
			size := gocv.GetTextSize(caption, gocv.FontHersheyPlain, 3, 2)
			pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
			gocv.PutText(&img, caption, pt, gocv.FontHersheyPlain, 3, red, 2)
			gocv.Rectangle(&img, r, red, 3)
		}

		// show the image in the window, and wait 100ms
		window.IMShow(img)
		window.WaitKey(100)
	}
}

func saveEntry(userID string) {
	log := hlp.GetLogger()
	handler, err := database.Create(&conf.Database, conf.Server.Deadline, log)
	if err != nil {
		log.Fatalf("cannot create handler %v", err)
	}
	ctx, cancel := ctxUtl.AddTimeoutToCtx(context.Background(), 5)
	defer cancel()
	ch := make(chan error, 1)
	handler.DbAddEntry(ctx, userID, ch)
}

func sendToSlack(userID string){

}