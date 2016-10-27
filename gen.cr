require "file_utils"
# replace with go

abort("Need a title") unless ARGV.size == 1
orig_title = ARGV[0]
#remove ruby dependency
title      = orig_title.dup.gsub(/\s+|['"\/]/,"_")
template   = File.join(File.dirname(__FILE__), "template.md")
date_fmt   = Time.now.to_s("%Y-%m-%d-%H-%M")

dest = "_posts/#{date_fmt}-#{title}.md"
FileUtils.cp(template, dest)

lines = File.read_lines(dest)
title_line = lines.each_with_index.find{|e,i| e =~ /^title:/}

if title_line
  lines[title_line[1]] = "title: #{orig_title}\n"

  File.open(dest,"w") do |f|
    f.print lines.join
  end
else
  raise "No title found"
end
