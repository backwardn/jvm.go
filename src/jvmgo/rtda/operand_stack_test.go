package rtda

import (
    "testing"
    "jvmgo/rtda/class"
)

func TestPushAndPop(t *testing.T) {
    stack := newOperandStack(6)
    stack.PushNull()
    stack.PushRef(class.NewArray(4, 4))
    stack.PushInt(-37)
    stack.PushLong(0xabcd1234ff)
    stack.PushFloat(3.14)
    stack.PushDouble(-2.71828)
    //stack.PushInt(0)

    if x:= stack.PopDouble(); x != -2.71828 {
        t.Errorf("double:%v", x)
    }
    if x:= stack.PopFloat(); x != 3.14 {
        t.Errorf("float:%v", x)
    }
    if x := stack.PopLong(); x != 0xabcd1234ff {
        t.Errorf("long:%v", x)
    }
    if x := stack.PopInt(); x != -37 {
        t.Errorf("int:%v", x)
    }
    if x := stack.PopRef(); x == nil {
        t.Errorf("nil!")
    }
    if x := stack.PopRef(); x != nil {
        t.Errorf("not nil: %v", x)
    }
}

func TestPopN(t *testing.T) {
    stack := newOperandStack(6)
    stack.PushInt(4)
    stack.PushInt(5)
    stack.PushInt(6)
    stack.PushInt(8)
    stack.PushInt(9)

    top3 := stack.PopN(3)
    if top3[0].(int32) != 6 || top3[1].(int32) != 8 || top3[2].(int32) != 9 {
        t.Errorf("top3: %v", top3)
    }

    if newTop := stack.PopInt(); newTop != 5 {
        t.Errorf("newTop: %v", newTop)
    }
}