# **********************************************************
# * Author           : ricky97gr
# * Email            : forgocode@163.com
# * Github           : https://github.com/ricky97gr
# * Create Time      : 2024-03-14 15:50
# * FileName         : prepare.sh
# * Description      : 
# **********************************************************

if [[ -z $(which docker) ]]; then
	echo "docker is not existed, must install it!"
	exit 0
else
	echo "docker is existed"
fi

check_image(){
	if [[ -z $(docker image list | grep mysql) ]];then
		echo "image mysql is not existed"
		pull_mysql
	else
		echo "image mysql is existed"
	fi
	if [[ -z $(docker image list | grep redis) ]];then
		echo "image redis is not existed"
		pull_redis
	else
		echo "image redis is existed"
	fi
	if [[ -z $(docker image list | grep nginx) ]];then
  		echo "image nginx is not existed"
  		pull_redis
  else
  		echo "image nginx is existed"
  fi

  if [[ -z $(docker image list | grep mongo) ]];then
    		echo "image mongo is not existed"
    		pull_redis
    else
    		echo "image redis is existed"
    fi
}

pull_mysql(){
	echo "start pull bitnami/mysql:latest"
	docker pull bitnami/mysql
	echo "pull bitnami/mysql:latest successfully!"
}

pull_redis(){
	echo "start pull bitnami/redis:latest"
	docker pull bitnami/redis
	echo "pull bitnami/redis:latest successfully!"
}

pull_mongo(){
	echo "start pull bitnami/mongodb:latest"
	docker pull bitnami/mongodb
	echo "pull bitnami/mongodb:latest successfully!"
}

pull_nginx(){
	echo "start pull bitnami/nginx:latest"
	docker pull bitnami/nginx
	echo "pull bitnami/nginx:latest successfully!"
}

run_mysql(){
	id=$(docker image list | grep mysql | awk -F ' ' '{print $3}')
	docker run --restart always --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -e TZ=Asiz/Shanghai -e MYSQL_DATABASE=test -d $id
}

run_redis(){
	id=$(docker image list | grep redis | awk -F ' ' '{print $3}')
	docker run --restart always --name redis -p 6379:6379  -e ALLOW_EMPTY_PASSWORD=yes -d $id 
}

check_image
run_mysql
run_redis

echo
echo "test env is ready"
echo "mysql:"
echo -e "database: \e[32mtest\e[0m"
echo -e "user:     \e[32mroot\e[0m"
echo -e "passwd:   \e[32m123456\e[0m"
echo 
echo "redis:"
echo -e "redis no password"
