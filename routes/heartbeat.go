package routes

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v3"
)

type ServerVitals struct {
	Status struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    struct {
		Raw     string `json:"raw"`
		Days    int    `json:"days"`
		Hours   int    `json:"hours"`
		Minutes int    `json:"minutes"`
		Seconds int    `json:"seconds"`
	} `json:"uptime"`
	Memory struct {
		Alloc      string `json:"alloc"`
		TotalAlloc string `json:"totalAlloc"`
		Sys        string `json:"sys"`
		NumGC      uint32 `json:"numGC"`
		HeapAlloc  string `json:"heapAlloc"`
		HeapSys    string `json:"heapSys"`
		Free       string `json:"free"`
		Used       string `json:"used"`
		Usage      string `json:"usage"`
	} `json:"memory"`
	System struct {
		Goroutines int    `json:"goroutines"`
		NumCPU     int    `json:"numCPU"`
		GoVersion  string `json:"goVersion"`
		OS         string `json:"os"`
		Arch       string `json:"arch"`
	} `json:"system"`
}

var startTime = time.Now()

// formatBytes converts bytes to human-readable format
func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// parseUptime breaks down the uptime into days, hours, minutes, and seconds
func parseUptime(d time.Duration) (days, hours, minutes, seconds int) {
	seconds = int(d.Seconds())
	minutes = seconds / 60
	seconds = seconds % 60
	hours = minutes / 60
	minutes = minutes % 60
	days = hours / 24
	hours = hours % 24
	return
}

// HeartbeatHandler returns server vitals and health status
func HeartbeatHandler(c fiber.Ctx) error {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	uptime := time.Since(startTime)
	days, hours, minutes, seconds := parseUptime(uptime)

	vitals := ServerVitals{
		Timestamp: time.Now(),
	}

	// Status information
	vitals.Status.Code = "healthy"
	vitals.Status.Message = "Server is running normally"

	// Uptime information
	vitals.Uptime.Raw = uptime.String()
	vitals.Uptime.Days = days
	vitals.Uptime.Hours = hours
	vitals.Uptime.Minutes = minutes
	vitals.Uptime.Seconds = seconds

	// Memory information
	vitals.Memory.Alloc = formatBytes(m.Alloc)
	vitals.Memory.TotalAlloc = formatBytes(m.TotalAlloc)
	vitals.Memory.Sys = formatBytes(m.Sys)
	vitals.Memory.NumGC = m.NumGC
	vitals.Memory.HeapAlloc = formatBytes(m.HeapAlloc)
	vitals.Memory.HeapSys = formatBytes(m.HeapSys)

	// Calculate free and used memory
	free := m.Sys - m.Alloc
	used := m.Alloc
	usagePercent := float64(used) / float64(m.Sys) * 100

	vitals.Memory.Free = formatBytes(free)
	vitals.Memory.Used = formatBytes(used)
	vitals.Memory.Usage = fmt.Sprintf("%.2f%%", usagePercent)

	// System information
	vitals.System.Goroutines = runtime.NumGoroutine()
	vitals.System.NumCPU = runtime.NumCPU()
	vitals.System.GoVersion = runtime.Version()
	vitals.System.OS = runtime.GOOS
	vitals.System.Arch = runtime.GOARCH

	return c.Status(fiber.StatusOK).JSON(vitals)
}
