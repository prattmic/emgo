#!/bin/sh

INTERFACE=stlink-v2-1
TARGET=stm32f3x
TRACECLKIN=72000000

. ../../../../../scripts/load-oocd.sh $@
