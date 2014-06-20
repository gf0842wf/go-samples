# -*- coding: utf-8- -*-

import socket
import struct
import json
import time

from array import array

# 加密解密因子(大小范围要选好, 然后双方协定好)
M1 = 1 << 19
IA1 = 2 << 20
IC1 = 3 << 21

def encrypt(data, encrypt_key=1):
    """加密解密
    >>> encrypt("abc", 2)
    'kde'
    >>> encrypt("kde", 2)
    'abc'
    >>> 
    """
    raw_data = array("B", data)
    for i in xrange(len(raw_data)):
        encrypt_key = IA1 * (encrypt_key % M1) + IC1
        raw_data[i] ^= (encrypt_key >> 20 & 0xff)
    return raw_data.tostring()


try:
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
except socket.error as e:
    print repr(e)

remote_addr = ("127.0.0.1", 8888)
sock.connect(remote_addr)         

msg = {"kind":"SYS", "type":"PRESHAKE"}
msg = json.dumps(msg)
length = len(msg)
data = struct.pack(">I%ds"%length, length, msg)
# print repr(data)
# '\x00\x00\x00\x06abcdef'
sock.sendall(data)                    

while True:
    buf = sock.recv(1024)
    print repr(buf)
    # print buf[4:]
    # print struct.unpack(">I", buf[:4])
    print json.loads(buf[4:])

    data = json.dumps({"kind":"SYS", "type":"ACKSHAKE", "result":{"code":0, "message":'0k'}})
    length = len(data)
    data = struct.pack(">I%ds"%length, length, data)
    sock.sendall(data)
    time.sleep(5)

sock.close()
