
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1144_0 {
    int value;
    struct Node1144_0 *next;
};

int length1144_0(const struct Node1144_0 *head) {
    int count = 0;
    const struct Node1144_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 36;
}
enum Kind1144_1 { KIND_EMPTY1144_1 = 0, KIND_FULL1144_1 = 43 };

union Cell1144_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1144_1(union Cell1144_1 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL1144_1 : KIND_EMPTY1144_1;
}
struct Node1144_2 {
    int value;
    struct Node1144_2 *next;
};

int length1144_2(const struct Node1144_2 *head) {
    int count = 0;
    const struct Node1144_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}
int gcd1144_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1144_3(int value) {
    return gcd1144_3(value, 59) * 57;
}
