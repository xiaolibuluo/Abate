
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind91_0 { KIND_EMPTY91_0 = 0, KIND_FULL91_0 = 24 };

union Cell91_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind91_0(union Cell91_0 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL91_0 : KIND_EMPTY91_0;
}
#if defined(__STDC_VERSION__)
int variant91_1(int value, ...) {
    return value + 31;
}
#else
int variant91_1(int value) {
    return value + 31;
}
#endif
int gcd91_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale91_2(int value) {
    return gcd91_2(value, 58) * 38;
}
enum Kind91_3 { KIND_EMPTY91_3 = 0, KIND_FULL91_3 = 45 };

union Cell91_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind91_3(union Cell91_3 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL91_3 : KIND_EMPTY91_3;
}
int scan91_4(const int *begin, const int *end) {
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
#if defined(__STDC_VERSION__)
int variant91_5(int value, ...) {
    return value + 59;
}
#else
int variant91_5(int value) {
    return value + 59;
}
#endif
enum Kind91_6 { KIND_EMPTY91_6 = 0, KIND_FULL91_6 = 66 };

union Cell91_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind91_6(union Cell91_6 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL91_6 : KIND_EMPTY91_6;
}
struct Node91_7 {
    int value;
    struct Node91_7 *next;
};

int length91_7(const struct Node91_7 *head) {
    int count = 0;
    const struct Node91_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
