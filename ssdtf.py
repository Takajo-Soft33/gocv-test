# -*- coding: utf-8 -*-

import os
import math
import random
import numpy as np
import tensorflow as tf
import cv2

slim = tf.contrib.slim

import sys
sys.path.append('ssdtf/')

from nets import ssd_vgg_300, ssd_common, np_methods
from preprocessing import ssd_vgg_preprocessing


gpu_options = tf.GPUOptions(allow_growth = True)
config = tf.ConfigProto(log_device_placement = False, gpu_options = gpu_options)
isess = tf.InteractiveSession(config = config)

net_shape = (300, 300)
data_format = 'NHWC'
img_input = tf.placeholder(tf.uint8, shape = (None, None, 3))

image_pre, labels_pre, bboxes_pre, bbox_img = ssd_vgg_preprocessing.preprocess_for_Eval(img_input, None, None, net_shape, data_format, resize = ssd_vgg_preprocessing.Resize.WARP_RESIZE))
image_4d = tf.expand_dims(image_pre, 0)

reuse = True if 'ssd_net' in locals() else None
ssd_net = ssd_vgg_300.SSDNet()
with slim.arg_scope(ssd_net.arg_scope(data_format = data_format)):
  predictions, localisations, _, _ = ssd_net.net(image_4d, is_training = False, reuse = reuse)

ckpt_filename = 'ssdtf/checkpoints/ssd_300_vgg.ckpt'
isess.run(tf.global_variables_initializer())
saver = tf.train.Saver()
saver.restore(isess, ckpt_filename)

ssd_anchors = ssd_net.anchors(net_shape)

def process_image(img, select_threshold=0.5, nms_threshold=.45, net_shape=(300, 300)):
rimg, rpredictions, rlocalisations, rbbox_img = isess.run([image_4d, predictions, localisations, bbox_img], feed_dict = {img_input: img})
