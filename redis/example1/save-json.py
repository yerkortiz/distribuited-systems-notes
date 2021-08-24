import redis
import json

r = redis.Redis(host='localhost', port=6379, db=0)

products= [
    {'name':'cpu', 'stock':2},
    {'name':'necronomicon', 'stock':3},
    {'name':'secret-product', 'stock':4},
]
json_images = json.dumps(products)
r.set('images', json_images)

# Read saved JSON str from Redis and unpack into python dict
unpacked_images = json.loads(r.get('images'))
products == unpacked_images
print(products)
