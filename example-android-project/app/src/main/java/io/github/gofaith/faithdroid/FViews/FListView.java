package io.github.gofaith.faithdroid.FViews;

import android.support.v7.widget.RecyclerView;
import android.view.ViewGroup;

import org.json.JSONObject;
import org.json.JSONTokener;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FListView extends FView implements AttrGettable,AttrSettable{
    private RecyclerView v;
    private MyAdapter adapter;
    public String fnOnGetView,fnOnBindData,fnOnGetCount;
    public FListView(UIController uiController, RecyclerView.LayoutManager lm) {
        parrentController=uiController;
        v = new RecyclerView(parrentController.activity);
        view=v;
        v.setLayoutManager(lm);
        adapter=new MyAdapter(this);
        v.setAdapter(adapter);
    }

    @Override
    public String getAttr(String attr) {
        if (attr == null) {
            return null;
        }
        switch (attr) {
            case "Visibility":
                return getVisibility();
        }
        return null;
    }

    @Override
    public void setAttr(String attr, String value) {
        if (attr == null) {
            return ;
        }
        switch (attr) {
            case "BackgroundColor":
                setBackgroundColor(value);
                break;
            case "Background":
                setBackground(value);
                break;
            case "Size":
                parseSize(parrentController.activity, v, value);
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
                // --------------------------------------------
            case "OnGetView":
                fnOnGetView=value;
                break;
            case "OnBindData":
                fnOnBindData=value;
                break;
            case "OnGetCount":
                    fnOnGetCount = value;
                break;
        }
        return ;
    }
    class MyAdapter extends RecyclerView.Adapter<MyAdapter.ViewHolder> {
        private FListView parentListView;
        public class ViewHolder extends RecyclerView.ViewHolder {
            public FView fView;
            public String str;
            public ViewHolder(FView v,String str) {
                super(v.view);
                fView = v;
                this.str =str;
            }
        }
        public MyAdapter(FListView fListView) {
            parentListView =fListView;
        }
        @Override
        public MyAdapter.ViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
            if (parentListView.fnOnGetView == null) {
                return new ViewHolder(null,"");
            }
            String viewGot= Faithdroid.triggerEventHandler(parentListView.fnOnGetView,"");

            try {
                JSONObject object = (JSONObject) (new JSONTokener(viewGot).nextValue());
                FView v = parentListView.parrentController.viewmap.get(object.getString("VID"));
                return new ViewHolder(v,viewGot);
            } catch (Exception e) {
                e.printStackTrace();
                return null;
            }
        }
        @Override
        public void onBindViewHolder(ViewHolder holder, int position) {
            if (parentListView.fnOnBindData == null) {
                return;
            }
            try {
                JSONObject jo = new JSONObject();
                jo.put("Str", holder.str);
                jo.put("Position", position);
                Faithdroid.triggerEventHandler(parentListView.fnOnBindData, jo.toString());
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
        @Override
        public int getItemCount() {
            if (parentListView.fnOnGetCount == null) {
                return 0;
            }
            try {
                return Integer.parseInt(Faithdroid.triggerEventHandler(parentListView.fnOnGetCount,""));
            } catch (Exception e) {
                e.printStackTrace();
            }
            return 0;
        }
    }

}
