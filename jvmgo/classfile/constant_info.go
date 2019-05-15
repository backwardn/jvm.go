package classfile

import (
	"math"

	"github.com/zxh0/jvm.go/jvmgo/jutil"
)

// Constant pool tags
const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

/*
cp_info {
    u1 tag;
    u1 info[];
}

CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantInfo interface {
	//readInfo(reader *ClassReader)
}
type ConstantInfo2 interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp *ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	switch tag {
	case CONSTANT_Integer:
		return int32(reader.readUint32())
	case CONSTANT_Float:
		return math.Float32frombits(reader.readUint32())
	case CONSTANT_Long:
		return int64(reader.readUint64())
	case CONSTANT_Double:
		return math.Float64frombits(reader.readUint64())
	case CONSTANT_Utf8:
		return readUtf8(reader)
	}

	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

// todo ugly code
func newConstantInfo(tag uint8, cp *ConstantPool) ConstantInfo2 {
	switch tag {

	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{cp: cp}
	default: // todo
		jutil.Panicf("BAD constant pool tag: %v", tag)
		return nil
	}
}
