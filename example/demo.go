package main

import "logmsg"

func printDebugMsgs(){

  logmsg.Print(logmsg.Debug02,"Debug 02 demo message - you should see this now")
  logmsg.Print(logmsg.Debug03,"Debug 03 demo message")
  logmsg.Print(logmsg.Debug01,"Debug 01 demo message")


}

func demoWhereCalled(){

  logmsg.Print(logmsg.Info,"If having a helper method that dumps an array, and you use logmsg.Print().  Notice the line # is inside this method")

  logmsg.PrintCallbacklevel(1, 
                            logmsg.Info,
                     "Now the line # is to where this methoed was called")

}

func main(){

  logmsg.Print(logmsg.Info,"Welcome to logmsg")

  // the default of logmsg is info level messaging
  // is INFO.  the debug messages should be suppressed

  printDebugMsgs()
 
  logmsg.Print(logmsg.Info,"You did not see the debug messages")
  logmsg.Print(logmsg.Info,"lowering log level to debug02")

  logmsg.SetLogLevel(logmsg.Debug02)

  printDebugMsgs()
  logmsg.Print(logmsg.Debug01,"Notice you did not see debug03 at all")

  // now creating a logfile instead of stdout

  logmsg.SetLogFile("demo.log")

  logmsg.Print(logmsg.Info, "Your message should be in the file");
  printDebugMsgs()

  // this is a demo of having the right line number to earlier 
  // in the stack with printing a message

  demoWhereCalled()

}
