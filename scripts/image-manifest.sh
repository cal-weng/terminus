#!/usr/bin/env bash



BASE_DIR=$(dirname $(realpath -s $0))

PACKAGE_MODULE=("frameworks" "libs" "apps" "third-party")
IMAGE_MANIFEST=".manifest/images.mf"

rm -rf .manifest
mkdir -p .manifest

# copy default base images
cp $BASE_DIR/../build/manifest/images ${IMAGE_MANIFEST}

TMP_MANIFEST=$(mktemp)
for mod in "${PACKAGE_MODULE[@]}";do
    echo "find images in ${mod} ..."
    ls ${mod} | while read app; do
        chart_path="${mod}/${app}/config"
        if [ -d $chart_path ]; then
            find $chart_path -type f -name *.yaml | while read p; do
                bash ${BASE_DIR}/yaml2prop.sh -f $p | while read l;do 
                    if [[ "$l" == *".image = "* ]]; then 
                        echo "$l"
                        echo "$l" >> ${TMP_MANIFEST}
                    fi;
                done
            done
        fi
    done
done

awk '{print $3}' ${TMP_MANIFEST} | sort | uniq | grep -v nitro | grep -v orion >> ${IMAGE_MANIFEST}

# patch
# fix backup server version
backup_version=$(egrep '{{ \$backupVersion := "(.*)" }}' frameworks/backup-server/config/cluster/deploy/backup_server.yaml | sed 's/{{ \$backupVersion := "\(.*\)" }}/\1/')
if [[ "$OSTYPE" == "darwin"* ]]; then
    bash -c "sed -i '' -e 's/backup-server:vvalue/backup-server:v$backup_version/' ${IMAGE_MANIFEST}"
else
    bash -c "sed -i 's/backup-server:vvalue/backup-server:v$backup_version/' ${IMAGE_MANIFEST}"
fi