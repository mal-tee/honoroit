// Code generated by 'ccgo fcntl/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o fcntl/fcntl_illumos_amd64.go -pkgname fcntl', DO NOT EDIT.

package fcntl

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
	AT_EACCESS                      = 0x4
	AT_FDCWD                        = 0xffd19553
	AT_REMOVEDIR                    = 0x1
	AT_SYMLINK_FOLLOW               = 0x2000
	AT_SYMLINK_NOFOLLOW             = 0x1000
	CLOCKS_PER_SEC                  = 1000000
	CLOCK_HIGHRES                   = 4
	CLOCK_MONOTONIC                 = 4
	CLOCK_PROCESS_CPUTIME_ID        = 5
	CLOCK_PROF                      = 2
	CLOCK_REALTIME                  = 3
	CLOCK_THREAD_CPUTIME_ID         = 2
	CLOCK_VIRTUAL                   = 1
	DIRECTIO_OFF                    = 0
	DIRECTIO_ON                     = 1
	DST_AUST                        = 2
	DST_AUSTALT                     = 10
	DST_CAN                         = 6
	DST_EET                         = 5
	DST_GB                          = 7
	DST_MET                         = 4
	DST_NONE                        = 0
	DST_RUM                         = 8
	DST_TUR                         = 9
	DST_USA                         = 1
	DST_WET                         = 3
	FD_CLOEXEC                      = 1
	FD_SETSIZE                      = 65536
	F_ALLOCSP                       = 10
	F_ALLOCSP64                     = 10
	F_BADFD                         = 46
	F_BLKSIZE                       = 19
	F_BLOCKS                        = 18
	F_CHKFL                         = 8
	F_COMPAT                        = 0x8
	F_DUP2FD                        = 9
	F_DUP2FD_CLOEXEC                = 36
	F_DUPFD                         = 0
	F_DUPFD_CLOEXEC                 = 37
	F_FLOCK                         = 53
	F_FLOCK64                       = 53
	F_FLOCKW                        = 54
	F_FLOCKW64                      = 54
	F_FREESP                        = 11
	F_FREESP64                      = 11
	F_GETFD                         = 1
	F_GETFL                         = 3
	F_GETLK                         = 14
	F_GETLK64                       = 14
	F_GETOWN                        = 23
	F_GETXFL                        = 45
	F_HASREMOTELOCKS                = 26
	F_ISSTREAM                      = 13
	F_MANDDNY                       = 0x10
	F_MDACC                         = 0x20
	F_NODNY                         = 0x0
	F_NPRIV                         = 16
	F_OFD_GETLK                     = 47
	F_OFD_GETLK64                   = 47
	F_OFD_SETLK                     = 48
	F_OFD_SETLK64                   = 48
	F_OFD_SETLKW                    = 49
	F_OFD_SETLKW64                  = 49
	F_PRIV                          = 15
	F_QUOTACTL                      = 17
	F_RDACC                         = 0x1
	F_RDDNY                         = 0x1
	F_RDLCK                         = 01
	F_REVOKE                        = 25
	F_RMACC                         = 0x4
	F_RMDNY                         = 0x4
	F_RWACC                         = 0x3
	F_RWDNY                         = 0x3
	F_SETFD                         = 2
	F_SETFL                         = 4
	F_SETLK                         = 6
	F_SETLK64                       = 6
	F_SETLK64_NBMAND                = 42
	F_SETLKW                        = 7
	F_SETLKW64                      = 7
	F_SETLK_NBMAND                  = 42
	F_SETOWN                        = 24
	F_SHARE                         = 40
	F_SHARE_NBMAND                  = 43
	F_UNLCK                         = 03
	F_UNLKSYS                       = 04
	F_UNSHARE                       = 41
	F_WRACC                         = 0x2
	F_WRDNY                         = 0x2
	F_WRLCK                         = 02
	ITIMER_PROF                     = 2
	ITIMER_REAL                     = 0
	ITIMER_REALPROF                 = 3
	ITIMER_VIRTUAL                  = 1
	MICROSEC                        = 1000000
	MILLISEC                        = 1000
	NANOSEC                         = 1000000000
	NBBY                            = 8
	O_ACCMODE                       = 6291459
	O_APPEND                        = 0x08
	O_CLOEXEC                       = 0x800000
	O_CREAT                         = 0x100
	O_DIRECT                        = 0x2000000
	O_DIRECTORY                     = 0x1000000
	O_DSYNC                         = 0x40
	O_EXCL                          = 0x400
	O_EXEC                          = 0x400000
	O_LARGEFILE                     = 0x2000
	O_NDELAY                        = 0x04
	O_NOCTTY                        = 0x800
	O_NOFOLLOW                      = 0x20000
	O_NOLINKS                       = 0x40000
	O_NONBLOCK                      = 0x80
	O_RDONLY                        = 0
	O_RDWR                          = 2
	O_RSYNC                         = 0x8000
	O_SEARCH                        = 0x200000
	O_SYNC                          = 0x10
	O_TRUNC                         = 0x200
	O_WRONLY                        = 1
	O_XATTR                         = 0x4000
	POSIX_FADV_DONTNEED             = 4
	POSIX_FADV_NOREUSE              = 5
	POSIX_FADV_NORMAL               = 0
	POSIX_FADV_RANDOM               = 1
	POSIX_FADV_SEQUENTIAL           = 2
	POSIX_FADV_WILLNEED             = 3
	P_MYID                          = -1
	REG_LABEL_BP                    = 2
	REG_LABEL_MAX                   = 8
	REG_LABEL_PC                    = 0
	REG_LABEL_R12                   = 4
	REG_LABEL_R13                   = 5
	REG_LABEL_R14                   = 6
	REG_LABEL_R15                   = 7
	REG_LABEL_RBX                   = 3
	REG_LABEL_SP                    = 1
	SEC                             = 1
	SEEK_DATA                       = 3
	SEEK_HOLE                       = 4
	TIMER_ABSTIME                   = 0x1
	TIMER_RELTIME                   = 0x0
	TIME_UTC                        = 0x1
	X_ALIGNMENT_REQUIRED            = 1
	X_AT_TRIGGER                    = 0x2
	X_BIT_FIELDS_LTOH               = 0
	X_BOOL_ALIGNMENT                = 1
	X_CHAR_ALIGNMENT                = 1
	X_CHAR_IS_SIGNED                = 0
	X_CLOCKID_T                     = 0
	X_CLOCK_T                       = 0
	X_COND_MAGIC                    = 0x4356
	X_DMA_USES_PHYSADDR             = 0
	X_DONT_USE_1275_GENERIC_NAMES   = 0
	X_DOUBLE_ALIGNMENT              = 8
	X_DOUBLE_COMPLEX_ALIGNMENT      = 8
	X_DTRACE_VERSION                = 1
	X_FCNTL_H                       = 0
	X_FILE_OFFSET_BITS              = 64
	X_FIRMWARE_NEEDS_FDISK          = 0
	X_FLOAT_ALIGNMENT               = 4
	X_FLOAT_COMPLEX_ALIGNMENT       = 4
	X_HAVE_CPUID_INSN               = 0
	X_IEEE_754                      = 0
	X_INT64_TYPE                    = 0
	X_INT_ALIGNMENT                 = 4
	X_ISO_CPP_14882_1998            = 0
	X_ISO_C_9899_1999               = 0
	X_ISO_C_9899_2011               = 0
	X_ISO_TIME_ISO_H                = 0
	X_LARGEFILE64_SOURCE            = 1
	X_LARGEFILE_SOURCE              = 1
	X_LITTLE_ENDIAN                 = 0
	X_LOCALE_T                      = 0
	X_LONGLONG_TYPE                 = 0
	X_LONG_ALIGNMENT                = 8
	X_LONG_DOUBLE_ALIGNMENT         = 16
	X_LONG_DOUBLE_COMPLEX_ALIGNMENT = 16
	X_LONG_LONG_ALIGNMENT           = 8
	X_LONG_LONG_ALIGNMENT_32        = 4
	X_LONG_LONG_LTOH                = 0
	X_LP64                          = 1
	X_MAX_ALIGNMENT                 = 16
	X_MULTI_DATAMODEL               = 0
	X_MUTEX_MAGIC                   = 0x4d58
	X_NBBY                          = 8
	X_NORETURN_KYWD                 = 0
	X_OFF_T                         = 0
	X_POINTER_ALIGNMENT             = 8
	X_PSM_MODULES                   = 0
	X_PTRDIFF_T                     = 0
	X_RESTRICT_KYWD                 = 0
	X_RTC_CONFIG                    = 0
	X_RWL_MAGIC                     = 0x5257
	X_SEMA_MAGIC                    = 0x534d
	X_SHORT_ALIGNMENT               = 2
	X_SIGEVENT                      = 0
	X_SIGSET_T                      = 0
	X_SIGVAL                        = 0
	X_SIZE_T                        = 0
	X_SOFT_HOSTID                   = 0
	X_SSIZE_T                       = 0
	X_STACK_GROWS_DOWNWARD          = 0
	X_STDC_C11                      = 0
	X_STDC_C99                      = 0
	X_SUNOS_VTOC_16                 = 0
	X_SUSECONDS_T                   = 0
	X_SYS_CCOMPILE_H                = 0
	X_SYS_FCNTL_H                   = 0
	X_SYS_FEATURE_TESTS_H           = 0
	X_SYS_INT_TYPES_H               = 0
	X_SYS_ISA_DEFS_H                = 0
	X_SYS_MACHTYPES_H               = 0
	X_SYS_NULL_H                    = 0
	X_SYS_SELECT_H                  = 0
	X_SYS_TIME_H                    = 0
	X_SYS_TIME_IMPL_H               = 0
	X_SYS_TYPES_H                   = 0
	X_TIMER_T                       = 0
	X_TIME_H                        = 0
	X_TIME_T                        = 0
	X_UID_T                         = 0
	X_XOPEN_VERSION                 = 3
	Sun                             = 1
	Unix                            = 1
)

