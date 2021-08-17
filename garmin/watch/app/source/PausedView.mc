using Toybox.WatchUi;


class PausedView extends WatchUi.View {

    function initialize() {
        View.initialize();
    }

    function onLayout(dc) {
        var height = dc.getHeight();
        var width = dc.getWidth();

		var discardButton = new WatchUi.Button({
            :behavior => :onClickDiscard,
            :stateHighlighted => Graphics.COLOR_RED,
            :stateDisabled => Graphics.COLOR_DK_RED,
            :stateDefault => Graphics.COLOR_DK_RED,
            :locX => 0,
            :locY => 0,
            :width => dc.getWidth(),
            :height => height / 2,
        });
        var discardText = new WatchUi.Text({
            :text => "DISCARD",
            :color => Graphics.COLOR_WHITE,
            :locX => width / 2,
            :locY => height / 3,
            :font => Graphics.FONT_MEDIUM,
            :justification => Graphics.TEXT_JUSTIFY_CENTER | Graphics.TEXT_JUSTIFY_VCENTER,
        });

        var saveButton = new WatchUi.Button({
            :behavior => :onClickSave,
            :stateHighlighted => Graphics.COLOR_GREEN,
            :stateDisabled => Graphics.COLOR_DK_GREEN,
            :stateDefault => Graphics.COLOR_DK_GREEN,
            :locX => 0,
            :locY => height - height / 2,
            :width => dc.getWidth(),
            :height => height / 2,
        });
        var saveText = new WatchUi.Text({
            :text => "SAVE",
            :color => Graphics.COLOR_WHITE,
            :locX => width / 2,
            :locY => height - height / 3,
            :font => Graphics.FONT_MEDIUM,
            :justification => Graphics.TEXT_JUSTIFY_CENTER | Graphics.TEXT_JUSTIFY_VCENTER,
        });
        setLayout([ discardButton, discardText, saveButton, saveText ]);
    }

    function onShow() {
    }

    function onUpdate(dc) {
        View.onUpdate(dc);
    }

    function onHide() {
    }
}
