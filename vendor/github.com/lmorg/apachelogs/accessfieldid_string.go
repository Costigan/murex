// Code generated by "stringer -type=AccessFieldId"; DO NOT EDIT

package apachelogs

import "fmt"

const _AccessFieldId_name = "AccFieldIpAccFieldUserIdAccFieldDateTimeAccFieldDateAccFieldTimeAccFieldMethodAccFieldUriAccFieldQueryStringAccFieldProtocolAccFieldStatusAccFieldSizeAccFieldReferrerAccFieldUserAgentAccFieldProcTimeAccFieldFileName"

var _AccessFieldId_index = [...]uint8{0, 10, 24, 40, 52, 64, 78, 89, 108, 124, 138, 150, 166, 183, 199, 215}

func (i AccessFieldId) String() string {
	i -= 1
	if i >= AccessFieldId(len(_AccessFieldId_index)-1) {
		return fmt.Sprintf("AccessFieldId(%d)", i+1)
	}
	return _AccessFieldId_name[_AccessFieldId_index[i]:_AccessFieldId_index[i+1]]
}