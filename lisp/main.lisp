(load "sonarsweep/main.lisp")
(defpackage :main
  (:use :cl)
  (:import-from sonarsweep :run-sonar-sweep)
)
(in-package :main)
(run-sonar-sweep)
