


head.js("js/freeboard_plugins.min.js",
    // *** Load more plugins here ***
    function(){
        $(function()
        { //DOM Ready
            freeboard.initialize(true);

            var hashpattern = window.location.hash.match(/(&|#)source=([^&]+)/);
            if (hashpattern !== null) {
                $.getJSON(hashpattern[2], function(data) {
                    freeboard.loadDashboard(data, function() {
                        freeboard.setEditing(false);
                    });
                });
            }

        });
    });

/*

loadDashboard
saveDashboard
deleteDashboard


 addDatasource
 deleteDatasource
 updateDatasource

 */