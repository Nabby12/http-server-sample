require "socket"

PAGES = {
    "/" => "static/index.html",
    "/showbanner" => "static/showbanner.html"
}

IMAGES = {
    "/banner" => "static/banner.png"
}

def handler()
    server = TCPServer.new 8080

    loop do
        session = server.accept

        request = []
        while (line = session.gets) && (line.chomp.length > 0)
            request << line.chomp
        end
        puts "finished reading"

        if request.empty? then
            print("!!! EMPTY !!!")
            session.close
            next
        end

        http_method, path, protocol = request[0].split(" ")
        puts http_method
        puts protocol

        contents = "HTTP/1.1 "
        if PAGES.keys.include? path
            contents << "200 OK\n\n"
            target = PAGES[path]
        else
            if IMAGES.keys.include? path
                contents << "200 OK\n\n"
                target = IMAGES[path]
            else
                contents << "404 Not Found\n\n"
                target = "static/404.html"
            end
        end

        file = File.open(target, "r")
        contents << file.read
        file.close

        session.puts contents

        session.close
    end
end

handler()
