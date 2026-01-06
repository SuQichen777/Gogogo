# PLG
PLG 是 Promtail、Loki、Grafana 的组合。
Promtail: 负责采集应用日志并打标签，推送到 Loki。
Loki: 负责存储与索引日志数据，提供查询接口。
Grafana: 负责可视化与查询日志。
## 内容解析
- `app.py`, `Dockerfile`, `logs`: 应用输出 JSON 日志，容器运行应用，`logs` 通过挂载持久化 `/var/log/app.log`。
- `promtail-config.yaml`: 采集 `/var/log/app.log`，打上 `coffee-web` 标签并推送到 Loki。
- `loki-config.yaml`: 单节点 Loki，使用文件系统存储与 boltdb-shipper 索引，关闭鉴权。
- `docker-compose.yaml`: 编排 app、promtail、loki、grafana，并暴露 3100/3000 端口。

## 使用步骤
在 `PLG-MiniLab` 目录下执行 `docker compose up -d`，确保 app、Promtail、Loki、Grafana 成功运行。然后访问：
```
http://localhost:3000
```
进入 Grafana 管理页面，在 Connection 中添加 Loki，链接为 `http://loki:3100`。

### Web可视化
- Promtail: http://localhost:9080/
- Grafana: http://localhost:3000

这是在对应yaml配置中选择监听端口和在`docker-compose.yaml`中转发端口的结果。如果需要改变监听端口，可以在对应文件中改变yaml配置，按需选择`docker-compose up -d --force-recreate [name]`