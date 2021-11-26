VUE_APP_BASE_URL = 'http://sincos.vip'


build:
	git clone https://gitee.com/Azerrroth/frontend-for-dorm-selection
	export VUE_APP_BASE_URL=$(VUE_APP_BASE_URL)
	cd frontend-for-dorm-selection && npm --registry https://registry.npm.taobao.org install && npm run build
	cp -r frontend-for-dorm-selection/dist .

copy:
	cp -r user/* dorm/
	cp -r user/* login/
	cp -r user/* register/
	cp -r user/* order/

clean:
	rm -rf dorm/*
	rm -rf login/*
	rm -rf register/*
	rm -rf order/*