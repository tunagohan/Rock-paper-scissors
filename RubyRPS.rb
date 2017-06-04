#encoding: utf-8
# 0:グー
# 1:チョキ
# 2:パー

puts "最初はグー！じゃんけん..."
puts "0:グー\n1:チョキ\n2:パー\n"
player = 4
while player > 3 do
    player = gets.to_i
    if player > 3 || integer_string?(player)==false then
        puts "それは受け付けんぞ！もっかいじゃ"
        puts "最初はグー！じゃんけん..."
    end
end
janken = ["グー","チョキ","パー"]
cpu = rand(3)

while player == cpu do
    puts "自分:#{janken[player]}\nCPU:#{janken[cpu]}"
    puts "あいこで..."
    player = gets.to_i
    cpu = rand(3)
end
res = player - cpu
if res == -1 || res == 2 then
    puts "自分:#{janken[player]}\nCPU:#{janken[cpu]}"
    puts "win!!"
else
    puts "自分:#{janken[player]}\nCPU:#{janken[cpu]}"
    puts "lose..."
end
