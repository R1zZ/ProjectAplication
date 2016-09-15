package main
import(
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"net/http"
)

var (
	images = make(map[string]image.Image)
)

func HandleRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, 
		`<html><body>
			<form method"post" enctype="multipart/form-data" action="/upload" name="upload">

				<label for="file">Choose image:</label>
				<input type="file" name="image" /><br/>
				<input type="submit" name="submit" value="Upload" />
			</form>
		</body></html>
`)
}

func HandleUpload(w http.ResponseWriter, r *http.Request){
	file, header, _ := r.FormFile("image")
	image, _, _ := image.Decode(file)
	images[header.Filename] = image
	http.Redirect(w, r, "/editor?name=+header.Filename, 303")
}

func HandleImage(w http.ResponseWriter, r *http.Request){
	imageName := r.FormValue("name")
	image := images[imageName]

jpeg.Encode(w, image, &jpeg.Options{Quality: jpeg.DefaultQuality})
}

func HandleEditor(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, ``)
}
