// Code generated by 'ccgo netdb/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o netdb/netdb_netbsd_amd64.go -pkgname netdb', DO NOT EDIT.

package netdb

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

const (
	AI_ADDRCONFIG                 = 0x00000400
	AI_CANONNAME                  = 0x00000002
	AI_MASK                       = 3087
	AI_NUMERICHOST                = 0x00000004
	AI_NUMERICSERV                = 0x00000008
	AI_PASSIVE                    = 0x00000001
	AI_SRV                        = 0x00000800
	EAI_ADDRFAMILY                = 1
	EAI_AGAIN                     = 2
	EAI_BADFLAGS                  = 3
	EAI_BADHINTS                  = 12
	EAI_FAIL                      = 4
	EAI_FAMILY                    = 5
	EAI_MAX                       = 15
	EAI_MEMORY                    = 6
	EAI_NODATA                    = 7
	EAI_NONAME                    = 8
	EAI_OVERFLOW                  = 14
	EAI_PROTOCOL                  = 13
	EAI_SERVICE                   = 9
	EAI_SOCKTYPE                  = 10
	EAI_SYSTEM                    = 11
	HOST_NOT_FOUND                = 1
	INT16_MAX                     = 32767
	INT16_MIN                     = -32768
	INT32_MAX                     = 2147483647
	INT32_MIN                     = -2147483648
	INT64_MAX                     = 9223372036854775807
	INT64_MIN                     = -9223372036854775808
	INT8_MAX                      = 127
	INT8_MIN                      = -128
	INTMAX_MAX                    = 9223372036854775807
	INTMAX_MIN                    = -9223372036854775808
	INTPTR_MAX                    = 9223372036854775807
	INTPTR_MIN                    = -9223372036854775808
	INT_FAST16_MAX                = 2147483647
	INT_FAST16_MIN                = -2147483648
	INT_FAST32_MAX                = 2147483647
	INT_FAST32_MIN                = -2147483648
	INT_FAST64_MAX                = 9223372036854775807
	INT_FAST64_MIN                = -9223372036854775808
	INT_FAST8_MAX                 = 2147483647
	INT_FAST8_MIN                 = -2147483648
	INT_LEAST16_MAX               = 32767
	INT_LEAST16_MIN               = -32768
	INT_LEAST32_MAX               = 2147483647
	INT_LEAST32_MIN               = -2147483648
	INT_LEAST64_MAX               = 9223372036854775807
	INT_LEAST64_MIN               = -9223372036854775808
	INT_LEAST8_MAX                = 127
	INT_LEAST8_MIN                = -128
	NETDB_INTERNAL                = -1
	NETDB_SUCCESS                 = 0
	NI_DGRAM                      = 0x00000010
	NI_MAXHOST                    = 1025
	NI_MAXSERV                    = 32
	NI_NAMEREQD                   = 0x00000004
	NI_NOFQDN                     = 0x00000001
	NI_NUMERICHOST                = 0x00000002
	NI_NUMERICSCOPE               = 0x00000040
	NI_NUMERICSERV                = 0x00000008
	NI_WITHSCOPEID                = 0x00000020
	NO_ADDRESS                    = 4
	NO_DATA                       = 4
	NO_RECOVERY                   = 3
	PRIX16                        = "X"
	PRIX32                        = "X"
	PRIX64                        = "lX"
	PRIX8                         = "X"
	PRIXFAST16                    = "X"
	PRIXFAST32                    = "X"
	PRIXFAST64                    = "lX"
	PRIXFAST8                     = "X"
	PRIXLEAST16                   = "X"
	PRIXLEAST32                   = "X"
	PRIXLEAST64                   = "lX"
	PRIXLEAST8                    = "X"
	PRIXMAX                       = "lX"
	PRIXPTR                       = "lX"
	PRId16                        = "d"
	PRId32                        = "d"
	PRId64                        = "ld"
	PRId8                         = "d"
	PRIdFAST16                    = "d"
	PRIdFAST32                    = "d"
	PRIdFAST64                    = "ld"
	PRIdFAST8                     = "d"
	PRIdLEAST16                   = "d"
	PRIdLEAST32                   = "d"
	PRIdLEAST64                   = "ld"
	PRIdLEAST8                    = "d"
	PRIdMAX                       = "ld"
	PRIdPTR                       = "ld"
	PRIi16                        = "i"
	PRIi32                        = "i"
	PRIi64                        = "li"
	PRIi8                         = "i"
	PRIiFAST16                    = "i"
	PRIiFAST32                    = "i"
	PRIiFAST64                    = "li"
	PRIiFAST8                     = "i"
	PRIiLEAST16                   = "i"
	PRIiLEAST32                   = "i"
	PRIiLEAST64                   = "li"
	PRIiLEAST8                    = "i"
	PRIiMAX                       = "li"
	PRIiPTR                       = "li"
	PRIo16                        = "o"
	PRIo32                        = "o"
	PRIo64                        = "lo"
	PRIo8                         = "o"
	PRIoFAST16                    = "o"
	PRIoFAST32                    = "o"
	PRIoFAST64                    = "lo"
	PRIoFAST8                     = "o"
	PRIoLEAST16                   = "o"
	PRIoLEAST32                   = "o"
	PRIoLEAST64                   = "lo"
	PRIoLEAST8                    = "o"
	PRIoMAX                       = "lo"
	PRIoPTR                       = "lo"
	PRIu16                        = "u"
	PRIu32                        = "u"
	PRIu64                        = "lu"
	PRIu8                         = "u"
	PRIuFAST16                    = "u"
	PRIuFAST32                    = "u"
	PRIuFAST64                    = "lu"
	PRIuFAST8                     = "u"
	PRIuLEAST16                   = "u"
	PRIuLEAST32                   = "u"
	PRIuLEAST64                   = "lu"
	PRIuLEAST8                    = "u"
	PRIuMAX                       = "lu"
	PRIuPTR                       = "lu"
	PRIx16                        = "x"
	PRIx32                        = "x"
	PRIx64                        = "lx"
	PRIx8                         = "x"
	PRIxFAST16                    = "x"
	PRIxFAST32                    = "x"
	PRIxFAST64                    = "lx"
	PRIxFAST8                     = "x"
	PRIxLEAST16                   = "x"
	PRIxLEAST32                   = "x"
	PRIxLEAST64                   = "lx"
	PRIxLEAST8                    = "x"
	PRIxMAX                       = "lx"
	PRIxPTR                       = "lx"
	PTRDIFF_MAX                   = 9223372036854775807
	PTRDIFF_MIN                   = -9223372036854775808
	SCNd16                        = "hd"
	SCNd32                        = "d"
	SCNd64                        = "ld"
	SCNd8                         = "hhd"
	SCNdFAST16                    = "d"
	SCNdFAST32                    = "d"
	SCNdFAST64                    = "ld"
	SCNdFAST8                     = "d"
	SCNdLEAST16                   = "hd"
	SCNdLEAST32                   = "d"
	SCNdLEAST64                   = "ld"
	SCNdLEAST8                    = "hhd"
	SCNdMAX                       = "ld"
	SCNdPTR                       = "ld"
	SCNi16                        = "hi"
	SCNi32                        = "i"
	SCNi64                        = "li"
	SCNi8                         = "hhi"
	SCNiFAST16                    = "i"
	SCNiFAST32                    = "i"
	SCNiFAST64                    = "li"
	SCNiFAST8                     = "i"
	SCNiLEAST16                   = "hi"
	SCNiLEAST32                   = "i"
	SCNiLEAST64                   = "li"
	SCNiLEAST8                    = "hhi"
	SCNiMAX                       = "li"
	SCNiPTR                       = "li"
	SCNo16                        = "ho"
	SCNo32                        = "o"
	SCNo64                        = "lo"
	SCNo8                         = "hho"
	SCNoFAST16                    = "o"
	SCNoFAST32                    = "o"
	SCNoFAST64                    = "lo"
	SCNoFAST8                     = "o"
	SCNoLEAST16                   = "ho"
	SCNoLEAST32                   = "o"
	SCNoLEAST64                   = "lo"
	SCNoLEAST8                    = "hho"
	SCNoMAX                       = "lo"
	SCNoPTR                       = "lo"
	SCNu16                        = "hu"
	SCNu32                        = "u"
	SCNu64                        = "lu"
	SCNu8                         = "hhu"
	SCNuFAST16                    = "u"
	SCNuFAST32                    = "u"
	SCNuFAST64                    = "lu"
	SCNuFAST8                     = "u"
	SCNuLEAST16                   = "hu"
	SCNuLEAST32                   = "u"
	SCNuLEAST64                   = "lu"
	SCNuLEAST8                    = "hhu"
	SCNuMAX                       = "lu"
	SCNuPTR                       = "lu"
	SCNx16                        = "hx"
	SCNx32                        = "x"
	SCNx64                        = "lx"
	SCNx8                         = "hhx"
	SCNxFAST16                    = "x"
	SCNxFAST32                    = "x"
	SCNxFAST64                    = "lx"
	SCNxFAST8                     = "x"
	SCNxLEAST16                   = "hx"
	SCNxLEAST32                   = "x"
	SCNxLEAST64                   = "lx"
	SCNxLEAST8                    = "hhx"
	SCNxMAX                       = "lx"
	SCNxPTR                       = "lx"
	SCOPE_DELIMITER               = 37
	SIG_ATOMIC_MAX                = 2147483647
	SIG_ATOMIC_MIN                = -2147483648
	SIZE_MAX                      = 18446744073709551615
	TRY_AGAIN                     = 2
	UINT16_MAX                    = 65535
	UINT32_MAX                    = 4294967295
	UINT64_MAX                    = 18446744073709551615
	UINT8_MAX                     = 255
	UINTMAX_MAX                   = 18446744073709551615
	UINTPTR_MAX                   = 18446744073709551615
	UINT_FAST16_MAX               = 4294967295
	UINT_FAST32_MAX               = 4294967295
	UINT_FAST64_MAX               = 18446744073709551615
	UINT_FAST8_MAX                = 4294967295
	UINT_LEAST16_MAX              = 65535
	UINT_LEAST32_MAX              = 4294967295
	UINT_LEAST64_MAX              = 18446744073709551615
	UINT_LEAST8_MAX               = 255
	WCHAR_MAX                     = 0x7fffffff
	WCHAR_MIN                     = -2147483648
	WINT_MAX                      = 0x7fffffff
	WINT_MIN                      = -2147483648
	X_AMD64_INT_CONST_H_          = 0
	X_AMD64_INT_FMTIO_H_          = 0
	X_AMD64_INT_LIMITS_H_         = 0
	X_AMD64_INT_MWGWTYPES_H_      = 0
	X_AMD64_INT_TYPES_H_          = 0
	X_AMD64_WCHAR_LIMITS_H_       = 0
	X_BSD_INT16_T_                = 0
	X_BSD_INT32_T_                = 0
	X_BSD_INT64_T_                = 0
	X_BSD_INT8_T_                 = 0
	X_BSD_INTPTR_T_               = 0
	X_BSD_UINT16_T_               = 0
	X_BSD_UINT32_T_               = 0
	X_BSD_UINT64_T_               = 0
	X_BSD_UINT8_T_                = 0
	X_BSD_UINTPTR_T_              = 0
	X_FILE_OFFSET_BITS            = 64
	X_INTTYPES_H_                 = 0
	X_LP64                        = 1
	X_NETBSD_SOURCE               = 1
	X_NETDB_H_                    = 0
	X_PATH_HEQUIV                 = "/etc/hosts.equiv"
	X_PATH_HOSTS                  = "/etc/hosts"
	X_PATH_NETWORKS               = "/etc/networks"
	X_PATH_PROTOCOLS              = "/etc/protocols"
	X_PATH_SERVICES               = "/etc/services"
	X_PATH_SERVICES_CDB           = "/var/db/services.cdb"
	X_PATH_SERVICES_DB            = "/var/db/services.db"
	X_SYS_ANSI_H_                 = 0
	X_SYS_CDEFS_ELF_H_            = 0
	X_SYS_CDEFS_H_                = 0
	X_SYS_COMMON_ANSI_H_          = 0
	X_SYS_COMMON_INT_LIMITS_H_    = 0
	X_SYS_COMMON_INT_MWGWTYPES_H_ = 0
	X_SYS_COMMON_INT_TYPES_H_     = 0
	X_SYS_INTTYPES_H_             = 0
	X_SYS_STDINT_H_               = 0
	X_X86_64_CDEFS_H_             = 0
)

