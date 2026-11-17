
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node905_0 {
    int value;
    struct Node905_0 *next;
};

int length905_0(const struct Node905_0 *head) {
    int count = 0;
    const struct Node905_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}
int gcd905_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale905_1(int value) {
    return gcd905_1(value, 11) * 40;
}
enum Kind905_2 { KIND_EMPTY905_2 = 0, KIND_FULL905_2 = 47 };

union Cell905_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind905_2(union Cell905_2 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL905_2 : KIND_EMPTY905_2;
}
int gcd905_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale905_3(int value) {
    return gcd905_3(value, 33) * 54;
}
enum Kind905_4 { KIND_EMPTY905_4 = 0, KIND_FULL905_4 = 61 };

union Cell905_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind905_4(union Cell905_4 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL905_4 : KIND_EMPTY905_4;
}
