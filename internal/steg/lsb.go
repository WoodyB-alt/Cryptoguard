package steg

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
)

// EmbedStringInPNG embeds a base64 string into the LSBs of a PNG image.
func EmbedStringInPNG(inputImgPath, outputImgPath, message string) error {
	// Step 1: Open the image
	inFile, err := os.Open(inputImgPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	img, err := png.Decode(inFile)
	if err != nil {
		return err
	}

	bounds := img.Bounds()
	rgbaImg := image.NewNRGBA(bounds)

	// Copy pixels to RGBA buffer
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgbaImg.Set(x, y, img.At(x, y))
		}
	}

	// Convert message to binary
	bits := toBits([]byte(message))
	if len(bits) > len(rgbaImg.Pix) {
		return errors.New("message too large to hide in image")
	}

	// Step 2: Embed bits into LSB of each byte
	for i := 0; i < len(bits); i++ {
		rgbaImg.Pix[i] = (rgbaImg.Pix[i] & 0xFE) | bits[i]
	}

	// Step 3: Write output PNG
	outFile, err := os.Create(outputImgPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, rgbaImg)
}

// toBits converts a byte slice into a slice of individual bits (0 or 1)
func toBits(data []byte) []byte {
	bits := make([]byte, 0, len(data)*8)
	for _, b := range data {
		for i := 7; i >= 0; i-- {
			bits = append(bits, (b>>i)&1)
		}
	}
	return bits
}

// ExtractStringFromPNG reads bits from the LSBs of the image and reconstructs the hidden message.
// It stops once it encounters 8 zero-bytes (used as end-of-message marker).
func ExtractStringFromPNG(imgPath string) (string, error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return "", err
	}

	bounds := img.Bounds()
	nrgba := image.NewNRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			nrgba.Set(x, y, img.At(x, y))
		}
	}

	// Read LSBs
	bits := make([]byte, 0, len(nrgba.Pix))
	for _, b := range nrgba.Pix {
		bits = append(bits, b&1)
	}

	// Convert bits to bytes
	data := fromBits(bits)

	// Detect end marker (8 null bytes in a row)
	end := findStringMarker(data, "<<<END>>>")
	if end == -1 {
		return "", fmt.Errorf("no end marker found; possibly corrupted or incomplete")
	}

	return string(data[:end]), nil
}

// findStringMarker finds a string in a byte slice and returns its index
func findStringMarker(data []byte, marker string) int {
	markerBytes := []byte(marker)
	for i := 0; i < len(data)-len(markerBytes); i++ {
		if string(data[i:i+len(markerBytes)]) == marker {
			return i
		}
	}
	return -1
}

// fromBits converts 8 bits at a time into bytes
func fromBits(bits []byte) []byte {
	data := make([]byte, 0, len(bits)/8)
	for i := 0; i+7 < len(bits); i += 8 {
		var b byte
		for j := 0; j < 8; j++ {
			b = (b << 1) | bits[i+j]
		}
		data = append(data, b)
	}
	return data
}

// findEndMarker looks for 8 null bytes (end of message)
func findEndMarker(data []byte) int {
	for i := 0; i < len(data)-8; i++ {
		if isAllZero(data[i : i+8]) {
			return i
		}
	}
	return -1
}

// isAllZero checks if all bytes are 0
func isAllZero(slice []byte) bool {
	for _, b := range slice {
		if b != 0 {
			return false
		}
	}
	return true
}
