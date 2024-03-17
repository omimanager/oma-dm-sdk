package omadm

type AuthenticationProvider interface {
	authenticateRequest()
}
