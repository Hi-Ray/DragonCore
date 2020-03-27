#!/bin/sh
lerna bootstrap --use-workspaces;
lerna exec --parallel -- yarn dev;