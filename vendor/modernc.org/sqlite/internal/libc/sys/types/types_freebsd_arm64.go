// Code generated by 'ccgo sys/types/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o sys/types/types_freebsd_amd64.go -pkgname types', DO NOT EDIT.

package types

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
	BIG_ENDIAN              = 4321
	BYTE_ORDER              = 1234
	FD_SETSIZE              = 1024
	LITTLE_ENDIAN           = 1234
	PDP_ENDIAN              = 3412
	X_ACCMODE_T_DECLARED    = 0
	X_BIG_ENDIAN            = 4321
	X_BLKCNT_T_DECLARED     = 0
	X_BLKSIZE_T_DECLARED    = 0
	X_BYTE_ORDER            = 1234
	X_CAP_IOCTL_T_DECLARED  = 0
	X_CAP_RIGHTS_T_DECLARED = 0
	X_CLOCKID_T_DECLARED    = 0
	X_CLOCK_T_DECLARED      = 0
	X_DEV_T_DECLARED        = 0
	X_FFLAGS_T_DECLARED     = 0
	X_FILE_OFFSET_BITS      = 64
	X_FSBLKCNT_T_DECLARED   = 0
	X_FTRUNCATE_DECLARED    = 0
	X_GID_T_DECLARED        = 0
	X_ID_T_DECLARED         = 0
	X_INO_T_DECLARED        = 0
	X_INT16_T_DECLARED      = 0
	X_INT32_T_DECLARED      = 0
	X_INT64_T_DECLARED      = 0
	X_INT8_T_DECLARED       = 0
	X_INTMAX_T_DECLARED     = 0
	X_INTPTR_T_DECLARED     = 0
	X_IN_ADDR_T_DECLARED    = 0
	X_IN_PORT_T_DECLARED    = 0
	X_KEY_T_DECLARED        = 0
	X_LITTLE_ENDIAN         = 1234
	X_LP64                  = 1
	X_LSEEK_DECLARED        = 0
	X_LWPID_T_DECLARED      = 0
	X_MACHINE_ENDIAN_H_     = 0
	X_MACHINE__LIMITS_H_    = 0
	X_MACHINE__TYPES_H_     = 0
	X_MMAP_DECLARED         = 0
	X_MODE_T_DECLARED       = 0
	X_MQD_T_DECLARED        = 0
	X_NLINK_T_DECLARED      = 0
	X_Nonnull               = 0
	X_Null_unspecified      = 0
	X_Nullable              = 0
	X_OFF64_T_DECLARED      = 0
	X_OFF_T_DECLARED        = 0
	X_PDP_ENDIAN            = 3412
	X_PID_T_DECLARED        = 0
	X_PTHREAD_T_DECLARED    = 0
	X_QUAD_HIGHWORD         = 1
	X_QUAD_LOWWORD          = 0
	X_RLIM_T_DECLARED       = 0
	X_SELECT_DECLARED       = 0
	X_SIGSET_T_DECLARED     = 0
	X_SIG_MAXSIG            = 128
	X_SIG_WORDS             = 4
	X_SIZE_T_DECLARED       = 0
	X_SSIZE_T_DECLARED      = 0
	X_SUSECONDS_T_DECLARED  = 0
	X_SYS_CDEFS_H_          = 0
	X_SYS_SELECT_H_         = 0
	X_SYS_TIMESPEC_H_       = 0
	X_SYS_TYPES_H_          = 0
	X_SYS__ENDIAN_H_        = 0
	X_SYS__PTHREADTYPES_H_  = 0
	X_SYS__SIGSET_H_        = 0
	X_SYS__STDINT_H_        = 0
	X_SYS__TIMESPEC_H_      = 0
	X_SYS__TIMEVAL_H_       = 0
	X_SYS__TYPES_H_         = 0
	X_TIMER_T_DECLARED      = 0
	X_TIME_T_DECLARED       = 0
	X_TRUNCATE_DECLARED     = 0
	X_UID_T_DECLARED        = 0
	X_UINT16_T_DECLARED     = 0
	X_UINT32_T_DECLARED     = 0
	X_UINT64_T_DECLARED     = 0
	X_UINT8_T_DECLARED      = 0
	X_UINTMAX_T_DECLARED    = 0
	X_UINTPTR_T_DECLARED    = 0
	X_USECONDS_T_DECLARED   = 0
	Unix                    = 1
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

