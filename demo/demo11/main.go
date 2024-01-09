package main

import "regexp"

func main() {

	/*

	- [ ]  在主机K8S上，K8S集群证书过期。
	- [ ]  主机centos，系统有未使用的块设备，处理成数据盘（假设该盘未来用于mysql应用）。
	    - [ ]  将主机ubuntu上的文件/data/bigfile同步到主机centos。
	    - [ ]  如果需要在文件/data/bigfile的第一行插入一条数据，描述处理过程。




	mv config config.old
	cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
	chown $(id -u):$(id -g) $HOME/.kube/config
	sudo chmod 644 $HOME/.kube/config


	docker ps | grep -v pause | grep -E "etcd|scheduler|controller|apiserver" | awk '{print $1}' | awk '{print "docker","restart",$1}' | bash




	 */

	//版本1.4.5-0（来自1.2.3.4版本）
	// ^\d
	regexp.Compile("/d")
}
