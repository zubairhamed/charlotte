APP.controller('app-controller', function($scope) {
    freeboard.initialize(true);
    freeboard.setEditing(false);

    $scope.onClickNewDataSource = onClickNewDataSource
    $scope.onClickRefreshDataSource = onClickRefreshDataSource
    $scope.onClickDeleteDataSource = onClickDeleteDataSource
    $scope.onClickNewDashboard = onClickNewDashboard
    $scope.onClickSaveDashboard = onClickSaveDashboard
    $scope.onClickLoadDashboard = onClickLoadDashboard
    $scope.onClickDeleteDashboard = onClickDeleteDashboard

    // TODO
    // Read current dashboard and if not, default
    // Load dashboard list
    // Populate datasources
});

function onClickNewDataSource() {
    alert("onClickNewDataSource")
}

function onClickRefreshDataSource() {
    alert("onClickRefreshDataSource")
}

function onClickDeleteDataSource() {
    alert("onClickDeleteDataSource")
}

function onClickNewDashboard() {
    alert("onClickNewDashboard")
}

function onClickSaveDashboard() {
    alert("onClickSaveDashboard")
}

function onClickLoadDashboard() {
    alert("onClickLoadDashboard")
}

function onClickDeleteDashboard() {
    alert("onClickDeleteDashboard")
}

// Service Methods
function listDashboards( ) {

}

function getDashboard(id) {

}

function saveDashboard(id, dash) {

}

function deleteDashboard(id) {

}
/*
// HTTP
 .controller('MovieController', function($scope, $http){

 $http.get("http://www.omdbapi.com/?t=" + $scope.search + "&tomatoes=true&plot=full")
 .then(function(response){ $scope.details = response.data; });

// Select Change
 <select ng-options="size as size.name for size in sizes" ng-model="item" ng-change="update()"></select>

 $scope.update = function() {
 $scope.item.size.code = $scope.selectedItem.code
 // use $scope.selectedItem.code and $scope.selectedItem.name here
 // for other stuff ...
 }


 <select name="repeatSelect" id="repeatSelect" ng-model="data.model">
    <option ng-repeat="option in data.availableOptions" value="{{option.id}}">{{option.name}}</option>
 </select>


 */