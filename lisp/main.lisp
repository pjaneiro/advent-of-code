(load "sonarsweep/main.lisp")
(load "dive/main.lisp")
(defpackage :main
  (:use :cl)
  (:import-from sonarsweep :run-sonar-sweep)
  (:import-from dive :run-dive)
)
(in-package :main)

(defun newline ()
  (format t "~C" #\linefeed)
)

(run-sonar-sweep)
(newline)
(run-dive)
(newline)
