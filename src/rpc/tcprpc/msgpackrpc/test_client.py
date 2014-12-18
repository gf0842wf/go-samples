from mprpc import RPCClient

client = RPCClient('127.0.0.1', 6000)
print client.call('add', 1000, 1000)
