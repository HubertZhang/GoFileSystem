__author__ = 'Hubert'

import requests
import json
import os
import random
import time

server = "http://localhost:4000/kv/"
server_admin = "http://localhost:4000/kvman/"
backup_admin = "http://localhost:5000/kvman/"

def start_primary():
    os.system("../start_server –p")

def start_backup():
    os.system("../start_server –b")

def stop_primary():
    r = requests.get(server_admin + "shutdown")
    assert (r.status_code == 200)
    return

def stop_backup():
    r = requests.get(backup_admin + "shutdown")
    assert (r.status_code == 200)
    return

def test_primary():
    for key in backup.keys():
        if (get(key)["value"]!=backup[key]):
            return False
    return True


def get(key=""):
    if (key == ""):
        return
    params={"key":key}
    r = requests.get(server + "get", params = params)
    return r.json()


def delete(key=""):
    if (key == ""):
        return
    payload = {'key': key}
    r = requests.post(server + "delete", data=json.dumps(payload))
    return r.json()


def insert(key="", value=""):
    if (key == ""):
        return
    payload = {'key': key, "value": value}
    r = requests.post(server + "insert", data=json.dumps(payload))
    return r.json()


def update(key="", value=""):
    if (key == ""):
        return
    payload = {'key': key, "value": value}
    r = requests.post(server + "update", data=json.dumps(payload))
    return r.json()

def dump():
    r = requests.get(server_admin + "dump")
    return r.json()

# 1.	Insert 10 pair, read it back – 5%
# 2.	Restart backup, on successful restart – 5%
# 3.	Delete 2 pair – without error return 5%
# 4.	Update 2 pair, read back the results 5%
# 5.	Restart primary, on successful restart – 5%
# 6.	Dump all key-values, and check with desired results – 35%


test_key_list = ["key1", "_key2", "^%!@#$%^&*()key3", "{key4", "key5_+=", "key6-=_+-[", "key7测试", "{key8=", "_key9\'\"",
            "]key10\\|"]
test_value_list = ["12421", "aslfjhalgha", "657468svca", "18726(^&(^(", "0chp3\"`", ")*HPB", "啦啦啦", "+++", "~!@#GX",
              "{\ndAFqw}"]

backup = dict()

def insert_test():
    for i in range(10):
        try:
            result = get(test_key_list[i])
            if result["success"]:
                print("Error when inserting \"{0}\"=\"{1}\", key exists".format(test_key_list[i], test_value_list[i]))
                break
            result = insert(test_key_list[i], test_value_list[i])
            if not result["success"]:
                print("Error when inserting \"{0}\"=\"{1}\", insert failed".format(test_key_list[i], test_value_list[i]))
                break
            result = get(test_key_list[i])
            if not result["success"] or result["value"]!=test_value_list[i]:
                print("Error when inserting \"{0}\"=\"{1}\", get error".format(test_key_list[i], test_value_list[i]))
            backup[test_key_list[i]] = test_value_list[i]
        except ValueError:
            print("Error when inserting \"{0}\"=\"{1}\", json failed".format(test_key_list[i], test_value_list[i]))
    print("Insert test finished")

def delete_test():
    for key in random.sample(backup.keys(), 2):
        result = delete(key)
        if (not result["success"]):
            print("Error when deleting \"{0}\"=\"{1}\"".format(key, backup[key]))
        del backup[key]
    print("delete test finished")

def update_test():
    for key in random.sample(backup.keys(), 2):
        result = update(key, "changed")
        if not result["success"]:
            print("Error when updating \"{0}\"=\"{1}\" to \"changed\"".format(key, backup[key]))
        backup[key] = "changed"
        result = get(key)
        if not result["success"] or result["value"]!="changed":
            print("Error when updating \"{0}\"=\"{1}\", value error".format(key, "changed"))
    print("update test finished")

def dump_test():
    result = dump()
    keys = backup.keys()
    for kv in result:
        if (kv[0] not in keys) or (backup[kv[0]])!=kv[1]:
            print("Error when checking \"{0}\"=\"{1}\", value error".format(key, "changed"))
    # print(result)
    print("Dump test finished")

def main():
    random.seed = time.clock()
    start_primary()
    start_backup()
    insert_test()
    delete_test()
    update_test()
    dump_test()
    # stop_backup()
    # start_backup()


main()