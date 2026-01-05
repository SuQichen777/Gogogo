# 核心app
import time
import random
import json

while True:
    log = {
        "service": "coffee-web",
        "level": "info",
        "msg": "order created"
    }

    # 有30%的概率会出错

    if random.random() < 0.3:
        log["level"] = "error"
        log["msg"] = "payment timeout"
    
    print(json.dumps(log), flush=True)
    time.sleep(1)