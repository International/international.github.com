require "fileutils"
# replace with go

orig_title = ARGV[0]
#remove ruby dependency
title      = orig_title.dup.gsub(/\s+|['"\/]/,"_")
template   = File.join(File.dirname(__FILE__), "template.md")
date_fmt   = Time.now.strftime("%Y-%m-%d-%H-%M")

dest = "_posts/#{date_fmt}-#{title}.md"
FileUtils.cp(template, dest)

lines = File.readlines(dest)
idx = lines.find_index{|e| e =~ /^title:/}
raise "No title found" unless idx
lines[idx] = "title: #{orig_title}\n"

File.open(dest,"w") do |f|
  f.print lines.join
end

