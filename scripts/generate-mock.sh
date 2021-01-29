#!/bin/bash

mockgen -source=internal/business/TaskRepository.go -destination=internal/mocks/TaskRepositoryMock.go -package=mocks

echo "Mock objects generated."