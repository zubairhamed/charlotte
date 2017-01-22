APP.factory("Thing", function($resource){
    return $resource('/service/things/:id', { id:'@_id' });
});
