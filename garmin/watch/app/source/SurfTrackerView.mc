using Toybox.WatchUi;
using Toybox.Time.Gregorian;
using Toybox.Time;

class SurfTrackerView extends WatchUi.View  {
    hidden var data;
    hidden var timer;
    hidden var mCurrentTime;
    hidden var mElapsedTime;

    function initialize(data) {
        View.initialize();
        self.data = data;

        self.timer = new Timer.Timer();
    }

    function onLayout(dc) {
		drawStaticComponents(dc);
    }

	function drawStaticComponents(dc) {
		dc.setColor(Graphics.COLOR_BLUE, Graphics.COLOR_WHITE);

	    // horizontal lines
	    var oneQuarterY = dc.getHeight() / 4;
	    var threeQuarterY = oneQuarterY * 3;
	    dc.drawLine(0, oneQuarterY, dc.getWidth(), oneQuarterY);
	    dc.drawLine(0, threeQuarterY, dc.getWidth(), threeQuarterY);
	}
	
	function drawTextCentered(dc, x, y, font, text) {
	    if (text != null) {
	      dc.drawText(x, y, font, text, Graphics.TEXT_JUSTIFY_CENTER | Graphics.TEXT_JUSTIFY_VCENTER);
	    }
	 }

    function onShow() {
        timer.start(method(:requestUpdate), 1000, true);
    }

    function requestUpdate() {
        WatchUi.requestUpdate();
    }

	function onUpdate(dc) {
		View.onUpdate(dc);
	
		drawStaticComponents(dc);
	
	    dc.setColor(Graphics.COLOR_LT_GRAY, Graphics.COLOR_BLACK);

		var centerX = dc.getWidth() / 2;
		var centerY = dc.getHeight() / 2;
		var oneQuarterY = dc.getHeight() / 4;
	    var threeQuarterY = oneQuarterY * 3;
	    drawTextCentered(dc, centerX, oneQuarterY - 20, Graphics.FONT_MEDIUM, data.getWaveCount().format("%d"));
	    drawTextCentered(dc, centerX, oneQuarterY - 40, Graphics.FONT_XTINY, "WAVES");
	
	    drawTextCentered(dc, centerX, centerY, Graphics.FONT_NUMBER_THAI_HOT, getCurrentTime());
	
		drawHrBackground(dc, 0, threeQuarterY, dc.getWidth(), threeQuarterY, Activity.getActivityInfo().currentHeartRate);
	    drawTextCentered(dc, centerX, threeQuarterY + 18, Graphics.FONT_SMALL, getElapsedTime());
	
	    drawLayout(dc);
	    return true;
	}
	
	function drawHrBackground(dc, x, y, width, height, hr) {
	    var color;
	    color = Graphics.COLOR_BLUE;
	
	    if (hr == null) {
	      hr = 0;
	    }
	
	    if (hr > 150) {
	      color = Graphics.COLOR_DK_RED;
	    } else if (hr > 135) {
	      color = Graphics.COLOR_ORANGE;
	    } else if (hr > 115) {
	      color = Graphics.COLOR_YELLOW;
	    } else if (hr > 95) {
	      color = Graphics.COLOR_GREEN;
	    }
	
	    dc.setColor(color, Graphics.COLOR_TRANSPARENT);
	    dc.fillRectangle(x, y, width, height);
	    dc.setColor(Graphics.COLOR_BLACK, Graphics.COLOR_TRANSPARENT);
  	}

    function onHide() {
        timer.stop();
    }

	function getCurrentTime() {
		return fmtCurrentTime(Gregorian.info(Time.now(), Time.FORMAT_SHORT));
	}

	function getElapsedTime() {
		var elapsed = Time.today().add(new Time.Duration(Activity.getActivityInfo().timerTime / 1000));
		return fmtElapsedTime(Gregorian.info(elapsed, Time.FORMAT_SHORT));
	}
    
    function fmtCurrentTime(time) {
    	var hour = time.hour;
	
	    if (hour > 12) {
	      hour -= 12;
	    } else if (hour == 0) {
	      hour += 12;
	    }
	
	    return "" + hour + ":" + time.min.format("%02d");
	}

    function fmtElapsedTime(time) {
        return Lang.format(
            "$1$:$2$:$3$",
            [ time.hour.format("%02d"), time.min.format("%02d"), time.sec.format("%02d")]
        );
    }
}
