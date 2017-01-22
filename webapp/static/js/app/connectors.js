angular.module('charlotte-app', []).controller('app-controller', function($scope) {
    $scope.OnViewConnector = OnViewConnector;

    $scope.connectors = [];
});

function OnViewConnector(id) {

}
