
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind816_0 { KIND_EMPTY816_0 = 0, KIND_FULL816_0 = 40 };

union Cell816_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind816_0(union Cell816_0 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL816_0 : KIND_EMPTY816_0;
}
int scan816_1(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 47;
}
int gcd816_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale816_2(int value) {
    return gcd816_2(value, 9) * 54;
}
int scan816_3(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 61;
}
struct Node816_4 {
    int value;
    struct Node816_4 *next;
};

int length816_4(const struct Node816_4 *head) {
    int count = 0;
    const struct Node816_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
int gcd816_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale816_5(int value) {
    return gcd816_5(value, 42) * 75;
}
