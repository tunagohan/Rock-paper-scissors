#encoding: utf-8
#文字が入力された場合0になるのを利用したもの
# 1:グー
# 2:チョキ
# 3:パー

janken = ["Retry","グー","チョキ","パー"]
puts "じゃんけんゲームへようこそ"
puts "1,2,3以外の値が入力された場合"
puts "強制的に「グー」にします"
puts "*----------------------*"
puts "最初はグー！じゃんけん..."
puts "1:グー\n2:チョキ\n3:パー\n"

player = gets.to_i
cpu = rand(1..3)

while player == 0 || player > 3 do
    puts "それは受け付けんぞ！もっかいじゃ"
    puts "最初はグー！じゃんけん...\n"
    player = gets.to_i
end

while player == cpu || player == 0 || player > 3 do
    if player == 0 || player > 3 then
        puts "それは受け付けんぞ！もっかいじゃ"
    end
    puts "自分:#{janken[player]}\nCPU:#{janken[cpu]}"
    puts "あいこで..."
    player = gets.to_i
    cpu = rand(1..3)
end

res = player - cpu

if res == -1 || res == 2 then
    puts "自分:#{janken[player]}\nCPU:#{janken[cpu]}"
    puts "win!!"
else
    puts "自分:#{janken[player]}\nCPU:#{janken[cpu]}"
    puts "lose..."
end
