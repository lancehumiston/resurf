using Toybox.Time;
using Toybox.System as Sys;
using Toybox.Time.Gregorian;

class SurfDataFields {
  /*
   * Constants
   */
  // Speed at which the rider is considered to be riding a wave i.e. not paddling
  static const maxPaddleSpeed = 3.75;
  // Number of consecutive speed readings required to change the `isRiding` flag
  static const jitterTolerance = 3;

  /*
   * Hidden fields
   */
  // Flag indicating that the user is riding a wave
  hidden var isRiding = false;
  // Collection of timestamps where the rider is above or below the 'maxPaddleSpeed'
  hidden var speedReadingTimes = [];

  /*
   * Public fields
   */
  var hr;
  var hrN;
  var speed = 0;
  var waves = 0;
  var time;

  function initialize() {}

  function toStr(o) {
    if (o != null) {
      return "" + o;
    } else {
      return "---";
    }
  }

  function fmtTime(clock) {
    var hour = clock.hour;

    if (hour > 12) {
      hour -= 12;
    } else if (hour == 0) {
      hour += 12;
    }

    return "" + hour + ":" + clock.min.format("%02d");
  }

  function getTimestamp() {
    var currentTime = Sys.getClockTime();

    return currentTime.hour.format("%02d") + ":" +
           currentTime.min.format("%02d") + ":" +
           currentTime.sec.format("%02d") + "Z" +
           (currentTime.timeZoneOffset / 3600).format("%02d") + ":" +
           (currentTime.timeZoneOffset % 3600).format("%02d");
  }

  function checkJitterTolerance(currentSpeed) {
    if (speedReadingTimes.size() >= jitterTolerance) {
      return true;
    }

    speedReadingTimes.add(getTimestamp());
    return false;
  }

  function compute(info) {
    hr = toStr(info.currentHeartRate);
    hrN = info.currentHeartRate;
    time = fmtTime(Sys.getClockTime());

    if (info.currentSpeed != null) {
      speed = info.currentSpeed;
    } else {
      return;
    }

    // Check if the rider is riding a wave
    if (speed > maxPaddleSpeed) {
      // Check if the rider's state was already riding a wave (`isRiding`)
      if (isRiding) {
        return;
      }

      // Check if the speed has been consistently high long enough to trust the
      // reading
      if (!checkJitterTolerance(speed)) {
        return;
      }

      waves++;
      isRiding = true;
      Sys.print(speedReadingTimes[0] + "|");
      speedReadingTimes = [];

      return;
    }

    // Check if the rider's state was already NOT riding a wave (`!isRiding`)
    if (!isRiding) {
      return;
    }

    // Check if the speed has been consistently low long enough to trust the reading
    if (!checkJitterTolerance(speed)) {
      return;
    }

    isRiding = false;
    Sys.print(speedReadingTimes[0] + "\r\n");
    speedReadingTimes = [];
  }
}