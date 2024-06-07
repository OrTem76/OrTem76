import time
import os
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from webdriver_manager.chrome import ChromeDriverManager
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys


driver = webdriver.Chrome(service=Service(ChromeDriverManager().install()))

url = "https://srdata.sandbox34.ru/"

try:
    driver.maximize_window()
    driver.get(url=url)
    time.sleep(5)
    
    current_title = driver.title
    print("Текущий заголовок:",current_title)
    
    
    assert url == "https://srdata.sandbox34.ru/", "Error in URL"
    #print(driver.page_source)
    #print(driver.find_element("id", "username"))
    #print(type(driver.find_element("id", "username")))
    email_input = driver.find_element(By.ID, 'username').click()
    #email_input.clear()
    email_input = driver.find_element(By.ID, 'username').click()
    driver.find_element(By.ID, 'username').send_keys("ssrdatatest@gmail.com")
    
    time.sleep(5)
    
    driver.find_element("id","current-password").click()
    #pass_input = driver.find_element(By.ID, 'current-password').click()
    #pass_input.clear()
    #pass_input = driver.find_element(By.ID, 'current-password').click()
    driver.find_element(By.ID, 'current-password').send_keys("Ssrdatatest2")
    
    time.sleep(5)
    
    button_eye = driver.find_element(By.XPATH, '/html/body/div/div/main/div/form/div[1]/div[2]/span/button').click()
     
    time.sleep(5)
     
    button_vhod = driver.find_element(By.XPATH, "/html/body/div/div/main/div/form/div[2]/button[2]").click()
    
    
    time.sleep(5)
    
except Exception as ex:
    print(ex)
    
finally:
    driver.close()
    driver.quit()