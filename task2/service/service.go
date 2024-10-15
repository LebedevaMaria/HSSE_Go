package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type Input struct {
	Input string `json:"input"`
}

type Output struct {
	Output string `json:"output"`
}

const (
	version = "1.0.0"
)

func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf(`%s`, version)))
}

func PostDecode(w http.ResponseWriter, r *http.Request) {
	var input Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(input.Input)
	if err != nil {
		http.Error(w, "Invalid base64 string", http.StatusBadRequest)
		return
	}

	output := Output{string(decoded)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func GetHardOp(w http.ResponseWriter, r *http.Request) {
	sleepDuration := time.Duration(10+rand.Intn(11)) * time.Second
	time.Sleep(sleepDuration)

	if rand.Intn(10) <= 5 {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	const shutdownTimeout = 15 * time.Second
	mux := http.NewServeMux()
	mux.HandleFunc("/version", GetVersion)
	mux.HandleFunc("/decode", PostDecode)
	mux.HandleFunc("/hard-op", GetHardOp)

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("err in listen: %s\n", err)
			return fmt.Errorf("failed to serve http server: %w", err)
		}
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			return err
		}
		return nil
	})

	err := group.Wait()
	if err != nil {
		return
	}
}
