# 20200605go_demo2
基于nats案例
项目介绍：通过nats来接收和发送数据，将接收到的数据进行处理，并根据tag_id去redis中获取对应的data_id，并将datas中很多数据拆分为一条一条的数据，存放到mysql中。
