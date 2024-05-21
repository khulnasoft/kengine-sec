PWD=$(shell pwd)

export ROOT_MAKEFILE_DIR=$(shell pwd)
export KHULNASOFT_AGENT_DIR=$(PWD)/khulnasoft_agent
export KHULNASOFT_ROUTER_DIR=$(PWD)/haproxy
export KHULNASOFT_TELEMETRY_DIR=$(PWD)/khulnasoft_telemetry
export KHULNASOFT_FILE_SERVER_DIR=$(PWD)/khulnasoft_file_server
export KHULNASOFT_FRONTEND_DIR=$(PWD)/khulnasoft_frontend
export SECRET_SCANNER_DIR=$(KHULNASOFT_AGENT_DIR)/plugins/SecretScanner
export MALWARE_SCANNER_DIR=$(KHULNASOFT_AGENT_DIR)/plugins/YaraHunter/
export PACKAGE_SCANNER_DIR=$(KHULNASOFT_AGENT_DIR)/plugins/package-scanner
export COMPLIANCE_SCANNER_DIR=$(KHULNASOFT_AGENT_DIR)/plugins/compliance
export CLOUD_SCANNER_DIR=$(KHULNASOFT_AGENT_DIR)/plugins/cloud-scanner
export KHULNASOFT_CTL=$(PWD)/khulnasoft_ctl
export KHULNASOFTD=$(PWD)/khulnasoft_bootstrapper
export KHULNASOFT_FARGATE_DIR=$(KHULNASOFT_AGENT_DIR)/agent-binary
export IMAGE_REPOSITORY?=docker.io/khulnasoft
export KE_IMG_TAG?=latest
export STEAMPIPE_IMG_TAG?=latest
export IS_DEV_BUILD?=false
export VERSION?=v2.2.1
export AGENT_BINARY_BUILD=$(KHULNASOFT_FARGATE_DIR)/build
export AGENT_BINARY_BUILD_RELATIVE=khulnasoft_agent/agent-binary/build
export AGENT_BINARY_DIST=$(KHULNASOFT_FARGATE_DIR)/dist
export AGENT_BINARY_DIST_RELATIVE=khulnasoft_agent/agent-binary/dist
export AGENT_BINARY_FILENAME=khulnasoft-agent-$(shell dpkg --print-architecture)-$(VERSION).tar.gz

default: bootstrap console_plugins agent console fargate-local

.PHONY: console
console: redis postgres kafka-broker router server worker ui file-server graphdb jaeger

.PHONY: console_plugins
console_plugins: secretscanner malwarescanner packagescanner compliancescanner

.PHONY: bootstrap
bootstrap:
	./bootstrap.sh

.PHONY: alpine_builder
alpine_builder:
	docker build --tag=$(IMAGE_REPOSITORY)/khulnasoft_builder_ce:$(KE_IMG_TAG) -f docker_builders/Dockerfile-alpine .

.PHONY: debian_builder
debian_builder:
	docker build --build-arg KE_IMG_TAG=${KE_IMG_TAG} --build-arg IMAGE_REPOSITORY=${IMAGE_REPOSITORY} --tag=$(IMAGE_REPOSITORY)/khulnasoft_glibc_builder_ce:$(KE_IMG_TAG) -f docker_builders/Dockerfile-debian .

.PHONY: bootstrap-agent-plugins
bootstrap-agent-plugins:
	(cd $(KHULNASOFT_AGENT_DIR)/plugins && make localinit)
	(cd $(PACKAGE_SCANNER_DIR) && bash bootstrap.sh)
	(cd $(SECRET_SCANNER_DIR) && bash bootstrap.sh)
	(cd $(MALWARE_SCANNER_DIR) && bash bootstrap.sh)
	(cd $(CLOUD_SCANNER_DIR) && bash bootstrap.sh)

.PHONY: agent
agent: debian_builder khulnasoftd console_plugins
	(cd $(KHULNASOFT_AGENT_DIR) &&\
	IMAGE_REPOSITORY=$(IMAGE_REPOSITORY) KE_IMG_TAG=$(KE_IMG_TAG) VERSION=$(VERSION) bash build.sh)

