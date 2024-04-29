#!/usr/bin/env bash



BASE_DIR=$(dirname $(realpath -s $0))
rm -rf ${BASE_DIR}/../.dist
DIST_PATH="${BASE_DIR}/../.dist/install-wizard" 
VERSION=$1

set -o pipefail
set -e

DIST_PATH=${DIST_PATH} bash ${BASE_DIR}/package.sh
cp ${BASE_DIR}/upgrade.sh ${DIST_PATH}/.

bash ${BASE_DIR}/image-manifest.sh

pushd ${BASE_DIR}/../.manifest
bash $BASE_DIR/save-images.sh images.mf
popd


pushd $DIST_PATH

rm -rf images
mv ${BASE_DIR}/../.manifest images

if [[ "$OSTYPE" == "darwin"* ]]; then
    TAR=gtar
    SED="sed -i '' -e"
else
    TAR=tar
    SED="sed -i"
fi

if [ ! -z $VERSION ]; then
    sh -c "$SED 's/#__VERSION__/${VERSION}/' wizard/config/settings/templates/terminus_cr.yaml"
    sh -c "$SED 's/#{{LATEST_VERSION}}/${VERSION}/' publicInstaller.latest"
    VERSION="v${VERSION}"
else
    VERSION="debug"
fi

$TAR --exclude=wizard/tools --exclude=.git -zcvf ${BASE_DIR}/../install-wizard-${VERSION}.tar.gz .

popd