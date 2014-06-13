var anybase = anybase || {};

function Disconnection(){
};

Disconnection.prototype.set = function(){
	console.log("[Disconnection][set]");
};
Disconnection.prototype.setWithPriority = function(){
	console.log("[Disconnection][setWithPriority]");
};
Disconnection.prototype.update = function(){
	console.log("[Disconnection][update]");
};
Disconnection.prototype.remove = function(){
	console.log("[Disconnection][remove]");
};
Disconnection.prototype.cancel = function(){
	console.log("[Disconnection][cancel]");
};

//Anybase prototype
//Construct a new Firebase reference from a full Firebase URL. See Creating Firebase References.
function Anybase(urlPath){
	console.log("[Anybase]" + urlPath);
	this.urlPath = urlPath;
	this.breadCrumb = urlPath;
}

//Anybase functions
//related to name, root, child etc.
Anybase.prototype.name = function(){
	console.log("[anybase][name]" + this.breadCrumb.split("::")[0]);
	return "JohnDoe";
};
Anybase.prototype.root = function(name){
	console.log("[anybase][root]" + name);
	this.breadCrumb = this.urlPath;
	return this;
};
Anybase.prototype.parent = function(name){
	console.log("[anybase][parent]" + name);
	return this;
};
Anybase.prototype.child = function(name){
	console.log("[anybase][child][name] " + name);
	var bCrumbArray = this.breadCrumb.split("::");
	if(bCrumbArray.indexOf(name) != -1){
		bCrumbArray = bCrumbArray.slice(0,bCrumbArray.indexOf(name) + 1);
		this.breadCrumb = bCrumbArray.join("::");
	}else{
		this.breadCrumb = bCrumbArray.join("::") + "::" + name;
	}
	console.log("[anybase][child] " + this.breadCrumb);
	return this;
};

//related to push set get etc!
Anybase.prototype.push = function(){
	console.log("[anybase][push] " + this.breadCrumb);
	return this;
};
Anybase.prototype.set = function(obj){
	console.log("[anybase][set] " + this.breadCrumb);
	console.log(obj);
};
Anybase.prototype.setWithPriority = function(obj){
	console.log("[anybase][setWithPriority]");
	console.log(obj);
};
Anybase.prototype.update = function(value){
	console.log("[anybase][update]");
};
Anybase.prototype.remove = function(){
	console.log("[anybase][remove]");
};

Anybase.prototype.on = function(event, callback, context){
	console.log("[anybase][on]" + event);
	return callback;
};
Anybase.prototype.onDisconnect = function(){
	console.log("[anybase][onDisconnect]");
	return new Disconnection();
};

Anybase.prototype.once = function(value, callback){
	console.log("[anybase][on]" + value);
};

//MISC functions
//Authenticates a Firebase client using the provided Authentication Token or Firebase Secret. See Custom Token Generation for details on authentication tokens. 
Anybase.prototype.auth = function(){
	console.log("[anybase][auth]");
};
Anybase.prototype.unauth = function(){
	console.log("[anybase][unauth]");
};

Anybase.prototype.toString = function(){
	console.log("[anybase][toString]");
};


showProps = function (obj, objName) {
  var result = "";
  for (var i in obj) {
    if (obj.hasOwnProperty(i)) {
        result += objName + "." + i + " = " + obj[i] + "\n";
    }
  }
  return result;
}