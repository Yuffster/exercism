(ns anagram
  (:require [clojure.string :as str]))

(defn norm-word [word]
  (str/lower-case word))

(defn norm-freq [word]
  (frequencies (norm-word word)))

(defn anagrams-for [word anagrams] 
  (filter #(and
            (not= (norm-word word) (norm-word %))
            (= (norm-freq word) (norm-freq %))) anagrams))
 
