why ? defer resp.Body.Close() or many tcp or http or grpc connections require us to do defer ...Close() ?


Recently while looking into this question, I came across a new term called FD(File descriptors)

In Go you usually do this in a http call

resp, err := http.Get("url.com")
if err != nil {
    return fmt....
}
defer resp.Body.Close()

when we do this kind of calls we basically use from an default trasnport config which is the following:

MaxIdleConns        = 100
MaxIdleConnsPerHost = 2
MaxConnsPerHost     = 0 

and with an unlimited total connection. 

Operating system like unix has an utility called file descriptors, what basically happens is this whenever we do an operation like following:
file read, tcp conn, http conn, grpc conn

Os will assign a file descriptor to that opened operation a basic number.

In go usually if we dont use close operation that connection will hang and if many of this same operation created we will have a memory leak or fd leak