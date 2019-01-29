// Code generated by protoc-gen-go. DO NOT EDIT.
// source: idemix.proto

package idemix

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ECP is an elliptic curve point specified by its coordinates
// ECP corresponds to an element of the first group (G1)
type ECP struct {
	X                    []byte   `protobuf:"bytes,1,opt,name=X,proto3" json:"X,omitempty"`
	Y                    []byte   `protobuf:"bytes,2,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (m *ECP) Reset()         { *m = ECP{} }
func (m *ECP) String() string { return proto.CompactTextString(m) }
func (*ECP) ProtoMessage()    {}
func (*ECP) Descriptor() ([]byte, []int) {
	return fileDescriptor_idemix_07b4f691f366623b, []int{0}
}

func (m *ECP) GetX() []byte {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *ECP) GetY() []byte {
	if m != nil {
		return m.Y
	}
	return nil
}

// ECP2 is an elliptic curve point specified by its coordinates
// ECP2 corresponds to an element of the second group (G2)
type ECP2 struct {
	XA                   []byte   `protobuf:"bytes,1,opt,name=XA,proto3" json:"XA,omitempty"`
	XB                   []byte   `protobuf:"bytes,2,opt,name=XB,proto3" json:"XB,omitempty"`
	YA                   []byte   `protobuf:"bytes,3,opt,name=YA,proto3" json:"YA,omitempty"`
	YB                   []byte   `protobuf:"bytes,4,opt,name=YB,proto3" json:"YB,omitempty"`
}

func (m *ECP2) Reset()         { *m = ECP2{} }
func (m *ECP2) String() string { return proto.CompactTextString(m) }
func (*ECP2) ProtoMessage()    {}
func (*ECP2) Descriptor() ([]byte, []int) {
	return fileDescriptor_idemix_07b4f691f366623b, []int{1}
}

func (m *ECP2) GetXA() []byte {
	if m != nil {
		return m.XA
	}
	return nil
}

func (m *ECP2) GetXB() []byte {
	if m != nil {
		return m.XB
	}
	return nil
}

func (m *ECP2) GetYA() []byte {
	if m != nil {
		return m.YA
	}
	return nil
}

func (m *ECP2) GetYB() []byte {
	if m != nil {
		return m.YB
	}
	return nil
}

// IssuerPublicKey specifies an issuer public key that consists of
// AttributeNames - a list of the attribute names of a credential issued by the issuer
// HSk, HRand, HAttrs, W, BarG1, BarG2 - group elements corresponding to the signing key, randomness, and attributes
// ProofC, ProofS compose a zero-knowledge proof of knowledge of the secret key
// Hash is a hash of the public key appended to it
type IssuerPublicKey struct {
	AttributeNames       []string `protobuf:"bytes,1,rep,name=AttributeNames,proto3" json:"AttributeNames,omitempty"`
	HSk                  *ECP     `protobuf:"bytes,2,opt,name=HSk,proto3" json:"HSk,omitempty"`
	HRand                *ECP     `protobuf:"bytes,3,opt,name=HRand,proto3" json:"HRand,omitempty"`
	HAttrs               []*ECP   `protobuf:"bytes,4,rep,name=HAttrs,proto3" json:"HAttrs,omitempty"`
	W                    *ECP2    `protobuf:"bytes,5,opt,name=W,proto3" json:"W,omitempty"`
	BarG1                *ECP     `protobuf:"bytes,6,opt,name=BarG1,proto3" json:"BarG1,omitempty"`
	BarG2                *ECP     `protobuf:"bytes,7,opt,name=BarG2,proto3" json:"BarG2,omitempty"`
	ProofC               []byte   `protobuf:"bytes,8,opt,name=ProofC,proto3" json:"ProofC,omitempty"`
	ProofS               []byte   `protobuf:"bytes,9,opt,name=ProofS,proto3" json:"ProofS,omitempty"`
	Hash                 []byte   `protobuf:"bytes,10,opt,name=Hash,proto3" json:"Hash,omitempty"`
}

func (m *IssuerPublicKey) Reset()         { *m = IssuerPublicKey{} }
func (m *IssuerPublicKey) String() string { return proto.CompactTextString(m) }
func (*IssuerPublicKey) ProtoMessage()    {}
func (*IssuerPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_idemix_07b4f691f366623b, []int{2}
}

func (m *IssuerPublicKey) GetAttributeNames() []string {
	if m != nil {
		return m.AttributeNames
	}
	return nil
}

func (m *IssuerPublicKey) GetHSk() *ECP {
	if m != nil {
		return m.HSk
	}
	return nil
}

func (m *IssuerPublicKey) GetHRand() *ECP {
	if m != nil {
		return m.HRand
	}
	return nil
}

func (m *IssuerPublicKey) GetHAttrs() []*ECP {
	if m != nil {
		return m.HAttrs
	}
	return nil
}

func (m *IssuerPublicKey) GetW() *ECP2 {
	if m != nil {
		return m.W
	}
	return nil
}

func (m *IssuerPublicKey) GetBarG1() *ECP {
	if m != nil {
		return m.BarG1
	}
	return nil
}

func (m *IssuerPublicKey) GetBarG2() *ECP {
	if m != nil {
		return m.BarG2
	}
	return nil
}

func (m *IssuerPublicKey) GetProofC() []byte {
	if m != nil {
		return m.ProofC
	}
	return nil
}

func (m *IssuerPublicKey) GetProofS() []byte {
	if m != nil {
		return m.ProofS
	}
	return nil
}

func (m *IssuerPublicKey) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// IssuerKey specifies an issuer key pair that consists of
// ISk - the issuer secret key and
// IssuerPublicKey - the issuer public key
type IssuerKey struct {
	ISk                  []byte           `protobuf:"bytes,1,opt,name=ISk,proto3" json:"ISk,omitempty"`
	IPk                  *IssuerPublicKey `protobuf:"bytes,2,opt,name=IPk,proto3" json:"IPk,omitempty"`
}

func (m *IssuerKey) Reset()         { *m = IssuerKey{} }
func (m *IssuerKey) String() string { return proto.CompactTextString(m) }
func (*IssuerKey) ProtoMessage()    {}
func (*IssuerKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_idemix_07b4f691f366623b, []int{3}
}

func (m *IssuerKey) GetISk() []byte {
	if m != nil {
		return m.ISk
	}
	return nil
}

func (m *IssuerKey) GetIPk() *IssuerPublicKey {
	if m != nil {
		return m.IPk
	}
	return nil
}

// TracingKey specifies a group tracing key that consists of
// TK1,TK2 - random integers
type TracingKey struct {
	TK1                  []byte   `protobuf:"bytes,1,opt,name=TK1,proto3" json:"TK1,omitempty"`
	TK2                  []byte   `protobuf:"bytes,2,opt,name=TK2,proto3" json:"TK2,omitempty"`
}

func (m *TracingKey) Reset()         { *m = TracingKey{} }
func (m *TracingKey) String() string { return proto.CompactTextString(m) }
func (*TracingKey) ProtoMessage()    {}
func (*TracingKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_idemix_07b4f691f366623b, []int{4}
}

func (m *TracingKey) GetTK1() []byte {
	if m != nil {
		return m.TK1
	}
	return nil
}

func (m *TracingKey) GetTK2() []byte {
	if m != nil {
		return m.TK2
	}
	return nil
}

// GroupPublicKey specifies a group public key that consists of
// u,v,h - group elements corresponding to the group signing key
type GroupPublicKey struct {
	U                    *ECP     `protobuf:"bytes,1,opt,name=U,proto3" json:"U,omitempty"`
	V                    *ECP     `protobuf:"bytes,2,opt,name=V,proto3" json:"V,omitempty"`
	H                    *ECP     `protobuf:"bytes,3,opt,name=H,proto3" json:"H,omitempty"`
}

func (m *GroupPublicKey) Reset()         { *m = GroupPublicKey{} }
func (m *GroupPublicKey) String() string { return proto.CompactTextString(m) }
func (*GroupPublicKey) ProtoMessage()    {}
func (*GroupPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_idemix_07b4f691f366623b, []int{5}
}

func (m *GroupPublicKey) GetU() *ECP {
	if m != nil {
		return m.U
	}
	return nil
}

func (m *GroupPublicKey) GetV() *ECP {
	if m != nil {
		return m.V
	}
	return nil
}

func (m *GroupPublicKey) GetH() *ECP {
	if m != nil {
		return m.H
	}
	return nil
}

// GroupKey specifies a group key that consists of
// TK - group tracing key
// GPK - group public key
type GroupKey struct {
	TK                   *TracingKey     `protobuf:"bytes,1,opt,name=TK,proto3" json:"TK,omitempty"`
	GPK                  *GroupPublicKey `protobuf:"bytes,2,opt,name=GPK,proto3" json:"GPK,omitempty"`
}

func (m *GroupKey) Reset()         { *m = GroupKey{} }
func (m *GroupKey) String() string { return proto.CompactTextString(m) }
func (*GroupKey) ProtoMessage()    {}
func (*GroupKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_idemix_07b4f691f366623b, []int{6}
}

func (m *GroupKey) GetTK() *TracingKey {
	if m != nil {
		return m.TK
	}
	return nil
}

func (m *GroupKey) GetGPK() *GroupPublicKey {
	if m != nil {
		return m.GPK
	}
	return nil
}

// UserKey specifies a user's private key pair that consists of
// UK1 - a random integer
// UK2 - an element corresponding to the private key
type UserKey struct {
	UK1                  []byte   `protobuf:"bytes,1,opt,name=UK1,proto3" json:"UK1,omitempty"`
	UK2                  *ECP     `protobuf:"bytes,2,opt,name=UK2,proto3" json:"UK2,omitempty"`
}

func (m *UserKey) Reset()         { *m = UserKey{} }
func (m *UserKey) String() string { return proto.CompactTextString(m) }
func (*UserKey) ProtoMessage()    {}
func (*UserKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_idemix_07b4f691f366623b, []int{7}
}

func (m *UserKey) GetUK1() []byte {
	if m != nil {
		return m.UK1
	}
	return nil
}

func (m *UserKey) GetUK2() *ECP {
	if m != nil {
		return m.UK2
	}
	return nil
}

// Credential specifies a credential object that consists of
// A, B, E, S - signature value
// Attrs - attribute values
type Credential struct {
	A                    *ECP     `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	B                    *ECP     `protobuf:"bytes,2,opt,name=B,proto3" json:"B,omitempty"`
	E                    []byte   `protobuf:"bytes,3,opt,name=E,proto3" json:"E,omitempty"`
	S                    []byte   `protobuf:"bytes,4,opt,name=S,proto3" json:"S,omitempty"`
	Attrs                [][]byte `protobuf:"bytes,5,rep,name=Attrs,proto3" json:"Attrs,omitempty"`
}

