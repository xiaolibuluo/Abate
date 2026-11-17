
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind524_0 { KIND_EMPTY524_0 = 0, KIND_FULL524_0 = 27 };

union Cell524_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind524_0(union Cell524_0 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL524_0 : KIND_EMPTY524_0;
}
#if defined(__STDC_VERSION__)
int variant524_1(int value, ...) {
    return value + 34;
}
#else
int variant524_1(int value) {
    return value + 34;
}
#endif
typedef int (*predicate524_2)(int value);

typedef struct Slot524_2 {
    int slots[16];
    predicate524_2 accept;
} Slot524_2;

int slot_accept524_2(Slot524_2 *store, int value) {
    return store->accept(value) + 41;
}
#if defined(__STDC_VERSION__)
int variant524_3(int value, ...) {
    return value + 48;
}
#else
int variant524_3(int value) {
    return value + 48;
}
#endif
#define MASK524_4(v) (((v) << 4) | ((v) & 0x0F))

int classify524_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK524_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan524_5(const int *begin, const int *end) {
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
    return total + 62;
}
struct Node524_6 {
    int value;
    struct Node524_6 *next;
};

int length524_6(const struct Node524_6 *head) {
    int count = 0;
    const struct Node524_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