type X__clock_t = X__int32_t
type X__critical_t = X__int64_t
type X__double_t = float64
type X__float_t = float32
type X__intfptr_t = X__int64_t
type X__intptr_t = X__int64_t
type X__intmax_t = X__int64_t
type X__int_fast8_t = X__int32_t
type X__int_fast16_t = X__int32_t
type X__int_fast32_t = X__int32_t
type X__int_fast64_t = X__int64_t
type X__int_least8_t = X__int8_t
type X__int_least16_t = X__int16_t
type X__int_least32_t = X__int32_t
type X__int_least64_t = X__int64_t
type X__ptrdiff_t = X__int64_t
type X__register_t = X__int64_t
type X__segsz_t = X__int64_t
type X__size_t = X__uint64_t
type X__ssize_t = X__int64_t
type X__time_t = X__int64_t
type X__uintfptr_t = X__uint64_t
type X__uintptr_t = X__uint64_t
type X__uintmax_t = X__uint64_t
type X__uint_fast8_t = X__uint32_t
type X__uint_fast16_t = X__uint32_t
type X__uint_fast32_t = X__uint32_t
type X__uint_fast64_t = X__uint64_t
type X__uint_least8_t = X__uint8_t
type X__uint_least16_t = X__uint16_t
type X__uint_least32_t = X__uint32_t
type X__uint_least64_t = X__uint64_t
type X__u_register_t = X__uint64_t
type X__vm_offset_t = X__uint64_t
type X__vm_paddr_t = X__uint64_t
type X__vm_size_t = X__uint64_t
type X___wchar_t = int32

type X__blksize_t = X__int32_t
type X__blkcnt_t = X__int64_t
type X__clockid_t = X__int32_t
type X__fflags_t = X__uint32_t
type X__fsblkcnt_t = X__uint64_t
type X__fsfilcnt_t = X__uint64_t
type X__gid_t = X__uint32_t
type X__id_t = X__int64_t
type X__ino_t = X__uint64_t
type X__key_t = int64
type X__lwpid_t = X__int32_t
type X__mode_t = X__uint16_t
type X__accmode_t = int32
type X__nl_item = int32
type X__nlink_t = X__uint64_t
type X__off_t = X__int64_t
type X__off64_t = X__int64_t
type X__pid_t = X__int32_t
type X__rlim_t = X__int64_t

type X__sa_family_t = X__uint8_t
type X__socklen_t = X__uint32_t
type X__suseconds_t = int64
type X__timer_t = uintptr
type X__mqd_t = uintptr
type X__uid_t = X__uint32_t
type X__useconds_t = uint32
type X__cpuwhich_t = int32
type X__cpulevel_t = int32
type X__cpusetid_t = int32
type X__daddr_t = X__int64_t

type X__ct_rune_t = int32
type X__rune_t = X__ct_rune_t
type X__wint_t = X__ct_rune_t

type X__char16_t = X__uint_least16_t
type X__char32_t = X__uint_least32_t

type X__max_align_t = struct {
	F__max_align1 int64
	F__max_align2 float64
}

type X__dev_t = X__uint64_t

type X__fixpt_t = X__uint32_t

type X__mbstate_t = struct {
	F__ccgo_pad1 [0]uint64
	F__mbstate8  [128]int8
}

type X__rman_res_t = X__uintmax_t

type X__va_list = X__builtin_va_list
type X__gnuc_va_list = X__va_list
type Pthread_once = struct {
	Fstate       int32
	F__ccgo_pad1 [4]byte
	Fmutex       Pthread_mutex_t
}

type Pthread_t = uintptr
type Pthread_attr_t = uintptr
type Pthread_mutex_t = uintptr
type Pthread_mutexattr_t = uintptr
type Pthread_cond_t = uintptr
type Pthread_condattr_t = uintptr
type Pthread_key_t = int32
type Pthread_once_t = Pthread_once
type Pthread_rwlock_t = uintptr
type Pthread_rwlockattr_t = uintptr
type Pthread_barrier_t = uintptr
type Pthread_barrierattr_t = uintptr
type Pthread_spinlock_t = uintptr

