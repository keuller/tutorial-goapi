#!/bin/bash

wrk -d30s -t50 -c500 http://localhost:8080/v1/tasks
