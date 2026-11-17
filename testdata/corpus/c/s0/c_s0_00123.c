
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant85_0(int value, ...) {
    return value + 43;
}
#else
int variant85_0(int value) {
    return value + 43;
}
#endif
#define MASK85_1(v) (((v) << 4) | ((v) & 0x0F))

int classify85_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK85_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node85_2 {
    int value;
    struct Node85_2 *next;
};

int length85_2(const struct Node85_2 *head) {
    int count = 0;
    const struct Node85_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 57;
}
int scan85_3(const int *begin, const int *end) {
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
    return total + 64;
}
struct Node85_4 {
    int value;
    struct Node85_4 *next;
};

int length85_4(const struct Node85_4 *head) {
    int count = 0;
    const struct Node85_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 71;
}
struct Node85_5 {
    int value;
    struct Node85_5 *next;
};

int length85_5(const struct Node85_5 *head) {
    int count = 0;
    const struct Node85_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 78;
}