type Pthread_addr_t = uintptr
type Pthread_startroutine_t = uintptr

type U_char = uint8
type U_short = uint16
type U_int = uint32
type U_long = uint64
type Ushort = uint16
type Uint = uint32

type Int8_t = X__int8_t

type Int16_t = X__int16_t

type Int32_t = X__int32_t

type Int64_t = X__int64_t

type Uint8_t = X__uint8_t

type Uint16_t = X__uint16_t

type Uint32_t = X__uint32_t

type Uint64_t = X__uint64_t

type Intptr_t = X__intptr_t
type Uintptr_t = X__uintptr_t
type Intmax_t = X__intmax_t
type Uintmax_t = X__uintmax_t

type U_int8_t = X__uint8_t
type U_int16_t = X__uint16_t
type U_int32_t = X__uint32_t
type U_int64_t = X__uint64_t

type U_quad_t = X__uint64_t
type Quad_t = X__int64_t
type Qaddr_t = uintptr

type Caddr_t = uintptr
type C_caddr_t = uintptr

type Blksize_t = X__blksize_t

type Cpuwhich_t = X__cpuwhich_t
type Cpulevel_t = X__cpulevel_t
type Cpusetid_t = X__cpusetid_t

type Blkcnt_t = X__blkcnt_t

type Clock_t = X__clock_t

type Clockid_t = X__clockid_t

type Critical_t = X__critical_t
type Daddr_t = X__daddr_t

type Dev_t = X__dev_t

type Fflags_t = X__fflags_t

type Fixpt_t = X__fixpt_t

type Fsblkcnt_t = X__fsblkcnt_t
type Fsfilcnt_t = X__fsfilcnt_t

type Gid_t = X__gid_t

type In_addr_t = X__uint32_t

type In_port_t = X__uint16_t

type Id_t = X__id_t

type Ino_t = X__ino_t

type Key_t = X__key_t

type Lwpid_t = X__lwpid_t

type Mode_t = X__mode_t

type Accmode_t = X__accmode_t

type Nlink_t = X__nlink_t

type Off_t = X__off_t

type Off64_t = X__off64_t

type Pid_t = X__pid_t

type Register_t = X__register_t

type Rlim_t = X__rlim_t

type Sbintime_t = X__int64_t

type Segsz_t = X__segsz_t

type Ssize_t = X__ssize_t

type Suseconds_t = X__suseconds_t

type Time_t = X__time_t

type Timer_t = X__timer_t

type Mqd_t = X__mqd_t

type U_register_t = X__u_register_t

type Uid_t = X__uid_t

type Useconds_t = X__useconds_t

type Cap_ioctl_t = uint64

type Kpaddr_t = X__uint64_t
type Kvaddr_t = X__uint64_t
type Ksize_t = X__uint64_t
type Kssize_t = X__int64_t

type Vm_offset_t = X__vm_offset_t
type Vm_ooffset_t = X__uint64_t
type Vm_paddr_t = X__vm_paddr_t
type Vm_pindex_t = X__uint64_t
type Vm_size_t = X__vm_size_t

type Rman_res_t = X__rman_res_t

type X__sigset = struct{ F__bits [4]X__uint32_t }

type X__sigset_t = X__sigset

type Timeval = struct {
	Ftv_sec  Time_t
	Ftv_usec Suseconds_t
}

type Timespec = struct {
	Ftv_sec  Time_t
	Ftv_nsec int64
}

type Itimerspec = struct {
	Fit_interval struct {
		Ftv_sec  Time_t
		Ftv_nsec int64
	}
	Fit_value struct {
		Ftv_sec  Time_t
		Ftv_nsec int64
	}
}

type X__fd_mask = uint64
type Fd_mask = X__fd_mask

type Sigset_t = X__sigset_t

type Fd_set1 = struct{ F__fds_bits [16]X__fd_mask }

type Fd_set = Fd_set1

var _ int8