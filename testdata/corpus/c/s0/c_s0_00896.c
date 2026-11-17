
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd617_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale617_0(int value) {
    return gcd617_0(value, 58) * 72;
}
int gcd617_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale617_1(int value) {
    return gcd617_1(value, 69) * 79;
}
enum Kind617_2 { KIND_EMPTY617_2 = 0, KIND_FULL617_2 = 86 };

union Cell617_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind617_2(union Cell617_2 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL617_2 : KIND_EMPTY617_2;
}
int gcd617_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale617_3(int value) {
    return gcd617_3(value, 30) * 93;
}
enum Kind617_4 { KIND_EMPTY617_4 = 0, KIND_FULL617_4 = 100 };

union Cell617_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind617_4(union Cell617_4 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL617_4 : KIND_EMPTY617_4;
}
struct Node617_5 {
    int value;
    struct Node617_5 *next;
};

int length617_5(const struct Node617_5 *head) {
    int count = 0;
    const struct Node617_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
