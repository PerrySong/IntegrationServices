// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package userpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Repo struct {
	GitTagsUrl           string               `protobuf:"bytes,1,opt,name=git_tags_url,json=gitTagsUrl,proto3" json:"git_tags_url,omitempty"`
	Description          string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Private              bool                 `protobuf:"varint,3,opt,name=private,proto3" json:"private,omitempty"`
	LanguagesUrl         string               `protobuf:"bytes,4,opt,name=languages_url,json=languagesUrl,proto3" json:"languages_url,omitempty"`
	StargazersUrl        string               `protobuf:"bytes,5,opt,name=stargazers_url,json=stargazersUrl,proto3" json:"stargazers_url,omitempty"`
	CommitsUrl           string               `protobuf:"bytes,6,opt,name=commits_url,json=commitsUrl,proto3" json:"commits_url,omitempty"`
	RepoCreatedAt        *timestamp.Timestamp `protobuf:"bytes,7,opt,name=repo_created_at,json=repoCreatedAt,proto3" json:"repo_created_at,omitempty"`
	RepoUpdatedAt        *timestamp.Timestamp `protobuf:"bytes,8,opt,name=repo_updated_at,json=repoUpdatedAt,proto3" json:"repo_updated_at,omitempty"`
	HomePage             string               `protobuf:"bytes,9,opt,name=home_page,json=homePage,proto3" json:"home_page,omitempty"`
	StargazersCount      int32                `protobuf:"varint,10,opt,name=stargazers_count,json=stargazersCount,proto3" json:"stargazers_count,omitempty"`
	LabelsUrl            string               `protobuf:"bytes,11,opt,name=labels_url,json=labelsUrl,proto3" json:"labels_url,omitempty"`
	Language             string               `protobuf:"bytes,12,opt,name=language,proto3" json:"language,omitempty"`
	Watchers             int32                `protobuf:"varint,13,opt,name=watchers,proto3" json:"watchers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Repo) Reset()         { *m = Repo{} }
func (m *Repo) String() string { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()    {}
func (*Repo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *Repo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repo.Unmarshal(m, b)
}
func (m *Repo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repo.Marshal(b, m, deterministic)
}
func (m *Repo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repo.Merge(m, src)
}
func (m *Repo) XXX_Size() int {
	return xxx_messageInfo_Repo.Size(m)
}
func (m *Repo) XXX_DiscardUnknown() {
	xxx_messageInfo_Repo.DiscardUnknown(m)
}

var xxx_messageInfo_Repo proto.InternalMessageInfo

func (m *Repo) GetGitTagsUrl() string {
	if m != nil {
		return m.GitTagsUrl
	}
	return ""
}

func (m *Repo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Repo) GetPrivate() bool {
	if m != nil {
		return m.Private
	}
	return false
}

func (m *Repo) GetLanguagesUrl() string {
	if m != nil {
		return m.LanguagesUrl
	}
	return ""
}

func (m *Repo) GetStargazersUrl() string {
	if m != nil {
		return m.StargazersUrl
	}
	return ""
}

func (m *Repo) GetCommitsUrl() string {
	if m != nil {
		return m.CommitsUrl
	}
	return ""
}

func (m *Repo) GetRepoCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.RepoCreatedAt
	}
	return nil
}

func (m *Repo) GetRepoUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.RepoUpdatedAt
	}
	return nil
}

func (m *Repo) GetHomePage() string {
	if m != nil {
		return m.HomePage
	}
	return ""
}

func (m *Repo) GetStargazersCount() int32 {
	if m != nil {
		return m.StargazersCount
	}
	return 0
}

func (m *Repo) GetLabelsUrl() string {
	if m != nil {
		return m.LabelsUrl
	}
	return ""
}

func (m *Repo) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Repo) GetWatchers() int32 {
	if m != nil {
		return m.Watchers
	}
	return 0
}

type GithubRepoListResponse struct {
	Repos                []*Repo  `protobuf:"bytes,1,rep,name=repos,proto3" json:"repos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GithubRepoListResponse) Reset()         { *m = GithubRepoListResponse{} }
