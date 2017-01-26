APP.controller('app-controller', function($scope, Thing) {
    var things = Thing.query(function(){
        console.log(things);
        $scope.things = things;
    });

    $scope.activeTab = 1;

    $scope.OnAdd = function(){
        $scope.dialogTitle = "Add New";
        $scope.thingId = guid();
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
        var t = Thing.get({"id": id}, function() {
            $scope.dialogTitle = "Edit " + t.id;
            $scope.thingId = t.id;
            $scope.thingSchema = t.schema;
            $scope.thingState = t.state;

            $('#dlgModal').modal('show')
        });
    }

    $scope.OnSearch = function() {

    }

    $scope.OnConfirmDelete = function(id) {
        $scope.deleteThing = id;
        $('#confirm-delete-thing').modal('show')
    }

    $scope.OnDelete = function(id) {
        alert(id + " deleted");
    }

    $scope.OnUpsert = function() {
        var t = new Thing();
        t.id = $scope.thingId;
        t.Schema = btoa($scope.thingSchema);
        t.State = btoa($scope.thingState);

        t.$save();
    }

    $scope.setTab = function (t) {
        $scope.activeTab = t;
    }

    $scope.printDate = function(d) {
        return new Date(d)
    }
});
