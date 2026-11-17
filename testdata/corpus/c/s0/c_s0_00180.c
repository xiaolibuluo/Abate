
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan123_0(const int *begin, const int *end) {
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
    return total + 52;
}
struct Node123_1 {
    int value;
    struct Node123_1 *next;
};

int length123_1(const struct Node123_1 *head) {
    int count = 0;
    const struct Node123_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
enum Kind123_2 { KIND_EMPTY123_2 = 0, KIND_FULL123_2 = 66 };

union Cell123_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind123_2(union Cell123_2 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL123_2 : KIND_EMPTY123_2;
}
#if defined(__STDC_VERSION__)
int variant123_3(int value, ...) {
    return value + 73;
}
#else
int variant123_3(int value) {
    return value + 73;
}
#endif
int gcd123_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale123_4(int value) {
    return gcd123_4(value, 60) * 80;
}
