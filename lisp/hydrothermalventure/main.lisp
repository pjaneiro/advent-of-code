(defpackage :hydrothermalventure
  (:use :cl)
  (:export :run-hydrothermalventure)
)
(in-package :hydrothermalventure)

(defstruct point
  x
  y
)

(defun delimiterp (c) (or (char= c #\Space) (char= c #\,)))

(defun split (string &key (delimiterp #'delimiterp))
  (loop :for beg = (position-if-not delimiterp string)
    :then (position-if-not delimiterp string :start (1+ end))
    :for end = (and beg (position-if delimiterp string :start beg))
    :when beg :collect (subseq string beg end)
    :while end))

(defun parse-vector (input)
  (setq result (make-list 2))
  (setf (nth 0 result) (make-point :x (parse-integer (nth 0 input)) :y (parse-integer (nth 1 input))))
  (setf (nth 1 result) (make-point :x (parse-integer (nth 3 input)) :y (parse-integer (nth 4 input))))
  result
)

(defun get-file (filename)
  (with-open-file (stream filename)
    (loop for line = (read-line stream nil)
      while line
      collect (parse-vector (split line))
    )
  )
)

(defun create-board (x y)
  (setq board (make-list x :initial-element '()))
  (loop for i from 0 to (- x 1)
    do (setf (nth i board) (make-list y :initial-element 0))
  )
  board
)

(defun maxes (data)
  (setq result (make-list 2 :initial-element 0))
  (dolist (cur data)
    (if (> (point-x (nth 0 cur)) (nth 0 result)) (setf (nth 0 result) (point-x (nth 0 cur))))
    (if (> (point-x (nth 1 cur)) (nth 0 result)) (setf (nth 0 result) (point-x (nth 1 cur))))
    (if (> (point-y (nth 0 cur)) (nth 1 result)) (setf (nth 1 result) (point-y (nth 0 cur))))
    (if (> (point-y (nth 1 cur)) (nth 1 result)) (setf (nth 1 result) (point-y (nth 1 cur))))
  )
  result
)

(defun mark-points (board fromX fromY toX toY)
  (if (= fromX toX) (progn
    (if (< fromY toY) (progn
        (loop for j from fromY to toY
          do (setf (nth j (nth fromX board)) (+ (nth j (nth fromX board)) 1))
        )
      ) (progn
        (loop for j from toY to fromY
          do (setf (nth j (nth fromX board)) (+ (nth j (nth fromX board)) 1))
        )
      )
    )
    (return-from mark-points)
  ))
  (if (= fromY toY) (progn
    (if (< fromX toX) (progn
        (loop for i from fromX to toX
          do (setf (nth fromY (nth i board)) (+ (nth fromY (nth i board)) 1))
        )
      ) (progn
        (loop for i from toX to fromX
          do (setf (nth fromY (nth i board)) (+ (nth fromY (nth i board)) 1))
        )
      )
    )
    (return-from mark-points)
  ))
  (if (< fromX toX) (progn
    (setq steps (- toX fromX))
    (if (< fromY toY) (progn
        (loop for i from 0 to steps
          do (setf (nth (+ fromY i) (nth (+ fromX i) board)) (+ (nth (+ fromY i) (nth (+ fromX i) board)) 1))
        )
      ) (progn
        (loop for i from 0 to steps
          do (setf (nth (- fromY i) (nth (+ fromX i) board)) (+ (nth (- fromY i) (nth (+ fromX i) board)) 1))
        )
      )
    )
  ) (progn
    (setq steps (- fromX toX))
    (if (< fromY toY) (progn
        (loop for i from 0 to steps
          do (setf (nth (+ fromY i) (nth (- fromX i) board)) (+ (nth (+ fromY i) (nth (- fromX i) board)) 1))
        )
      ) (progn
        (loop for i from 0 to steps
          do (setf (nth (- fromY i) (nth (- fromX i) board)) (+ (nth (- fromY i) (nth (- fromX i) board)) 1))
        )
      )
    )
  ))
)

(defun count-points (board)
  (setq result 0)
  (dolist (row board)
    (dolist (elem row)
      (if (>= elem 2) (setq result (+ result 1)))
    )
  )
  result
)

(defun challenge1 (inputData)
  (setq localMaxes (maxes inputData))
  (setq board (create-board (+ 1 (nth 0 localMaxes)) (+ 1 (nth 1 localMaxes))))
  (dolist (cur inputData)
    (setq fromX (point-x (nth 0 cur)))
    (setq fromY (point-y (nth 0 cur)))
    (setq toX (point-x (nth 1 cur)))
    (setq toY (point-y (nth 1 cur)))
    (if (or (= fromX toX) (= fromY toY)) (mark-points board fromX fromY toX toY))
  )
  (count-points board)
)

(defun challenge2 (inputData)
  (setq localMaxes (maxes inputData))
  (setq board (create-board (+ 1 (nth 0 localMaxes)) (+ 1 (nth 1 localMaxes))))
  (dolist (cur inputData)
    (setq fromX (point-x (nth 0 cur)))
    (setq fromY (point-y (nth 0 cur)))
    (setq toX (point-x (nth 1 cur)))
    (setq toY (point-y (nth 1 cur)))
    (mark-points board fromX fromY toX toY)
  )
  (count-points board)
)

(defun run-hydrothermalventure ()
  (format t "Day 5 - Hydrothermal Venture~%")
  (setq fileData (get-file "hydrothermalventure/input.txt"))
  (format t "Challenge 1: ~d~%" (challenge1 fileData))
  (format t "Challenge 2: ~d~%" (challenge2 fileData))
)
