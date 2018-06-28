#!/usr/bin/python

import requests


class Telegram(object):

	def __init__(self, bot_token, callback_chat_id = ''):
		self.BOT_TOKEN = bot_token
		self.baseurl = 'https://api.telegram.org/bot{}/'.format(self.BOT_TOKEN)
		if callback_chat_id:
			self.callback_chat_id = callback_chat_id
		self.validCommands = ( '/help', '/start' )


	def sendTxtMessage(self, content, chat_id):
		data = {}
		data['chat_id'] = chat_id
		data['text'] = content 
		r = requests.post(self.baseurl + 'sendMessage', data = data)
		assert r.status_code == 200


	def sendHTMLMessage(self, content, chat_id):
		data = {}
		data['chat_id'] = chat_id
		data['text'] = content
		data['parse_mode'] = 'HTML'
		r = requests.post(self.baseurl + 'sendMessage', data = data)
		assert r.status_code == 200

	def sendPhotoFile(self, filepath, chat_id):
		f = open(filepath, 'rb')
		data = {}
		data['chat_id'] = chat_id
		files = { 'photo' : open(filepath, 'rb') }
		r = requests.post(self.baseurl + 'sendPhoto', data = data, files = files)
		assert r.status_code == 200

	def sendAudioFile(self, filepath, chat_id):
		f = open(filepath, 'rb')
		data = {}
		data['chat_id'] = chat_id
		files = { 'audio' : open(filepath, 'rb') }
		r = requests.post(self.baseurl + 'sendAudio', data = data, files = files)
		assert r.status_code == 200

	def sendPhotoID(self, photo_id, chat_id):
		data = {}
		data['chat_id'] = chat_id
		data['photo'] = photo_id
		r = requests.post(self.baseurl + 'sendPhoto', data = data)
		assert r.status_code == 200


	def getUpdates(self):
		r = requests.get(self.baseurl + 'getUpdates')
		self.updates = r.json()
		
	
	def processUpdate(self, command):
		cmd = command.lower().strip('/')
		if '/{}'.format(cmd) in self.validCommands:
			c = eval('self.{}'.format(cmd))
			c()