type Ptrdiff_t = int64

type Size_t = uint64

type Wchar_t = int32

type X__int128_t = struct {
	Flo int64
	Fhi int64
}
type X__uint128_t = struct {
	Flo uint64
	Fhi uint64
}

type X__builtin_va_list = uintptr
type X__float128 = float64

type X__int8_t = int8
type X__uint8_t = uint8
type X__int16_t = int16
type X__uint16_t = uint16
type X__int32_t = int32
type X__uint32_t = uint32
type X__int64_t = int64
type X__uint64_t = uint64

type X__intptr_t = int64
type X__uintptr_t = uint64

type X__caddr_t = uintptr
type X__gid_t = X__uint32_t
type X__in_addr_t = X__uint32_t
type X__in_port_t = X__uint16_t
type X__mode_t = X__uint32_t
type X__off_t = X__int64_t
type X__pid_t = X__int32_t
type X__sa_family_t = X__uint8_t
type X__socklen_t = uint32
type X__uid_t = X__uint32_t
type X__fsblkcnt_t = X__uint64_t
type X__fsfilcnt_t = X__uint64_t
type X__wctrans_t = uintptr
type X__wctype_t = uintptr

type X__mbstate_t = struct {
	F__mbstateL  X__int64_t
	F__ccgo_pad1 [120]byte
}

