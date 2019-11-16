package network

// This is the driver. It will feature a goroutine that will read off packets, aggregate them into large buckets, and then
// forward them out.