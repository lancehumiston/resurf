using Toybox.WatchUi;
using Toybox.Time;
using Toybox.System;
using Toybox.Timer;
using Toybox.Attention;

class SurfTrackerDelegate extends WatchUi.InputDelegate {
    hidden const DOUBLE_PRESS_THRESHOLD = 400;

    hidden var data;

    hidden var lastKeyPressMillis;
    hidden var buttonTimer;
    
    function initialize(data) {
    	InputDelegate.initialize();
    	BehaviorDelegate.initialize();

        self.data = data;
        data.start();

        buttonTimer = new Timer.Timer();
        lastKeyPressMillis = System.getTimer() - DOUBLE_PRESS_THRESHOLD;
    }

    /* Disable all screen touch events */
    function onHold(clickEvent) {
        return true;
    }
    function onRelease(clickEvent) {
        return true;
    }
    function onSelectable(selectableEvent){
        return true;
    }
    function onSwipe(swipeEvent) {
        return true;
    }
    function onTap(clickEvent){
        return true;
    }

    /*
     * Enter single press => Record wave
     * Enter double press => Pause / end
     */
    function onKey(keyEvent) {
        if (keyEvent.getKey() == KEY_ENTER) {
            var now = System.getTimer();
            if (now - lastKeyPressMillis < DOUBLE_PRESS_THRESHOLD) {
                buttonTimer.stop();
                WatchUi.switchToView(new PausedView(), new PausedDelegate(data), WatchUi.SLIDE_LEFT);
            } else  {
                recordWave();
            }
            lastKeyPressMillis = now;
            return true;
        }

        return true;
    }

    function recordWave() {
        data.addWave();
        Attention.vibrate([new Attention.VibeProfile(100, 100)]);
        WatchUi.requestUpdate();
    }
}
