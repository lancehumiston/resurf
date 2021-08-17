using Toybox.ActivityRecording;
using Toybox.Position;
using Toybox.Time;
using Toybox.Time.Gregorian;
using Toybox.Sensor;


// Rename this to `Session` and all constructor params of Data
class Data
{
    hidden var session = null;
    hidden var running = false;
    hidden var surfingTimeStart;
    hidden var waveCount = 0;

    function start() {
    	Sensor.setEnabledSensors( [Sensor.SENSOR_HEARTRATE] ); 
    
        if (session == null) {
            session = ActivityRecording.createSession({
                :name => "ReSurf",
                :sport => ActivityRecording.SPORT_SURFING
            });
            surfingTimeStart = Time.now();
        }
        session.start();
        running = true;
    }

    function stop() {
        running = false;
        Sensor.unregisterSensorDataListener();
        return session.stop();
    }

    function save() {
        return session.save();
    }
    
    function discard() {
        return session.discard();
    }

    function addWave() {
        waveCount += 1;
        session.addLap();
        
        System.print("-|" + getTimestamp() + "\r\n");
    }
    
    function getTimestamp() {
	    var secondsPerHour = 3600;
	    var timeZoneOffsetSeconds = System.getClockTime().timeZoneOffset;
	    var currentTime = Gregorian.info(Time.now(), Time.FORMAT_SHORT);
	
	    return currentTime.year.format("%02d") + "-" +
	           currentTime.month.format("%02d") + "-" +
	           currentTime.day.format("%02d") + "T" +
	           currentTime.hour.format("%02d") + ":" +
	           currentTime.min.format("%02d") + ":" +
	           currentTime.sec.format("%02d") +
	           (timeZoneOffsetSeconds / secondsPerHour).format("%+03d") + ":" +
	           (timeZoneOffsetSeconds % secondsPerHour).format("%02d");
	}

    function getWaveCount() {
        return waveCount;
    }

    function getSurfingTime() {
        return Time.now().subtract(surfingTimeStart);
    }
}
