using Toybox.WatchUi;


class PausedDelegate extends WatchUi.BehaviorDelegate {
    hidden var data;

    function initialize(data) {
        BehaviorDelegate.initialize();
        self.data = data;
    }

    function onKey(keyEvent) {
        if (keyEvent.getKey() == KEY_ENTER) {
            backToSurfTracker();
            return true;
        }
        return false;
    }

    function onBack() {
        backToSurfTracker();
        return true;
    }

    hidden function backToSurfTracker() {
        WatchUi.switchToView(new SurfTrackerView(data), new SurfTrackerDelegate(data), WatchUi.SLIDE_RIGHT);
    }

    function onClickSave() {
        var result = data.save();
        if (result) {
            WatchUi.popView(WatchUi.SLIDE_IMMEDIATE);
       }
       return result;
    }
    
    function onClickDiscard() {
        var result = data.discard();
        if (result) {
            WatchUi.popView(WatchUi.SLIDE_IMMEDIATE);
       }
       return result;
    }
}
