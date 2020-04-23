import json
import sys

with open('/tmp/test.log', 'a') as fp:
    fp.write(sys.stdin.read() + '\n')

print('{"test1": "test2"}')