package faithdroid

type FImageView struct {
	FBaseView
}

func ImageView(a *Activity) *FImageView {
	v := &FImageView{}
	v.VID = NewToken()
	v.ClassName = "ImageView"
	v.UI = a.UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (vh *ViewHolder) GetImageViewByItemId(id string) *FImageView {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FImageView); ok {
			return bt
		}
	}
	return nil
}

func (v *FImageView) Size(w, h int) *FImageView {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}
func (v *FImageView) SetId(s string)*FImageView {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FImageView) SetItemId(parent *FListView, id string) *FImageView {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetImageViewById(id string) *FImageView {
	if v, ok := GlobalVars.IdMap[id].(*FImageView); ok {
		return v
	}
	return nil
}

func (v *FImageView) Background(s string) *FImageView {
	v.FBaseView.Background(s)
	return v
}
func (v *FImageView) BackgroundColor(s int) *FImageView {
	v.FBaseView.BackgroundColor(s)
	return v
}
func (v *FImageView) CachedBackground(s string) *FImageView {
	v.FBaseView.CachedBackground(s)
	return v
}

func (v *FImageView) Visible() *FImageView {
	v.FBaseView.Visible()
	return v
}
func (v *FImageView) Invisible() *FImageView {
	v.FBaseView.Invisible()
	return v
}
func (v *FImageView) Gone() *FImageView {
	v.FBaseView.Gone()
	return v
}

func (v *FImageView) Padding(left, top, right, bottom int) *FImageView {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FImageView) PaddingLeft(dp int) *FImageView {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FImageView) PaddingTop(dp int) *FImageView {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FImageView) PaddingRight(dp int) *FImageView {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FImageView) PaddingBottom(dp int) *FImageView {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FImageView) PaddingAll(dp int) *FImageView {
	v.FBaseView.Padding(dp, dp, dp, dp)
	return v
}

func (v *FImageView) Margin(left, top, right, bottom int) *FImageView {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FImageView) MarginLeft(dp int) *FImageView {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FImageView) MarginTop(dp int) *FImageView {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FImageView) MarginRight(dp int) *FImageView {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FImageView) MarginBottom(dp int) *FImageView {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FImageView) MarginAll(dp int) *FImageView {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FImageView) LayoutGravity(gravity int) *FImageView {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FImageView) Elevation(dp float32) *FImageView {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FImageView) Assign(fb **FImageView) *FImageView {
	*fb = v
	return v
}
func (v *FImageView) LayoutWeight(f int) *FImageView {
	v.FBaseView.LayoutWeight(f)
	return v
}

// --------------------------------------------------------
func (v *FImageView) Src(s string) *FImageView {
	if len(s) > len("https://") && s[:len("http")] == "http" {
		go DownloadNetFile(s, "/data/data/"+GlobalVars.UIs[v.UI].GetPkg()+"/tmp/", func(f string) {
			fnID := NewToken()
			GlobalVars.EventHandlersMap[fnID] = func(string) string {
				GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Src", "file://"+f)
				return ""
			}
			GlobalVars.UIs[v.UI].RunOnUIThread(fnID)
		})
		return v
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Src", s)
	return v
}
func (v *FImageView) CachedSrc(s string) *FImageView {
	if len(s) > len("https://") && s[:len("http")] == "http" {
		go CacheNetFile(s, "/data/data/"+GlobalVars.UIs[v.UI].GetPkg()+"/cacheDir/", func(f string) {
			fnID := NewToken()
			GlobalVars.EventHandlersMap[fnID] = func(string) string {
				GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Src", "file://"+f)
				return ""
			}
			GlobalVars.UIs[v.UI].RunOnUIThread(fnID)
		})
		return v
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Src", s)
	return v
}