const (
	B_FALSE   = 0
	B_TRUE    = 1
	X_B_FALSE = 0
	X_B_TRUE  = 1
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

type X_label_t = struct{ Fval [8]int64 }

type Label_t = X_label_t

type Lock_t = uint8

type Int8_t = int8
type Int16_t = int16
type Int32_t = int32
type Int64_t = int64

type Uint8_t = uint8
type Uint16_t = uint16
type Uint32_t = uint32
type Uint64_t = uint64

type Intmax_t = int64
type Uintmax_t = uint64

type Intptr_t = int64
type Uintptr_t = uint64

type Int_fast8_t = int8
type Int_fast16_t = int32
type Int_fast32_t = int32
type Int_fast64_t = int64

type Uint_fast8_t = uint8
type Uint_fast16_t = uint32
type Uint_fast32_t = uint32
type Uint_fast64_t = uint64

type Int_least8_t = int8
type Int_least16_t = int16
type Int_least32_t = int32
type Int_least64_t = int64

type Uint_least8_t = uint8
type Uint_least16_t = uint16
type Uint_least32_t = uint32
type Uint_least64_t = uint64

type Longlong_t = int64
type U_longlong_t = uint64

type T_scalar_t = int32
type T_uscalar_t = uint32

type Uchar_t = uint8
type Ushort_t = uint16
type Uint_t = uint32
type Ulong_t = uint64

type Caddr_t = uintptr
type Daddr_t = int64
type Cnt_t = int16

type Pfn_t = uint64
type Pgcnt_t = uint64
type Spgcnt_t = int64

type Use_t = uint8
type Sysid_t = int16
type Index_t = int16
type Timeout_id_t = uintptr
type Bufcall_id_t = uintptr

type Off_t = int64

type Off64_t = int64

type Ino_t = uint64
type Blkcnt_t = int64
type Fsblkcnt_t = uint64
type Fsfilcnt_t = uint64

type Ino64_t = uint64
type Blkcnt64_t = int64
type Fsblkcnt64_t = uint64
type Fsfilcnt64_t = uint64

type Blksize_t = int32

type Boolean_t = uint32

type Pad64_t = int64
type Upad64_t = uint64

type Pad128_t = struct {
	F_q          float64
	F__ccgo_pad1 [8]byte
}

type Upad128_t = struct {
	F_q          float64
	F__ccgo_pad1 [8]byte
}

type Offset_t = int64
type U_offset_t = uint64
type Len_t = uint64
type Diskaddr_t = uint64

type Lloff_t = struct{ F_f int64 }

type Lldaddr_t = struct{ F_f int64 }

type K_fltset_t = uint32

type Id_t = int32

type Lgrp_id_t = int32

type Useconds_t = uint32

type Suseconds_t = int64

type Major_t = uint32
type Minor_t = uint32

type Pri_t = int16

type Cpu_flag_t = uint16

type O_mode_t = uint16
type O_dev_t = int16
type O_uid_t = uint16
type O_gid_t = uint16
type O_nlink_t = int16
type O_pid_t = int16
type O_ino_t = uint16

type Key_t = int32
type Mode_t = uint32

type Uid_t = uint32

type Gid_t = uint32

type Datalink_id_t = uint32
type Vrid_t = uint32

type Taskid_t = int32
type Projid_t = int32
type Poolid_t = int32
type Zoneid_t = int32
type Ctid_t = int32

type Pthread_t = uint32
type Pthread_key_t = uint32

type X_pthread_mutex = struct {
	F__pthread_mutex_flags struct {
		F__pthread_mutex_flag1   uint16
		F__pthread_mutex_flag2   uint8
		F__pthread_mutex_ceiling uint8
		F__pthread_mutex_type    uint16
		F__pthread_mutex_magic   uint16
	}
	F__pthread_mutex_lock struct {
		F__ccgo_pad1            [0]uint64
		F__pthread_mutex_lock64 struct{ F__pthread_mutex_pad [8]uint8 }
	}
	F__pthread_mutex_data uint64
}

type Pthread_mutex_t = X_pthread_mutex

type X_pthread_cond = struct {
	F__pthread_cond_flags struct {
		F__pthread_cond_flag  [4]uint8
		F__pthread_cond_type  uint16
		F__pthread_cond_magic uint16
	}
	F__pthread_cond_data uint64
}

type Pthread_cond_t = X_pthread_cond

type X_pthread_rwlock = struct {
	F__pthread_rwlock_readers  int32
	F__pthread_rwlock_type     uint16
	F__pthread_rwlock_magic    uint16
	F__pthread_rwlock_mutex    Pthread_mutex_t
	F__pthread_rwlock_readercv Pthread_cond_t
	F__pthread_rwlock_writercv Pthread_cond_t
}

type Pthread_rwlock_t = X_pthread_rwlock

type Pthread_barrier_t = struct {
	F__pthread_barrier_count    uint32
	F__pthread_barrier_current  uint32
	F__pthread_barrier_cycle    uint64
	F__pthread_barrier_reserved uint64
	F__pthread_barrier_lock     Pthread_mutex_t
	F__pthread_barrier_cond     Pthread_cond_t
}

type Pthread_spinlock_t = Pthread_mutex_t

type X_pthread_attr = struct{ F__pthread_attrp uintptr }

type Pthread_attr_t = X_pthread_attr

type X_pthread_mutexattr = struct{ F__pthread_mutexattrp uintptr }

type Pthread_mutexattr_t = X_pthread_mutexattr

type X_pthread_condattr = struct{ F__pthread_condattrp uintptr }

type Pthread_condattr_t = X_pthread_condattr

type X_once = struct{ F__pthread_once_pad [4]uint64 }

type Pthread_once_t = X_once

type X_pthread_rwlockattr = struct{ F__pthread_rwlockattrp uintptr }

type Pthread_rwlockattr_t = X_pthread_rwlockattr

type Pthread_barrierattr_t = struct{ F__pthread_barrierattrp uintptr }

type Dev_t = uint64

type Nlink_t = uint32
type Pid_t = int32

type Ssize_t = int64

type Time_t = int64

type Clock_t = int64

type Clockid_t = int32

type Timer_t = int32

type Unchar = uint8
type Ushort = uint16
type Uint = uint32
type Ulong = uint64

type U_char = uint8
type U_short = uint16
type U_int = uint32
type U_long = uint64
type X_quad = struct{ Fval [2]int32 }

type Quad_t = X_quad
type Quad = Quad_t

type Timespec = struct {
	Ftv_sec  int64
	Ftv_nsec int64
}

type Timespec_t = Timespec

type Timestruc_t = Timespec

type Itimerspec = struct {
	Fit_interval struct {
		Ftv_sec  int64
		Ftv_nsec int64
	}
	Fit_value struct {
		Ftv_sec  int64
		Ftv_nsec int64
	}
}

type Itimerspec_t = Itimerspec

type Timeval = struct {
	Ftv_sec  int64
	Ftv_usec int64
}

type Timezone = struct {
	Ftz_minuteswest int32
	Ftz_dsttime     int32
}

type Itimerval = struct {
	Fit_interval struct {
		Ftv_sec  int64
		Ftv_usec int64
	}
	Fit_value struct {
		Ftv_sec  int64
		Ftv_usec int64
	}
}

type Hrtime_t = int64

type Tm = struct {
	Ftm_sec   int32
	Ftm_min   int32
	Ftm_hour  int32
	Ftm_mday  int32
	Ftm_mon   int32
	Ftm_year  int32
	Ftm_wday  int32
	Ftm_yday  int32
	Ftm_isdst int32
}

type Sigval = struct {
	F__ccgo_pad1 [0]uint64
	Fsival_int   int32
	F__ccgo_pad2 [4]byte
}

type Sigevent = struct {
	Fsigev_notify int32
	Fsigev_signo  int32
	Fsigev_value  struct {
		F__ccgo_pad1 [0]uint64
		Fsival_int   int32
		F__ccgo_pad2 [4]byte
	}
	Fsigev_notify_function   uintptr
	Fsigev_notify_attributes uintptr
	F__sigev_pad2            int32
	F__ccgo_pad1             [4]byte
}

type Locale_t = uintptr

type Sigset_t = struct{ F__sigbits [4]uint32 }

type Fd_mask = int64
type Fds_mask = int64

type Fd_set1 = struct{ Ffds_bits [1024]int64 }

type Fd_set = Fd_set1

type Flock = struct {
	Fl_type      int16
	Fl_whence    int16
	F__ccgo_pad1 [4]byte
	Fl_start     int64
	Fl_len       int64
	Fl_sysid     int32
	Fl_pid       int32
	Fl_pad       [4]int64
}

type Flock_t = Flock

type Flock64 = struct {
	Fl_type      int16
	Fl_whence    int16
	F__ccgo_pad1 [4]byte
	Fl_start     int64
	Fl_len       int64
	Fl_sysid     int32
	Fl_pid       int32
	Fl_pad       [4]int64
}

type Flock64_t = Flock64

type Fshare = struct {
	Ff_access int16
	Ff_deny   int16
	Ff_id     int32
}

type Fshare_t = Fshare

var _ int8