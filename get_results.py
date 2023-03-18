from selenium import webdriver
from selenium.webdriver.common.by import By
import time

driver = webdriver.Chrome()
driver.implicitly_wait(30)
driver.get("https://tttv.click-tt.de/cgi-bin/WebObjects/nuLigaTTDE.woa/wa/home")

driver.find_element(By.XPATH, '//*[@id="content-col2"]/form/fieldset/input[1]').send_keys('hoffelmann@gmail.com')
driver.find_element(By.XPATH, '//*[@id="content-col2"]/form/fieldset/input[4]').click()
driver.find_element(By.XPATH, '//*[@id="password"]').send_keys('AJqE3JsTLDEbx6N')
driver.find_element(By.XPATH, '//*[@id="submit_0"]').click()
driver.find_element(By.XPATH, '//*[@id="tabs"]/ul/li[7]/a').click()
driver.find_element(By.XPATH, '//*[@id="content-row1"]/table[1]/tbody/tr[6]/td[1]/a[2]').click()
time.sleep(10)
driver.find_element(By.XPATH, '//*[@id="content-row1"]/form/p[2]/a').click()
input('Press ENTER to finish')
