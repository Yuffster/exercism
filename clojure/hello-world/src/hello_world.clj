(ns hello-world)

(defn hello
  ([] (hello "World"))
  ([name] (apply str ["Hello, " name "!"])))
