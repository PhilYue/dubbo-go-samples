// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package pkg

import (
	"bytes"
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// Attributes:
//  - ID
//  - Name
//  - Age
//  - Address
//  - Mobile
type User struct {
	ID      int32  `thrift:"id,1,required" json:"id"`
	Name    string `thrift:"name,2,required" json:"name"`
	Age     int32  `thrift:"Age,3,required" json:"Age"`
	Address string `thrift:"address,4,required" json:"address"`
	Mobile  string `thrift:"mobile,5,required" json:"mobile"`
}

func NewUser() *User {
	return &User{}
}

func (p *User) GetID() int32 {
	return p.ID
}

func (p *User) GetName() string {
	return p.Name
}

func (p *User) GetAge() int32 {
	return p.Age
}

func (p *User) GetAddress() string {
	return p.Address
}

func (p *User) GetMobile() string {
	return p.Mobile
}
func (p *User) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(context.TODO()); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetID bool = false
	var issetName bool = false
	var issetAge bool = false
	var issetAddress bool = false
	var issetMobile bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(context.TODO())
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetID = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetName = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
			issetAge = true
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
			issetAddress = true
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
			issetMobile = true
		default:
			if err := iprot.Skip(context.TODO(), fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(context.TODO()); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(context.TODO()); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetID {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ID is not set"))
	}
	if !issetName {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Name is not set"))
	}
	if !issetAge {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Age is not set"))
	}
	if !issetAddress {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Address is not set"))
	}
	if !issetMobile {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Mobile is not set"))
	}
	return nil
}

func (p *User) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(context.TODO()); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *User) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(context.TODO()); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *User) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(context.TODO()); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Age = v
	}
	return nil
}

func (p *User) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(context.TODO()); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Address = v
	}
	return nil
}

func (p *User) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(context.TODO()); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Mobile = v
	}
	return nil
}

func (p *User) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(context.TODO(), "User"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(context.TODO()); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(context.TODO()); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *User) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(context.TODO(), "id", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err)
	}
	if err := oprot.WriteI32(context.TODO(), int32(p.ID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(context.TODO()); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err)
	}
	return err
}

func (p *User) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(context.TODO(), "name", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:name: ", p), err)
	}
	if err := oprot.WriteString(context.TODO(), string(p.Name)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.name (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(context.TODO()); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:name: ", p), err)
	}
	return err
}

func (p *User) writeField3(oprot thrift.TProtocol) (err error) {
	ctx := context.TODO()
	if err := oprot.WriteFieldBegin(ctx, "Age", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:Age: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.Age)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Age (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:Age: ", p), err)
	}
	return err
}

func (p *User) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(context.TODO(), "address", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:address: ", p), err)
	}
	if err := oprot.WriteString(context.TODO(), string(p.Address)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.address (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(context.TODO()); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:address: ", p), err)
	}
	return err
}

func (p *User) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(context.TODO(), "mobile", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:mobile: ", p), err)
	}
	if err := oprot.WriteString(context.TODO(), string(p.Mobile)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.mobile (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(context.TODO()); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:mobile: ", p), err)
	}
	return err
}

func (p *User) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("User(%+v)", *p)
}
