module Sublist (Sublist(..), sublist) where

data Sublist = Equal | Sublist | Superlist | Unequal deriving (Show, Eq)

sublist :: Eq a => [a] -> [a] -> Sublist
sublist [] []      = Equal
sublist _ []       = Superlist
sublist [] _       = Sublist
sublist a b
    | a == b        = Equal
    | subsetOf a b  = Sublist
    | subsetOf b a  = Superlist
    | otherwise     = Unequal

-- True if a is a subset of b.
subsetOf :: Eq a => [a] -> [a] -> Bool
subsetOf a [] = False
subsetOf a b
    | length a > length b    = False -- Can't be a superset if it's bigger.
    | a == take (length a) b = True  -- Could be a prefix.
    | otherwise              = subsetOf a (tail b)