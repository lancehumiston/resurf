using Toybox.WatchUi as Ui;
using Toybox.Graphics;
using Toybox.System as Sys;

class SurfDataView extends Ui.DataField {
  hidden var fields;

  function initialize() {
    DataField.initialize();
    fields = new SurfDataFields();
  }

  function onLayout(dc) {}

  function onShow() {}

  function onHide() {}

  function drawLayout(dc) {
    dc.setColor(Graphics.COLOR_BLUE, Graphics.COLOR_WHITE);

    // horizontal lines
    dc.drawLine(0, 71, 240, 71);
    dc.drawLine(0, 132, 240, 132);
    dc.drawLine(0, 198, 240, 198);
  }

  function onUpdate(dc) {
    dc.setColor(Graphics.COLOR_BLACK, Graphics.COLOR_WHITE);
    dc.clear();

    drawHrBackground(dc, 0, 0, 240, 71, fields.hrN);
    drawTextCentered(dc, 120, 45, Graphics.FONT_NUMBER_MEDIUM, fields.hr);
    drawTextCentered(dc, 120, 18, Graphics.FONT_XTINY, "HR");

    drawTextCentered(dc, 120, 107, Graphics.FONT_NUMBER_MEDIUM, fields.speed);
    drawTextCentered(dc, 120, 80, Graphics.FONT_XTINY, "SPEED");

    drawTextCentered(dc, 120, 154, Graphics.FONT_NUMBER_MEDIUM, fields.waves);
    drawTextCentered(dc, 120, 186, Graphics.FONT_XTINY, "WAVES");

    drawTextCentered(dc, 120, 210, Graphics.FONT_TINY, fields.time);

    drawLayout(dc);
    return true;
  }

  function drawHrBackground(dc, x, y, width, height, hr) {
    var color;
    color = Graphics.COLOR_BLUE;

    if (hr == null) {
      hr = 0;
    }

    if (hr > 170) {
      color = Graphics.COLOR_RED;
    } else if (hr > 150) {
      color = Graphics.COLOR_ORANGE;
    } else if (hr > 130) {
      color = Graphics.COLOR_YELLOW;
    } else if (hr > 110) {
      color = Graphics.COLOR_GREEN;
    }

    dc.setColor(color, Graphics.COLOR_TRANSPARENT);
    dc.fillRectangle(x, y, width, height);
    dc.setColor(Graphics.COLOR_BLACK, Graphics.COLOR_TRANSPARENT);
  }

  function compute(info) {
    fields.compute(info);
    return;
  }

  function drawTextCentered(dc, x, y, font, text) {
    if (text != null) {
      dc.drawText(x, y, font, text, Graphics.TEXT_JUSTIFY_CENTER | Graphics.TEXT_JUSTIFY_VCENTER);
    }
  }
}