# golang-export-iOS-battery
A program written in Go that extracts current battery information from a jailbroken iOS device that has bstat installed. The tool works by using ssh to get into the device and running the above program, the return value of that gets written into a local file that one can further use to regex the desired values out of. This is the go equivalent to the C implementation, only needs work to make it work for docker too.
