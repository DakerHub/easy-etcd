#!/bin/sh

funStopContainer() {
  cidTmp=$1
  echo "stop container $cidTmp"
  $(docker stop $cidTmp)
}

funRemoveContainer() {
  cidTmp=$1
  echo "remove container $cidTmp"
  $(docker rm $cidTmp)
}

funRemoveImages() {
  iidTmp=$1
  echo "remove image $iidTmp"
  $(docker rmi $iidTmp)
}

CID_LIST=$(docker ps -a | grep $1 | awk '{print $1}')
echo "clear the containers"
if [ "$CID_LIST" ]; then
  echo "$CID_LIST"
  for cid in $CID_LIST; do
    funStopContainer $cid || true
    funRemoveContainer $cid || true
  done
else
  echo "nothing to clear"
fi

IID_LIST=$(docker images | grep $1 | awk '{print $3}')
echo "remove the images"
if [ "$IID_LIST" ]; then
  echo "$IID_LIST"
  for iid in $IID_LIST; do
    funRemoveImages $iid || true
  done
else
  echo "nothing to remove"
fi
