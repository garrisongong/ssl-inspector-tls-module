package tls

// kTLSCipher is a placeholder to tell the record layer to skip wrapping.
type kTLSCipher struct{}

func (c *Conn) EnableKTLS() error {
	// Enable kernel TLS if possible
	return c.enableKernelTLS(c.cipherSuite, c.in.key, c.out.key, c.in.iv, c.out.iv)
}
