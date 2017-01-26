APP.factory("Thing", function($resource){
    return $resource('/service/things/:id', { id:'@id' });
});
