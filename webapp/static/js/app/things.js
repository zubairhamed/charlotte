APP.controller('app-controller', function($scope, Thing) {
    $scope.things = [];
    $scope.activeTab = 1;

    $scope.OnAdd = function(){
        $scope.dialogTitle = "Add New";
        $scope.thingName = guid();
        $scope.thingSchema = `{
    "@context": [
        "https://w3c.github.io/wot/w3c-wot-td-context.jsonld",
        "https://w3c.github.io/wot/w3c-wot-common-context.jsonld"
    ],
    "@type": "Sensor",
    "interactions": []
}`;
        $scope.thingState = `{}`;

        $('#dlgModal').modal('show')
    }

    $scope.OnEdit = function(id) {

    }

    $scope.OnSearch = function() {

    }

    $scope.OnDelete = function(id) {

    }

    $scope.OnSave = function() {
        alert("Saving..");
        var t = new Thing();
        t.Name = $scope.thingName;
        t.Schema = $scope.thingSchema;
        t.State = $scope.thingState;

        t.$save();
    }

    $scope.setTab = function (t) {
        $scope.activeTab = t;
    }
});
