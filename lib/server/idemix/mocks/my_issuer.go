// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import (
	db "github.com/hyperledger/fabric-ca/lib/server/db"
	idemix "github.com/hyperledger/fabric-ca/lib/server/idemix"

	mock "github.com/stretchr/testify/mock"
)

// MyIssuer is an autogenerated mock type for the MyIssuer type
type MyIssuer struct {
	mock.Mock
}

// Config provides a mock function with given fields:
func (_m *MyIssuer) Config() *idemix.Config {
	ret := _m.Called()

	var r0 *idemix.Config
	if rf, ok := ret.Get(0).(func() *idemix.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*idemix.Config)
		}
	}

	return r0
}

// CredDBAccessor provides a mock function with given fields:
func (_m *MyIssuer) CredDBAccessor() idemix.CredDBAccessor {
	ret := _m.Called()

	var r0 idemix.CredDBAccessor
	if rf, ok := ret.Get(0).(func() idemix.CredDBAccessor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(idemix.CredDBAccessor)
		}
	}

	return r0
}

// DB provides a mock function with given fields:
func (_m *MyIssuer) DB() db.FabricCADB {
	ret := _m.Called()

	var r0 db.FabricCADB
	if rf, ok := ret.Get(0).(func() db.FabricCADB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.FabricCADB)
		}
	}

	return r0
}

// HomeDir provides a mock function with given fields:
func (_m *MyIssuer) HomeDir() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IdemixLib provides a mock function with given fields:
func (_m *MyIssuer) IdemixLib() idemix.Lib {
	ret := _m.Called()

	var r0 idemix.Lib
	if rf, ok := ret.Get(0).(func() idemix.Lib); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(idemix.Lib)
		}
	}

	return r0
}

// IssuerCredential provides a mock function with given fields:
func (_m *MyIssuer) IssuerCredential() idemix.IssuerCredential {
	ret := _m.Called()

	var r0 idemix.IssuerCredential
	if rf, ok := ret.Get(0).(func() idemix.IssuerCredential); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(idemix.IssuerCredential)
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *MyIssuer) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NonceManager provides a mock function with given fields:
func (_m *MyIssuer) NonceManager() idemix.NonceManager {
	ret := _m.Called()

	var r0 idemix.NonceManager
	if rf, ok := ret.Get(0).(func() idemix.NonceManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(idemix.NonceManager)
		}
	}

	return r0
}

// RevocationAuthority provides a mock function with given fields:
func (_m *MyIssuer) RevocationAuthority() idemix.RevocationAuthority {
	ret := _m.Called()

	var r0 idemix.RevocationAuthority
	if rf, ok := ret.Get(0).(func() idemix.RevocationAuthority); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(idemix.RevocationAuthority)
		}
	}

	return r0
}
