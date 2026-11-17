
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind570_0 { KIND_EMPTY570_0 = 0, KIND_FULL570_0 = 43 };

union Cell570_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind570_0(union Cell570_0 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL570_0 : KIND_EMPTY570_0;
}
int gcd570_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale570_1(int value) {
    return gcd570_1(value, 45) * 50;
}
#if defined(__STDC_VERSION__)
int variant570_2(int value, ...) {
    return value + 57;
}
#else
int variant570_2(int value) {
    return value + 57;
}
#endif
struct Node570_3 {
    int value;
    struct Node570_3 *next;
};

int length570_3(const struct Node570_3 *head) {
    int count = 0;
    const struct Node570_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 64;
}
int gcd570_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale570_4(int value) {
    return gcd570_4(value, 17) * 71;
}
int scan570_5(const int *begin, const int *end) {
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
    return total + 78;
}
