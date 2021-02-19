defmodule GiftopotamusWeb.PageController do
  use GiftopotamusWeb, :controller

  alias Giftopotamus.Groups

  def index(conn, _params) do
    user_id = get_session(conn, :user_id)
    groups = Groups.list_user_groups(user_id)

    render(conn, "index.html", groups: groups)
  end
end
