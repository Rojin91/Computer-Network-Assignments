import smtplib
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart

# Email configuration
sender_email = "rojinpuri@gmail.com"
receiver_email = "purirojin@gmail.com"
password = "Rojin345"

# Create the email
subject = "Test Email"
body = "This is a test email sent from Python."

msg = MIMEMultipart()
msg['From'] = sender_email
msg['To'] = receiver_email
msg['Subject'] = subject

msg.attach(MIMEText(body, 'plain'))

# SMTP server configuration
smtp_server = "smtp.gmail.com"
smtp_port = 587

try:
    # Connect to the SMTP server
    server = smtplib.SMTP(smtp_server, smtp_port)
    server.starttls()  # Secure the connection

    # Login to the email account
    server.login(sender_email, password)

    # Send the email
    text = msg.as_string()
    server.sendmail(sender_email, receiver_email, text)

    print("Email sent successfully!")

except Exception as e:
    print(f"Error: {e}")

finally:
    server.quit()
