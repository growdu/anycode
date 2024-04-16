#!/bin/bash
case  $1 in
    "start")
    echo "start ok"
    exit 0
    ;;
    "stop")
    echo "stop ok"
    exit 0
    ;;
    "restart")
    echo "restart ok"
    exit 0
    ;;
    "upgrade")
    echo "upgrade ok"
    exit 0
    ;;
    "change")
    echo "change ok"
    exit 0
    ;;
    *)
     exit 1
esac