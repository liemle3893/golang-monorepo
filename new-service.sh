#!/bin/bash
set -e
CUR_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
TPL_DIR="${CUR_DIR}/template"
usage() {
  echo "---"
  echo "Usage: "
  echo "./new-service.sh <service-name>"
}

# $1: Variable value
# $2: Variable name
function ensureNonEmpty() {
  if [ "x$1" == "x" ]; then
    echo "ERR: Variable must be set. $2"
    usage
    exit 1
  fi
  echo "$1"
}

buildDirStructure() {
  svc_name=$(ensureNonEmpty $1 "Service name")
  mkdir -p "${CUR_DIR}/services/${svc_name}"
  mkdir -p "${CUR_DIR}/services/${svc_name}/internal"
  mkdir -p "${CUR_DIR}/services/${svc_name}/cmd"
  mkdir -p "${CUR_DIR}/services/${svc_name}/cmd/migrate"
  mkdir -p "${CUR_DIR}/services/${svc_name}/cmd/run"
}

genFile() {
  local svc_name
  local src_file
  local dest_file
  svc_name=$(ensureNonEmpty $1 "Service name")
  src_file=$(ensureNonEmpty $2 "Source file")
  dest_file=$(ensureNonEmpty $3 "Destination file")
  sed "s/__SVC_NAME__/${svc_name}/g" "${src_file}" >"${dest_file}"
}

genDockerFile() {
  genFile $1 "${TPL_DIR}/Dockerfile.tpl" "${CUR_DIR}/services/${svc_name}/dockerfile.${svc_name}"
}

genMakeFile() {
  genFile $1 "${TPL_DIR}/Makefile.tpl" "${CUR_DIR}/services/${svc_name}/Makefile"
}

genMainFile() {
  genFile $1 "${TPL_DIR}/main.go.tpl" "${CUR_DIR}/services/${svc_name}/main.go"
}

genDeployScriptFile() {
  genFile $1 "${TPL_DIR}/deploy.sh.tpl" "${CUR_DIR}/services/${svc_name}/deploy.sh"
}

SVC_NAME=$(ensureNonEmpty $1 "Service name")

buildDirStructure "${SVC_NAME}"
genDockerFile "${SVC_NAME}"
genMakeFile "${SVC_NAME}"
genDeployScriptFile "${SVC_NAME}"
# Go Project
genMainFile "${SVC_NAME}"
### Gen Root command file
genFile "${SVC_NAME}" "${TPL_DIR}/cmd/root.go.tpl" "${CUR_DIR}/services/${svc_name}/cmd/root.go"
### Gen Migrate File
genFile "${SVC_NAME}" "${TPL_DIR}/cmd/migrate/migrate.go.tpl" "${CUR_DIR}/services/${svc_name}/cmd/migrate/migrate.go"
### Gen Run command File
genFile "${SVC_NAME}" "${TPL_DIR}/cmd/run/run.go.tpl" "${CUR_DIR}/services/${svc_name}/cmd/run/run.go"
