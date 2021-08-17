using Toybox.Application;
using Toybox.WatchUi;
using Toybox.Position;


class SurfTrackerApp extends Application.AppBase {
    hidden var data;

    function initialize() {
        AppBase.initialize();
        data = new Data();
    }

    function onStart(state) {
    }

    function onStop(state) {
    }

    function getInitialView() {
        return [
        	new SurfTrackerView(data),
        	new SurfTrackerDelegate(data)
        ];
    }
}
