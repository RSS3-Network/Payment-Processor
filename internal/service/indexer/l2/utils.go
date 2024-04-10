package l2

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"math/big"
	"strings"
)

// encodeInput encodes the input data according to the contract ABI
func (s *server) encodeInput(contractABI, methodName string, args ...interface{}) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return nil, err
	}

	encodedArgs, err := parsedABI.Pack(methodName, args...)
	if err != nil {
		return nil, err
	}

	return encodedArgs, nil
}

func scaleGwei(in *big.Int) {
	in.Mul(in, big.NewInt(1e18))
}
