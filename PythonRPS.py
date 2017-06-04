# !/usr/bin/env python
# -*- coding: utf-8 -*-

import random

janken = {1:"ぐー",2:"チョキ",3:"パー"}

print "じゃんけんゲームへようこそ"
print "1,2,3以外の値が入力された場合"
print "強制的に「グー」にします"
print "*----------------------*"
print "最初はグー！じゃんけん..."
print "1:グー\n2:チョキ\n3:パー\n"

res = False
win =[0]*2

cpu = random.randint(1,3)
