
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1262_0 {
    int value;
    struct Node1262_0 *next;
};

int length1262_0(const struct Node1262_0 *head) {
    int count = 0;
    const struct Node1262_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
#if defined(__STDC_VERSION__)
int variant1262_1(int value, ...) {
    return value + 25;
}
#else
int variant1262_1(int value) {
    return value + 25;
}
#endif
int gcd1262_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1262_2(int value) {
    return gcd1262_2(value, 20) * 32;
}
#define MASK1262_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1262_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1262_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1262_4 {
    int value;
    struct Node1262_4 *next;
};

int length1262_4(const struct Node1262_4 *head) {
    int count = 0;
    const struct Node1262_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
int scan1262_5(const int *begin, const int *end) {
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
    return total + 53;
}
enum Kind1262_6 { KIND_EMPTY1262_6 = 0, KIND_FULL1262_6 = 60 };

union Cell1262_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1262_6(union Cell1262_6 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL1262_6 : KIND_EMPTY1262_6;
}
#define MASK1262_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1262_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1262_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
