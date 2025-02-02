package op

import "github.com/go-ole/go-ole"

/* 获取鼠标位置.

参数定义:

x int类型 返回X坐标

y int类型 返回Y坐标

返回值:

int类型 0失败，1成功 */
func (com *Opsoft) GetCursorPos(x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.op.CallMethod("GetCursorPos", &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

/*
获取指定的按键状态.(前台信息,不是后台)

参数定义:

vkCode int类型 虚拟按键码

返回值:

int类型 0弹起，1按下
*/
func (com *Opsoft) GetKeyState(vkCode int) int {
	ret, _ := com.op.CallMethod("GetKeyState", vkCode)
	return int(ret.Val)
}

/*
按住指定的虚拟键码

参数定义:

vkCode int 虚拟按键码

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) KeyDown(vkCode string) int {
	ret, _ := com.op.CallMethod("KeyDown", vkCode)
	return int(ret.Val)
}

/*
按住指定的虚拟键码

参数定义:

keyStr string 字符串描述的键码. 大小写无所谓.

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) KeyDownChar(keyStr string) int {
	ret, _ := com.op.CallMethod("KeyDownChar", keyStr)
	return int(ret.Val)
}

/*
按住指定的虚拟键码

参数定义:

vkCode int 虚拟按键码

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) KeyPress(vkCode int) int {
	ret, _ := com.op.CallMethod("KeyPress", vkCode)
	return int(ret.Val)
}

/*
按住指定的虚拟键码

参数定义:

keyStr string 字符串描述的键码. 大小写无所谓.

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) KeyPressChar(keyStr string) int {
	ret, _ := com.op.CallMethod("KeyPressChar", keyStr)
	return int(ret.Val)
}

/*
弹起来虚拟键

参数定义:

vkCode int 虚拟按键码.

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) KeyUp(vkCode int) int {
	ret, _ := com.op.CallMethod("KeyUp", vkCode)
	return int(ret.Val)
}

/*
弹起来虚拟键

参数定义:

keyStr string 字符串描述的键码. 大小写无所谓.

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) KeyUpChar(keyStr string) int {
	ret, _ := com.op.CallMethod("KeyUpChar", keyStr)
	return int(ret.Val)
}

/*
按下鼠标左键

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) LeftClick() int {
	ret, _ := com.op.CallMethod("LeftClick")
	return int(ret.Val)
}

/*
双击鼠标左键

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) LeftDoubleClick() int {
	ret, _ := com.op.CallMethod("LeftDoubleClick")
	return int(ret.Val)
}

/*
按住鼠标左键

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) LeftDown() int {
	ret, _ := com.op.CallMethod("LeftDown")
	return int(ret.Val)
}

/*
弹起鼠标左键

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) LeftUp() int {
	ret, _ := com.op.CallMethod("LeftUp")
	return int(ret.Val)
}

/*
按下鼠标中键

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) MiddleClick() int {
	ret, _ := com.op.CallMethod("MiddleClick")
	return int(ret.Val)
}

/*
按住鼠标中键

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) MiddleDown() int {
	ret, _ := com.op.CallMethod("MiddleDown")
	return int(ret.Val)
}

/*
弹起鼠标中键

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) MiddleUp() int {
	ret, _ := com.op.CallMethod("MiddleUp")
	return int(ret.Val)
}

/*
鼠标相对于上次的位置移动rx,ry.

参数定义:

rx int 相对于上次的X偏移.
ry int 相对于上次的Y偏移.

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) MoveR(rx, ry int) int {
	ret, _ := com.op.CallMethod("MoveR", rx, ry)
	return int(ret.Val)
}

/*
把鼠标移动到目的点(x,y)

参数定义:

x int X坐标.
y int Y坐标.

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) MoveTo(x, y int) int {
	ret, _ := com.op.CallMethod("MoveTo", x, y)
	return int(ret.Val)
}

/*
把鼠标移动到目的范围内的任意一点

参数定义:

x int X坐标
y int Y坐标
w int 宽度(从x计算起)
h int 高度(从y计算起)

返回值:

string 返回要移动到的目标点. 格式为x,y. 比如MoveToEx 100,100,10,10,返回值可能是101,102
*/
func (com *Opsoft) MoveToEx(x, y, w, h int) string {
	ret, _ := com.op.CallMethod("MoveToEx", x, y, w, h)
	return ret.ToString()
}

/*
按下鼠标右键

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) RightClick() int {
	ret, _ := com.op.CallMethod("RightClick")
	return int(ret.Val)
}

/*
按住鼠标右键

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) RightDown() int {
	ret, _ := com.op.CallMethod("RightDown")
	return int(ret.Val)
}

/*
弹起鼠标右键

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) RightUp() int {
	ret, _ := com.op.CallMethod("RightUp")
	return int(ret.Val)
}

/*
等待指定的按键按下 (前台,不是后台)

参数定义:

vkCode int 虚拟按键码,当此值为0，表示等待任意按键。 鼠标左键是1,鼠标右键时2,鼠标中键是4.
timeOut int 等待多久,单位毫秒. 如果是0，表示一直等待

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) WaitKey(vkCode, timeOut int) int {
	ret, _ := com.op.CallMethod("SetSimMode", vkCode, timeOut)
	return int(ret.Val)
}

/*
滚轮向下滚

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) WheelDown() int {
	ret, _ := com.op.CallMethod("WheelDown")
	return int(ret.Val)
}

/*
滚轮向上滚

返回值:

int类型 0成功，1失败
*/
func (com *Opsoft) WheelUp() int {
	ret, _ := com.op.CallMethod("WheelUp")
	return int(ret.Val)
}
