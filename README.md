#位置公共服务
与业务解耦的一些公共服务，GO语言实现

##环境部署
    sudo apt-get install golang
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


