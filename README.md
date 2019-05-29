# crawl_movie
基于beego框架的豆瓣电影信息爬虫项目

## 目录结构描述  
├── conf                    // 配置文件  
├── controllers                       
│   ├── crawlMovie.go      //进行豆瓣电影的爬取       
├── models           
│   ├── movie_info.go    //创建关于电影的结构体，和各种爬取电影信息的方法  
│   ├── redis.go    //将爬取过的电影页面存入redis数据库中     
├── routers                        
│   ├── router.go   //创建路由               
├── main.go   //主程序                          
├── README.md   

## 项目详情描述
利用beego框架爬取豆瓣电影网页关于电影的信息，并将其存入到mysql数据库中，将爬取过的网页做好标记，存入redis数据库中，防止二次爬取。  

           