.PHONY: agent-binary
agent-binary: agent agent-binary-tar

.PHONY: agent-binary-tar
agent-binary-tar:
	mkdir -p $(AGENT_BINARY_DIST) $(AGENT_BINARY_BUILD)
	ID=$$(docker create $(IMAGE_REPOSITORY)/khulnasoft_agent_ce:$(KE_IMG_TAG)); \
	(cd $(KHULNASOFT_FARGATE_DIR) &&\
	CONTAINER_ID=$$ID VERSION=$(VERSION) AGENT_BINARY_BUILD=$(AGENT_BINARY_BUILD) AGENT_BINARY_DIST=$(AGENT_BINARY_DIST) AGENT_BINARY_FILENAME=$(AGENT_BINARY_FILENAME) bash copy-bin-from-agent.sh); \
	docker rm -v $$ID

.PHONY: fargate-local
fargate-local: agent-binary-tar
	(cd $(KHULNASOFT_AGENT_DIR) &&\
	IMAGE_REPOSITORY=$(IMAGE_REPOSITORY) KE_IMG_TAG=$(KE_IMG_TAG) VERSION=$(VERSION) AGENT_BINARY_BUILD_RELATIVE=$(AGENT_BINARY_BUILD_RELATIVE) AGENT_BINARY_FILENAME=$(AGENT_BINARY_FILENAME) bash build-fargate-local-bin.sh)

.PHONY: fargate
fargate:
	mkdir -p $(AGENT_BINARY_BUILD)
	(cd $(KHULNASOFT_AGENT_DIR) &&\
	IMAGE_REPOSITORY=$(IMAGE_REPOSITORY) KE_IMG_TAG=$(KE_IMG_TAG) VERSION=$(VERSION) AGENT_BINARY_BUILD=$(AGENT_BINARY_BUILD) AGENT_BINARY_BUILD_RELATIVE=$(AGENT_BINARY_BUILD_RELATIVE) bash build-fargate.sh)

.PHONY: khulnasoftd
khulnasoftd: alpine_builder bootstrap bootstrap-agent-plugins
	(cd $(KHULNASOFTD) && make prepare)
	cp $(KHULNASOFTD)/khulnasoft_bootstrapper $(KHULNASOFT_AGENT_DIR)/khulnasoftd

.PHONY: redis
redis:
	(cd khulnasoft_redis && docker build --tag=$(IMAGE_REPOSITORY)/khulnasoft_redis_ce:$(KE_IMG_TAG) .)

.PHONY: postgres
postgres:
	docker build --tag=$(IMAGE_REPOSITORY)/khulnasoft_postgres_ce:$(KE_IMG_TAG) -f khulnasoft_postgres/Dockerfile ./khulnasoft_postgres

.PHONY: kafka-broker
kafka-broker:
	docker build -t $(IMAGE_REPOSITORY)/khulnasoft_kafka_broker_ce:$(KE_IMG_TAG) -f ./khulnasoft_kafka/kafka-broker-Dockerfile ./khulnasoft_kafka

.PHONY: router
router:
	docker build --build-arg is_dev_build=$(IS_DEV_BUILD) -t $(IMAGE_REPOSITORY)/khulnasoft_router_ce:$(KE_IMG_TAG) $(KHULNASOFT_ROUTER_DIR)

.PHONY: file-server
file-server:
	docker build -t $(IMAGE_REPOSITORY)/khulnasoft_file_server_ce:$(KE_IMG_TAG) $(KHULNASOFT_FILE_SERVER_DIR)

.PHONY: server
server: debian_builder
	(cd ./khulnasoft_server && VERSION=$(VERSION) make image)

.PHONY: worker
worker: debian_builder agent-binary-tar
	(cd ./khulnasoft_worker && VERSION=$(VERSION) AGENT_BINARY_DIST_RELATIVE=$(AGENT_BINARY_DIST_RELATIVE) make image)

