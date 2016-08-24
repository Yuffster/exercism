module Accumulate (accumulate) where

accumulate :: (a -> b) -> [a] -> [b]
accumulate fn l = [fn x | x <- l]