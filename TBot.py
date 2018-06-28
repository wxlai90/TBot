#!/usr/bin/python

import requests


class Telegram(object):

	def __init__(self, bot_token, callback_chat_id = ''):
		self.BOT_TOKEN = bot_token
		self.baseurl = 'https://api.telegram.org/bot{}/'.format(self.BOT_TOKEN)
		if callback_chat_id:
			self.callback_chat_id = callback_chat_id
		self.validCommands = ( '/help', '/start' )
		self.update_offset = '' #init


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

	def sendPhotoFile(self, filepath, chat_id, **kargs):
		f = open(filepath, 'rb')
		data = {}
		data['chat_id'] = chat_id
		files = { 'photo' : open(filepath, 'rb') }
		if 'caption' in kargs:
			files['caption'] = kargs['caption']
		r = requests.post(self.baseurl + 'sendPhoto', data = data, files = files)
		assert r.status_code == 200

	def sendAudioFile(self, filepath, chat_id, **kargs):
		f = open(filepath, 'rb')
		data = {}
		data['chat_id'] = chat_id
		files = { 'audio' : open(filepath, 'rb') }
		if 'caption' in kargs:
			files['caption'] = kargs['caption']
		r = requests.post(self.baseurl + 'sendAudio', data = data, files = files)
		assert r.status_code == 200

	def sendPhotoID(self, photo_id, chat_id, **kargs):
		data = {}
		data['chat_id'] = chat_id
		data['photo'] = photo_id
		if 'caption' in kargs:
			data['caption'] = kargs['caption']
		r = requests.post(self.baseurl + 'sendPhoto', data = data)
		assert r.status_code == 200


	def getUpdates(self):
		'''Gets updates via polling'''
		if self.update_offset:
			update_offset = int(self.update_offset + 1)
			r = requests.get(self.baseurl + 'getUpdates', data = { 'offset':update_offset })
		else:
			r = requests.get(self.baseurl + 'getUpdates')
		self.updates = r.json()
		updates = self.parseUpdates()
		if updates:
			self.update_offset = updates[-1]['update_id']
		return updates

	def parseUpdates(self):
		updates_list = []
		for i in self.updates['result']:
			updates = {}
			updates['update_id'] = i['update_id']
			updates['from_chat_id'] = i['message']['from']['id']
			try:
				updates['first_name'] = i['message']['from']['first_name']
			except KeyError as e:
				print 'Key not found in JSON update ({})'.format(e)
			try:
				updates['username'] = i['message']['from']['username']
			except KeyError as e:
				print 'Key not found in JSON update ({})'.format(e)
			try:
				updates['text'] = i['message']['text']
			except KeyError as e:
				print 'Key not found in JSON update ({})'.format(e)
			try:
				updates['photo'] = i['message']['photo'][-1]['file_id']
			except KeyError as e:
				print 'Key not found in JSON update ({})'.format(e)
			try:
				updates['caption'] = i['message']['caption']
			except KeyError as e:
				print 'Key not found in JSON update ({})'.format(e)
			updates_list.append(updates)
		return updates_list


	def processUpdate(self, command):
		cmd = command.lower().strip('/')
		if '/{}'.format(cmd) in self.validCommands:
			c = eval('self.{}'.format(cmd))
			c()
