.PHONY: all-ui list-versions clean-artifacts \
	build-v1-small build-v1-medium build-v1-large \
	test-v1-small test-v1-medium test-v1-large \
	package-v1-small package-v1-medium package-v1-large \
	build-v2-small build-v2-medium build-v2-large \
	test-v2-small test-v2-medium test-v2-large \
	package-v2-small package-v2-medium package-v2-large \
	build-v3-small build-v3-medium build-v3-large \
	test-v3-small test-v3-medium test-v3-large \
	package-v3-small package-v3-medium package-v3-large

# Aggregate target with no recipe (intentional for parser/UI testing)
all-ui: build-v1-small test-v1-small package-v1-small \
	build-v2-medium test-v2-medium package-v2-medium \
	build-v3-large test-v3-large package-v3-large

list-versions:
	@echo "versions: v1 v2 v3"

clean-artifacts:
	@echo "cleaning fake artifacts"
	@rm -rf ./tmp/fake-artifacts
	@find ./tmp -name '*.fake' -delete

build-v1-small:
	@echo "[v1/small] build"
build-v1-medium:
	@echo "[v1/medium] build"
	@echo "using medium profile"
build-v1-large:
	@echo "[v1/large] build"
	@echo "step 1: prepare"
	@echo "step 2: compile"
	@echo "step 3: emit artifacts"

test-v1-small:
	@echo "[v1/small] test"
test-v1-medium:
	@echo "[v1/medium] test"
	@echo "running integration subset"
test-v1-large:
	@echo "[v1/large] test"
	@echo "running unit + integration"
	@echo "collecting coverage"
	@echo "publishing fake report"

package-v1-small:
	@echo "[v1/small] package"
package-v1-medium:
	@echo "[v1/medium] package"
	@echo "creating tar.gz"
package-v1-large:
	@echo "[v1/large] package"
	@echo "creating tar.gz"
	@echo "creating zip"
	@echo "creating checksum"

build-v2-small:
	@echo "[v2/small] build"
build-v2-medium:
	@echo "[v2/medium] build"
	@echo "using medium profile"
build-v2-large:
	@echo "[v2/large] build"
	@echo "step 1: prepare"
	@echo "step 2: compile"
	@echo "step 3: emit artifacts"

test-v2-small:
	@echo "[v2/small] test"
test-v2-medium:
	@echo "[v2/medium] test"
	@echo "running integration subset"
test-v2-large:
	@echo "[v2/large] test"
	@echo "running unit + integration"
	@echo "collecting coverage"
	@echo "publishing fake report"

package-v2-small:
	@echo "[v2/small] package"
package-v2-medium:
	@echo "[v2/medium] package"
	@echo "creating tar.gz"
package-v2-large:
	@echo "[v2/large] package"
	@echo "creating tar.gz"
	@echo "creating zip"
	@echo "creating checksum"

build-v3-small:
	@echo "[v3/small] build"
build-v3-medium:
	@echo "[v3/medium] build"
	@echo "using medium profile"
build-v3-large:
	@echo "[v3/large] build"
	@echo "step 1: prepare"
	@echo "step 2: compile"
	@echo "step 3: emit artifacts"

test-v3-small:
	@echo "[v3/small] test"
test-v3-medium:
	@echo "[v3/medium] test"
	@echo "running integration subset"
test-v3-large:
	@echo "[v3/large] test"
	@echo "running unit + integration"
	@echo "collecting coverage"
	@echo "publishing fake report"

package-v3-small:
	@echo "[v3/small] package"
package-v3-medium:
	@echo "[v3/medium] package"
	@echo "creating tar.gz"
package-v3-large:
	@echo "[v3/large] package"
	@echo "creating tar.gz"
	@echo "creating zip"
	@echo "creating checksum"
