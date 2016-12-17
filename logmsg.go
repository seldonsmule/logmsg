//
//
// logmsg package
//
// logmsg extends the go log package to provide the concept of having multiple
// levels of logging.  This allows the developer to keep debug messages in the
// code and with low overhead have them skipped until a more detailed level
// of logging is desired.
// also provides a single open a logfile instead of writing to stdout (default)
//
//

package logmsg

import "fmt"
import "log"
import "sync"
import "os"



type LogLevel int


// Error levels that the user can set
const (
	Critical LogLevel = 1 + iota
	Error
	Warning
	Info
	Debug01
	Debug02
	Debug03
    LevelEnd
)

// internal structure of keep all our stuff
// 
type MyLogMsg struct {
	mu sync.Mutex // needed only for changing the log level or other global items
	currentLevel LogLevel
	levelNames [LevelEnd]string
	logFileName string
	setup bool
    fileHandle *os.File
	ourLog *log.Logger

}

//
// New()
// Allows a users to creat a new instance of logmsg
//
func New(level LogLevel, outputFilename string) *MyLogMsg{
	
	//fmt.Println("filename len: ", len(outputFilename))
	
	return &MyLogMsg{currentLevel: level, logFileName: outputFilename, setup:false}
}

// our internal global instance 
var internalLogMsg = New(Info, "")

//
// init()
// sets up our inital (default) settings.  uses the bool setup to make sure it only happens once
//
func (l *MyLogMsg) init(){
	
	if l.setup{
	  return
	}
	
	
	l.mu.Lock()

	if l.ourLog == nil{
	  l.ourLog = log.New(os.Stdout, "", log.LstdFlags)

	}

    //l.ourLog.SetOutput(os.Stdout)		
	l.ourLog.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	
	//log.SetOutput(os.Stdout)		
	//log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	
	

	
	l.levelNames[Critical] = "CRITICAL"
	l.levelNames[Error] = "ERROR"
	l.levelNames[Warning] = "WARNING"
	l.levelNames[Info] = "INFO"
	l.levelNames[Debug01] = "DEBUG01"
	l.levelNames[Debug02] = "DEBUG02"
	l.levelNames[Debug03] = "DEBUG03"
	
	l.setup = true
	
	l.mu.Unlock()
	
}

//
// SetLogFileName(name string) bool
// 
// Sets up writint go a file.  If there is already a file opened, that will be close 1st (allowing for file rollover)
//
//  name - New name of the logfile.  Will be used relative to where the program is running unless an explicit path name is used
//
//  Returns a true|false depending on success.  Up to the caller to abort/panic
//
func (l* MyLogMsg) SetLogFile(name string) bool{
	
	var err error
	
	
	l.logFileName = name

	if(l.fileHandle != nil){
		l.fileHandle.Close()
	}
	
	l.fileHandle, err = os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
    
	if err != nil {
      //log.Fatalln("Failed to open log file :", err)
	  var msg string
	  msg = fmt.Sprint(err)
	  l.ourLog.Print(Critical, "Error: ", msg)
	  return false
    }else{
		//l.ourLog.SetOutput(l.fileHandle)
		l.ourLog = log.New(l.fileHandle, "", 0)
		l.ourLog.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
		//l.ourLog.Print(Info, "Opened logfile[", name, "]")
	}
	
	return true
	
}

//
//
// Print(level LogLevel, v...)
// 
//   Used to print a message - actaully calls a lower level Print (PrintCallbacklevel) to do this
//
//   level - log level
// 
//   v... - Variable data
//
// 
func (l *MyLogMsg) Print(level LogLevel, v ...interface{}){
	
	
	l.PrintCallbacklevel(1, level, v...)

}

//
//
// PrintCallbacklevel(callback int, level LogLevel, v...)
// 
//   Used to print a message - 
//
//   callback - allows the caller to say how far into the code they are so we can skip back
//              up the stack to get a more helpful filename/line number
//   level - log level
// 
//   v... - Variable data
//
// 
func (l *MyLogMsg) PrintCallbacklevel(callback int, level LogLevel, v ...interface{}){
	var msg string
	
	l.init()
	
	msg = fmt.Sprint(v...)
	
	l.ourLog.Output(callback+3, fmt.Sprintln(l.levelNames[level],"|", msg))
	  
}

//
// SetLogLevel(level LogLevel)
//
// Lets the caller change the log level
//
// level - new loglevel
//

func (l *MyLogMsg) SetLogLevel(level LogLevel){
	l.mu.Lock()
	l.currentLevel = level
	l.mu.Unlock()
}


// public/externally callable methods

//
//
// Print(level LogLevel, v...)
// 
//   Used to print a message - actaully calls a lower level Print (PrintCallbacklevel) to do this
//
//   level - log level
// 
//   v... - Variable data
//
// 
func Print(level LogLevel, v ...interface{}){
	
	if level > internalLogMsg.currentLevel{
		return
	}
    internalLogMsg.Print(level, v...)	
}

//
//
// PrintCallbacklevel(callback int, level LogLevel, v...)
// 
//   Used to print a message - 
//
//   callback - allows the caller to say how far into the code they are so we can skip back
//              up the stack to get a more helpful filename/line number
//   level - log level
// 
//   v... - Variable data
//
// 
func PrintCallbacklevel(callback int, level LogLevel, v ...interface{}){
	
	if level > internalLogMsg.currentLevel{
		return
	}
    internalLogMsg.PrintCallbacklevel(callback, level, v...)	
}


//
// SetLogLevel(level LogLevel)
//
// Lets the caller change the log level
//
// level - new loglevel
//

func SetLogLevel(level LogLevel){
	internalLogMsg.SetLogLevel(level)
}

//
// SetLogFileName(name string) bool
// 
// Sets up writint go a file.  If there is already a file opened, that will be close 1st (allowing for file rollover)
//
//  name - New name of the logfile.  Will be used relative to where the program is running unless an explicit path name is used
//
//  Returns a true|false depending on success.  Up to the caller to abort/panic
//
func SetLogFile(name string) bool{
	return internalLogMsg.SetLogFile(name)
}
