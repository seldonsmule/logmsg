//
// demos the use of the logmsg package
//

package logmsg

import "testing"
import "log"
//import "logmsg"

func TestBasicLogging(t *testing.T){

  Print(Info, "Hello cruel world - info log msg - to standard out")	
  Print(Debug02, "Simple debug02 message") // will not see default is INFO level
  
  SetLogLevel(Debug03)

  Print(Debug02, "Should see this error message now")
	
}

func TestOpenFileAndUse(t *testing.T){
	
	
	if !SetLogFile("/tmp/logmsg_test.log") {
	  t.Error("Unable to create/open logfile")
	}
	
	Print(Info,"does this write to the file?")
	SetLogLevel(Error)
	
	Print(Info, "You should not see me in /tmp/logmsg_test.lgo")
	Print(Debug02, "This should be skipped from printing as well")
	Print(Error, "Example Error level message")
	Print(Critical, "Example Critical level message")

}


func TestOpenFailFails(t *testing.T){
	if !SetLogFile("/logmsg_test.log"){
		return
	}
	
	// if here we somehow succeeded which is bad.
	t.Error("Should not have been able to open /tmp/logmsg_test.log")
}

func myTestGoRoutine1(c1 chan string){
	
	for i:=0; i < 100; i++ {
	  log.Println("hello")
	  c1 <- "hello"
	  Print(Info, "In myTestGoRoutine1: ", i)		
	}
	
}

func myTestGoRoutine2(c1 chan string, c2 chan string){
	
	for i := 0; i < 100; i++ {
	  msg := <- c1 
	  log.Println(msg , "and goodby")
	  Print(Info, "In myTestGoRoutine2: ", i)
	}
	
	c2 <- "done"
	
}
func TestGoRoutine(t *testing.T){
	
	c1 := make(chan string)
	c2 := make(chan string)

	if !SetLogFile("/tmp/logmsg_test_goroutine.log") {
	  t.Error("Unable to create/open logfile")
	}
	
	SetLogLevel(Debug02)

	go myTestGoRoutine1(c1)
	go myTestGoRoutine2(c1, c2)
	
	final_msg := <- c2
	
	log.Println("Final_msg: ", final_msg)
	
}
