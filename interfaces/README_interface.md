# Interfaces: Notes

Notes from working through the `order` / `user` example in [`order/order.go`](../order/order.go) and [`user/user.go`](../user/user.go).

## 1. Compile-time check that a type implements an interface

```go
var _ order.UserProvider = User{}       // checks that User (value) implements it
var _ order.UserProvider = (*User)(nil) // checks that *User (pointer) implements it
```

- If a method is missing, the build fails with `User does not implement order.UserProvider (missing method GetEmail)`.
- `_` throws the value away, so there's no runtime cost.
- Use `T{}` when the methods have value receivers. It's the stricter check.
- Use `(*T)(nil)` when the methods have pointer receivers. It doesn't allocate anything.

Common mistakes:

- `(User)(nil)` fails with `cannot convert nil to type User`. `nil` only works for pointer, interface, slice, map, channel and function types, not structs.
- The left side must be an **interface** type. `order.Order` is a struct, so it can't be used there.

## 2. The problem: an import cycle

`Order` needs an owner. The obvious way:

```go
// order/order.go
import "myapp/user"

type Order struct {
	ID    int
	Owner user.User
}
```

But `user` already imports `order`, because `PlaceOrder` returns `order.Order`:

```
user  ──imports──▶ order
order ──imports──▶ user     ✗ import cycle not allowed
```

## 3. The fix: the consumer defines an interface

```go
// order/order.go (imports nothing)
type UserProvider interface {
	GetEmail() string
}

type Order struct {
	ID    int
	Owner UserProvider
}
```

```go
// user/user.go
import "myapp/order"

var _ order.UserProvider = User{}

func (u User) GetEmail() string { return u.Email }

func (u User) PlaceOrder(orderID int) order.Order {
	return order.Order{ID: orderID, Owner: u}
}
```

Step by step:

1. `Order` needs an owner, but it only cares about the owner's email.
2. So `order` defines a small interface for exactly that: `UserProvider { GetEmail() string }`.
3. `Owner` has the type `UserProvider` instead of `user.User`, so `order` never imports `user`.
4. `User` has a `GetEmail()` method, so it satisfies `UserProvider` automatically. Go has no `implements` keyword.
5. Only `user` imports `order`. The dependency goes one way, so there's no cycle.

```
        ┌──────────────── order ────────────────┐
        │  type UserProvider interface {         │
        │      GetEmail() string                 │
        │  }                                     │
        │  Order.Owner UserProvider              │
        └────────────────────▲───────────────────┘
                             │ imports and satisfies
        ┌──────────────── user ─────────────────┐
        │  User.GetEmail()                       │
        │  User.PlaceOrder() → order.Order       │
        └────────────────────────────────────────┘
```

## 4. Why this is dependency inversion

- **Normal flow:** the code that uses something depends on the concrete thing it uses. Order uses User, so Order would import User.
- **Inverted:** the code that uses something defines the abstraction it needs (`UserProvider`), and the concrete type conforms to it. The dependency arrow now points from User toward Order.
- **Go idiom:** define interfaces where they're *used*, not where they're *implemented*.
- **Keep them small:** ask only for what you use, here one method. That's the interface segregation principle.

This is a *partial* inversion. `user` still depends on the concrete `order.Order`. That's fine: only the direction that caused the cycle needed inverting.

## 5. What `order` can do with `Owner`

`order` can only call the methods the interface lists. Here that's `GetEmail()`, used for things like confirmations, shipping notices and receipts:

```go
func (o Order) Confirm() string {
	return fmt.Sprintf("Order #%d confirmed, mail sent to %s", o.ID, o.Owner.GetEmail())
}
```

- `o.Owner.ID` and `o.Owner.Email` don't compile, because `order` doesn't know `User` exists.
- An interface gives access to **behavior** (methods), not **properties** (fields).
- If Order later needs more, add a method such as `GetName() string` to `UserProvider`. Every owner type must then implement it, or the build fails.

Benefits:

- Any type with `GetEmail()` can own an order: `Admin`, `Guest`, `Company`.
- Tests can use a fake owner without touching `user`:

  ```go
  type fakeOwner struct{}
  func (fakeOwner) GetEmail() string { return "test@example.com" }
  ```

## 6. What is actually stored in `Owner`

`Owner: u` passes the full `User` struct. An interface value holds two things:

```
Owner (UserProvider)
 ├── type:  user.User
 └── value: copy of u  {ID: 1, Email: "a@b.com"}
```

- The full struct is stored, but `order` can only call `GetEmail()` on it.
- It's a **copy**. Changing `u.Email` afterward doesn't change the order.
- Pass `&u` to share the same user instead. `*User` also satisfies `UserProvider`, because a pointer type gets the value-receiver methods too.
- A type assertion (`o.Owner.(user.User)`) could get the `User` back, but it needs `import "myapp/user"`, which brings the cycle back.

## Summary

`order` sees the user only through `UserProvider`, never `User` itself. That breaks the import cycle and inverts the dependency.
