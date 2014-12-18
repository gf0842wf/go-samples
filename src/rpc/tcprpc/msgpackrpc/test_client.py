from mprpc import RPCClient

client = RPCClient('127.0.0.1', 6000)
print client.call('echo', 1000, 100000000000000)
