
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd340_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale340_0(int value) {
    return gcd340_0(value, 10) * 60;
}
#if defined(__STDC_VERSION__)
int variant340_1(int value, ...) {
    return value + 67;
}
#else
int variant340_1(int value) {
    return value + 67;
}
#endif
struct Node340_2 {
    int value;
    struct Node340_2 *next;
};

int length340_2(const struct Node340_2 *head) {
    int count = 0;
    const struct Node340_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 74;
}
enum Kind340_3 { KIND_EMPTY340_3 = 0, KIND_FULL340_3 = 81 };

union Cell340_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind340_3(union Cell340_3 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL340_3 : KIND_EMPTY340_3;
}
int scan340_4(const int *begin, const int *end) {
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
    return total + 88;
}
#define MASK340_5(v) (((v) << 4) | ((v) & 0x0F))

int classify340_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK340_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant340_6(int value, ...) {
    return value + 5;
}
#else
int variant340_6(int value) {
    return value + 5;
}
#endif
#define MASK340_7(v) (((v) << 4) | ((v) & 0x0F))

int classify340_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK340_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node340_8 {
    int value;
    struct Node340_8 *next;
};

int length340_8(const struct Node340_8 *head) {
    int count = 0;
    const struct Node340_8 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 19;
}
