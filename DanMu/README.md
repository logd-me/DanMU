这个项目，运用了微服务的思想，通过Redis做消息队列，当用户发送弹幕时，通过http发送给admin，admin进行推送，每个连接了的ws是订阅了该视频的推送，收到及时的弹幕。同时，当推送成功后，还会将弹幕压入缓存，当用户登录的时候就可以看见这些历史弹幕。（最近在实习，太忙了，基本实现）

以下是实验代码的实现（不是该源代码）

1.当用户登录：查看到历史弹幕：

![image-20210802170950713](C:\Users\HUAWEI\AppData\Roaming\Typora\typora-user-images\image-20210802170950713.png)

2.当用户发送弹幕，通过redis的publish来实现

![image-20210802171307018](C:\Users\HUAWEI\AppData\Roaming\Typora\typora-user-images\image-20210802171307018.png)