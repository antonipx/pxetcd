#!/bin/bash -e
# {{.Origin}}

IP_ADDRESSES=("{{.IP1}}" "{{.IP2}}" "{{.IP3}}")
ENCRYPTION="{{.Encryption}}"
INITIAL_TOKEN="{{.InitialToken}}"
NODE_PREFIX="{{.Prefix}}"
CLIENT_PORT="{{.ClientPort}}"
PEER_PORT="{{.PeerPort}}"
DIRECTORY="{{.Directory}}"
VERSION="{{.Version}}"
USER="{{.Username}}"
ETCD_BIN="{{.EtcdBin}}"

check() { [ -x "$(command -v $1)" ] || { echo "Error: $1 is not installed, aborting"; exit 1; } }
for cmd in ip install curl awk tar cat systemctl useradd; do check $cmd; done

for ((n=0; n<${#IP_ADDRESSES[@]}; n++)); do
    cluster+=( "${NODE_PREFIX}${n}=http://${IP_ADDRESSES[${n}]}:${PEER_PORT}" )
    for locaddr in $(ip addr | awk '{ if($1 == "inet") { split($2, i, "/"); print i[1]; } }'); do
        [ ${IP_ADDRESSES[$n]} == $locaddr ] && { myidx=$n; }
    done
done

ARGS=" \
--name ${NODE_PREFIX}${myidx?:Error: This nodes IP address not found on IP address list} \
--listen-peer-urls http://0.0.0.0:${PEER_PORT} \
--listen-client-urls http://0.0.0.0:${CLIENT_PORT} \
--advertise-client-urls http://${IP_ADDRESSES[$myidx]}:${CLIENT_PORT} \
--initial-advertise-peer-urls http://${IP_ADDRESSES[$myidx]}:${PEER_PORT} \
--initial-cluster-token ${INITIAL_TOKEN} \
--initial-cluster $(IFS=,; echo "${cluster[*]}") \
--initial-cluster-state new \
--data-dir ${DIRECTORY}/data \
"

mkdir -pm 755 ${DIRECTORY}/bin
mkdir -pm 755 ${DIRECTORY}/data
useradd -c "Portworx Etcd" -d ${DIRECTORY} -s /bin/false ${USER}

if [ $USE_EXISTING -eq 1 ]
then

curl -fsSL https://github.com/coreos/etcd/releases/download/${VERSION}/etcd-${VERSION}-linux-amd64.tar.gz | \
    tar -xvz --strip=1 -f - -C ${DIRECTORY}/bin etcd-${VERSION}-linux-amd64/etcdctl etcd-${VERSION}-linux-amd64/etcd

chown -R ${USER}:${USER} ${DIRECTORY}

cat > /etc/systemd/system/pxetcd.service << EOF
[Unit]
Description=Etcd Key Value Store for Portworx
After=network.target
Wants=network-online.target
 
[Service]
User=${USER}
Type=notify
PermissionsStartOnly=true
ExecStart=${DIRECTORY}/bin/etcd ${ARGS}
Restart=always
RestartSec=10s
LimitNOFILE=40000
TimeoutStartSec=0
 
[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable pxetcd
systemctl start pxetcd --no-block
 
