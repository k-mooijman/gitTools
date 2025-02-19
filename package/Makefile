PACKAGE         = fleet
PACKAGE_VERSION = $(PACKAGE)_$(VERSION)
PKG_DIR         = tmp/$(PACKAGE_VERSION)
DEBIAN_DIR      = $(PKG_DIR)/DEBIAN
CONF_DIR        = $(PKG_DIR)/etc
BIN_DIR         = $(PKG_DIR)/usr/bin
SYSTEMD_DIR     = $(PKG_DIR)/usr/lib/systemd/system
CONFIGS         = default/fleet
BINS            = ../fleet
SERVICES        = $(notdir $(wildcard systemd/*))

CONTROL_FILE  = $(DEBIAN_DIR)/control
DEB_FILES     = $(addprefix $(DEBIAN_DIR)/, postinst prerm conffiles)
CONFIG_FILES  = $(addprefix $(CONF_DIR)/, $(CONFIGS))
BIN_FILES     = $(addprefix $(BIN_DIR)/, $(notdir $(BINS)))
SERVICE_FILES = $(addprefix $(SYSTEMD_DIR)/, $(SERVICES))

all: ../$(PACKAGE_VERSION).deb

$(CONTROL_FILE): deb/control
	@mkdir -p $(@D)
	sed 's/^Version: .*/Version: '$(DEB_VERSION)'/' $< > $@

$(DEB_FILES): $(DEBIAN_DIR)/%: deb/%
	@mkdir -p $(@D)
	cp $< $@

$(CONFIG_FILES): $(CONF_DIR)/%: %
	install -Dm0600 $< $@

$(BIN_FILES): $(BIN_DIR)/%: ../%
	install -Dm0755 $< $@

$(SERVICE_FILES): $(SYSTEMD_DIR)/%: systemd/%
	install -Dm0644 $< $@

tmp/$(PACKAGE_VERSION).deb: $(CONTROL_FILE) $(DEB_FILES) $(CONFIG_FILES) $(SERVICE_FILES) $(BIN_FILES)
	dpkg-deb --build -Zgzip $(PKG_DIR)

../$(PACKAGE_VERSION).deb: tmp/$(PACKAGE_VERSION).deb
	cp $< $@

clean:
	rm -rf tmp
	rm -f ../*.deb

.PHONY: all clean
