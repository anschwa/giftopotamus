defmodule GiftopotamusWeb.PageController do
  use GiftopotamusWeb, :controller

  alias Giftopotamus.Groups

  def index(conn, _params) do
    groups = Groups.list_user_groups(conn.assigns.current_user)
    render(conn, "index.html", groups: groups)
  end
end
