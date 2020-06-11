package nicepay

type Environment int

const (
	_ Environment = iota
	Development
	Production
)

var envApiEndpoint = map[Environment]string{
	Development: "https://dev.nicepay.co.id",
	Production:  "https://api.nicepay.co.id",
}

func (e Environment) String() string {
	for k, v := range envApiEndpoint {
		if k == e {
			return v
		}
	}
	return ""
}
