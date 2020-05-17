from scipy import io
import numpy as np

import sensor_pb2

LIDAR_FILE_TEMPLATE = '/Users/misiu-dev/src/go/robot/doc/assignments/ECE276A_P2_2020/lidar/train_lidar{:d}.mat'
OUT_FILE_TEMPLATE = '/Users/misiu-dev/src/go/robot/data/lidar_data_{}.proto'


def lidar_to_protobuf_msg(file_name):
    """
    Read in matlab datafile, serialize it to protobuf.

    Notes on data types, extracted from the notes etc.
    -------------------
    lidar_data type: array where each element is a dictionary with a form of
                     't','pose','res','rpy','scan'

    x['ts']: timestamps (Absolute time)
    x['delta_pose']: relative odometry between last reading [x, y, theta]
                     (+x: forward, +y: left, +theta: counterclockwise rotation around z)
    x['scan']: 1x1081 lidar scan data, range -135° to 135°.
    """
    data = io.loadmat(file_name)

    acc = []

    for idx, m in enumerate(data['lidar'][0]):
        if idx % 2000 == 0:
            print("doing {} out of {}".format(idx, len(data['lidar'][0])))
        t = m[0][0][0][0][0]

        nn = len(m[0][0])
        if (nn != 3):
            raise ValueError("different length!")
        delta_pose = m[0][0][nn-1][0]
        scan = m[0][0][nn-2][0]

        # make lidar scan message
        scan_msg = sensor_pb2.Lidar2DScan()
        scan_msg.timestamp = t
        for range in scan:
            scan_msg.scan.append(range)
        scan_msg.n_scans = len(scan)
        scan_msg.min_range = -135. / 360. * 2 * np.pi
        scan_msg.max_range = 135. / 360. * 2 * np.pi
        scan_msg.scan_resolution = (scan_msg.max_range - scan_msg.min_range) / (scan_msg.n_scans - 1)

        # make odom update message
        if len(delta_pose.shape) == 1:
            odom_msg = sensor_pb2.Odom2DUpdate(
                timestamp=t,
                x=delta_pose[0],
                y=delta_pose[1],
                theta=delta_pose[2],
            )
        else:
            odom_msg = sensor_pb2.Odom2DUpdate(
                timestamp=t,
                x=delta_pose[0][0],
                y=delta_pose[0][1],
                theta=delta_pose[0][2]
            )

        # make the summary message
        # nikolai_msg = nikolai_acc_msgs.scans.add()
        nikolai_msg = sensor_pb2.NikolaiScan(
            odom=odom_msg,
            lidar=scan_msg
        )
        acc.append(nikolai_msg)

    nikolai_acc_msgs = sensor_pb2.NikolaiManyScans()
    nikolai_acc_msgs.scans.extend(acc)

    return nikolai_acc_msgs


def write_all_files_to_serialized():
    for i in range(5):
        in_fname = LIDAR_FILE_TEMPLATE.format(i)
        out_fname = OUT_FILE_TEMPLATE.format(i)
        print("Converting {} to protobuf...".format(in_fname))
        lidar_msg = lidar_to_protobuf_msg(in_fname)
        with open(out_fname, 'wb') as f:
            print("Writing data to {}".format(out_fname))
            f.write(lidar_msg.SerializeToString())
        print("done!")


if __name__ == '__main__':
    write_all_files_to_serialized()



