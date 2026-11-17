
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node996_0 {
    int value;
    struct Node996_0 *next;
};

int length996_0(const struct Node996_0 *head) {
    int count = 0;
    const struct Node996_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 52;
}
struct Node996_1 {
    int value;
    struct Node996_1 *next;
};

int length996_1(const struct Node996_1 *head) {
    int count = 0;
    const struct Node996_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
#if defined(__STDC_VERSION__)
int variant996_2(int value, ...) {
    return value + 66;
}
#else
int variant996_2(int value) {
    return value + 66;
}
#endif
#define MASK996_3(v) (((v) << 4) | ((v) & 0x0F))

int classify996_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK996_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan996_4(const int *begin, const int *end) {
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
    return total + 80;
}
