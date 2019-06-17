# go-rsyslog-elastic-apm
从[rsyslog](https://www.rsyslog.com)接收日志，并转发至[Elastic APM Server](https://www.elastic.co/guide/en/apm/server/current/setting-up-and-running.html)

# 功能

* 以本地应用的方式启动，以标准输入的形式接受rsyslog日志内容，转发到[Elastic APM Server](https://www.elastic.co/guide/en/apm/server/current/setting-up-and-running.html)
* 以远程服务的方式启动，伪装为rsyslog server，以TCP/UDP的形式接受rsyslog客户端发来的日志内容，转发到[Elastic APM Server](https://www.elastic.co/guide/en/apm/server/current/setting-up-and-running.html)
