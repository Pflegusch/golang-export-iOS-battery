# Description
A program written in C that extracts current battery information from a jailbroken iOS device that has bstat installed. The tool works by using ssh to get into the device and running the above program, the return value of that gets written into a local file that one can further use to regex the desired values out of. This is the go equivalent to the C implementation, only needs work to make it work for docker too.

# Dependencies
Install the dependencies with ``go get github.com/melbahja/goph``

# Build and Run
Build the webserver with ``go build webserver`` and the program with ``go build battery``. Then run the webserver with ``./webserver`` and the programm with ``./battery <ip-addr> <path-to-private-sshkey>``. Make sure to have a directory called ``static`` in the battery binary root with a file named ``data.txt`` in it. 