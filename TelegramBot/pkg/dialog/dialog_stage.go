package dialog

type Stage string

const (
	Secret   Stage = "Secret"
	Verify   Stage = "Verify"
	Service  Stage = "Service"
	Login    Stage = "Login"
	Password Stage = "Password"
)
