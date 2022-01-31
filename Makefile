FRONTEND_PATH = $(PWD)/frontend
BACKEND_PATH = $(PWD)/backend

dev-fe:
	@if [ -d "$(FRONTEND_PATH)" ]; then cd $(FRONTEND_PATH) && yarn dev; fi


