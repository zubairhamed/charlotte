angular.module('charlotte-app', []).controller('dashboard-controller', function($scope) {
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
