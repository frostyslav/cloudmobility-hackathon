package model

import "sync"

// FuncCreateResponse - response for /create_func_from_repo.
// swagger:model FuncCreateResponse
type FuncCreateResponse struct {
	// Request ID
	ID string `json:"id"`
}

// FuncCreateResponse - response for /create_func_from_repo.
// swagger:model FuncCreateResponse
type FuncStatusResponse struct {
	// Request ID
	ID     string `json:"id"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

type FuncCreateRequest struct {
	Repo     Repo   `json:"repo"`
	Language string `json:"language"`
	Code     string `json:"code"`
}

type FuncStatusRequest struct {
	ID string `json:"id"`
}

type Repo struct {
	URL  string `json:"url"`
	Tag  string `json:"tag"`
	Path string `json:"path"`
}

type Hash struct {
	data map[string]*request
	mu   sync.RWMutex
}

type request struct {
	status string
	url    string
}

func (hash *Hash) SetStatus(key, value string) {
	hash.mu.Lock()
	hash.data[key] = &request{status: value}
	hash.mu.Unlock()
}

func (hash *Hash) SetURL(key, url string) {
	hash.mu.Lock()
	hash.data[key] = &request{url: url}
	hash.mu.Unlock()
}

func (hash *Hash) GetStatus(key string) string {
	hash.mu.RLock()
	defer hash.mu.RUnlock()
	val, ok := hash.data[key]
	if ok {
		return val.status
	}

	return ""
}

func (hash *Hash) New() *Hash {
	hash.data = make(map[string]*request)
	return hash
}
