import smtplib
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart

def send_email(subject, body, from_email, to_email, smtp_server, smtp_port, password):
    # Create the email
    msg = MIMEMultipart()
    msg['From'] = from_email
    msg['To'] = to_email
    msg['Subject'] = subject

    # Attach the body text
    msg.attach(MIMEText(body, 'plain'))

    try:
        # Establish connection to the SMTP server
        server = smtplib.SMTP(smtp_server, smtp_port)
        server.starttls()  # Secure the connection
        server.login(from_email, password)  # Log in to the server
        
        # Send the email
        server.sendmail(from_email, to_email, msg.as_string())
        
        print("Email sent successfully!")
    
    except Exception as e:
        print(f"Error: {e}")
    
    finally:
        # Quit the SMTP server connection
        server.quit()

# User inputs
subject = "Test Email"
body = input("Enter the message to send: ")
from_email = "lama8050@gmail.com"
to_email = "connectwithsaange@gmail.com"
smtp_server = "smtp.gmail.com"
smtp_port = 587  # TLS SMTP port
password = input("Enter your email password: ")

send_email(subject, body, from_email, to_email, smtp_server, smtp_port, password)
