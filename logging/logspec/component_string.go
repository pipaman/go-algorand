// Code generated by "stringer -type=Component"; DO NOT EDIT.

package logspec

import "strconv"

const _Component_name = "AgreementCatchupNetworkLedgerFrontendnumComponents"

var _Component_index = [...]uint8{0, 9, 16, 23, 29, 37, 50}

func (i Component) String() string {
	if i < 0 || i >= Component(len(_Component_index)-1) {
		return "Component(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Component_name[_Component_index[i]:_Component_index[i+1]]
}
