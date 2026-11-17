
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node428_0 {
    int value;
    struct Node428_0 *next;
};

int length428_0(const struct Node428_0 *head) {
    int count = 0;
    const struct Node428_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 40;
}
#if defined(__STDC_VERSION__)
int variant428_1(int value, ...) {
    return value + 47;
}
#else
int variant428_1(int value) {
    return value + 47;
}
#endif
int gcd428_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale428_2(int value) {
    return gcd428_2(value, 38) * 54;
}
enum Kind428_3 { KIND_EMPTY428_3 = 0, KIND_FULL428_3 = 61 };

union Cell428_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind428_3(union Cell428_3 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL428_3 : KIND_EMPTY428_3;
}
int scan428_4(const int *begin, const int *end) {
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
    return total + 68;
}
int scan428_5(const int *begin, const int *end) {
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
    return total + 75;
}
