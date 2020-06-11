package nicepay

type BankCode string

const (
	BankMandiri BankCode = "BMRI"
	BankMaybank          = "IBBK"
	BankPermata          = "BBBA"
	BankBCA              = "CENA"
	BankBNI              = "BNIN"
	BankKEBHANA          = "HNBN"
	BankBRI              = "BRIN"
	BankCIMB             = "BNIA"
	BankDanamon          = "BDIN"
	BankOther            = "OTHR"
)

type MitraCode string

const (
	MitraCVSAlfamart     = "ALMA"
	MitraCVSIndomaret    = "INDO"
	MitraClickPayMandiri = "MDRC"
	MitraClickPayBCA     = "BCAC"
	MitraClickPayCIMB    = "CIMC"
	MitraEwalletMandiri  = "MDRE"
	MitraSakuku          = "BCAE"
	MitraAkulaku         = "AKLP"
	MitraKredivo         = "KDVI"
	MitraOvo             = "OVOE"
)
