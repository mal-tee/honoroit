// Code generated by 'ccgo pwd/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o pwd/pwd_darwin_amd64.go -pkgname pwd', DO NOT EDIT.

package pwd

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
	X_BSD_I386__TYPES_H_                   = 0
	X_BSD_MACHINE__TYPES_H_                = 0
	X_CDEFS_H_                             = 0
	X_DARWIN_FEATURE_64_BIT_INODE          = 1
	X_DARWIN_FEATURE_ONLY_UNIX_CONFORMANCE = 1
	X_DARWIN_FEATURE_UNIX_CONFORMANCE      = 3
	X_FILE_OFFSET_BITS                     = 64
	X_FORTIFY_SOURCE                       = 2
	X_GID_T                                = 0
	X_LP64                                 = 1
	X_MASTERPASSWD                         = "master.passwd"
	X_MP_DB                                = "pwd.db"
	X_Nonnull                              = 0
	X_Null_unspecified                     = 0
	X_Nullable                             = 0
	X_PASSWD                               = "passwd"
	X_PASSWORD_CHGNOW                      = -1
	X_PASSWORD_EFMT1                       = 95
	X_PASSWORD_LEN                         = 128
	X_PASSWORD_NOCHG                       = 0x04
	X_PASSWORD_NOEXP                       = 0x08
	X_PASSWORD_NOGID                       = 0x02
	X_PASSWORD_NOUID                       = 0x01
	X_PASSWORD_WARNDAYS                    = 14
	X_PATH_MASTERPASSWD                    = "/etc/master.passwd"
	X_PATH_MASTERPASSWD_LOCK               = "/etc/ptmp"
	X_PATH_MP_DB                           = "/etc/pwd.db"
	X_PATH_PASSWD                          = "/etc/passwd"
	X_PATH_PWD                             = "/etc"
	X_PATH_PWD_MKDB                        = "/usr/sbin/pwd_mkdb"
	X_PATH_SMP_DB                          = "/etc/spwd.db"
	X_PWD_H_                               = 0
	X_PW_KEYBYNAME                         = 49
	X_PW_KEYBYNUM                          = 50
	X_PW_KEYBYUID                          = 51
	X_SIZE_T                               = 0
	X_SMP_DB                               = "spwd.db"
	X_SYS__PTHREAD_TYPES_H_                = 0
	X_SYS__TYPES_H_                        = 0
	X_UID_T                                = 0
	X_UUID_STRING_T                        = 0
	X_UUID_T                               = 0
	X_UUID_UUID_H                          = 0
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

var X__darwin_check_fd_set_overflow uintptr

type X__int8_t = int8
type X__uint8_t = uint8
type X__int16_t = int16
type X__uint16_t = uint16
type X__int32_t = int32
type X__uint32_t = uint32
type X__int64_t = int64
type X__uint64_t = uint64

type X__darwin_intptr_t = int64
type X__darwin_natural_t = uint32

type X__darwin_ct_rune_t = int32

type X__mbstate_t = struct {
	F__ccgo_pad1 [0]uint64
	F__mbstate8  [128]int8
}

type X__darwin_mbstate_t = X__mbstate_t

type X__darwin_ptrdiff_t = int64

type X__darwin_size_t = uint64

type X__darwin_va_list = X__builtin_va_list

type X__darwin_wchar_t = int32

type X__darwin_rune_t = X__darwin_wchar_t

type X__darwin_wint_t = int32

type X__darwin_clock_t = uint64
type X__darwin_socklen_t = X__uint32_t
type X__darwin_ssize_t = int64
type X__darwin_time_t = int64

type X__darwin_blkcnt_t = X__int64_t
type X__darwin_blksize_t = X__int32_t
type X__darwin_dev_t = X__int32_t
type X__darwin_fsblkcnt_t = uint32
type X__darwin_fsfilcnt_t = uint32
type X__darwin_gid_t = X__uint32_t
type X__darwin_id_t = X__uint32_t
type X__darwin_ino64_t = X__uint64_t
type X__darwin_ino_t = X__darwin_ino64_t
type X__darwin_mach_port_name_t = X__darwin_natural_t
type X__darwin_mach_port_t = X__darwin_mach_port_name_t
type X__darwin_mode_t = X__uint16_t
type X__darwin_off_t = X__int64_t
type X__darwin_pid_t = X__int32_t
type X__darwin_sigset_t = X__uint32_t
type X__darwin_suseconds_t = X__int32_t
type X__darwin_uid_t = X__uint32_t
type X__darwin_useconds_t = X__uint32_t
type X__darwin_uuid_t = [16]uint8
type X__darwin_uuid_string_t = [37]int8

type X__darwin_pthread_handler_rec = struct {
	F__routine uintptr
	F__arg     uintptr
	F__next    uintptr
}

type X_opaque_pthread_attr_t = struct {
	F__sig    int64
	F__opaque [56]int8
}

type X_opaque_pthread_cond_t = struct {
	F__sig    int64
	F__opaque [40]int8
}

type X_opaque_pthread_condattr_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type X_opaque_pthread_mutex_t = struct {
	F__sig    int64
	F__opaque [56]int8
}

type X_opaque_pthread_mutexattr_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type X_opaque_pthread_once_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type X_opaque_pthread_rwlock_t = struct {
	F__sig    int64
	F__opaque [192]int8
}

type X_opaque_pthread_rwlockattr_t = struct {
	F__sig    int64
	F__opaque [16]int8
}

type X_opaque_pthread_t = struct {
	F__sig           int64
	F__cleanup_stack uintptr
	F__opaque        [8176]int8
}

type X__darwin_pthread_attr_t = X_opaque_pthread_attr_t
type X__darwin_pthread_cond_t = X_opaque_pthread_cond_t
type X__darwin_pthread_condattr_t = X_opaque_pthread_condattr_t
type X__darwin_pthread_key_t = uint64
type X__darwin_pthread_mutex_t = X_opaque_pthread_mutex_t
type X__darwin_pthread_mutexattr_t = X_opaque_pthread_mutexattr_t
type X__darwin_pthread_once_t = X_opaque_pthread_once_t
type X__darwin_pthread_rwlock_t = X_opaque_pthread_rwlock_t
type X__darwin_pthread_rwlockattr_t = X_opaque_pthread_rwlockattr_t
type X__darwin_pthread_t = uintptr

type X__darwin_nl_item = int32
type X__darwin_wctrans_t = int32
type X__darwin_wctype_t = X__uint32_t

type Gid_t = X__darwin_gid_t

type Uid_t = X__darwin_uid_t

type Passwd = struct {
	Fpw_name   uintptr
	Fpw_passwd uintptr
	Fpw_uid    Uid_t
	Fpw_gid    Gid_t
	Fpw_change X__darwin_time_t
	Fpw_class  uintptr
	Fpw_gecos  uintptr
	Fpw_dir    uintptr
	Fpw_shell  uintptr
	Fpw_expire X__darwin_time_t
}

type Uuid_t = X__darwin_uuid_t

type Uuid_string_t = X__darwin_uuid_string_t

var _ int8