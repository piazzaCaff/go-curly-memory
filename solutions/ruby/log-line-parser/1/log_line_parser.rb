class LogLineParser
  def initialize(line)
    @line = line
  end

  def message
    @line.split(":")[1].strip
  end

  def log_level
   @line.split(":")[0].delete("[]").downcase
  end

  def reformat
    first = @line.split(":")[0].strip.downcase
    second = @line.split(":")[1].strip
    
    level = first.gsub('[', '(').gsub(']', ')')
  
    return second + " " + level
  
  end
end
