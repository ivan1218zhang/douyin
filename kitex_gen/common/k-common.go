// Code generated by Kitex v0.4.4. DO NOT EDIT.


package common

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
)

func (p *BaseResp) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_BaseResp[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *BaseResp) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.StatusCode = v

	}
	return offset, nil
}

func (p *BaseResp) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.StatusMessage = v

	}
	return offset, nil
}

// for compatibility
func (p *BaseResp) FastWrite(buf []byte) int {
	return 0
}

func (p *BaseResp) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "BaseResp")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *BaseResp) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("BaseResp")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *BaseResp) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "status_code", thrift.I64, 1)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.StatusCode)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *BaseResp) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "status_message", thrift.STRING, 2)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.StatusMessage)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *BaseResp) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("status_code", thrift.I64, 1)
	l += bthrift.Binary.I64Length(p.StatusCode)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *BaseResp) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("status_message", thrift.STRING, 2)
	l += bthrift.Binary.StringLengthNocopy(p.StatusMessage)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *User) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.BOOL {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_User[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *User) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Id = v

	}
	return offset, nil
}

func (p *User) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Name = v

	}
	return offset, nil
}

func (p *User) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.FollowCount = v

	}
	return offset, nil
}

func (p *User) FastReadField4(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.FollowerCount = v

	}
	return offset, nil
}

func (p *User) FastReadField5(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadBool(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.IsFollow = v

	}
	return offset, nil
}

// for compatibility
func (p *User) FastWrite(buf []byte) int {
	return 0
}

func (p *User) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "User")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField5(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *User) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("User")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *User) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "id", thrift.I64, 1)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.Id)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *User) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "name", thrift.STRING, 2)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.Name)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *User) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "follow_count", thrift.I64, 3)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.FollowCount)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *User) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "follower_count", thrift.I64, 4)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.FollowerCount)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *User) fastWriteField5(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "is_follow", thrift.BOOL, 5)
	offset += bthrift.Binary.WriteBool(buf[offset:], p.IsFollow)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *User) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("id", thrift.I64, 1)
	l += bthrift.Binary.I64Length(p.Id)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *User) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("name", thrift.STRING, 2)
	l += bthrift.Binary.StringLengthNocopy(p.Name)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *User) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("follow_count", thrift.I64, 3)
	l += bthrift.Binary.I64Length(p.FollowCount)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *User) field4Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("follower_count", thrift.I64, 4)
	l += bthrift.Binary.I64Length(p.FollowerCount)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *User) field5Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("is_follow", thrift.BOOL, 5)
	l += bthrift.Binary.BoolLength(p.IsFollow)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Video) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField6(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 7:
			if fieldTypeId == thrift.BOOL {
				l, err = p.FastReadField7(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 8:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField8(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Video[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Video) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Id = v

	}
	return offset, nil
}

func (p *Video) FastReadField2(buf []byte) (int, error) {
	offset := 0
	tmp := NewUser()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Author = tmp
	return offset, nil
}

func (p *Video) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.PlayUrl = v

	}
	return offset, nil
}

func (p *Video) FastReadField4(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.CoverUrl = v

	}
	return offset, nil
}

func (p *Video) FastReadField5(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.FavoriteCount = v

	}
	return offset, nil
}

func (p *Video) FastReadField6(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.CommentCount = v

	}
	return offset, nil
}

func (p *Video) FastReadField7(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadBool(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.IsFavorite = v

	}
	return offset, nil
}

func (p *Video) FastReadField8(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Title = v

	}
	return offset, nil
}

// for compatibility
func (p *Video) FastWrite(buf []byte) int {
	return 0
}

func (p *Video) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "Video")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField5(buf[offset:], binaryWriter)
		offset += p.fastWriteField6(buf[offset:], binaryWriter)
		offset += p.fastWriteField7(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField8(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *Video) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("Video")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
		l += p.field7Length()
		l += p.field8Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *Video) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "id", thrift.I64, 1)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.Id)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Video) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "Author", thrift.STRUCT, 2)
	offset += p.Author.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Video) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "PlayUrl", thrift.STRING, 3)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.PlayUrl)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Video) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "CoverUrl", thrift.STRING, 4)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.CoverUrl)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Video) fastWriteField5(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "favorite_count", thrift.I64, 5)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.FavoriteCount)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Video) fastWriteField6(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "comment_count", thrift.I64, 6)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.CommentCount)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Video) fastWriteField7(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "is_favorite", thrift.BOOL, 7)
	offset += bthrift.Binary.WriteBool(buf[offset:], p.IsFavorite)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Video) fastWriteField8(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "title", thrift.STRING, 8)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.Title)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Video) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("id", thrift.I64, 1)
	l += bthrift.Binary.I64Length(p.Id)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Video) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("Author", thrift.STRUCT, 2)
	l += p.Author.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Video) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("PlayUrl", thrift.STRING, 3)
	l += bthrift.Binary.StringLengthNocopy(p.PlayUrl)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Video) field4Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("CoverUrl", thrift.STRING, 4)
	l += bthrift.Binary.StringLengthNocopy(p.CoverUrl)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Video) field5Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("favorite_count", thrift.I64, 5)
	l += bthrift.Binary.I64Length(p.FavoriteCount)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Video) field6Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("comment_count", thrift.I64, 6)
	l += bthrift.Binary.I64Length(p.CommentCount)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Video) field7Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("is_favorite", thrift.BOOL, 7)
	l += bthrift.Binary.BoolLength(p.IsFavorite)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Video) field8Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("title", thrift.STRING, 8)
	l += bthrift.Binary.StringLengthNocopy(p.Title)

	l += bthrift.Binary.FieldEndLength()
	return l
}
