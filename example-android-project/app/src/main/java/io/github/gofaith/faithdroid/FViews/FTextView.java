package io.github.gofaith.faithdroid.FViews;

import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.widget.TextView;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FTextView extends FView implements AttrGettable,AttrSettable {
    private TextView v;

    public FTextView(UIController c) {
        parrentController=c;
        v = new TextView(parrentController.activity);
        view=v;
    }
    @Override
    public String getAttr(String attr) {
        switch (attr) {
            case "Text":
                return v.getText().toString();
        }
        return "";
    }

    @Override
    public void setAttr(String attr, final String value) {
        switch (attr) {
            case "BackgroundColor":
                setBackgroundColor(v,value);
                break;
            case "Background":
                setBackground(v,value);
                break;
            case "Size":
                parseSize(parrentController.activity, v, value);
                break;
            case "Text":
                if (value!=null)
                    v.setText(value);
                break;
            case "OnClick":
                v.setOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View view) {
                        Faithdroid.triggerEventHandler(value,"");
                    }
                });
                break;
        }
    }
}
