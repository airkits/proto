# how to gen pb code

## win10

```
go get -u google.golang.org/protobuf/protoc-gen-go
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
go install google.golang.org/protobuf/protoc-gen-go


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