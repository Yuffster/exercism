(ns bob
   (:require [clojure.string :as string]))

(defn- uppercase? [phrase]
  (and (re-find #"[A-Z]+" phrase)
       (= phrase (string/upper-case phrase))))

(defn- question? [phrase]
  (string/ends-with? phrase "?"))

(defn- yelling? [phrase] (uppercase? phrase))

(defn- silence? [phrase]
  (or (= phrase "")
      (re-find #"^\W+$" phrase)))

(defn response-for [phrase]
  (cond
    (yelling? phrase) "Whoa, chill out!"
    (question? phrase) "Sure."
    (silence? phrase) "Fine. Be that way!"
    :else "Whatever."))
