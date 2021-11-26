VUE_APP_BASE_URL = 'http://sincos.vip:8080'


build:
	git clone https://gitee.com/Azerrroth/frontend-for-dorm-selection
	export VUE_APP_BASE_URL=$(VUE_APP_BASE_URL)
	cd frontend-for-dorm-selection && npm --registry https://registry.npm.taobao.org install && npm run build
	cp -r frontend-for-dorm-selection/dist/* .