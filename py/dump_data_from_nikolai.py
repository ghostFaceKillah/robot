from scipy import io


def get_lidar(file_name):
    """
    lidar_data type: array where each element is a dictionary with a form of
                     't','pose','res','rpy','scan'

    x['ts']: timestamps (Absolute time)
    x['delta_pose']: relative odometry between last reading [x, y, theta]
                     (+x: forward, +y: left, +theta: counterclockwise rotation around z)
    x['scan']: 1x1081 lidar scan data, range -135° to 135°.
    """
    data = io.loadmat(file_name+".mat")
    lidar = []
    for m in data['lidar'][0]:
        tmp = {}
        tmp['t']= m[0][0][0]
        nn = len(m[0][0])
        if (nn != 3):
            raise ValueError("different length!")
        tmp['delta_pose'] = m[0][0][nn-1]
        tmp['scan'] = m[0][0][nn-2]
        lidar.append(tmp)
    return lidar