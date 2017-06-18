# -*- coding: utf-8 -*-
import time
import random

count = 0
hand_dict = {"1":1, "2":2, "3":3}
hand_print_dict = {1:"グー", 2:"チョキ", 3:"パー"}

print "*------------------------*"
print "じゃんけんゲームへようこそ！"
print "3本どちらかが勝てば終了！"
print "*------------------------*"

time.sleep(0.5)

# 初期化
flg_draw = False
win = [0]*2

while count == 0:

    cpu = random.randint(1, 3)

    while 1:
        if flg_draw:
            print "あいこで..."
            time.sleep(0.5)
        else:
            print "じゃんけん..."
            time.sleep(0.5)

        player = raw_input("(何を出しますか?)\n1:グー\n2:チョキ\n3:パー\n>>> ")
        if player not in ("1", "2", "3"):
            print "不正な文字です"
            time.sleep(0.5)
        else:
            phand = hand_dict[player]
            break

    if flg_draw:
        print "しょ!"
    else:
        print "ぽん!"

    print "《あなた%s》vs《あいて:%s》" % (hand_print_dict[phand], hand_print_dict[cpu])
    time.sleep(0.5)

    if (3 + hand_dict[player] - cpu) % 3 == 2:
        flg_draw = False
        win[0] += 1
        print "【勝ち】"
        print "*--------------------------*"
        time.sleep(0.5)
    elif (3 + hand_dict[player] - cpu) % 3 == 1:
        flg_draw = False
        win[1] += 1
        print "【負け】"
        print "*--------------------------*"
        time.sleep(0.5)
    else:
        flg_draw = True
        print "【あいこ】"
        print "*--------------------------*"
        time.sleep(0.5)
        continue

    print "《あなた:%s勝》《あいて:%s勝》" % (win[0], win[1])
    print "*--------------------------*"
    print "続けるの？[終わる：q][続ける：y]"
    cont = raw_input()
    if cont not in ("q","y"):
        print "不正な文字です"
    elif cont == "q":
        count = 1
    else:
        print "やったー！\nじゃあいっくよー？"


print "*------------------------*"
print "遊んでくれてありがとう！"
time.sleep(1.5)
