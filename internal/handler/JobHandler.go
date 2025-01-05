package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vignesh/file_handler_service/internal/pool"
)

type JobRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func JobHandler(workerPool *pool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var req JobRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Add job to worker pool
		err := workerPool.AddJob(func() error {
			log.Printf("Processing job ID: %s, Name: %s\n", req.ID, req.Name)
			return nil // Simulated successful job processing
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusTooManyRequests)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Job submitted successfully"))
	}
}

type BatchJobRequest struct {
	Jobs []JobRequest `json:"jobs"`
}

func BatchJobHandler(workerPool *pool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var batchReq BatchJobRequest
		if err := json.NewDecoder(r.Body).Decode(&batchReq); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate jobs
		if len(batchReq.Jobs) == 0 {
			http.Error(w, "No jobs provided", http.StatusBadRequest)
			return
		}

		// Add each job to the worker pool
		for _, req := range batchReq.Jobs {
			job := req // Avoid closure issues with loop variables
			err := workerPool.AddJob(func() error {
				log.Printf("Processing job ID: %s, Name: %s\n", job.ID, job.Name)
				return nil
			})

			if err != nil {
				http.Error(w, err.Error(), http.StatusTooManyRequests)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Batch jobs submitted successfully"))
	}
}
