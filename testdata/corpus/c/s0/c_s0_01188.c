
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind815_0 { KIND_EMPTY815_0 = 0, KIND_FULL815_0 = 27 };

union Cell815_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind815_0(union Cell815_0 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL815_0 : KIND_EMPTY815_0;
}
#define MASK815_1(v) (((v) << 4) | ((v) & 0x0F))

int classify815_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK815_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan815_2(const int *begin, const int *end) {
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
    return total + 41;
}
#define MASK815_3(v) (((v) << 4) | ((v) & 0x0F))

int classify815_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK815_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant815_4(int value, ...) {
    return value + 55;
}
#else
int variant815_4(int value) {
    return value + 55;
}
#endif
struct Node815_5 {
    int value;
    struct Node815_5 *next;
};

int length815_5(const struct Node815_5 *head) {
    int count = 0;
    const struct Node815_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 62;
}
int gcd815_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale815_6(int value) {
    return gcd815_6(value, 46) * 69;
}
