#!/usr/bin/env bash

HTPASSWD_FILE="./htpass"
USERNAME="consoledeveloper"
USERPASS="developer"
HTPASSWD_SECRET="htpasswd-consoledeveloper-secret"

OC_USERS_LIST="$(oc get users)"
if echo "${OC_USERS_LIST}" | grep -q "${USERNAME}"; then
    echo -e "\n\033[0;32m \xE2\x9C\x94 User consoledeveloper already exists \033[0m\n"
    exit;
fi
htpasswd -cb $HTPASSWD_FILE $USERNAME $USERPASS

oc get secret $HTPASSWD_SECRET -n openshift-config &> /dev/null

oc create secret generic ${HTPASSWD_SECRET} --from-file=htpasswd=${HTPASSWD_FILE} -n openshift-config

oc apply -f - <<EOF
apiVersion: config.openshift.io/v1
kind: OAuth
metadata:
  name: cluster
spec:
  identityProviders:
  - name: consoledeveloper
    challenge: true
    login: true
    mappingMethod: claim
    type: HTPasswd
    htpasswd:
      fileData:
        name: ${HTPASSWD_SECRET}
EOF

sleep 10s
oc create clusterrolebinding ${USERNAME}_role1 --clusterrole=self-provisioner --user=${USERNAME}
oc create clusterrolebinding ${USERNAME}_role2 --clusterrole=view --user=${USERNAME}
sleep 15s
echo -e "\n\e[1;35m User consoledeveloper created with the password developer. Type the below\e[0m \n"
echo -e "\n\e[1;32m oc login -u\e[3m \e[1;36mconsoledeveloper\e[0m \e[1;32m-p\e[3m \e[1;36mdeveloper\e[0m \n"