#!/bin/bash

rsync -avC --exclude-from=rsync.excludes ./page/ /var/www/www.7de.se/go-learn/SE/
