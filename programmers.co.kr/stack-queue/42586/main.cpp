#include <string>
#include <vector>

using namespace std;

vector<int> solution(vector<int> progresses, vector<int> speeds) {
	vector<int> answer;
	vector<int> days;

	int size = progresses.size();

	for (int i = 0; i < size; i++) {
		int day = (100 - progresses[i]) / speeds[i];

		if ((100 - progresses[i]) % speeds[i] > 0) {
			day++;
		}

		days.push_back(day);
	}

	int count = 0;
	int temp = 0;
	size = days.size();
	for (int i = 0; i < size; i++) {
		if (i != 0 && temp >= days[i]) {
			answer[answer.size() - 1]++;
		}
		else {
			answer.push_back(1);
			temp = days[i];
		}
	}
	return answer;
}