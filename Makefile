.PHONY: build

apply: create_namespace
	helmfile --file deploy/helmfile.kafka.yaml $(MAKE_HELM_OPTIONS) apply

destroy:
	helmfile --file deploy/helmfile.kafka.yaml $(MAKE_HELM_OPTIONS) destroy

update_rules:
	cd make && python3 update_rules.py

update_engine:
	cd deploy && helmfile -l name=sigma-engine destroy && helmfile -l name=sigma-engine apply

build:
	cd build && docker-compose build --no-cache $(image) 

create_namespace:
	bash make/create_namespace.sh || true

test: create_namespace
	bash make/test.sh
