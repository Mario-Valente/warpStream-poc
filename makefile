install:
	@echo "Installing dependencies..."
	chmod +x ./scripts/init_run.sh
	./scripts/init_run.sh

deploy:
	@echo "Deploying application..."
	chmod +x ./scripts/deploy_apps.sh
	./scripts/deploy_apps.sh


