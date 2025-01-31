// Code generated by 'ccgo errno/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o errno/errno_illumos_amd64.go -pkgname errno', DO NOT EDIT.

package errno

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
	E2BIG              = 7
	EACCES             = 13
	EADDRINUSE         = 125
	EADDRNOTAVAIL      = 126
	EADV               = 68
	EAFNOSUPPORT       = 124
	EAGAIN             = 11
	EALREADY           = 149
	EBADE              = 50
	EBADF              = 9
	EBADFD             = 81
	EBADMSG            = 77
	EBADR              = 51
	EBADRQC            = 54
	EBADSLT            = 55
	EBFONT             = 57
	EBUSY              = 16
	ECANCELED          = 47
	ECHILD             = 10
	ECHRNG             = 37
	ECOMM              = 70
	ECONNABORTED       = 130
	ECONNREFUSED       = 146
	ECONNRESET         = 131
	EDEADLK            = 45
	EDEADLOCK          = 56
	EDESTADDRREQ       = 96
	EDOM               = 33
	EDQUOT             = 49
	EEXIST             = 17
	EFAULT             = 14
	EFBIG              = 27
	EHOSTDOWN          = 147
	EHOSTUNREACH       = 148
	EIDRM              = 36
	EILSEQ             = 88
	EINPROGRESS        = 150
	EINTR              = 4
	EINVAL             = 22
	EIO                = 5
	EISCONN            = 133
	EISDIR             = 21
	EL2HLT             = 44
	EL2NSYNC           = 38
	EL3HLT             = 39
	EL3RST             = 40
	ELIBACC            = 83
	ELIBBAD            = 84
	ELIBEXEC           = 87
	ELIBMAX            = 86
	ELIBSCN            = 85
	ELNRNG             = 41
	ELOCKUNMAPPED      = 72
	ELOOP              = 90
	EMFILE             = 24
	EMLINK             = 31
	EMSGSIZE           = 97
	EMULTIHOP          = 74
	ENAMETOOLONG       = 78
	ENETDOWN           = 127
	ENETRESET          = 129
	ENETUNREACH        = 128
	ENFILE             = 23
	ENOANO             = 53
	ENOBUFS            = 132
	ENOCSI             = 43
	ENODATA            = 61
	ENODEV             = 19
	ENOENT             = 2
	ENOEXEC            = 8
	ENOLCK             = 46
	ENOLINK            = 67
	ENOMEM             = 12
	ENOMSG             = 35
	ENONET             = 64
	ENOPKG             = 65
	ENOPROTOOPT        = 99
	ENOSPC             = 28
	ENOSR              = 63
	ENOSTR             = 60
	ENOSYS             = 89
	ENOTACTIVE         = 73
	ENOTBLK            = 15
	ENOTCONN           = 134
	ENOTDIR            = 20
	ENOTEMPTY          = 93
	ENOTRECOVERABLE    = 59
	ENOTSOCK           = 95
	ENOTSUP            = 48
	ENOTTY             = 25
	ENOTUNIQ           = 80
	ENXIO              = 6
	EOPNOTSUPP         = 122
	EOVERFLOW          = 79
	EOWNERDEAD         = 58
	EPERM              = 1
	EPFNOSUPPORT       = 123
	EPIPE              = 32
	EPROTO             = 71
	EPROTONOSUPPORT    = 120
	EPROTOTYPE         = 98
	ERANGE             = 34
	EREMCHG            = 82
	EREMOTE            = 66
	ERESTART           = 91
	EROFS              = 30
	ESHUTDOWN          = 143
	ESOCKTNOSUPPORT    = 121
	ESPIPE             = 29
	ESRCH              = 3
	ESRMNT             = 69
	ESTALE             = 151
	ESTRPIPE           = 92
	ETIME              = 62
	ETIMEDOUT          = 145
	ETOOMANYREFS       = 144
	ETXTBSY            = 26
	EUNATCH            = 42
	EUSERS             = 94
	EWOULDBLOCK        = 11
	EXDEV              = 18
	EXFULL             = 52
	X_ERRNO_H          = 0
	X_FILE_OFFSET_BITS = 64
	X_LP64             = 1
	X_SYS_ERRNO_H      = 0
	Sun                = 1
	Unix               = 1
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

var _ int8