func (m *Credential) Reset()         { *m = Credential{} }
func (m *Credential) String() string { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()    {}
func (*Credential) Descriptor() ([]byte, []int) {
	return fileDescriptor_idemix_07b4f691f366623b, []int{8}
}

func (m *Credential) GetA() *ECP {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *Credential) GetB() *ECP {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *Credential) GetE() []byte {
	if m != nil {
		return m.E
	}
	return nil
}

func (m *Credential) GetS() []byte {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *Credential) GetAttrs() [][]byte {
	if m != nil {
		return m.Attrs
	}
	return nil
}

// CredRequest specifies a credential request object that consists of
// Nym - a pseudonym, which is a commitment to the user secret
// IssuerNonce - a random nonce provided by the issuer
// ProofC, ProofS1, ProofS2 - a zero-knowledge proof of knowledge of the
// user secret inside Nym
type CredRequest struct {
	Nym                  *ECP     `protobuf:"bytes,1,opt,name=Nym,proto3" json:"Nym,omitempty"`
	IssuerNonce          []byte   `protobuf:"bytes,2,opt,name=IssuerNonce,proto3" json:"IssuerNonce,omitempty"`
	ProofC               []byte   `protobuf:"bytes,3,opt,name=ProofC,proto3" json:"ProofC,omitempty"`
	ProofS1              []byte   `protobuf:"bytes,4,opt,name=ProofS1,proto3" json:"ProofS1,omitempty"`
	ProofS2              []byte   `protobuf:"bytes,5,opt,name=ProofS2,proto3" json:"ProofS2,omitempty"`
}

func (m *CredRequest) Reset()         { *m = CredRequest{} }
func (m *CredRequest) String() string { return proto.CompactTextString(m) }
func (*CredRequest) ProtoMessage()    {}
func (*CredRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_idemix_07b4f691f366623b, []int{9}
}

func (m *CredRequest) GetNym() *ECP {
	if m != nil {
		return m.Nym
	}
	return nil
}

func (m *CredRequest) GetIssuerNonce() []byte {
	if m != nil {
		return m.IssuerNonce
	}
	return nil
}

func (m *CredRequest) GetProofC() []byte {
	if m != nil {
		return m.ProofC
	}
	return nil
}

func (m *CredRequest) GetProofS1() []byte {
	if m != nil {
		return m.ProofS1
	}
	return nil
}

func (m *CredRequest) GetProofS2() []byte {
	if m != nil {
		return m.ProofS2
	}
	return nil
}

// Signature specifies a signature object that consists of
// APrime, ABar, BPrime, Proof* - randomized credential signature values
// and a zero-knowledge proof of knowledge of a credential
// and the corresponding user secret together with the attribute values
// Nonce - a fresh nonce used for the signature
// Nym - a fresh pseudonym (a commitment to to the user secert)
// ProofSRNym - a zero-knowledge proof of knowledge of the
// user secret inside Nym
type Signature struct {
	APrime               *ECP     `protobuf:"bytes,1,opt,name=APrime,proto3" json:"APrime,omitempty"`
	ABar                 *ECP     `protobuf:"bytes,2,opt,name=ABar,proto3" json:"ABar,omitempty"`
	BPrime               *ECP     `protobuf:"bytes,3,opt,name=BPrime,proto3" json:"BPrime,omitempty"`
	T1                   *ECP     `protobuf:"bytes,4,opt,name=T1,proto3" json:"T1,omitempty"`
	T2                   *ECP     `protobuf:"bytes,5,opt,name=T2,proto3" json:"T2,omitempty"`
	T3                   *ECP     `protobuf:"bytes,6,opt,name=T3,proto3" json:"T3,omitempty"`
	ProofC               []byte   `protobuf:"bytes,7,opt,name=ProofC,proto3" json:"ProofC,omitempty"`
	ProofSSk             []byte   `protobuf:"bytes,8,opt,name=ProofSSk,proto3" json:"ProofSSk,omitempty"`
	ProofSE              []byte   `protobuf:"bytes,9,opt,name=ProofSE,proto3" json:"ProofSE,omitempty"`
	ProofSR2             []byte   `protobuf:"bytes,10,opt,name=ProofSR2,proto3" json:"ProofSR2,omitempty"`
	ProofSR3             []byte   `protobuf:"bytes,11,opt,name=ProofSR3,proto3" json:"ProofSR3,omitempty"`
	ProofSSPrime         []byte   `protobuf:"bytes,12,opt,name=ProofSSPrime,proto3" json:"ProofSSPrime,omitempty"`
	ProofSalpha          []byte   `protobuf:"bytes,13,opt,name=ProofSalpha,proto3" json:"ProofSalpha,omitempty"`
	ProofSbeta           []byte   `protobuf:"bytes,14,opt,name=ProofSbeta,proto3" json:"ProofSbeta,omitempty"`
	ProofSx              []byte   `protobuf:"bytes,15,opt,name=ProofSx,proto3" json:"ProofSx,omitempty"`
	ProofSdelta1         []byte   `protobuf:"bytes,16,opt,name=ProofSdelta1,proto3" json:"ProofSdelta1,omitempty"`
	ProofSdelta2         []byte   `protobuf:"bytes,17,opt,name=ProofSdelta2,proto3" json:"ProofSdelta2,omitempty"`
	ProofSAttrs          [][]byte `protobuf:"bytes,18,rep,name=ProofSAttrs,proto3" json:"ProofSAttrs,omitempty"`
	Nonce                []byte   `protobuf:"bytes,19,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	Nym                  *ECP     `protobuf:"bytes,20,opt,name=Nym,proto3" json:"Nym,omitempty"`
	ProofSRNym           []byte   `protobuf:"bytes,21,opt,name=ProofSRNym,proto3" json:"ProofSRNym,omitempty"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_idemix_07b4f691f366623b, []int{10}
}

func (m *Signature) GetAPrime() *ECP {
	if m != nil {
		return m.APrime
	}
	return nil
}

func (m *Signature) GetABar() *ECP {
	if m != nil {
		return m.ABar
	}
	return nil
}

func (m *Signature) GetBPrime() *ECP {
	if m != nil {
		return m.BPrime
	}
	return nil
}

func (m *Signature) GetT1() *ECP {
	if m != nil {
		return m.T1
	}
	return nil
}

func (m *Signature) GetT2() *ECP {
	if m != nil {
		return m.T2
	}
	return nil
}

func (m *Signature) GetT3() *ECP {
	if m != nil {
		return m.T3
	}
	return nil
}

func (m *Signature) GetProofC() []byte {
	if m != nil {
		return m.ProofC
	}
	return nil
}

func (m *Signature) GetProofSSk() []byte {
	if m != nil {
		return m.ProofSSk
	}
	return nil
}

func (m *Signature) GetProofSE() []byte {
	if m != nil {
		return m.ProofSE
	}
	return nil
}

func (m *Signature) GetProofSR2() []byte {
	if m != nil {
		return m.ProofSR2
	}
	return nil
}

func (m *Signature) GetProofSR3() []byte {
	if m != nil {
		return m.ProofSR3
	}
	return nil
}

func (m *Signature) GetProofSSPrime() []byte {
	if m != nil {
		return m.ProofSSPrime
	}
	return nil
}

func (m *Signature) GetProofSalpha() []byte {
	if m != nil {
		return m.ProofSalpha
	}
	return nil
}

func (m *Signature) GetProofSbeta() []byte {
	if m != nil {
		return m.ProofSbeta
	}
	return nil
}

func (m *Signature) GetProofSx() []byte {
	if m != nil {
		return m.ProofSx
	}
	return nil
}

func (m *Signature) GetProofSdelta1() []byte {
	if m != nil {
		return m.ProofSdelta1
	}
	return nil
}

func (m *Signature) GetProofSdelta2() []byte {
	if m != nil {
		return m.ProofSdelta2
	}
	return nil
}

func (m *Signature) GetProofSAttrs() [][]byte {
	if m != nil {
		return m.ProofSAttrs
	}
	return nil
}

func (m *Signature) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *Signature) GetNym() *ECP {
	if m != nil {
		return m.Nym
	}
	return nil
}

