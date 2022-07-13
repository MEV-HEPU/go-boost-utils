// Code generated by fastssz. DO NOT EDIT.
// Hash: 2c44f04492ec6dca9d1d1df6e7fcc2cb49c20f93d1b3470423ad3749b1f11b4a
package types

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SigningData object
func (s *SigningData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SigningData object to a target array
func (s *SigningData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Root'
	dst = append(dst, s.Root[:]...)

	// Field (1) 'Domain'
	dst = append(dst, s.Domain[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the SigningData object
func (s *SigningData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 64 {
		return ssz.ErrSize
	}

	// Field (0) 'Root'
	copy(s.Root[:], buf[0:32])

	// Field (1) 'Domain'
	copy(s.Domain[:], buf[32:64])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SigningData object
func (s *SigningData) SizeSSZ() (size int) {
	size = 64
	return
}

// HashTreeRoot ssz hashes the SigningData object
func (s *SigningData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SigningData object with a hasher
func (s *SigningData) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Root'
	hh.PutBytes(s.Root[:])

	// Field (1) 'Domain'
	hh.PutBytes(s.Domain[:])

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the forkData object
func (f *forkData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(f)
}

// MarshalSSZTo ssz marshals the forkData object to a target array
func (f *forkData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'CurrentVersion'
	dst = append(dst, f.CurrentVersion[:]...)

	// Field (1) 'GenesisValidatorsRoot'
	dst = append(dst, f.GenesisValidatorsRoot[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the forkData object
func (f *forkData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 36 {
		return ssz.ErrSize
	}

	// Field (0) 'CurrentVersion'
	copy(f.CurrentVersion[:], buf[0:4])

	// Field (1) 'GenesisValidatorsRoot'
	copy(f.GenesisValidatorsRoot[:], buf[4:36])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the forkData object
func (f *forkData) SizeSSZ() (size int) {
	size = 36
	return
}

// HashTreeRoot ssz hashes the forkData object
func (f *forkData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(f)
}

// HashTreeRootWith ssz hashes the forkData object with a hasher
func (f *forkData) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'CurrentVersion'
	hh.PutBytes(f.CurrentVersion[:])

	// Field (1) 'GenesisValidatorsRoot'
	hh.PutBytes(f.GenesisValidatorsRoot[:])

	hh.Merkleize(indx)
	return
}
