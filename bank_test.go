package nicepay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBankMitra(t *testing.T) {
	assert.EqualValues(t, BankMandiri, "BMRI")
	assert.EqualValues(t, BankMaybank, "IBBK")
	assert.EqualValues(t, BankPermata, "BBBA")
	assert.EqualValues(t, BankBCA, "CENA")
	assert.EqualValues(t, BankBNI, "BNIN")
	assert.EqualValues(t, BankKEBHANA, "HNBN")
	assert.EqualValues(t, BankBRI, "BRIN")
	assert.EqualValues(t, BankCIMB, "BNIA")
	assert.EqualValues(t, BankDanamon, "BDIN")
	assert.EqualValues(t, BankOther, "OTHR")

	assert.EqualValues(t, MitraCVSAlfamart, "ALMA")
	assert.EqualValues(t, MitraCVSIndomaret, "INDO")
	assert.EqualValues(t, MitraClickPayMandiri, "MDRC")
	assert.EqualValues(t, MitraClickPayBCA, "BCAC")
	assert.EqualValues(t, MitraClickPayCIMB, "CIMC")
	assert.EqualValues(t, MitraEwalletMandiri, "MDRE")
	assert.EqualValues(t, MitraSakuku, "BCAE")
	assert.EqualValues(t, MitraAkulaku, "AKLP")
	assert.EqualValues(t, MitraKredivo, "KDVI")
	assert.EqualValues(t, MitraOvo, "OVOE")
}
