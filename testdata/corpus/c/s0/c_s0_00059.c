
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant38_0(int value, ...) {
    return value + 14;
}
#else
int variant38_0(int value) {
    return value + 14;
}
#endif
enum Kind38_1 { KIND_EMPTY38_1 = 0, KIND_FULL38_1 = 21 };

union Cell38_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind38_1(union Cell38_1 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL38_1 : KIND_EMPTY38_1;
}
#if defined(__STDC_VERSION__)
int variant38_2(int value, ...) {
    return value + 28;
}
#else
int variant38_2(int value) {
    return value + 28;
}
#endif
#define MASK38_3(v) (((v) << 4) | ((v) & 0x0F))

int classify38_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK38_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind38_4 { KIND_EMPTY38_4 = 0, KIND_FULL38_4 = 42 };

union Cell38_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind38_4(union Cell38_4 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL38_4 : KIND_EMPTY38_4;
}
struct Node38_5 {
    int value;
    struct Node38_5 *next;
};

int length38_5(const struct Node38_5 *head) {
    int count = 0;
    const struct Node38_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 49;
}
#if defined(__STDC_VERSION__)
int variant38_6(int value, ...) {
    return value + 56;
}
#else
int variant38_6(int value) {
    return value + 56;
}
#endif
int scan38_7(const int *begin, const int *end) {
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
    return total + 63;
}
