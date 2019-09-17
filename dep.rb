require 'fileutils'

def dep(name, ver)
  dir = "#{__dir__}/vendor/#{name}"
  FileUtils.mkdir_p(File.dirname(dir))

  Dir.chdir(File.dirname(dir)) {
    # hmm...
    checkout_name = name.gsub('golang.org/x/', 'github.com/golang/')
    system("git submodule add -f https://#{checkout_name}")
  } unless File.directory?(dir)

  Dir.chdir(dir) {
    system("git checkout #{ver}")
  }
end

name = nil

File.readlines("#{__dir__}/Gopkg.lock").each { |l|
  next unless l.include?(?=)
  k, v = l.split(?=)
  name = v.strip.delete(?") if k.strip == 'name'
  dep(name, v.strip.delete(?")) if name != 'github.com/exercism/cli' && k.strip == 'revision'
}
