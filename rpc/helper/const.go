package helper

type EncType int8

const (
	ENCTYPE_PLAIN   EncType = iota
	ENCTYPE_DEFLATE EncType = iota
)

func (p EncType) String() string {
	switch p {
	case ENCTYPE_PLAIN:
		return "plain"
	case ENCTYPE_DEFLATE:
		return "deflate"
	}
	return "Unknown"
}

type Status int

const (
	STATUS_UNKNOWERROR  Status = iota
	STATUS_OK           Status = iota
	STATUS_UPDATED      Status = iota
	STATUS_NOTREADY     Status = iota
	STATUS_SEQERROR     Status = iota
	STATUS_VERSIONERROR Status = iota
)

func (p Status) String() string {
	switch p {
	case STATUS_OK:
		return "Ok"
	case STATUS_UPDATED:
		return "Updated"
	case STATUS_NOTREADY:
		return "NotReady"
	case STATUS_UNKNOWERROR:
		return "UnknownError"
	case STATUS_SEQERROR:
		return "SeqError"
	case STATUS_VERSIONERROR:
		return "VersionError"
	}
	return "Unknown"
}

/*
type RequestType int8

const (
	REQUESTTYPE_GETVERSION RequestType = iota
	REQUESTTYPE_TOTAL      RequestType = iota
	REQUESTTYPE_REAL       RequestType = iota
)

func (p RequestType) String() string {
	switch p {
	case REQUESTTYPE_GETVERSION:
		return "GetVersion"
	case REQUESTTYPE_TOTAL:
		return "total"
	case REQUESTTYPE_REAL:
		return "real"
	}
	return "Unknown"
}*/
