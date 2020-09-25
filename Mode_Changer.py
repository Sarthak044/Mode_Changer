import subprocess

def process():
    #Selecting the mode
    mode=input("Select the Mode\n1.Monitor Mode\n2.Managed Mode\n")

    if mode=='1':
        print("[+]Changing the mode to Monitor")
        subprocess.call(["ifconfig", interface, "down"])
        subprocess.call(["iwconfig", interface, "mode", "monitor"])
        subprocess.call(["ifconfig", interface, "up"])
    elif mode=='2':
        print("[+]Changing the Mode to Managed")
        subprocess.call(["ifconfig", interface, "down"])
        subprocess.call(["iwconfig", interface, "mode", "managed"])
        subprocess.call(["ifconfig", interface, "up"])
    else:
        print("\t**INVALID**\n\tEnter The correct OPTION")

#Selecting the Interface

print("[+]Make sure your Wireless Adapter supports Monitor mode/Packet Injection")
interface=input("Enter the Interface\n")

if interface=='eth0':
    print("Not a Wireless Adapter")
    print("BYE BYE!")
else:
    process()    



