# goredis

Redis有单机、哨兵、集群3种部署方式。不同部署，使用的Redis库不一样，造成一定的不便。

因此，本库集成网络上一些比较流行的库，使之能够在3种部署方式下，使用统一的API。

本库只是简单的把redigo、go-sentinel、redisc做下集成。

### 依赖的库

https://github.com/garyburd/redigo

https://github.com/FZambia/go-sentinel

https://github.com/mna/redisc


### TODO:

  - 自识别3种redis部署方式，并自动创建对应的访问实例
