package io.github.gofaith.faithdroid.FViews;

import android.graphics.drawable.Drawable;
import android.widget.ImageView;

import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.Toolkit;
import io.github.gofaith.faithdroid.UI.UIController;

public class FImageView extends FView implements AttrSettable, AttrGettable {
    public ImageView v;

    public FImageView(UIController controller) {
        parentController=controller;
        v = new ImageView(parentController.activity);
        view=v;
    }
    @Override
    public String getAttr(String attr) {
        switch (attr) {
            case "Visibility":
                return getVisibility();
            // ------------------------------------------
        }
        return "";
    }

    @Override
    public void setAttr(String attr, String value) {
        if (value==null)
            return;
        switch (attr) {
            case "BackgroundColor":
                setBackgroundColor(value);
                break;
            case "Background":
                setBackground(value);
                break;
            case "Size":
                parseSize(value);
                break;
            case "Visibility":
                setVisibility(value);
                break;
            case "Padding":
                setPadding(value);
                break;
            case "Margin":
                setMargin(value);
                break;
            case "LayoutGravity":
                setLayoutGravity(value);
                break;
            case "Elevation":
                setElevation(value);
                break;
            case "LayoutWeight":
                setLayoutWeight(value);
                break;
            // -------------------------------------------------------------------
            case "Src":
                Drawable drawable = Toolkit.file2Drawable(parentController.activity, value);
                if (drawable == null) {
                    return;
                }
                v.setImageDrawable(drawable);
                break;
        }
    }
}
