
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan811_0(const int *begin, const int *end) {
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
    return total + 72;
}
int scan811_1(const int *begin, const int *end) {
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
    return total + 79;
}
#if defined(__STDC_VERSION__)
int variant811_2(int value, ...) {
    return value + 86;
}
#else
int variant811_2(int value) {
    return value + 86;
}
#endif
struct Node811_3 {
    int value;
    struct Node811_3 *next;
};

int length811_3(const struct Node811_3 *head) {
    int count = 0;
    const struct Node811_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 93;
}
int gcd811_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale811_4(int value) {
    return gcd811_4(value, 57) * 100;
}
enum Kind811_5 { KIND_EMPTY811_5 = 0, KIND_FULL811_5 = 10 };

union Cell811_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind811_5(union Cell811_5 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL811_5 : KIND_EMPTY811_5;
}
