FRONTEND_PATH = $(PWD)/frontend
BACKEND_PATH = $(PWD)/backend

dev:
	@if [ -d "$(FRONTEND_PATH)" ]; then cd $(FRONTEND_PATH) && yarn dev; fi


