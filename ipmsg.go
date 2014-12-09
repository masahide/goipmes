package ipmsg

/*  IP Messenger Communication Protocol version 1.2 define  */
/*  macro  */
func GetMode(command uint32) uint32 {
	return command & 0x000000ff
}
func GetOpt(command uint32) uint32 {
	return command & 0xffffff00
}

/*  header  */
var (
	IPMSG_VERSION = uint32(0x0001)

	/*  command  */
	IPMSG_NOOPERATION     = uint32(0x00000000)
	IPMSG_BR_ENTRY        = uint32(0x00000001)
	IPMSG_BR_EXIT         = uint32(0x00000002)
	IPMSG_ANSENTRY        = uint32(0x00000003)
	IPMSG_BR_ABSENCE      = uint32(0x00000004)
	IPMSG_BR_ISGETLIST    = uint32(0x00000010)
	IPMSG_OKGETLIST       = uint32(0x00000011)
	IPMSG_GETLIST         = uint32(0x00000012)
	IPMSG_ANSLIST         = uint32(0x00000013)
	IPMSG_BR_ISGETLIST2   = uint32(0x00000018)
	IPMSG_SENDMSG         = uint32(0x00000020)
	IPMSG_RECVMSG         = uint32(0x00000021)
	IPMSG_READMSG         = uint32(0x00000030)
	IPMSG_DELMSG          = uint32(0x00000031)
	IPMSG_ANSREADMSG      = uint32(0x00000032)
	IPMSG_GETINFO         = uint32(0x00000040)
	IPMSG_SENDINFO        = uint32(0x00000041)
	IPMSG_GETABSENCEINFO  = uint32(0x00000050)
	IPMSG_SENDABSENCEINFO = uint32(0x00000051)
	IPMSG_GETFILEDATA     = uint32(0x00000060)
	IPMSG_RELEASEFILES    = uint32(0x00000061)
	IPMSG_GETDIRFILES     = uint32(0x00000062)
	IPMSG_GETPUBKEY       = uint32(0x00000072)
	IPMSG_ANSPUBKEY       = uint32(0x00000073)

	/*  option for all command  */
	IPMSG_ABSENCEOPT    = uint32(0x00000100)
	IPMSG_SERVEROPT     = uint32(0x00000200)
	IPMSG_DIALUPOPT     = uint32(0x00010000)
	IPMSG_FILEATTACHOPT = uint32(0x00200000)
	IPMSG_ENCRYPTOPT    = uint32(0x00400000)

	/*  option for send command  */
	IPMSG_SENDCHECKOPT = uint32(0x00000100)
	IPMSG_SECRETOPT    = uint32(0x00000200)
	IPMSG_BROADCASTOPT = uint32(0x00000400)
	IPMSG_MULTICASTOPT = uint32(0x00000800)
	IPMSG_NOPOPUPOPT   = uint32(0x00001000)
	IPMSG_AUTORETOPT   = uint32(0x00002000)
	IPMSG_RETRYOPT     = uint32(0x00004000)
	IPMSG_PASSWORDOPT  = uint32(0x00008000)
	IPMSG_NOLOGOPT     = uint32(0x00020000)
	IPMSG_NEWMUTIOPT   = uint32(0x00040000)
	IPMSG_NOADDLISTOPT = uint32(0x00080000)
	IPMSG_READCHECKOPT = uint32(0x00100000)
	IPMSG_SECRETEXOPT  = IPMSG_READCHECKOPT | IPMSG_SECRETOPT

	/* encryption flags for encrypt command */
	IPMSG_RSA_512      = uint32(0x00000001)
	IPMSG_RSA_1024     = uint32(0x00000002)
	IPMSG_RSA_2048     = uint32(0x00000004)
	IPMSG_RC2_40       = uint32(0x00001000)
	IPMSG_RC2_128      = uint32(0x00004000)
	IPMSG_RC2_256      = uint32(0x00008000)
	IPMSG_BLOWFISH_128 = uint32(0x00020000)
	IPMSG_BLOWFISH_256 = uint32(0x00040000)
	IPMSG_SIGN_MD5     = uint32(0x10000000)

	/* compatibilty for Win beta version */
	IPMSG_RC2_40OLD       = uint32(0x00000010) // for beta1-4 only
	IPMSG_RC2_128OLD      = uint32(0x00000040) // for beta1-4 only
	IPMSG_BLOWFISH_128OLD = uint32(0x00000400) // for beta1-4 only
	IPMSG_RC2_40ALL       = IPMSG_RC2_40 | IPMSG_RC2_40OLD
	IPMSG_RC2_128ALL      = IPMSG_RC2_128 | IPMSG_RC2_128OLD
	IPMSG_BLOWFISH_128ALL = IPMSG_BLOWFISH_128 | IPMSG_BLOWFISH_128OLD
)

