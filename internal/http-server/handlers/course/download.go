package course

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func DownloadCourseHandler(logger *zap.SugaredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract courseID from URL parameters
		courseIDStr := chi.URLParam(r, "id")
		courseID, err := strconv.Atoi(courseIDStr)
		if err != nil {
			logger.Error(err, "Invalid course ID")
			http.Error(w, "Invalid course ID", http.StatusBadRequest)
			return
		}

		// Construct the URL for /course/{id}
		host := r.Host
		if host == "" {
			host = "localhost:8080" // Default host if not available
		}
		reqURLStr := fmt.Sprintf("http://%s/course/%d", host, courseID)
		reqURL, err := url.Parse(reqURLStr)
		if err != nil {
			logger.Error(err, "Failed to parse URL")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Create a client request to /course/{id}
		req, err := http.NewRequest("GET", reqURL.String(), nil)
		if err != nil {
			logger.Error(err, "Failed to create request")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Add a custom header to indicate this is a download request
		req.Header.Set("X-Download-Request", "true")

		// Make the client request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logger.Error(err, "Failed to make request to /course/{id}")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Check if the request was successful
		if resp.StatusCode != http.StatusOK {
			logger.Error("Failed to fetch course content with status:", resp.StatusCode)
			http.Error(w, "Failed to fetch course content", resp.StatusCode)
			return
		}

		// Read the response body
		htmlContent, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Error(err, "Failed to read response body")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Create a ZIP archive containing the HTML content
		var buf bytes.Buffer
		zipWriter := zip.NewWriter(&buf)

		// Add the HTML content to the archive
		file, err := zipWriter.Create("course.html")
		if err != nil {
			logger.Error(err, "Failed to create zip file")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		_, err = file.Write(htmlContent)
		if err != nil {
			logger.Error(err, "Failed to write to zip file")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Close the ZIP writer to finalize the archive
		err = zipWriter.Close()
		if err != nil {
			logger.Error(err, "Failed to close zip writer")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Set headers for download
		w.Header().Set("Content-Disposition", "attachment; filename=course.zip")
		w.Header().Set("Content-Type", "application/zip")

		// Write the ZIP archive to the response
		_, err = buf.WriteTo(w)
		if err != nil {
			logger.Error(err, "Failed to write response")
			return
		}
	}
}
