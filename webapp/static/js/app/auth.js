angular.module('charlotte-app', []).controller('app-controller', function($scope) {
    $scope.OnAddNewAuth = OnAddNewAuth;
    $scope.OnViewEditAuth = OnViewEditAuth;
    $scope.OnDeleteAuth = OnDeleteAuth;

    $scope.auths = [
        {
            "id": "auth-1",
            "type": "preshared-key",
            "added": "2017.01.01",
            "updated": "2017.01.03"
        }
    ];
});

function OnAddNewAuth() {
    var $scope = angular.element($("#authModal")).scope();

    $scope.dialogTitle = "Add New";
    $scope.dlgAuthType = "psk";

    $('#authModal').modal('show')
}

function OnViewEditAuth(id) {
    var $scope = angular.element($("#authModal")).scope();

    // Get auth entry from remote
    $scope.dialogTitle = "Edit " + id;
    $('#authModal').modal('show')
        // Show Dialog
}

function OnDeleteAuth(id) {
    var $scope = angular.element($("#authModal")).scope();

    // COnfirm
    // Send Delete
    // Feedback
}