func GetFiletype(command uint32) uint32 {
	return command & 0x000000ff
}
func GetFileattr(command uint32) uint32 {
	return command & 0xffffff00
}

var (
	/* file types for fileattach command */
	IPMSG_FILE_REGULAR   = uint32(0x00000001)
	IPMSG_FILE_DIR       = uint32(0x00000002)
	IPMSG_FILE_RETPARENT = uint32(0x00000003) // return parent directory
	IPMSG_FILE_SYMLINK   = uint32(0x00000004)
	IPMSG_FILE_CDEV      = uint32(0x00000005) // for UNIX
	IPMSG_FILE_BDEV      = uint32(0x00000006) // for UNIX
	IPMSG_FILE_FIFO      = uint32(0x00000007) // for UNIX
	IPMSG_FILE_RESFORK   = uint32(0x00000010) // for Mac

	/* file attribute options for fileattach command */
	IPMSG_FILE_RONLYOPT    = uint32(0x00000100)
	IPMSG_FILE_HIDDENOPT   = uint32(0x00001000)
	IPMSG_FILE_EXHIDDENOPT = uint32(0x00002000) // for MacOS X
	IPMSG_FILE_ARCHIVEOPT  = uint32(0x00004000)
	IPMSG_FILE_SYSTEMOPT   = uint32(0x00008000)

	/* extend attribute types for fileattach command */
	IPMSG_FILE_UID          = uint32(0x00000001)
	IPMSG_FILE_USERNAME     = uint32(0x00000002) // uid by string
	IPMSG_FILE_GID          = uint32(0x00000003)
	IPMSG_FILE_GROUPNAME    = uint32(0x00000004) // gid by string
	IPMSG_FILE_PERM         = uint32(0x00000010) // for UNIX
	IPMSG_FILE_MAJORNO      = uint32(0x00000011) // for UNIX devfile
	IPMSG_FILE_MINORNO      = uint32(0x00000012) // for UNIX devfile
	IPMSG_FILE_CTIME        = uint32(0x00000013) // for UNIX
	IPMSG_FILE_MTIME        = uint32(0x00000014)
	IPMSG_FILE_ATIME        = uint32(0x00000015)
	IPMSG_FILE_CREATETIME   = uint32(0x00000016)
	IPMSG_FILE_CREATOR      = uint32(0x00000020) // for Mac
	IPMSG_FILE_FILETYPE     = uint32(0x00000021) // for Mac
	IPMSG_FILE_FINDERINFO   = uint32(0x00000022) // for Mac
	IPMSG_FILE_ACL          = uint32(0x00000030)
	IPMSG_FILE_ALIASFNAME   = uint32(0x00000040) // alias fname
	IPMSG_FILE_UNICODEFNAME = uint32(0x00000041) // UNICODE fname

	FILELIST_SEPARATOR = '\a'
	HOSTLIST_SEPARATOR = '\a'
	HOSTLIST_DUMMY     = "\b"

	/*  end of IP Messenger Communication Protocol version 1.2 define  */

	// General define
	MAX_SOCKBUF  = 65536
	MAX_UDPBUF   = 16384
	MAX_CRYPTLEN = (MAX_UDPBUF - MAX_BUF) / 2
	MAX_BUF      = 1024
	MAX_NAMEBUF  = 50
	MAX_LANGBUF  = 10
	MAX_LISTBUF  = MAX_NAMEBUF*3 + 50
	MAX_ANSLIST  = 100
)
