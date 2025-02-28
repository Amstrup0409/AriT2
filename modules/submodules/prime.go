package submodules

import (
	"errors"
	"fmt"
	"math"
	"math/bits"

	"arit/cli/parser"
	u "arit/modules/util"
)

type Prime struct{}

func (*Prime) Name() string {
	return "Prime"
}
func (*Prime) Keys() []string {
	return []string{"prime", "p"}
}
func (*Prime) Description() string {
	return "Module for primeness in arit"
}

func (p *Prime) Parse(cmd parser.Command) (any, error) {
	switch cmd.Func {
	case "is", "full":
		n, err := u.SingleInt64(cmd.Args)
		if err != nil {
			return nil, err
		}
		return p.isprime(n)

	case "mersenne":
		n, err := u.SingleInt64(cmd.Args)
		if err != nil {
			return nil, err
		}
		return p.isprime(n)
	case "factors", "fac":
		n, err := u.SingleInt64(cmd.Args)
		if err != nil {
			return nil, err
		}
		return p.factors2(n)
	default:
		n, err := u.SingleInt64(cmd.Args)
		if err != nil {
			return nil, err
		}
		return p.isprime(n)
	}
}

// Returns whether given number p is prime
// should be AKS at some point
func (mod *Prime) isprime(p int64) (bool, error) {
	if p < 1 {
		return false, errors.New("negative numbers cannot be prime")
	}

	if p < 100 {
		return smallPrime[uint64(p)], nil
	}

	psq := int64(math.Ceil(math.Sqrt(float64(p))))

	var i int64 = 2

	for i < psq {
		if (p % i) == 0 {
			return mod.isprime(i)
		}

		i++
	}

	return true, nil
}

// Returns whether given number p is a mersenne prime
func (mod *Prime) mersenne(p int64) (bool, error) {
	if p < 1 {
		return false, errors.New("negative numbers cannot be prime")
	}

	if p_ := (p + 1) & p; p_ != 0 {
		return false, nil
	}

	return mod.isprime(p)
}

// Returns list of non-distinct aliquot parts of p
// Which is fancy talk for "prime factors"
func (*Prime) factors2(p int64) ([]uint32, error) {
	if p > 1<<32 {
		return []uint32{}, fmt.Errorf("ask someone else")
	}

	if p < 1 {
		return []uint32{}, fmt.Errorf("cannot factorize negative numbers")
	}

	if p < (1<<32)-1 {
		return _read(factorFile, factorTableFile, p)
	}

	return []uint32{}, fmt.Errorf("now yet impl for p larger than 2^16")

	lg2 := bits.LeadingZeros64(uint64(p))

	// i thought this was very clever
	factors := make([]uint32, lg2)
	idx := 0

	rem := p

	for rem&1 == 0 {
		factors[idx] = 2
		idx++
		rem >>= 1
	}

	for rem%2 == 0 {
		factors[idx] = 2
		idx++
		rem >>= 1
	}

	return factors, nil
}
