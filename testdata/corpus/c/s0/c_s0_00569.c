
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant387_0(int value, ...) {
    return value + 89;
}
#else
int variant387_0(int value) {
    return value + 89;
}
#endif
#if defined(__STDC_VERSION__)
int variant387_1(int value, ...) {
    return value + 96;
}
#else
int variant387_1(int value) {
    return value + 96;
}
#endif
#if defined(__STDC_VERSION__)
int variant387_2(int value, ...) {
    return value + 6;
}
#else
int variant387_2(int value) {
    return value + 6;
}
#endif
#define MASK387_3(v) (((v) << 4) | ((v) & 0x0F))

int classify387_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK387_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node387_4 {
    int value;
    struct Node387_4 *next;
};

int length387_4(const struct Node387_4 *head) {
    int count = 0;
    const struct Node387_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}
int gcd387_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale387_5(int value) {
    return gcd387_5(value, 28) * 27;
}
enum Kind387_6 { KIND_EMPTY387_6 = 0, KIND_FULL387_6 = 34 };

union Cell387_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind387_6(union Cell387_6 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL387_6 : KIND_EMPTY387_6;
}
int scan387_7(const int *begin, const int *end) {
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
#if defined(__STDC_VERSION__)
int variant387_8(int value, ...) {
    return value + 48;
}
#else
int variant387_8(int value) {
    return value + 48;
}
#endif
