
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant180_0(int value, ...) {
    return value + 17;
}
#else
int variant180_0(int value) {
    return value + 17;
}
#endif
enum Kind180_1 { KIND_EMPTY180_1 = 0, KIND_FULL180_1 = 24 };

union Cell180_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind180_1(union Cell180_1 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL180_1 : KIND_EMPTY180_1;
}
struct Node180_2 {
    int value;
    struct Node180_2 *next;
};

int length180_2(const struct Node180_2 *head) {
    int count = 0;
    const struct Node180_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}
int gcd180_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale180_3(int value) {
    return gcd180_3(value, 21) * 38;
}
enum Kind180_4 { KIND_EMPTY180_4 = 0, KIND_FULL180_4 = 45 };

union Cell180_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind180_4(union Cell180_4 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL180_4 : KIND_EMPTY180_4;
}
int scan180_5(const int *begin, const int *end) {
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
