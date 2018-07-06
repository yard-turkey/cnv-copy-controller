unset BINARIES TESTS DOCKER_IMAGES DOCKER_REPO DOCKER_TAG CONTROLLER IMPORTER CLONER

CONTROLLER="cdi-controller"
IMPORTER="cdi-importer"
CLONER="cdi-cloner"

BINARIES="cmd/${CONTROLLER} cmd/${IMPORTER}"
TESTS="cmd/ pkg/ test/"
DOCKER_IMAGES="${CONTROLLER} ${IMPORTER} ${CLONER}"
DOCKER_REPO=${DOCKER_REPO:-kubevirt}
DOCKER_TAG=${DOCKER_TAG:-latest}
