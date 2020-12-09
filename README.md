# simple-docker



一、chapter2 linux Namespace 的6中类型类型

1、MNT Namespace  提供磁盘挂载点和文件系统的隔离能力

2、IPC Namespace  提供进程间通信的隔离能力
    进程间通信主要指（共享内存，信号量，消息队列）

3、Net Namespace  提供网路隔离能力

4、UTS Namespace 提供主机名hostname隔离能力

5、PID Namespace 提供进程隔离能力

6、User Namespace 提供用户隔离能力 （一个进程的User ID 和Group ID 在User Namespace 内外可以是不同的）

二、chapter3  使用Cgroup 构造简单的容器，具备基本的Namespace隔离，并且确定了基本的开发架构

1、构造容器基本流程
创建容器 ——> run ——> 配置namespace 创建父进程 ——> init 
 ——> 挂载文件系统替换 init进程 ——> 完成容器创建
