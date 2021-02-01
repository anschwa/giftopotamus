defmodule GiftopotamusWeb.PageController do
  use GiftopotamusWeb, :controller

  def index(conn, _params) do
    render(conn, "index.html")
  end
end
