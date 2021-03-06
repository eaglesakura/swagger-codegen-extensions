part of client_petstore.api;

class Pet {
  
  int id = null;
  

  Category category = null;
  

  String name = null;
  

  List<String> photoUrls = [];
  

  List<Tag> tags = [];
  
/* pet status in the store */
  String status = null;
  //enum statusEnum {  available,  pending,  sold,  };
  Pet();

  @override
  String toString() {
    return 'Pet[id=$id, category=$category, name=$name, photoUrls=$photoUrls, tags=$tags, status=$status, ]';
  }

  Pet.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
      if(json.containsKey("id")) {
              id =
                      json['id']
              ;
      }
      if(json.containsKey("category")) {
              category =
                  
                      
                      new Category.fromJson(json['category'])
;
      }
      if(json.containsKey("name")) {
              name =
                      json['name']
              ;
      }
      if(json.containsKey("photoUrls")) {
              photoUrls =
                      (json['photoUrls'] as List).map((item) => item as String).toList()
              ;
      }
      if(json.containsKey("tags")) {
              tags =
                  Tag.listFromJson(json['tags'])
;
      }
      if(json.containsKey("status")) {
              status =
                      json['status']
              ;
      }
  }


  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'category': category,
      'name': name,
      'photoUrls': photoUrls,
      'tags': tags,
      'status': status
     };
  }

  static List<Pet> listFromJson(List<dynamic> json) {
    return json == null ? new List<Pet>() : json.map((value) => new Pet.fromJson(value)).toList();
  }

  static Map<String, Pet> mapFromJson(Map<String, Map<String, dynamic>> json) {
    var map = new Map<String, Pet>();
    if (json != null && json.length > 0) {
      json.forEach((String key, Map<String, dynamic> value) => map[key] = new Pet.fromJson(value));
    }
    return map;
  }
}

