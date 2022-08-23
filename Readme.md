# how to gen pb code

## win10

```
go install github.com/golang/protobuf/protoc-gen-go@v1.4.3
download protoc from https://github.com/protocolbuffers/protobuf/releases

```

## linux

```
wget https://dl.google.com/go/go1.18.1.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.18.1.linux-amd64.tar.gz
echo "PATH=$PATH:/usr/local/go/bin" >> ~/.profile
source ~/.profile

curl -fsSL https://deb.nodesource.com/setup_14.x | sudo -E bash -
sudo apt-get install -y nodejs
go install github.com/golang/protobuf/protoc-gen-go@v1.4.3


// install protoc

## clone pro
git clone --recursive https://gitee.com/jeremyjj/protobuf.git 
​
## 安装依赖 
sudo apt-get install autoconf automake libtool curl make g++ unzip libffi-dev -y

cd protobuf
​
./autogen.sh 
​
./configure 
​
sudo make -j32 
​
sudo make install 
​
## 刷新共享库 
sudo ldconfig
```