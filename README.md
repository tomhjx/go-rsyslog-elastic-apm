# go-rsyslog-elastic-apm
接受rsyslog转发至elastic apm server

# 功能

* 以本地应用的方式启动，以标准输入的形式接受rsyslog日志内容，转发到elastic apm server
* 以远程服务的方式启动，伪装为rsyslog server，以TCP/UDP的形式接受rsyslog客户端发来的日志内容，转发到elastic apm server