.PHONY: jaeger
jaeger:
	docker build -t $(IMAGE_REPOSITORY)/khulnasoft_telemetry_ce:$(KE_IMG_TAG) $(KHULNASOFT_TELEMETRY_DIR)

.PHONY: graphdb
graphdb:
	docker build -f ./khulnasoft_neo4j/Dockerfile --build-arg IMAGE_REPOSITORY=$(IMAGE_REPOSITORY) --build-arg KE_IMG_TAG=$(KE_IMG_TAG) -t $(IMAGE_REPOSITORY)/khulnasoft_neo4j_ce:$(KE_IMG_TAG) ./khulnasoft_neo4j

.PHONY: ui
ui:
	git log --format="%h" -n 1 > $(KHULNASOFT_FRONTEND_DIR)/console_version.txt && \
	echo $(VERSION) > $(KHULNASOFT_FRONTEND_DIR)/product_version.txt && \
	docker run --rm --entrypoint=bash -v $(KHULNASOFT_FRONTEND_DIR):/app node:18-bullseye-slim -c "cd /app && corepack enable && corepack prepare pnpm@7.17.1 --activate && PLAYWRIGHT_SKIP_BROWSER_DOWNLOAD=true pnpm install --frozen-lockfile --prefer-offline && ENABLE_ANALYTICS=true pnpm run build" && \
	docker build -f $(KHULNASOFT_FRONTEND_DIR)/Dockerfile -t $(IMAGE_REPOSITORY)/khulnasoft_ui_ce:$(KE_IMG_TAG) $(KHULNASOFT_FRONTEND_DIR) && \
	rm -rf $(KHULNASOFT_FRONTEND_DIR)/console_version.txt $(KHULNASOFT_FRONTEND_DIR)/product_version.txt

.PHONY: secretscanner
secretscanner: bootstrap-agent-plugins
	docker build --tag=$(IMAGE_REPOSITORY)/khulnasoft_secret_scanner_ce:$(KE_IMG_TAG) -f $(SECRET_SCANNER_DIR)/Dockerfile $(SECRET_SCANNER_DIR)

.PHONY: malwarescanner
malwarescanner: bootstrap-agent-plugins
	docker build --tag=$(IMAGE_REPOSITORY)/khulnasoft_malware_scanner_ce:$(KE_IMG_TAG) -f $(MALWARE_SCANNER_DIR)/Dockerfile $(MALWARE_SCANNER_DIR)

.PHONY: packagescanner
packagescanner: bootstrap-agent-plugins
	docker build --tag=$(IMAGE_REPOSITORY)/khulnasoft_package_scanner_ce:$(KE_IMG_TAG) -f $(PACKAGE_SCANNER_DIR)/Dockerfile $(PACKAGE_SCANNER_DIR)

.PHONY: compliancescanner
compliancescanner:
	docker build --tag=$(IMAGE_REPOSITORY)/khulnasoft_compliance_scanner_ce:$(KE_IMG_TAG) -f $(COMPLIANCE_SCANNER_DIR)/Dockerfile $(COMPLIANCE_SCANNER_DIR)

.PHONY: cloudscanner
cloudscanner: debian_builder khulnasoftd 
	(cd $(KHULNASOFT_AGENT_DIR) &&\
	IMAGE_REPOSITORY=$(IMAGE_REPOSITORY) KE_IMG_TAG=$(KE_IMG_TAG) VERSION=$(VERSION) bash build_cs_agent.sh)

