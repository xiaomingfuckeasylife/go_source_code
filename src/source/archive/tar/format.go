package tar


const (
	formatUnknown = (1 << iota) / 2

	formatV7

	formatGNU

	formatSTAR

	formatUSTAR

	formatPAX
)


const  (
	magicGNU , versionGUN = "ustar"," \x00"
	magincUSTAR , versionUSTAR = "ustar\x00","00"
	trailerSTAR ="tar\x00"
)

const (
	blockSize = 512
	nameSize = 100
	prefixSize = 155
)

var zeroBlock block

type block [blockSize]byte

func (b *block) V7() *headerV7 {return (*headerV7)(b)}
func (b *block) GNU() *headerGNU {return (*headerGNU)(b)}
func (b *block) STAR() *headerSTAR{return (*headerSTAR)(b)}
func (b *block) USTAR() *headerUSTAR{return (*headerUSTAR)(b)}
func (b *block) Sparse() sparseArray{return (sparseArray)(b[:])}

type headerV7 [blockSize]byte

func (h *headerV7) Name() []byte {return h[000:][:100]}
func (h *headerV7) Mode() []byte {return h[100:][:8]}
func (h *headerV7) UID() []byte {return h[108:][:8]}
func (h *headerV7) GID() []byte {return h[116:][:8]}
func (h *headerV7) Size() []byte {return h[124:][:12]}
func (h *headerV7) ModTime() []byte {return h[136:][:12]}
func (h *headerV7) Chksum() []byte {return h[148:][:8]}
func (h *headerV7) TypeFlag() []byte {return h[136:][:12]}
func (h *headerV7) LinkName() []byte {return h[157:][:100]}


type headerGNU [blockSize]byte

func (h *headerGNU) V7() *headerV7       { return (*headerV7)(h) }
func (h *headerGNU) Magic() []byte       { return h[257:][:6] }
func (h *headerGNU) Version() []byte     { return h[263:][:2] }
func (h *headerGNU) UserName() []byte    { return h[265:][:32] }
func (h *headerGNU) GroupName() []byte   { return h[297:][:32] }
func (h *headerGNU) DevMajor() []byte    { return h[329:][:8] }
func (h *headerGNU) DevMinor() []byte    { return h[337:][:8] }
func (h *headerGNU) AccessTime() []byte  { return h[345:][:12] }
func (h *headerGNU) ChangeTime() []byte  { return h[357:][:12] }
func (h *headerGNU) Sparse() sparseArray { return (sparseArray)(h[386:][:24*4+1]) }
func (h *headerGNU) RealSize() []byte    { return h[483:][:12] }

type headerSTAR [blockSize]byte

func (h *headerSTAR) V7() *headerV7      { return (*headerV7)(h) }
func (h *headerSTAR) Magic() []byte      { return h[257:][:6] }
func (h *headerSTAR) Version() []byte    { return h[263:][:2] }
func (h *headerSTAR) UserName() []byte   { return h[265:][:32] }
func (h *headerSTAR) GroupName() []byte  { return h[297:][:32] }
func (h *headerSTAR) DevMajor() []byte   { return h[329:][:8] }
func (h *headerSTAR) DevMinor() []byte   { return h[337:][:8] }
func (h *headerSTAR) Prefix() []byte     { return h[345:][:131] }
func (h *headerSTAR) AccessTime() []byte { return h[476:][:12] }
func (h *headerSTAR) ChangeTime() []byte { return h[488:][:12] }
func (h *headerSTAR) Trailer() []byte    { return h[508:][:4] }

type headerUSTAR [blockSize]byte

func (h *headerUSTAR) V7() *headerV7     { return (*headerV7)(h) }
func (h *headerUSTAR) Magic() []byte     { return h[257:][:6] }
func (h *headerUSTAR) Version() []byte   { return h[263:][:2] }
func (h *headerUSTAR) UserName() []byte  { return h[265:][:32] }
func (h *headerUSTAR) GroupName() []byte { return h[297:][:32] }
func (h *headerUSTAR) DevMajor() []byte  { return h[329:][:8] }
func (h *headerUSTAR) DevMinor() []byte  { return h[337:][:8] }
func (h *headerUSTAR) Prefix() []byte    { return h[345:][:155] }

type sparseArray []byte

func (s sparseArray) Entry(i int) sparseNode { return (sparseNode)(s[i*24:]) }
func (s sparseArray) IsExtended() []byte     { return s[24*s.MaxEntries():][:1] }
func (s sparseArray) MaxEntries() int        { return len(s) / 24 }

type sparseNode []byte

func (s sparseNode) Offset() []byte   { return s[00:][:12] }
func (s sparseNode) NumBytes() []byte { return s[12:][:12] }

