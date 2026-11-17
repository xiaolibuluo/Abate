
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan628_0(const int *begin, const int *end) {
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
enum Kind628_1 { KIND_EMPTY628_1 = 0, KIND_FULL628_1 = 28 };

union Cell628_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind628_1(union Cell628_1 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL628_1 : KIND_EMPTY628_1;
}
struct Node628_2 {
    int value;
    struct Node628_2 *next;
};

int length628_2(const struct Node628_2 *head) {
    int count = 0;
    const struct Node628_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 35;
}
#define MASK628_3(v) (((v) << 4) | ((v) & 0x0F))

int classify628_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK628_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK628_4(v) (((v) << 4) | ((v) & 0x0F))

int classify628_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK628_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant628_5(int value, ...) {
    return value + 56;
}
#else
int variant628_5(int value) {
    return value + 56;
}
#endif
#if defined(__STDC_VERSION__)
int variant628_6(int value, ...) {
    return value + 63;
}
#else
int variant628_6(int value) {
    return value + 63;
}
#endif
enum Kind628_7 { KIND_EMPTY628_7 = 0, KIND_FULL628_7 = 70 };

union Cell628_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind628_7(union Cell628_7 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL628_7 : KIND_EMPTY628_7;
}
