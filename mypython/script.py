import camelcase
c = camelcase.CamelCase()
try:
  print("Start typing: ")
  while True:
    txt = input("")
    print(c.hump(txt))
except KeyboardInterrupt:
  print("Exiting")
