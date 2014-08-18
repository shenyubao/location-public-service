#搜车公共服务
与业务解耦的一些公共服务，GO语言实现

##环境部署
    访问[下载地址](http://golang.org/dl/),版本选择V1.1以上,假定你想安装的目录为`/opt/`.
    解压缩tar.gz到安装目录下:`sudo tar -C /opt/ -xzf go$VERSION.$OS-$ARCH.tar.gz`.
    比如你下载的是go1.2.1.linux-amd64.tar.gz,执行`sudo tar -C /opt/ -xzf go1.2.1.linux-amd64.tar.gz`.
    设置PATH,`export GOROOT=/opt/go`,`export PATH=$PATH:$GOROOT/bin`
    go get github.com/mattn/go-sqlite3
    go get github.com/beego/bee
    go get github.com/astaxie/beego
##项目启动
    cd souche-public-service
    bee run
    
##接口说明
###1. 根据手机号获取归属地
    URL： /Mobile/GetLocation

    PARAM:
    * mobile  get  notnull  15903060489
    
    RETURN:
    {
      "Code": 100000,
      "Success": true,
      "Message": "success",
      "Data": {
        "Mobile": "1590306",
        "Province": "河南",
        "City": "新乡",
       "Opera": "移动",
       "Fullname": "河南新乡[移动]"
     }
    }