.PHONY: openapi
openapi: server
	docker run --rm \
	--entrypoint=/usr/local/bin/khulnasoft_server \
	-v $(PWD):/app $(IMAGE_REPOSITORY)/khulnasoft_server_ce:$(KE_IMG_TAG) \
	--export-api-docs-path /app/openapi.yaml

	rm -rf golang_khulnasoft_sdk/client/*

	docker pull openapitools/openapi-generator-cli:latest
	docker run --rm \
	-v $(PWD):/local openapitools/openapi-generator-cli:latest generate \
	-i /local/openapi.yaml \
	-g go \
	-o /local/golang_khulnasoft_sdk/client \
	-p isGoSubmodule=true \
	-p packageName=client \
	--git-repo-id golang_khulnasoft_sdk \
	--git-user-id khulnasoft

	rm openapi.yaml
	cd $(PWD)/golang_khulnasoft_sdk/client && rm -rf ./test && sed -i 's/go 1.18/go 1.20/g' go.mod && go mod tidy -v && cd -

.PHONY: cli
cli: bootstrap
	(cd $(KHULNASOFT_CTL) && make clean && make all)

.PHONY: publish
publish: publish-redis publish-postgres publish-kafka publish-router publish-file-server publish-server publish-worker publish-ui publish-agent publish-cluster-agent publish-packagescanner publish-secretscanner publish-malwarescanner publish-graphdb publish-jaeger

.PHONY: publish-redis
publish-redis:
	docker push $(IMAGE_REPOSITORY)/khulnasoft_redis_ce:$(KE_IMG_TAG)

.PHONY: publish-postgres
publish-postgres:
	docker push $(IMAGE_REPOSITORY)/khulnasoft_postgres_ce:$(KE_IMG_TAG)

.PHONY: publish-kafka
publish-kafka:
	docker push $(IMAGE_REPOSITORY)/khulnasoft_kafka_broker_ce:$(KE_IMG_TAG)

.PHONY: publish-router
publish-router:
	docker push $(IMAGE_REPOSITORY)/khulnasoft_router_ce:$(KE_IMG_TAG)

.PHONY: publish-file-server
publish-file-server:
	docker push $(IMAGE_REPOSITORY)/khulnasoft_file_server_ce:$(KE_IMG_TAG)

.PHONY: publish-server
publish-server:
	docker push $(IMAGE_REPOSITORY)/khulnasoft_server_ce:$(KE_IMG_TAG)

.PHONY: publish-worker
publish-worker:
	docker push $(IMAGE_REPOSITORY)/khulnasoft_worker_ce:$(KE_IMG_TAG)

.PHONY: publish-ui
publish-ui:
	docker push $(IMAGE_REPOSITORY)/khulnasoft_ui_ce:$(KE_IMG_TAG)

.PHONY: publish-agent
publish-agent:
	docker push $(IMAGE_REPOSITORY)/khulnasoft_agent_ce:$(KE_IMG_TAG)

.PHONY: publish-cluster-agent
publish-cluster-agent:
	docker push $(IMAGE_REPOSITORY)/khulnasoft_cluster_agent_ce:$(KE_IMG_TAG)

.PHONY: publish-packagescanner
publish-packagescanner:
	docker push $(IMAGE_REPOSITORY)/khulnasoft_package_scanner_ce:$(KE_IMG_TAG)

.PHONY: publish-secretscanner
publish-secretscanner:
	docker push $(IMAGE_REPOSITORY)/khulnasoft_secret_scanner_ce:$(KE_IMG_TAG)

.PHONY: publish-malwarescanner
publish-malwarescanner:
	docker push $(IMAGE_REPOSITORY)/khulnasoft_malware_scanner_ce:$(KE_IMG_TAG)

.PHONY: publish-graphdb
publish-graphdb:
	docker push $(IMAGE_REPOSITORY)/khulnasoft_neo4j_ce:$(KE_IMG_TAG)

.PHONY: publish-jaeger
publish-jaeger:
	docker push $(IMAGE_REPOSITORY)/khulnasoft_telemetry_ce:$(KE_IMG_TAG)

.PHONY: clean
clean:
	-(cd $(KHULNASOFT_AGENT_DIR) && make clean)
	-(cd $(KHULNASOFT_FARGATE_DIR) && rm -rf khulnasoft-agent-bin-$(VERSION)*)
	-(cd $(ROOT_MAKEFILE_DIR)/khulnasoft_server && make clean)
	-(cd $(ROOT_MAKEFILE_DIR)/khulnasoft_worker && make clean)
	-(cd $(KHULNASOFTD) && make clean && rm $(KHULNASOFT_AGENT_DIR)/khulnasoftd)
	-rm -rf $(AGENT_BINARY_DIST)/* $(AGENT_BINARY_BUILD)/*