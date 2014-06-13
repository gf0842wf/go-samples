# -*- coding: utf-8- -*-

import socket
import struct

try:
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
except socket.error as e:
    print repr(e)

remote_addr = ("127.0.0.1", 8888)
sock.connect(remote_addr)         

msg = "abcdef"
length = len(msg)
data = struct.pack(">I%ds"%length, length, msg)
print repr(data)
# '\x00\x00\x00\x06abcdef'
sock.sendall(data*3)                    

buf = sock.recv(1024)
print buf

sock.close()
