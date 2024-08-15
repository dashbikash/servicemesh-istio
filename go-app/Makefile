# Makefile

# Variables

DOCKER_IMAGE_NAME_DASHBOARD = bikashpdash/goapp-dashboard
DOCKER_IMAGE_NAME_INVENTORY = bikashpdash/goapp-inventory
DOCKER_IMAGE_NAME_SALES = bikashpdash/goapp-sales

DASHBOARD_PORT=8080
INVENTORY_PORT=8081
SALES_PORT=8082

docker-build:
	docker build -f Dockerfile.dashboard -t $(DOCKER_IMAGE_NAME_DASHBOARD) .
	# docker build -f Dockerfile.inventory -t $(DOCKER_IMAGE_NAME_INVENTORY) .
	# docker build -f Dockerfile.sales -t $(DOCKER_IMAGE_NAME_SALES) .

docker-push: 
	docker login -u bikashpdash
	docker push $(DOCKER_IMAGE_NAME)





