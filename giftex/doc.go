/*
   Package giftex provides an implementation of the Kuhn-Munkres
   assignment algorithm for creating gift exchanges.

   Kuhn-Munkres, or the "Hungarian Algorithm" allows us to easily
   determine whether an assignment exists that can satisfy all
   constraints of the given gift exchange.

   Once the set of constraints are verified, a final assignment is
   made at random because a completely deterministic gift exchange
   would spoil the fun.
*/
package giftex
