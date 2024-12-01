package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Struktur untuk menyimpan status idempotency
type IdempotencyStore struct {
	mu    sync.Mutex
	store map[string]IdempotencyData
}

type IdempotencyData struct {
	Response    interface{}
	CreatedAt   time.Time
	ExpiryAfter time.Duration
}

func NewIdempotencyStore() *IdempotencyStore {
	return &IdempotencyStore{
		store: make(map[string]IdempotencyData),
	}
}

// Menyimpan data idempotency
func (s *IdempotencyStore) Save(key string, data IdempotencyData) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = data
}

// Mendapatkan data idempotency
func (s *IdempotencyStore) Get(key string) (IdempotencyData, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, exists := s.store[key]
	if !exists || time.Since(data.CreatedAt) > data.ExpiryAfter {
		// Jika sudah expired, hapus data
		delete(s.store, key)
		return IdempotencyData{}, false
	}
	return data, true
}

var idempotencyStore = NewIdempotencyStore()

// Middleware Idempotent
func IdempotentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idempotencyKey := c.GetHeader("Idempotency-Key")
		if idempotencyKey == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Idempotency-Key header is required"})
			c.Abort()
			return
		}

		// Periksa apakah key sudah diproses sebelumnya
		if data, exists := idempotencyStore.Get(idempotencyKey); exists {
			c.JSON(http.StatusOK, gin.H{"message": "Request already processed", "data": data.Response})
			c.Abort()
			return
		}

		// Simpan key untuk sementara
		idempotencyStore.Save(idempotencyKey, IdempotencyData{
			Response:    "processing",
			CreatedAt:   time.Now(),
			ExpiryAfter: 10 * time.Minute,
		})
		c.Set("Idempotency-Key", idempotencyKey)

		c.Next()
	}
}
