package squares

import (
	"crypto/md5"
	"fmt"
	"image"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	key string
)

func init() {
	h := md5.New()
	io.WriteString(h, "hello")
	key = fmt.Sprintf("%x", h.Sum(nil)[:])
}

func TestGrid(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	Grid(img, colorTheme[0], colorTheme[1])
}

func TestGridSVG(t *testing.T) {
	rec := httptest.NewRecorder()
	GridSVG(rec, colorTheme[0], colorTheme[1], 10)
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}
}

func TestSquares(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	Squares(img, key, colorTheme)
}

func TestSquaresSVG(t *testing.T) {
	rec := httptest.NewRecorder()
	SquaresSVG(rec, key, colorTheme, 10)
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}

}
