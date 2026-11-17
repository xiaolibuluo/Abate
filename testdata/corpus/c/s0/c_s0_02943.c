
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1990_0 {
    int value;
    struct Node1990_0 *next;
};

int length1990_0(const struct Node1990_0 *head) {
    int count = 0;
    const struct Node1990_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
struct Node1990_1 {
    int value;
    struct Node1990_1 *next;
};

int length1990_1(const struct Node1990_1 *head) {
    int count = 0;
    const struct Node1990_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 80;
}
struct Node1990_2 {
    int value;
    struct Node1990_2 *next;
};

int length1990_2(const struct Node1990_2 *head) {
    int count = 0;
    const struct Node1990_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
#define MASK1990_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1990_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1990_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1990_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1990_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1990_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1990_5(const int *begin, const int *end) {
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
    return total + 11;
}
#if defined(__STDC_VERSION__)
int variant1990_6(int value, ...) {
    return value + 18;
}
#else
int variant1990_6(int value) {
    return value + 18;
}
#endif
