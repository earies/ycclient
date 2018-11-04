# OpenConfig

#./ycclient -config ../../configs/config-old.toml -get -module openconfig-telemetry -revision 2017-08-24 -org openconfig -vendor openconfig
#./ycclient -config ../../configs/config.toml -get -module openconfig-aaa -revision 2017-09-18 -org openconfig -vendor openconfig
#./ycclient -config ../../configs/config.toml -get -module openconfig-acl -revision 2017-05-26 -org openconfig -vendor openconfig

# Cisco

#./ycclient -config ../../configs/config.toml -get -module Cisco-IOS-XR-man-netconf-cfg -revision 2016-03-15 -org cisco -vendor cisco


# IETF
./ycclient -config ../../configs/config.toml -get -module ietf-netconf-notifications -revision 2012-02-06 -org ietf -vendor ietf
# Below is an example of union (bool/string) that needs to be accounted for
./ycclient -config ../../configs/config.toml -get -module ietf-netconf-server -revision 2018-10-22 -org ietf -vendor ietf