func (m *Signature) GetProofSRNym() []byte {
	if m != nil {
		return m.ProofSRNym
	}
	return nil
}

// NymSignature specifies a signature object that signs a message
// with respect to a pseudonym. It differs from the standard idemix.signature in the fact that
// the  standard signature object also proves that the pseudonym is based on a secret certified by
// a CA (issuer), whereas NymSignature only proves that the the owner of the pseudonym
// signed the message
type NymSignature struct {
	// ProofC is the Fiat-Shamir challenge of the ZKP
	ProofC []byte `protobuf:"bytes,1,opt,name=ProofC,proto3" json:"ProofC,omitempty"`
	// ProofSSK is the s-value proving knowledge of the user secret key
	ProofSSk []byte `protobuf:"bytes,2,opt,name=ProofSSk,proto3" json:"ProofSSk,omitempty"`
	// ProofSRNym is the s-value proving knowledge of the pseudonym secret
	ProofSRNym []byte `protobuf:"bytes,3,opt,name=ProofSRNym,proto3" json:"ProofSRNym,omitempty"`
	// Nonce is a fresh nonce used for the signature
	Nonce                []byte   `protobuf:"bytes,4,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
}

func (m *NymSignature) Reset()         { *m = NymSignature{} }
func (m *NymSignature) String() string { return proto.CompactTextString(m) }
func (*NymSignature) ProtoMessage()    {}
func (*NymSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_idemix_07b4f691f366623b, []int{11}
}

func (m *NymSignature) GetProofC() []byte {
	if m != nil {
		return m.ProofC
	}
	return nil
}

func (m *NymSignature) GetProofSSk() []byte {
	if m != nil {
		return m.ProofSSk
	}
	return nil
}

func (m *NymSignature) GetProofSRNym() []byte {
	if m != nil {
		return m.ProofSRNym
	}
	return nil
}

func (m *NymSignature) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

// K specifies a user's private key that is decrypted by CA
type K struct {
	Ax                   *ECP     `protobuf:"bytes,1,opt,name=Ax,proto3" json:"Ax,omitempty"`
}

func (m *K) Reset()         { *m = K{} }
func (m *K) String() string { return proto.CompactTextString(m) }
func (*K) ProtoMessage()    {}
func (*K) Descriptor() ([]byte, []int) {
	return fileDescriptor_idemix_07b4f691f366623b, []int{12}
}

func (m *K) GetAx() *ECP {
	if m != nil {
		return m.Ax
	}
	return nil
}

func init() {
	proto.RegisterType((*ECP)(nil), "ECP")
	proto.RegisterType((*ECP2)(nil), "ECP2")
	proto.RegisterType((*IssuerPublicKey)(nil), "IssuerPublicKey")
	proto.RegisterType((*IssuerKey)(nil), "IssuerKey")
	proto.RegisterType((*TracingKey)(nil), "TracingKey")
	proto.RegisterType((*GroupPublicKey)(nil), "GroupPublicKey")
	proto.RegisterType((*GroupKey)(nil), "GroupKey")
	proto.RegisterType((*UserKey)(nil), "UserKey")
	proto.RegisterType((*Credential)(nil), "Credential")
	proto.RegisterType((*CredRequest)(nil), "CredRequest")
	proto.RegisterType((*Signature)(nil), "Signature")
	proto.RegisterType((*NymSignature)(nil), "NymSignature")
	proto.RegisterType((*K)(nil), "K")
}

func init() { proto.RegisterFile("idemix.proto", fileDescriptor_idemix_07b4f691f366623b) }

var fileDescriptor_idemix_07b4f691f366623b = []byte{
	// 760 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xdb, 0x6e, 0xda, 0x4a,
	0x14, 0xd5, 0xd8, 0xe6, 0xb6, 0xe1, 0x90, 0x9c, 0x49, 0x4e, 0x34, 0x27, 0xad, 0x2a, 0x32, 0x0f,
	0x55, 0x9e, 0x50, 0x31, 0x5f, 0x60, 0x23, 0x1a, 0x52, 0x57, 0xc8, 0xb2, 0x4d, 0x02, 0x7d, 0x33,
	0xc1, 0x49, 0xac, 0x70, 0x49, 0x8d, 0x91, 0xc8, 0x77, 0xf4, 0xb5, 0x3f, 0xd7, 0x3f, 0xa9, 0xe6,
	0xe2, 0xcb, 0xb8, 0xea, 0x9b, 0xd7, 0x5a, 0x33, 0xfb, 0xba, 0x46, 0x86, 0x4e, 0xbc, 0x8a, 0x36,
	0xf1, 0xb1, 0xff, 0x9a, 0xec, 0xd2, 0x1d, 0xbd, 0x02, 0x7d, 0x3c, 0x72, 0x71, 0x07, 0xd0, 0x9c,
	0xa0, 0x1e, 0xba, 0xee, 0x78, 0x68, 0xce, 0xd0, 0x82, 0x68, 0x02, 0x2d, 0xe8, 0x67, 0x30, 0xc6,
	0x23, 0xd7, 0xc4, 0x5d, 0xd0, 0xe6, 0x96, 0x3c, 0xa4, 0xcd, 0x2d, 0x8e, 0x6d, 0x79, 0x4c, 0x9b,
	0xdb, 0x0c, 0x2f, 0x2c, 0xa2, 0x0b, 0xbc, 0xe0, 0xfa, 0xc2, 0x26, 0x86, 0xc4, 0x36, 0xfd, 0xa9,
	0xc1, 0xc9, 0xed, 0x7e, 0x7f, 0x88, 0x12, 0xf7, 0xb0, 0x5c, 0xc7, 0x0f, 0x4e, 0xf4, 0x86, 0x3f,
	0x42, 0xd7, 0x4a, 0xd3, 0x24, 0x5e, 0x1e, 0xd2, 0x68, 0x1a, 0x6e, 0xa2, 0x3d, 0x41, 0x3d, 0xfd,
	0xba, 0xe5, 0x55, 0x58, 0x7c, 0x01, 0xfa, 0xc4, 0x7f, 0xe1, 0xc9, 0xda, 0xa6, 0xd1, 0x1f, 0x8f,
	0x5c, 0x8f, 0x11, 0xf8, 0x12, 0x6a, 0x13, 0x2f, 0xdc, 0xae, 0x78, 0xda, 0x4c, 0x11, 0x14, 0x7e,
	0x0f, 0xf5, 0x09, 0x0b, 0xb3, 0x27, 0x46, 0x4f, 0xcf, 0x45, 0xc9, 0xe1, 0x33, 0x40, 0xf7, 0xa4,
	0xc6, 0x6f, 0xd5, 0x98, 0x60, 0x7a, 0xe8, 0x9e, 0x85, 0xb3, 0xc3, 0xe4, 0x66, 0x40, 0xea, 0xe5,
	0x70, 0x9c, 0xca, 0x34, 0x93, 0x34, 0xaa, 0x9a, 0x89, 0x2f, 0xa0, 0xee, 0x26, 0xbb, 0xdd, 0xe3,
	0x88, 0x34, 0x79, 0xbb, 0x12, 0xe5, 0xbc, 0x4f, 0x5a, 0x25, 0xde, 0xc7, 0x18, 0x8c, 0x49, 0xb8,
	0x7f, 0x26, 0xc0, 0x59, 0xfe, 0x4d, 0x2d, 0x68, 0x89, 0xe9, 0xb0, 0xb9, 0x9c, 0x82, 0x7e, 0xeb,
	0xbf, 0xc8, 0x61, 0xb3, 0x4f, 0x4c, 0x41, 0xbf, 0x75, 0xb3, 0x09, 0x9c, 0xf6, 0x2b, 0x83, 0xf4,
	0x98, 0x48, 0x3f, 0x01, 0x04, 0x49, 0xf8, 0x10, 0x6f, 0x9f, 0x64, 0x8c, 0xc0, 0x19, 0x64, 0x31,
	0x02, 0x67, 0x20, 0x18, 0x53, 0xae, 0x8c, 0x7d, 0xd2, 0xaf, 0xd0, 0xbd, 0x49, 0x76, 0x87, 0xd7,
	0x62, 0x23, 0x18, 0xd0, 0x8c, 0xdf, 0xc9, 0x5a, 0x44, 0x33, 0xc6, 0xdd, 0x29, 0xb3, 0x47, 0x77,
	0x8c, 0x9b, 0x28, 0x53, 0x47, 0x13, 0xfa, 0x05, 0x9a, 0x3c, 0x1a, 0x8b, 0xf3, 0x0e, 0xb4, 0xc0,
	0x91, 0x81, 0xda, 0xfd, 0xa2, 0x2c, 0x4f, 0x0b, 0x1c, 0x7c, 0x05, 0xfa, 0x8d, 0xeb, 0xc8, 0x90,
	0x27, 0x7d, 0xb5, 0x04, 0x8f, 0x69, 0x74, 0x08, 0x8d, 0xd9, 0x3e, 0x1f, 0xc6, 0xac, 0x68, 0x64,
	0xe6, 0x0c, 0x98, 0x1d, 0x66, 0xb2, 0x91, 0xdc, 0x0e, 0x33, 0xc7, 0xa4, 0x8f, 0x00, 0xa3, 0x24,
	0x5a, 0x45, 0xdb, 0x34, 0x0e, 0xd7, 0xac, 0x44, 0x4b, 0x6d, 0xc5, 0x62, 0x9c, 0xad, 0xb6, 0x62,
	0x33, 0xbb, 0x8f, 0xa5, 0x6f, 0xd1, 0x98, 0x21, 0x5f, 0xba, 0x16, 0xf9, 0xf8, 0x1c, 0x6a, 0xc2,
	0x43, 0xb5, 0x9e, 0x7e, 0xdd, 0xf1, 0x04, 0xa0, 0x3f, 0x10, 0xb4, 0x59, 0x22, 0x2f, 0xfa, 0x7e,
	0x88, 0xf6, 0x29, 0xab, 0x67, 0xfa, 0xb6, 0x51, 0x72, 0x31, 0x02, 0xf7, 0xa0, 0x2d, 0x16, 0x35,
	0xdd, 0x6d, 0x1f, 0x22, 0x39, 0xf8, 0x32, 0x55, 0x72, 0x8e, 0xae, 0x38, 0x87, 0x40, 0x43, 0x78,
	0x65, 0x20, 0x6b, 0xc9, 0x60, 0xa1, 0x98, 0xdc, 0xbe, 0xb9, 0x62, 0xd2, 0x5f, 0x06, 0xb4, 0xfc,
	0xf8, 0x69, 0x1b, 0xa6, 0x87, 0x24, 0x62, 0xf6, 0xb7, 0xdc, 0x24, 0xde, 0x44, 0x4a, 0x59, 0x92,
	0xc3, 0x04, 0x0c, 0xcb, 0x0e, 0x13, 0x65, 0x14, 0x9c, 0x61, 0xf7, 0x6c, 0x71, 0xaf, 0xbc, 0x5d,
	0xc9, 0xe1, 0x73, 0xd0, 0x02, 0x51, 0x52, 0xa6, 0x68, 0xc1, 0x80, 0xb3, 0xa6, 0x7c, 0x4d, 0x19,
	0x6b, 0x72, 0x76, 0xa8, 0x3c, 0x25, 0x2d, 0x18, 0x96, 0x3a, 0x6e, 0x28, 0x1d, 0x5f, 0x42, 0x53,
	0x34, 0xe2, 0xbf, 0xc8, 0x57, 0x94, 0xe3, 0xa2, 0xe7, 0xb1, 0x7c, 0x48, 0x19, 0x2c, 0x6e, 0x79,
	0xa6, 0x7c, 0x4d, 0x39, 0x2e, 0x69, 0x43, 0xd2, 0x56, 0xb4, 0x21, 0xa6, 0xd0, 0x91, 0xd1, 0x45,
	0xaf, 0x1d, 0xae, 0x2b, 0x1c, 0xdb, 0x9e, 0xc0, 0xe1, 0xfa, 0xf5, 0x39, 0x24, 0xff, 0x88, 0xed,
	0x95, 0x28, 0xfc, 0x01, 0x40, 0xc0, 0x65, 0x94, 0x86, 0xa4, 0xcb, 0x0f, 0x94, 0x98, 0xa2, 0xee,
	0x23, 0x39, 0x29, 0xd7, 0x7d, 0x2c, 0xf2, 0xaf, 0xa2, 0x75, 0x1a, 0x0e, 0xc8, 0x69, 0x39, 0xbf,
	0xe0, 0x2a, 0x67, 0x4c, 0xf2, 0xef, 0x1f, 0x67, 0xcc, 0xa2, 0x46, 0xe1, 0x52, 0xcc, 0x5d, 0x5a,
	0xa6, 0x98, 0x83, 0x85, 0xfb, 0xce, 0xf8, 0xf5, 0x5a, 0xe6, 0x3b, 0xee, 0xd8, 0xf3, 0xaa, 0x63,
	0xf3, 0x8e, 0x3c, 0x26, 0xff, 0x57, 0xee, 0x88, 0x31, 0xf4, 0x08, 0x9d, 0xe9, 0xdb, 0xa6, 0x70,
	0x59, 0xb1, 0x4d, 0xf4, 0xd7, 0x6d, 0x6a, 0x95, 0x6d, 0xaa, 0x39, 0xf4, 0x6a, 0x8e, 0xa2, 0x62,
	0xa3, 0x54, 0x31, 0xfd, 0x1f, 0x90, 0xc3, 0x2c, 0x65, 0x1d, 0x15, 0x43, 0x6b, 0xd6, 0xd1, 0x6e,
	0x7e, 0xab, 0x8b, 0x9f, 0xda, 0xb2, 0xce, 0xff, 0x6a, 0xc3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x67, 0xb5, 0x2b, 0x60, 0xe5, 0x06, 0x00, 0x00,
}
