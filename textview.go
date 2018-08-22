package faithdroid

type FTextView struct {
	FBaseView
}

func (v *FTextView) getVID() string {
	return v.vID
}
func (v *FTextView) id(s string) *FTextView {
	idMap[s] = v
	return v
}
func getTextViewById(id string) *FTextView {
	if v, ok := idMap[id].(*FTextView); ok {
		return v
	}
	return nil
}
func textview(a *Activity) *FTextView {
	v := &FTextView{}
	v.vID = newToken()
	v.className = "TextView"
	v.ui = a.ui
	a.ui.NewView(v.className, v.vID)
	return v
}
func (v *FTextView) size(w, h int) *FTextView {
	i := []int{w, h}
	v.ui.ViewSetAttr(v.vID, "Size", jsonArray(i))
	return v
}
func (v *FTextView) text(s string) *FTextView {
	v.ui.ViewSetAttr(v.vID, "Text", s)
	return v
}
func (v *FTextView) getText() string {
	return v.ui.ViewGetAttr(v.vID, "Text")
}

func (v *FTextView) background(s string) *FTextView {
	v.FBaseView.background(s)
	return v
}
func (v *FTextView) backgroundColor(s string) *FTextView {
	v.FBaseView.backgroundColor(s)
	return v
}
