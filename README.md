# logmsg
Dynamic log messaging levels for a go application

Allows for a go application to have log messaging that range from debug to critical events.  These messages can be dynamically turned on/off (supressed).  Allows for including of detailed diagnostic messages in production code and not impact production, but can be turned on when required.

This is just the package to provide the logging.  It will be up to the consuming application to implement the appropiate listener to invoke the logmsg.SetLogLevel() method and change the logging level.


