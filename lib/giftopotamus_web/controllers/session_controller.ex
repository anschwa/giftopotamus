defmodule GiftopotamusWeb.SessionController do
  use GiftopotamusWeb, :controller

  alias Giftopotamus.Auth

  def new(conn, _) do
    render(conn, "new.html")
  end

  def create(conn, %{"user" => %{"name" => name}}) do
    case Auth.authenticate_user(name) do
      {:ok, user} ->
        conn
        |> put_flash(:info, "Welcome Back!")
        |> put_session(:user_id, user.id)
        |> configure_session(renew: true)
        |> redirect(to: "/")

      {:error, :unauthorized} ->
        conn
        |> put_flash(:error, "Unknown user ")
        |> redirect(to: Routes.session_path(conn, :new))
    end
  end

  def delete(conn, _) do
    conn
    |> configure_session(drop: true)
    |> redirect(to: "/")
  end
end
