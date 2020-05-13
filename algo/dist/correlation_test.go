package dist

// map correlation here
// re-implemented p2_utils

/*

def test_mapCorrelation():
  dataIn = io.loadmat("lidar/train_lidar0.mat")
  angles = np.array([np.arange(-135,135.25,0.25)*np.pi/180.]).T
  ranges = np.double(dataIn['lidar'][0][110]['scan'][0][0]).T

  # take valid indices
  indValid = np.logical_and((ranges < 30),(ranges> 0.1))
  ranges = ranges[indValid]
  angles = angles[indValid]

  # init MAP
  MAP = {}
  MAP['res']   = 0.05 #meters
  MAP['xmin']  = -20  #meters
  MAP['ymin']  = -20
  MAP['xmax']  =  20
  MAP['ymax']  =  20
  MAP['sizex']  = int(np.ceil((MAP['xmax'] - MAP['xmin']) / MAP['res'] + 1)) #cells
  MAP['sizey']  = int(np.ceil((MAP['ymax'] - MAP['ymin']) / MAP['res'] + 1))
  MAP['map'] = np.zeros((MAP['sizex'],MAP['sizey']),dtype=np.int8) #DATA TYPE: char or int8

  # xy position in the sensor frame
  xs0 = np.array([ranges*np.cos(angles)])
  ys0 = np.array([ranges*np.sin(angles)])

  # convert position in the map frame here
  Y = np.concatenate([np.concatenate([xs0,ys0],axis=0),np.zeros(xs0.shape)],axis=0)

  # convert from meters to cells
  xis = np.ceil((xs0 - MAP['xmin']) / MAP['res'] ).astype(np.int16)-1
  yis = np.ceil((ys0 - MAP['ymin']) / MAP['res'] ).astype(np.int16)-1

  # build an arbitrary map
  indGood = np.logical_and(np.logical_and(np.logical_and((xis > 1), (yis > 1)), (xis < MAP['sizex'])), (yis < MAP['sizey']))
  MAP['map'][xis[0][indGood[0]],yis[0][indGood[0]]]=1

  x_im = np.arange(MAP['xmin'],MAP['xmax']+MAP['res'],MAP['res']) #x-positions of each pixel of the map
  y_im = np.arange(MAP['ymin'],MAP['ymax']+MAP['res'],MAP['res']) #y-positions of each pixel of the map

  x_range = np.arange(-0.2,0.2+0.05,0.05)
  y_range = np.arange(-0.2,0.2+0.05,0.05)

  print("Testing map_correlation...")
  c = mapCorrelation(MAP['map'],x_im,y_im,Y[0:3,:],x_range,y_range)

  c_ex = np.array([[3,4,8,162,270,132,18,1,0],
		  [25  ,1   ,8   ,201  ,307 ,109 ,5  ,1   ,3],
		  [314 ,198 ,91  ,263  ,366 ,73  ,5  ,6   ,6],
		  [130 ,267 ,360 ,660  ,606 ,87  ,17 ,15  ,9],
		  [17  ,28  ,95  ,618  ,668 ,370 ,271,136 ,30],
		  [9   ,10  ,64  ,404  ,229 ,90  ,205,308 ,323],
		  [5   ,16  ,101 ,360  ,152 ,5   ,1  ,24  ,102],
		  [7   ,30  ,131 ,309  ,105 ,8   ,4  ,4   ,2],
		  [16  ,55  ,138 ,274  ,75  ,11  ,6  ,6   ,3]])

  if np.sum(c==c_ex) == np.size(c_ex):
	  print("...Test passed.")
  else:
	  print("...Test failed. Close figures to continue tests.")


  fig = plt.figure(figsize=(18,6))

  #plot original lidar points
  ax1 = fig.add_subplot(131)
  plt.plot(xs0,ys0,'.k')
  plt.scatter(0,0,s=30,c='r')
  plt.xlabel("x")
  plt.ylabel("y")
  plt.title("Laser reading (red being robot location)")
  plt.axis('equal')

  #plot map
  ax2 = fig.add_subplot(132)
  plt.imshow(MAP['map'],cmap="hot")
  plt.title("Occupancy map")

  #plot correlation
  ax3 = fig.add_subplot(133,projection='3d')
  X, Y = np.meshgrid(np.arange(0,9), np.arange(0,9))
  ax3.plot_surface(X,Y,c,linewidth=0,cmap=plt.cm.jet, antialiased=False,rstride=1, cstride=1)
  plt.title("Correlation coefficient map")

  plt.show()




*/