type X__va_list = X__builtin_va_list

type Int8_t = X__int8_t

type Uint8_t = X__uint8_t

type Int16_t = X__int16_t

type Uint16_t = X__uint16_t

type Int32_t = X__int32_t

type Uint32_t = X__uint32_t

type Int64_t = X__int64_t

type Uint64_t = X__uint64_t

type Intptr_t = X__intptr_t

type Uintptr_t = X__uintptr_t

type Int_least8_t = int8
type Uint_least8_t = uint8
type Int_least16_t = int16
type Uint_least16_t = uint16
type Int_least32_t = int32
type Uint_least32_t = uint32
type Int_least64_t = int64
type Uint_least64_t = uint64

type Int_fast8_t = int32
type Uint_fast8_t = uint32
type Int_fast16_t = int32
type Uint_fast16_t = uint32
type Int_fast32_t = int32
type Uint_fast32_t = uint32
type Int_fast64_t = int64
type Uint_fast64_t = uint64

type Intmax_t = int64
type Uintmax_t = uint64

type Imaxdiv_t = struct {
	Fquot Intmax_t
	Frem  Intmax_t
}

type Locale_t = uintptr

type Socklen_t = X__socklen_t

type Hostent = struct {
	Fh_name      uintptr
	Fh_aliases   uintptr
	Fh_addrtype  int32
	Fh_length    int32
	Fh_addr_list uintptr
}

type Netent = struct {
	Fn_name     uintptr
	Fn_aliases  uintptr
	Fn_addrtype int32
	Fn_net      Uint32_t
}

type Servent = struct {
	Fs_name      uintptr
	Fs_aliases   uintptr
	Fs_port      int32
	F__ccgo_pad1 [4]byte
	Fs_proto     uintptr
}

type Protoent = struct {
	Fp_name      uintptr
	Fp_aliases   uintptr
	Fp_proto     int32
	F__ccgo_pad1 [4]byte
}

type Addrinfo = struct {
	Fai_flags     int32
	Fai_family    int32
	Fai_socktype  int32
	Fai_protocol  int32
	Fai_addrlen   X__socklen_t
	F__ccgo_pad1  [4]byte
	Fai_canonname uintptr
	Fai_addr      uintptr
	Fai_next      uintptr
}

var _ int8