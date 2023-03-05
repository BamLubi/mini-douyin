#!/usr/bin/env bash

# 启动服务
cd cmd/user
sh build.sh
nohup sh output/bootstrap.sh > output/nohup.log 2>&1 &
pid_user=`ps -ef | grep userservice | grep -v grep | awk 'NR==1 {print $2}'`
echo "run userservice pid ${pid_user}"

cd ../video
sh build.sh
nohup sh output/bootstrap.sh > output/nohup.log 2>&1 &
pid_video=`ps -ef | grep videoservice | grep -v grep | awk 'NR==1 {print $2}'`
echo "run videoservice pid ${pid_video}"

cd ../socity
sh build.sh
nohup sh output/bootstrap.sh > output/nohup.log 2>&1 &
pid_socity=`ps -ef | grep socityservice | grep -v grep | awk 'NR==1 {print $2}'`
echo "run socityservice pid ${pid_socity}"

# 等待用户输入以关闭客户端
while true; do
    read -p "输入'Y'以关闭所有微服务: " input
    if [ "$input" = "Y" ]; then
        echo "stop userservice pid ${pid_user}"
        kill ${pid_user}

        echo "stop videoservice pid ${pid_video}"
        kill ${pid_video}

        echo "stop socityservice pid ${pid_socity}"
        kill ${pid_socity}
        break
    fi
done