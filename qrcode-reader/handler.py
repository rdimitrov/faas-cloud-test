import qrtools

def handle(req):
    """handle a request to the function
    Args:
        req (str): request body
    """
    if req:
        with open("/tmp/qrcode.png", "wb") as f:
            f.write(req)
        qr = qrtools.QR()
        qr.decode("/tmp/qrcode.png")
        return "Your QR says:", qr.data
    else:
        return "No QR image sent :)"