class BaseModel {
  late final int albumId;
  late final int id;
  late final String title;
  late final String url;
  late final String thumbnailUrl;

  BaseModel.fromJson({required Map<String, dynamic> data})
      : albumId = data["albumId"],
        id = data["id"],
        title = data["title"],
        url = data["url"],
        thumbnailUrl = data["thumbnailUrl"];
}

class UserModel {
  final int userID;
  final String loginID;
  final String fristName;
  final String lastName;
  final String email;

  UserModel.fromJSON(Map<String, dynamic> json)
      : userID = json["id"],
        loginID = json["login_id"],
        fristName = json["first_name"],
        lastName = json["last_name"],
        email = json["email"];
}

class RestaurantModel {
  final int restaurantID;
  final String userID;
  final String name;
  final String restaurantType;
  final String origin;
  final List<String> menuTypes;

  RestaurantModel.fromJSON(Map<String, dynamic> json)
      : restaurantID = json["id"],
        userID = json["user_id"],
        name = json["name"],
        restaurantType = json["restaurant_type"],
        origin = json["origin"],
        menuTypes = json["menu_types"];
}

class MenuModel {
  final int restaurantID;
  final String name;
  final int price;
  final String menuType;
  final List<String> images;
  final String contents;
  final String allergies;

  MenuModel.fromJSON(Map<String, dynamic> json)
      : restaurantID = json["id"],
        name = json["name"],
        price = json["price"],
        menuType = json["menu_type"],
        images = json["images"],
        contents = json["contents"],
        allergies = json["allergies"];
}
