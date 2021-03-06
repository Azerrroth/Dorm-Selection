VUE_APP_BASE_URL = http://127.0.0.1


build:
	echo "Start building frontent pages"
	-git clone https://gitee.com/Azerrroth/frontend-for-dorm-selection
	export VUE_APP_BASE_URL=$(VUE_APP_BASE_URL) && cd frontend-for-dorm-selection && npm --registry https://registry.npm.taobao.org install && npm run build
	cp -r frontend-for-dorm-selection/dist .

serve:
	-git clone https://gitee.com/Azerrroth/Dorm-selection-mid-layer.git node-server
	docker-compose up

sync:
	git clone https://gitee.com/Azerrroth/Dorm-selection-backend.git user
	rm -rf user/config
	rm -rf user/.git
	cp -r user/* dorm/
	cp -r user/* login/
	cp -r user/* register/
	cp -r user/* order/

clean:
	-docker-compose down
	rm -rf user/*
	rm -rf dorm/*
	rm -rf login/*
	rm -rf register/*
	rm -rf order/*
	-rm -rf node-server
	-rm -rf frontend-for-dorm-selection

cleanall:
	make clean
	-docker rmi $$(docker images | grep "ms-" | awk '{print $1}')
