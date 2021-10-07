TAG := $(git log -1 --pretty=%!H)

run: 
#docker build -t ${TAG} .
#docker run -p 8081:8081 ${TAG}
	@echo hello