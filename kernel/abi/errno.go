/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package abi

// Errno represents a Linux errno value.
//
// Adapted from gvisor/pkg/abi/linux/errno/errno.go, Apache-2.0. The numeric
// values follow include/uapi/asm-generic/errno-base.h and errno.h.
type Errno uint32

// Errno values from include/uapi/asm-generic/errno-base.h and errno.h.
const (
	NOERRNO = iota
	EPERM
	ENOENT
	ESRCH
	EINTR
	EIO
	ENXIO
	E2BIG
	ENOEXEC
	EBADF
	ECHILD // 10
	EAGAIN
	ENOMEM
	EACCES
	EFAULT
	ENOTBLK
	EBUSY
	EEXIST
	EXDEV
	ENODEV
	ENOTDIR // 20
	EISDIR
	EINVAL
	ENFILE
	EMFILE
	ENOTTY
	ETXTBSY
	EFBIG
	ENOSPC
	ESPIPE
	EROFS // 30
	EMLINK
	EPIPE
	EDOM
	ERANGE
	// Errno values from include/uapi/asm-generic/errno.h.
	EDEADLK
	ENAMETOOLONG
	ENOLCK
	ENOSYS
	ENOTEMPTY
	ELOOP  // 40
	_      // Skip for EWOULDBLOCK = EAGAIN.
	ENOMSG // 42
	EIDRM
	ECHRNG
	EL2NSYNC
	EL3HLT
	EL3RST
	ELNRNG
	EUNATCH
	ENOCSI
	EL2HLT // 50
	EBADE
	EBADR
	EXFULL
	ENOANO
	EBADRQC
	EBADSLT
	_ // Skip for EDEADLOCK = EDEADLK.
	EBFONT
	ENOSTR // 60
	ENODATA
	ETIME
	ENOSR
	ENONET
	ENOPKG
	EREMOTE
	ENOLINK
	EADV
	ESRMNT
	ECOMM // 70
	EPROTO
	EMULTIHOP
	EDOTDOT
	EBADMSG
	EOVERFLOW
	ENOTUNIQ
	EBADFD
	EREMCHG
	ELIBACC
	ELIBBAD // 80
	ELIBSCN
	ELIBMAX
	ELIBEXEC
	EILSEQ
	ERESTART
	ESTRPIPE
	EUSERS
	ENOTSOCK
	EDESTADDRREQ
	EMSGSIZE // 90
	EPROTOTYPE
	ENOPROTOOPT
	EPROTONOSUPPORT
	ESOCKTNOSUPPORT
	EOPNOTSUPP
	EPFNOSUPPORT
	EAFNOSUPPORT
	EADDRINUSE
	EADDRNOTAVAIL
	ENETDOWN // 100
	ENETUNREACH
	ENETRESET
	ECONNABORTED
	ECONNRESET
	ENOBUFS
	EISCONN
	ENOTCONN
	ESHUTDOWN
	ETOOMANYREFS
	ETIMEDOUT // 110
	ECONNREFUSED
	EHOSTDOWN
	EHOSTUNREACH
	EALREADY
	EINPROGRESS
	ESTALE
	EUCLEAN
	ENOTNAM
	ENAVAIL
	EISNAM // 120
	EREMOTEIO
	EDQUOT
	ENOMEDIUM
	EMEDIUMTYPE
	ECANCELED
	ENOKEY
	EKEYEXPIRED
	EKEYREVOKED
	EKEYREJECTED
	EOWNERDEAD // 130
	ENOTRECOVERABLE
	ERFKILL
	EHWPOISON
)

// errno values derived from other errno values.
const (
	// EWOULDBLOCK is an alias for EAGAIN.
	EWOULDBLOCK = EAGAIN

	// EDEADLOCK is an alias for EDEADLK.
	EDEADLOCK = EDEADLK
)

// errno values for internal restart signaling.
const (
	// ERESTARTSYS indicates the syscall should be restarted unless interrupted
	// by a signal delivered without SA_RESTART.
	ERESTARTSYS = 512

	// ERESTARTNOINTR indicates the syscall should always be restarted.
	ERESTARTNOINTR = 513

	// ERESTARTNOHAND indicates the syscall should be restarted unless
	// interrupted by a signal delivered to a user handler.
	ERESTARTNOHAND = 514

	// ERestartBlock indicates the syscall should be restarted using a custom
	// restart function. Upstream names this ERESTART_RESTARTBLOCK.
	ERestartBlock = 516
)
