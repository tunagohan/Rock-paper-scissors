#encoding: utf-8
#文字が入力された場合0になるのを利用したもの
# 1:グー
# 2:チョキ
# 3:パー
win = 0
lose = 0
point = 0
continue = 0

janken = ["Retry","グー","チョキ","パー"]

puts "じゃんけんゲームへようこそ"
puts "1,2,3以外の値が入力された場合"
puts "強制的に「グー」にします"
puts "*----------------------*"

while continue == 0 do
    cont_res = 0
    puts "最初はグー！じゃんけん..."
    puts "1:グー\n2:チョキ\n3:パー\n"
    puts "*----------------------*"
    print "何を出す？："
    player = gets.to_i
    cpu = rand(1..3)

    while player == 0 || player > 3 do
        puts "それは受け付けんぞ！もっかいじゃ"
        puts "最初はグー！じゃんけん...\n"
        print "何を出す？："
        player = gets.to_i
        puts "*----------------------*"
    end

    while player == cpu || player == 0 || player > 3 do
        if player == 0 || player > 3 then
            puts "それは受け付けんぞ！もっかいじゃ"
        end
        puts "自分:#{janken[player]}\nCPU:#{janken[cpu]}"
        puts "*----------------------*"
        puts "あいこで..."
        print "何を出す？："
        player = gets.to_i
        cpu = rand(1..3)
    end

    res = player - cpu

    if res == -1 || res == 2 then
        puts "自分:#{janken[player]}\nCPU:#{janken[cpu]}"
        puts "win!!"
        win += 1
        puts "*----------------------*"
    else
        puts "自分:#{janken[player]}\nCPU:#{janken[cpu]}"
        puts "lose..."
        lose += 1
        puts "*----------------------*"
    end

    point = win - lose
    puts "*---*:Score:*---*"
    puts "WIN:#{win}回"
    puts "LOSE:#{lose}回"
    puts "Point:#{point}ポイント!\n"
    puts "*---*-------*---*"

    while cont_res == 0 do
        puts "続ける場合は[y]を押してね！"
        puts "やめるなら[n]を押してね..."
        print "どうする？："
        get_res = gets.chomp
        case get_res
        when /^[yY]/, "yes", "YES", "No"
            puts "*----------------------*"
            puts "OK続けよう！"
            continue = 0
            cont_res = 1
        when /^[nN]/, /^$/, "No", "no"
            puts "遊んでくれてありがとう！"
            puts "*----------------------*"
            continue = 1
            cont_res = 1
        else
            puts "[y][n]を入力してね？"
            cont_res = 0
        end
    end
end