func (m *GithubRepoListResponse) String() string { return proto.CompactTextString(m) }
func (*GithubRepoListResponse) ProtoMessage()    {}
func (*GithubRepoListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *GithubRepoListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GithubRepoListResponse.Unmarshal(m, b)
}
func (m *GithubRepoListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GithubRepoListResponse.Marshal(b, m, deterministic)
}
func (m *GithubRepoListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GithubRepoListResponse.Merge(m, src)
}
func (m *GithubRepoListResponse) XXX_Size() int {
	return xxx_messageInfo_GithubRepoListResponse.Size(m)
}
func (m *GithubRepoListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GithubRepoListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GithubRepoListResponse proto.InternalMessageInfo

func (m *GithubRepoListResponse) GetRepos() []*Repo {
	if m != nil {
		return m.Repos
	}
	return nil
}

type GithubRepoListRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GithubRepoListRequest) Reset()         { *m = GithubRepoListRequest{} }
func (m *GithubRepoListRequest) String() string { return proto.CompactTextString(m) }
func (*GithubRepoListRequest) ProtoMessage()    {}
func (*GithubRepoListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *GithubRepoListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GithubRepoListRequest.Unmarshal(m, b)
}
func (m *GithubRepoListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GithubRepoListRequest.Marshal(b, m, deterministic)
}
func (m *GithubRepoListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GithubRepoListRequest.Merge(m, src)
}
func (m *GithubRepoListRequest) XXX_Size() int {
	return xxx_messageInfo_GithubRepoListRequest.Size(m)
}
func (m *GithubRepoListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GithubRepoListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GithubRepoListRequest proto.InternalMessageInfo

func (m *GithubRepoListRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserGithubBasicRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGithubBasicRequest) Reset()         { *m = UserGithubBasicRequest{} }
func (m *UserGithubBasicRequest) String() string { return proto.CompactTextString(m) }
func (*UserGithubBasicRequest) ProtoMessage()    {}
func (*UserGithubBasicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserGithubBasicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGithubBasicRequest.Unmarshal(m, b)
}
func (m *UserGithubBasicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGithubBasicRequest.Marshal(b, m, deterministic)
}
func (m *UserGithubBasicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGithubBasicRequest.Merge(m, src)
}
func (m *UserGithubBasicRequest) XXX_Size() int {
	return xxx_messageInfo_UserGithubBasicRequest.Size(m)
}
func (m *UserGithubBasicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGithubBasicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserGithubBasicRequest proto.InternalMessageInfo

func (m *UserGithubBasicRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserGithubBasicResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Bio                  string   `protobuf:"bytes,5,opt,name=bio,proto3" json:"bio,omitempty"`
	Company              string   `protobuf:"bytes,6,opt,name=company,proto3" json:"company,omitempty"`
	ReposUrl             string   `protobuf:"bytes,7,opt,name=repos_url,json=reposUrl,proto3" json:"repos_url,omitempty"`
	Email                string   `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	Location             string   `protobuf:"bytes,9,opt,name=location,proto3" json:"location,omitempty"`
	PublicRepos          int32    `protobuf:"varint,10,opt,name=public_repos,json=publicRepos,proto3" json:"public_repos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGithubBasicResponse) Reset()         { *m = UserGithubBasicResponse{} }
func (m *UserGithubBasicResponse) String() string { return proto.CompactTextString(m) }
func (*UserGithubBasicResponse) ProtoMessage()    {}
func (*UserGithubBasicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserGithubBasicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGithubBasicResponse.Unmarshal(m, b)
}
func (m *UserGithubBasicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGithubBasicResponse.Marshal(b, m, deterministic)
}
func (m *UserGithubBasicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGithubBasicResponse.Merge(m, src)
}
func (m *UserGithubBasicResponse) XXX_Size() int {
	return xxx_messageInfo_UserGithubBasicResponse.Size(m)
}
func (m *UserGithubBasicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGithubBasicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserGithubBasicResponse proto.InternalMessageInfo

func (m *UserGithubBasicResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserGithubBasicResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserGithubBasicResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *UserGithubBasicResponse) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *UserGithubBasicResponse) GetBio() string {
	if m != nil {
		return m.Bio
	}
	return ""
}

func (m *UserGithubBasicResponse) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *UserGithubBasicResponse) GetReposUrl() string {
	if m != nil {
		return m.ReposUrl
	}
	return ""
}

func (m *UserGithubBasicResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserGithubBasicResponse) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *UserGithubBasicResponse) GetPublicRepos() int32 {
	if m != nil {
		return m.PublicRepos
	}
	return 0
}

type UserHasTokenRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserHasTokenRequest) Reset()         { *m = UserHasTokenRequest{} }
func (m *UserHasTokenRequest) String() string { return proto.CompactTextString(m) }
func (*UserHasTokenRequest) ProtoMessage()    {}
func (*UserHasTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *UserHasTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserHasTokenRequest.Unmarshal(m, b)
}
func (m *UserHasTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserHasTokenRequest.Marshal(b, m, deterministic)
}
func (m *UserHasTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserHasTokenRequest.Merge(m, src)
}
func (m *UserHasTokenRequest) XXX_Size() int {
	return xxx_messageInfo_UserHasTokenRequest.Size(m)
}
func (m *UserHasTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserHasTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserHasTokenRequest proto.InternalMessageInfo

func (m *UserHasTokenRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserHasTokenResponse struct {
	HasToken             bool     `protobuf:"varint,1,opt,name=has_token,json=hasToken,proto3" json:"has_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserHasTokenResponse) Reset()         { *m = UserHasTokenResponse{} }
func (m *UserHasTokenResponse) String() string { return proto.CompactTextString(m) }
func (*UserHasTokenResponse) ProtoMessage()    {}
func (*UserHasTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *UserHasTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserHasTokenResponse.Unmarshal(m, b)
}
func (m *UserHasTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserHasTokenResponse.Marshal(b, m, deterministic)
}
func (m *UserHasTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserHasTokenResponse.Merge(m, src)
}
func (m *UserHasTokenResponse) XXX_Size() int {
	return xxx_messageInfo_UserHasTokenResponse.Size(m)
}
func (m *UserHasTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserHasTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserHasTokenResponse proto.InternalMessageInfo

func (m *UserHasTokenResponse) GetHasToken() bool {
	if m != nil {
		return m.HasToken
	}
	return false
}

func init() {
	proto.RegisterType((*Repo)(nil), "user.Repo")
	proto.RegisterType((*GithubRepoListResponse)(nil), "user.GithubRepoListResponse")
	proto.RegisterType((*GithubRepoListRequest)(nil), "user.GithubRepoListRequest")
	proto.RegisterType((*UserGithubBasicRequest)(nil), "user.UserGithubBasicRequest")
	proto.RegisterType((*UserGithubBasicResponse)(nil), "user.UserGithubBasicResponse")
	proto.RegisterType((*UserHasTokenRequest)(nil), "user.UserHasTokenRequest")
	proto.RegisterType((*UserHasTokenResponse)(nil), "user.UserHasTokenResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xfd, 0x12, 0x27, 0xad, 0x33, 0x4e, 0xda, 0x6a, 0xbf, 0x52, 0x8c, 0xdb, 0xaa, 0x26, 0xa8,
	0x22, 0x5c, 0x52, 0xa9, 0xbd, 0x71, 0x6b, 0x7b, 0x28, 0x48, 0x80, 0x90, 0x69, 0x2f, 0x5c, 0xac,
	0xb5, 0xb3, 0x75, 0x56, 0xd8, 0x5e, 0xb3, 0xbb, 0x2e, 0x82, 0xdf, 0xc1, 0x5f, 0x45, 0xe2, 0x88,
	0x66, 0xd7, 0x4e, 0x42, 0x9b, 0x4a, 0xdc, 0x3c, 0x6f, 0xdf, 0xbc, 0xdd, 0x79, 0xf3, 0x0c, 0x50,
	0x2b, 0x26, 0xa7, 0x95, 0x14, 0x5a, 0x90, 0x1e, 0x7e, 0x07, 0x47, 0x99, 0x10, 0x59, 0xce, 0x4e,
	0x0c, 0x96, 0xd4, 0xb7, 0x27, 0x9a, 0x17, 0x4c, 0x69, 0x5a, 0x54, 0x96, 0x36, 0xfe, 0xed, 0x40,
	0x2f, 0x62, 0x95, 0x20, 0x21, 0x0c, 0x33, 0xae, 0x63, 0x4d, 0x33, 0x15, 0xd7, 0x32, 0xf7, 0x3b,
	0x61, 0x67, 0x32, 0x88, 0x20, 0xe3, 0xfa, 0x9a, 0x66, 0xea, 0x46, 0xe6, 0x24, 0x04, 0x6f, 0xc6,
	0x54, 0x2a, 0x79, 0xa5, 0xb9, 0x28, 0xfd, 0xae, 0x21, 0xac, 0x42, 0xc4, 0x87, 0xcd, 0x4a, 0xf2,
	0x3b, 0xaa, 0x99, 0xef, 0x84, 0x9d, 0x89, 0x1b, 0xb5, 0x25, 0x79, 0x01, 0xa3, 0x9c, 0x96, 0x59,
	0x4d, 0x33, 0x66, 0xe5, 0x7b, 0xa6, 0x7b, 0xb8, 0x00, 0xf1, 0x82, 0x63, 0xd8, 0x52, 0x9a, 0xca,
	0x8c, 0xfe, 0x60, 0xd2, 0xb2, 0xfa, 0x86, 0x35, 0x5a, 0xa2, 0x48, 0x3b, 0x02, 0x2f, 0x15, 0x45,
	0xc1, 0xb5, 0xe5, 0x6c, 0xd8, 0x87, 0x36, 0x10, 0x12, 0x2e, 0x60, 0x5b, 0xb2, 0x4a, 0xc4, 0xa9,
	0x64, 0x54, 0xb3, 0x59, 0x4c, 0xb5, 0xbf, 0x19, 0x76, 0x26, 0xde, 0x69, 0x30, 0xb5, 0x76, 0x4c,
	0x5b, 0x3b, 0xa6, 0xd7, 0xad, 0x1d, 0xd1, 0x08, 0x5b, 0x2e, 0x6d, 0xc7, 0xb9, 0x5e, 0x68, 0xd4,
	0xd5, 0xac, 0xd5, 0x70, 0xff, 0x4d, 0xe3, 0xc6, 0x76, 0x9c, 0x6b, 0xb2, 0x0f, 0x83, 0xb9, 0x28,
	0x58, 0x5c, 0xd1, 0x8c, 0xf9, 0x03, 0xf3, 0x4c, 0x17, 0x81, 0x8f, 0x34, 0x63, 0xe4, 0x15, 0xec,
	0xac, 0x0c, 0x9b, 0x8a, 0xba, 0xd4, 0x3e, 0x84, 0x9d, 0x49, 0x3f, 0xda, 0x5e, 0xe2, 0x97, 0x08,
	0x93, 0x43, 0x80, 0x9c, 0x26, 0x2c, 0xb7, 0xf3, 0x7a, 0x46, 0x68, 0x60, 0x11, 0x1c, 0x37, 0x00,
	0xb7, 0xb5, 0xd1, 0x1f, 0xda, 0x5b, 0xda, 0x1a, 0xcf, 0xbe, 0x51, 0x9d, 0xce, 0x99, 0x54, 0xfe,
	0xc8, 0xa8, 0x2f, 0xea, 0xf1, 0x6b, 0xd8, 0xbb, 0xe2, 0x7a, 0x5e, 0x27, 0xb8, 0xff, 0x77, 0x5c,
	0xe9, 0x88, 0xa9, 0x4a, 0x94, 0x8a, 0x91, 0x10, 0xfa, 0x38, 0x89, 0xf2, 0x3b, 0xa1, 0x33, 0xf1,
	0x4e, 0x61, 0x6a, 0x72, 0x85, 0xb4, 0xc8, 0x1e, 0x8c, 0x5f, 0xc2, 0x93, 0xfb, 0xbd, 0x5f, 0x6b,
	0xa6, 0x34, 0xd9, 0x82, 0x2e, 0x9f, 0x99, 0xf0, 0xf4, 0xa2, 0x2e, 0x9f, 0x8d, 0x27, 0xb0, 0x77,
	0xa3, 0x98, 0xb4, 0xe4, 0x0b, 0xaa, 0x78, 0xfa, 0x18, 0xf3, 0x67, 0x17, 0x9e, 0x3e, 0xa0, 0x36,
	0x0f, 0x5a, 0x72, 0x1d, 0xe4, 0xe2, 0x58, 0xf8, 0xa4, 0x92, 0x16, 0xac, 0xc9, 0xe1, 0xa2, 0x26,
	0x3b, 0xe0, 0xa0, 0x4d, 0x8e, 0x81, 0xf1, 0x13, 0xfd, 0xa3, 0x77, 0x54, 0x53, 0xb9, 0x92, 0xbc,
	0x81, 0x45, 0xd0, 0xbf, 0x1d, 0x70, 0x12, 0x2e, 0x9a, 0xac, 0xe1, 0x27, 0xe6, 0x38, 0x15, 0x45,
	0x45, 0xcb, 0xef, 0x4d, 0xba, 0xda, 0x12, 0x57, 0x6a, 0x0c, 0x30, 0x4a, 0x9b, 0xf6, 0x66, 0x03,
	0xa0, 0xd0, 0x2e, 0xf4, 0x59, 0x41, 0x79, 0x6e, 0x92, 0x32, 0x88, 0x6c, 0x61, 0xd6, 0x23, 0x52,
	0x6a, 0xfe, 0x99, 0x26, 0x04, 0x6d, 0x4d, 0x9e, 0xc3, 0xb0, 0xaa, 0x93, 0x9c, 0xa7, 0xb1, 0xf5,
	0xdb, 0x06, 0xc0, 0xb3, 0x58, 0x64, 0x9c, 0x3e, 0x86, 0xff, 0xd1, 0x95, 0x37, 0x54, 0x5d, 0x8b,
	0x2f, 0xac, 0x7c, 0xe8, 0x9e, 0x71, 0x64, 0x7c, 0x06, 0xbb, 0x7f, 0xd3, 0x1a, 0xe7, 0x30, 0x83,
	0x54, 0xc5, 0x1a, 0x41, 0x43, 0x77, 0x23, 0x77, 0xde, 0x90, 0x4e, 0x7f, 0x75, 0xc0, 0xc3, 0xae,
	0x4f, 0x4c, 0xde, 0xf1, 0x94, 0x91, 0x0f, 0x30, 0xba, 0x62, 0xda, 0x2e, 0xe0, 0x6d, 0x79, 0x2b,
	0xc8, 0x81, 0xdd, 0xfc, 0xfa, 0x0d, 0x06, 0x87, 0x8f, 0x9c, 0xda, 0xab, 0xc7, 0xff, 0x91, 0xf7,
	0xb0, 0xb5, 0xd0, 0x33, 0xd3, 0x90, 0x7d, 0xdb, 0xb2, 0x36, 0x3b, 0xc1, 0xc1, 0xfa, 0xc3, 0x85,
	0xdc, 0x15, 0x0c, 0x57, 0x67, 0x24, 0xcf, 0x96, 0xf7, 0xdf, 0xb3, 0x27, 0x08, 0xd6, 0x1d, 0xb5,
	0x42, 0x17, 0xee, 0xe7, 0x0d, 0x3c, 0xae, 0x92, 0x64, 0xc3, 0xfc, 0xc5, 0x67, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x0e, 0x9c, 0xdf, 0xdf, 0x3a, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetGithubInfo(ctx context.Context, in *UserGithubBasicRequest, opts ...grpc.CallOption) (*UserGithubBasicResponse, error)
	GetGithubRepos(ctx context.Context, in *GithubRepoListRequest, opts ...grpc.CallOption) (*GithubRepoListResponse, error)
	UserHasToken(ctx context.Context, in *UserHasTokenRequest, opts ...grpc.CallOption) (*UserHasTokenResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetGithubInfo(ctx context.Context, in *UserGithubBasicRequest, opts ...grpc.CallOption) (*UserGithubBasicResponse, error) {
	out := new(UserGithubBasicResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetGithubInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetGithubRepos(ctx context.Context, in *GithubRepoListRequest, opts ...grpc.CallOption) (*GithubRepoListResponse, error) {
	out := new(GithubRepoListResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetGithubRepos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserHasToken(ctx context.Context, in *UserHasTokenRequest, opts ...grpc.CallOption) (*UserHasTokenResponse, error) {
	out := new(UserHasTokenResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UserHasToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetGithubInfo(context.Context, *UserGithubBasicRequest) (*UserGithubBasicResponse, error)
	GetGithubRepos(context.Context, *GithubRepoListRequest) (*GithubRepoListResponse, error)
	UserHasToken(context.Context, *UserHasTokenRequest) (*UserHasTokenResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetGithubInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGithubBasicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetGithubInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetGithubInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetGithubInfo(ctx, req.(*UserGithubBasicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetGithubRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GithubRepoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetGithubRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetGithubRepos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetGithubRepos(ctx, req.(*GithubRepoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserHasToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserHasTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserHasToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UserHasToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserHasToken(ctx, req.(*UserHasTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGithubInfo",
			Handler:    _UserService_GetGithubInfo_Handler,
		},
		{
			MethodName: "GetGithubRepos",
			Handler:    _UserService_GetGithubRepos_Handler,
		},
		{
			MethodName: "UserHasToken",
			Handler:    _UserService_UserHasToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
