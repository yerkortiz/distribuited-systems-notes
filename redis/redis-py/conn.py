import redis

r = redis.Redis(host='localhost', port=6379, db=0)

r.set('hello', 'world') 
value = r.get('hello')
print(value) 