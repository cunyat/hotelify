#!/bin/bash

pushd internal
pushd rooms
mockery --case=snake --outpkg=storagemocks --output=adapters/storagemocks --name=Repository  --dir domain/room/
popd
pushd common
mockery --case=snake --outpkg=commandmocks --output=adapters/commandmocks --name=Bus --dir domain/command
mockery --case=snake --outpkg=querymocks --output=adapters/querymocks --name=Bus --dir domain/query
mockery --case=snake --outpkg=eventmocks --output=adapters/eventmocks --name=Bus --dir domain/event
popd
popd

