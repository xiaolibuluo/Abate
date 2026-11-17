
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1889_0(const int *begin, const int *end) {
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
    return total + 21;
}
#if defined(__STDC_VERSION__)
int variant1889_1(int value, ...) {
    return value + 28;
}
#else
int variant1889_1(int value) {
    return value + 28;
}
#endif
struct Node1889_2 {
    int value;
    struct Node1889_2 *next;
};

int length1889_2(const struct Node1889_2 *head) {
    int count = 0;
    const struct Node1889_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 35;
}
int scan1889_3(const int *begin, const int *end) {
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
    return total + 42;
}
int gcd1889_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1889_4(int value) {
    return gcd1889_4(value, 39) * 49;
}
enum Kind1889_5 { KIND_EMPTY1889_5 = 0, KIND_FULL1889_5 = 56 };

union Cell1889_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1889_5(union Cell1889_5 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL1889_5 : KIND_EMPTY1889_5;
}
#define MASK1889_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1889_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1889_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
