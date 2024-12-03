filename = "input.txt"


def is_report_safe(report: list) -> bool:
    total_nums = len(report)
    decreases = 0
    # Is the list strictly decreasing
    for i, j in zip(report, report[1:]):
        if int(i) > int(j) and 1 <= abs(int(i) - int(j)) <= 3:
            decreases += 1
        else:
            break

    if decreases == (total_nums - 1):
        return True

    increases = 0
    # Is the list strictly increasing
    for i, j in zip(report, report[1:]):
        if int(i) < int(j) and 1 <= abs(int(i) - int(j)) <= 3:
            increases += 1
        else:
            break

    if increases == (total_nums - 1):
        return True

    return False


safe_reports = 0

with open(filename) as file:
    for line in file:
        report = line.rstrip().split()
        safe = is_report_safe(report)
        if safe:
            safe_reports += 1

print(f"Part 1: Count of safe reports is: {safe_reports}")

safe_reports = 0
with open(filename) as file:
    for line in file:
        report = line.rstrip().split()
        safe_dec = is_report_safe(report)
        safe_inc = is_report_safe(report)
        if safe_dec or safe_inc:
            safe_reports += 1
        else:
            for i in range(len(report)):
                subset = report[:i] + report[i+1:]
                safe_dec = is_report_safe(subset)
                safe_inc = is_report_safe(subset)
                if safe_dec or safe_inc:
                    safe_reports += 1
                    break

print(f"Part 2: Count of safe reports is: {safe_reports}")
